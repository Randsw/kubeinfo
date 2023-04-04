# Build the kubeinfo binary
FROM golang:1.19 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY logger/ logger/
COPY prometheus-exporter/ prometheus-exporter/
COPY k8sclient/ k8sclient/
COPY KubeApiResponseStruct/ KubeApiResponseStruct/
COPY cmd/ cmd/
COPY main.go main.go

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o kubeinfo main.go

# Use distroless as minimal base image to package the kubeinfo binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/kubeinfo .
USER 65532:65532

ENTRYPOINT ["/kubeinfo"]