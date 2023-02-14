package main

import (
	"fmt"
	"os"

	"github.com/goMLLibrary/core/graph"
	"github.com/goMLLibrary/core/mnist"
	"github.com/goMLLibrary/core/neuralNetwork"
)

func main() {
	// ニューラルネットワーク層をまとめるレイヤーの作成
	layers := neuralNetwork.NewDefaultNeuralNetworkLayers()

	// 1層目
	layers.Add(neuralNetwork.NewAffine(28*28, 10