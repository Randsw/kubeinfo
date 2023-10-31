# Build the kubeinfo binary
FROM golang:1.21 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY . .

RUN ls -lah

RUN make test

RUN make pwd
# Build
RUN make build

# Use distroless as minimal base image to package the kubeinfo binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/kubeinfo .
USER 65532:65532

ENTRYPOINT ["/kubeinfo"]