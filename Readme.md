<h1 align="center" style="border-bottom: none;">Kubeinfo</h1>
<h3 align="center">Get information about resources in kubernetes cluster</h3>
</p>
<p align="center">
<a href="" alt="GoVersion">
        <img src="https://img.shields.io/github/go-mod/go-version/randsw/kubeinfo" /></a>
<a href="" alt="Release">
        <img src="https://img.shields.io/github/v/release/randsw/kubeinfo" /></a>
<a href="" alt="LastTag">
        <img src="https://img.shields.io/github/v/tag/randsw/kubeinfo" /></a>
<a href="" alt="LastCommit">
        <img src="https://img.shields.io/github/last-commit/randsw/kubeinfo" /></a>
<a href="" alt="Build">
        <img src="https://img.shields.io/github/actions/workflow/status/randsw/kubeinfo/build.yaml" /></a>
<a href="" alt="Build">
        <img src="https://img.shields.io/github/actions/workflow/status/randsw/kubeinfo/test_release.yaml?label=tests" /></a>
<a href="" alt="Build">
        <img src="https://img.shields.io/github/actions/workflow/status/randsw/kubeinfo/helm-chart-test.yaml?label=helm-chart-test" /></a></p>

## Introduction

:warning: **Attention**
**This project created for educational purposes and serves as proof-of-concept. Another goal is to learn how to work with Kubernetes CRD and CR using client-go library. As a CRD for information i choose my favorite GitOps tool - [FluxCD](https://fluxcd.io/)**

## Content

1. [Requirements](#requirements)
2. [Overview](#overview)
   * [Metrics](#metrics)
   * [Health](#health)
   * [Swagger](#swagger)
   * [Configuration](#configuration)
3. [Examples](#examples)
   * [Docker](#docker)
   * [Kubernetes](#kubernetes)

## Requirements

[Helm](https://helm.sh/docs/intro/install/) version **3.10.1** or newer
[Kubernetes](https://kubernetes.io) version **1.23.1** or newer
[Fluxcd](https://fluxcd.io/) version **0.40.1** or newer

To build the container image locally you need [Docker](https://docs.docker.com/engine/install/) with **BuildKit** enabled — the [`Dockerfile`](Dockerfile) uses `# syntax=docker/dockerfile:1` and cache mounts, which require it. BuildKit is the default since Docker **23.0**; on older versions export `DOCKER_BUILDKIT=1`. The **buildx** plugin (bundled since Docker **20.10**) is only needed for multi-platform builds.

Building is not required to run kubeinfo — prebuilt images are published to the GitHub Container Registry, see [Docker](#docker).

## Overview

This application listen at port **8080** and provides several endpoints
| endpoint            | HTTP Verb             |return       | return value|
|---------------------|-----------------------|-------------|-------------|
| /                   | GET                   | All of the following | [struct](https://github.com/Randsw/kubeinfo/blob/e27d51c9f40db18d2ad56718052483812042f68d/KubeApiResponseStruct/kubeapiresponsestruct.go#L53-L60) |
| /nodes              | GET                   | Overall number of nodes and number of nodes by role(control-plane or worker) | [struct](https://github.com/Randsw/kubeinfo/blob/e27d51c9f40db18d2ad56718052483812042f68d/KubeApiResponseStruct/kubeapiresponsestruct.go#L10-L16) |
| /namespaces         | GET                   | Overall number of namespaces | [struct](https://github.com/Randsw/kubeinfo/blob/e27d51c9f40db18d2ad56718052483812042f68d/KubeApiResponseStruct/kubeapiresponsestruct.go#L18-L20) |
| /pods               | GET                   | Overall number of pods and number of pods by phase | [struct](https://github.com/Randsw/kubeinfo/blob/e27d51c9f40db18d2ad56718052483812042f68d/KubeApiResponseStruct/kubeapiresponsestruct.go#L22-L25)
| /ingresses          | GET                   | Overall number of ingresses | [struct](https://github.com/Randsw/kubeinfo/blob/e27d51c9f40db18d2ad56718052483812042f68d/KubeApiResponseStruct/kubeapiresponsestruct.go#L27-L29) |
| /fluxkustomizations | GET                   | Overall number of fluxkustomizations and number of fluxkustomizations by status | [struct](https://github.com/Randsw/kubeinfo/blob/e27d51c9f40db18d2ad56718052483812042f68d/KubeApiResponseStruct/kubeapiresponsestruct.go#L37-L40) |
| /fluxhelmreleases   | GET                   | Overall number of fluxhelmreleases and number of fluxhelmreleases by status | [struct](https://github.com/Randsw/kubeinfo/blob/e27d51c9f40db18d2ad56718052483812042f68d/KubeApiResponseStruct/kubeapiresponsestruct.go#L48-L51) |

### Metrics

Prometheus metrcis exposed on port **8080** via `/metrics` endpoint.

### Health

Healthcheck provided by `/healthz` endpoint on port **8080**. Along with the `OK` status it reports the build metadata baked in at compile time — the image version (`tag`), the commit (`hash`) and the build date (`date`). See [Build the image locally](#build-the-image-locally) for how these values are injected.

### Swagger

An interactive OpenAPI (Swagger) UI is served at `/swagger/` on port **8080**.

### Configuration

The application is configured entirely through environment variables:

| variable      | default | description                                                    |
|---------------|---------|----------------------------------------------------------------|
| `API_PORT`    | `8080`  | Port the HTTP server listens on                                |
| `API_ADDRESS` | `""`    | Address to bind to (empty means all interfaces)                |
| `KUBECONFIG`  | unset   | Path to the kubeconfig used when running outside a cluster     |

## Examples

### Docker

#### Run the published image

Prebuilt images are published to the GitHub Container Registry on every tag push. The binary listens on port **8080**, so expose it to reach the endpoints described in the [Overview](#overview):

`docker run --rm -p 8080:8080 ghcr.io/randsw/kubeinfo:latest`

The container has no kubeconfig of its own, so to inspect a cluster reachable from your host, mount your kubeconfig read-only and point `KUBECONFIG` at it:

`docker run --rm -p 8080:8080 -e KUBECONFIG=/kubeconfig -v ~/.kube/config:/kubeconfig:ro ghcr.io/randsw/kubeinfo:latest`

> Inside a Kubernetes cluster the application uses the in-cluster service account credentials instead of a kubeconfig — see [Kubernetes](#kubernetes).

#### Build the image locally

To build the image with the application run

`docker build -t kubeinfo:<your-tag> .`

Then push the container to your favorite registry.

The [`Dockerfile`](Dockerfile) is a multi-stage build worth knowing about if you plan to tweak it:

| Aspect | What it does |
|--------|--------------|
| Multi-arch | The builder runs `--platform=$BUILDPLATFORM` and cross-compiles with `GOOS`/`GOARCH`, so `linux/arm64` images are built natively instead of emulating the Go toolchain under QEMU |
| Static binary | `CGO_ENABLED=0` produces a binary with no libc dependency, which is what makes the distroless runtime work |
| Minimal runtime | `gcr.io/distroless/static:nonroot` — no shell, no package manager, no libc to exploit |
| Non-root | Runs as UID/GID `65532` |
| Layer caching | `go mod download` and the source copies are separate layers with BuildKit cache mounts, so source-only changes don't re-download dependencies |

To bake version metadata into the binary (the same values CI passes) use build args:

`docker build -t kubeinfo:<your-tag> --build-arg VERSION=<version> --build-arg COMMIT=<commit-sha> --build-arg BUILD_DATE=<date> .`

These values are reported by the `/healthz` endpoint. The build never shells out to git — `.git` is excluded from the build context by [`.dockerignore`](.dockerignore) — so passing `COMMIT` is the only way to get an accurate hash into the image. Without any build args the defaults (`dev`, `unknown`) are used, so a plain `docker build` still works.

To build for several platforms at once, use buildx:

`docker buildx build --platform linux/amd64,linux/arm64 -t kubeinfo:<your-tag> .`

### Kubernetes

To deploy kubeinfo in kubernetes cluster use helm package manager

The chart deploys the image built by the [`Dockerfile`](Dockerfile) — by default `ghcr.io/randsw/kubeinfo:latest`. Since the image is distroless and runs as UID/GID `65532` with no shell, set `securityContext`/`podSecurityContext` accordingly if you override them (for example `runAsNonRoot: true`, `runAsUser: 65532`).

First clone this repo

`git clone https://github.com/Randsw/kubeinfo.git`

Then go to folder `helm-chart/kubeinfo-backend`

`cd helm-chart/kubeinfo-backend`

And install helm chart in your kubernetes cluster

`helm upgrade --install --namespace <your-namepsace> --create-namespace <your-release-name> .`

Don\`t forget edit `values.yaml` for your needs, specially rigth image name and other settings or just provide your values file

`helm upgrade --install --namespace <your-namepsace> --create-namespace -f <your-values-file> <your-release-name> .`

If this code helped solve your problem or helped point you in the right direction, I will be very happy!
