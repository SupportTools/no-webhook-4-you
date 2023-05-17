# no-webhook-4-you

[![Build Status](https://drone.support.tools/api/badges/SupportTools/no-webhook-4-you/status.svg?ref=refs/heads/main)](https://drone.support.tools/SupportTools/no-webhook-4-you)

## Description
 The No Webhook 4 You is a Go application that provides as workaround to Rancher's webhook being redeployed automatically even if you disable it. This tool watches for the webhook creation and deletes it if it is found.

## Features
- Watches for changes in ValidatingWebhookConfigurations and MutatingWebhookConfigurations in a Kubernetes cluster.
- Deletes specific configurations with the name "rancher.cattle.io".

# Prerequisites
- Go programming language (version 1.16 or later)
- Kubernetes cluster
- Access to the Kubernetes cluster with sufficient privileges

# Installation

## To Kubernetes cluster
1. Run the following command to install the application to the Kubernetes cluster:
```bash
kubectl apply -f https://raw.githubusercontent.com/SupportTools/no-webhook-4-you/main/deploy/no-webhook-4-you.yaml
```

## From source
1. Clone the repository
2. Run `go build` to build the binary
3. Run `./no-webhook-4-you` to run the application

## From Docker
1. Run `docker build -t no-webhook-4-you .` to build the Docker image
2. Run `docker run -d no-webhook-4-you` to run the application

# Configuration
The application does not require any additional configuration. It uses the default Kubernetes client configuration provided by the kubeconfig file or in-cluster configuration.

# Contributing
Contributions are welcome! If you find any issues or have suggestions for improvement, please open an issue or submit a pull request.

# License
This project is licensed under the Apache License 2.0. See [LICENSE](LICENSE) for details.