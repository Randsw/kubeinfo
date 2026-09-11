# syntax=docker/dockerfile:1

# Build the kubeinfo binary.
# The builder runs natively on the build platform and cross-compiles to TARGETOS/TARGETARCH,
# so multi-arch builds don't have to emulate the Go toolchain under QEMU.
FROM --platform=$BUILDPLATFORM golang:1.27 AS builder

ARG TARGETOS
ARG TARGETARCH

# Version metadata baked into handlers.tag/hash/date and served by /healthz.
# CI passes the real values; the defaults keep a plain `docker build` working.
ARG VERSION=dev
ARG COMMIT=unknown
ARG BUILD_DATE=unknown

WORKDIR /workspace

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Copy only the sources the build needs so unrelated file changes don't invalidate this layer
COPY main.go ./
COPY cmd/ cmd/
COPY docs/ docs/
COPY handlers/ handlers/
COPY k8sclient/ k8sclient/
COPY KubeApiResponseStruct/ KubeApiResponseStruct/
COPY logger/ logger/
COPY prometheus-exporter/ prometheus-exporter/

# Build a fully static binary: CGO is off so the result can run on the libc-free
# distroless static runtime, and -trimpath keeps build paths out of the binary.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} \
    go build -trimpath -ldflags "-s -w \
        -X github.com/randsw/kubeinfo/handlers.tag=${VERSION} \
        -X github.com/randsw/kubeinfo/handlers.hash=${COMMIT} \
        -X github.com/randsw/kubeinfo/handlers.date=${BUILD_DATE}" \
    -o /out/kubeinfo .

# Use distroless as minimal base image to package the kubeinfo binary: no shell,
# no package manager, and no libc to exploit.
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot

COPY --link --from=builder /out/kubeinfo /kubeinfo

# The distroless nonroot tag already sets this UID/GID; restating it keeps the
# non-root guarantee explicit for image scanners and reviewers.
USER 65532:65532

ENTRYPOINT ["/kubeinfo"]
