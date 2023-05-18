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
3. [Examples](#examples)
   * [Docker](#docker)
   * [Kubernetes](#kubernetes)

## Requirements

[Helm](https://helm.sh/docs/intro/install/) version **3.10.1** or newer
[Kubernetes](https://kubernetes.io) version **1.23.1** or newer
[Fluxcd](https://fluxcd.io/) version **0.40.1** or newer

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

Healthcheck provided by `/healthz` endpoint on port **8080**

## Examples

### Docker

To build image with application run

`docker build-t kubeinfo:<your-tag> .`

Then push container to your favorite registry

### Kubernetes

To deploy kubeinfo in kubernetes cluster use helm package manager

First clone this repo

`git clone https://github.com/Randsw/kubeinfo.git`

Then go to folder `helm-chart/kubeinfo-backend`

`cd helm-chart/kubeinfo-backend`

And install helm chart in your kubernetes cluster

`helm upgrade --install --namespace <your-namepsace> --create-namespace <your-release-name> .`

Don\`t forget edit `values.yaml` for your needs, specially rigth image name and other settings or just provide your values file

`helm upgrade --install --namespace <your-namepsace> --create-namespace -f <your-values-file> <your-release-name> .`

If this code helped solve your problem or helped point you in the right direction, I will be very happy!
