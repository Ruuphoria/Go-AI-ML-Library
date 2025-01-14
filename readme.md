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
$ cd go/src/github.com/Go-AI-ML-Library
$ bash install/setup_container.sh
$ source ~/.bash_profile
```

## Run Sample TVM code

If you want to use the sample TVM code, you'll need to create a docker container.
All the following steps can be implemented inside docker container.

### Navigate to AI-ML-Library root path

```bash
$ cd go/src/github.com/Go-AI-ML-Library
```

### Get onnx model

```bash
$ cd python
$ python3 tvm_wrapper.py
$ cd ../
```

### Build sample code

```bash
$ export GOPATH=/home/development/go
$ bash build_tvm_sample.sh
```

### Run Inference

```bash
$ go run tvm_sample/GotvmSample
```