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
	layers.Add(neuralNetwork.NewAffine(28*28, 1000))
	layers.Add(neuralNetwork.NewRelu())

	// 2層目
	layers.Add(neuralNetwork.NewAffine(1000, 1000))
	layers.Add(neuralNetwork.NewRelu())

	// 3層目
	layers.Add(neuralNetwork.NewAffine(1000, 10))

	// MNISTデータセットを格納するためのフォルダを作成
	os.Mkdir("data", 0777)

	// MNISTのデータセットを取得
	train, test, err := mnist.LoadData("data")
	if err != nil {
		fmt.Printf("Can't get mnist train and test data! \n")
		os.Exit(-1)
	}

	// 学習時の各種パラメーターの設定
	batchSize := 100
	iterationCount := 1800
	iteracionCountPerEpoch := int(train.Count() / batchSize)

	// 学習時