# Go-AI-ML-Library

## Overview

Go-AI-ML-Library is a comprehensive library housing various neural network components implemented in Golang.

## Environment

* Go 1.12(using Go Modules)
* Docker

## Build & Run

* Make sure to use Go Modules

```bash
$ export GO111MODULE=on
```

* To run the mnist sample use the following command:

```bash
$ go run ./main/MnistSample.go
```
## Layer

### Activation

* Relu
* Sigmoid
* Tanh
* SoftmaxWithCrossEntropy

### Neural Network Cells

* Affine

### Optimizer

* Stochastic Gradient Descent (SGD)

## Docker

### Build Container

* Create docker container and attach it

```bash
$ bash install/run_container.sh
```

* Set up the environment inside docker container:

```bash
$ 