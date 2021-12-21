
package model

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"

	"github.com/goMLLibrary/core/neuralNetwork"
	"gonum.org/v1/gonum/mat"
)

// WriteNNLayers : ニューラルネットワークの情報をファイルに書き出す
func WriteNNLayers(modelPath string, nnLayers *neuralNetwork.NeuralNetworkLayers) error {
	// レイヤー情報を保存用のモデル情報に書き換える
	nnModel, err := convertNNModel(nnLayers)
	if err != nil {
		return err
	}

	// モデル情報をbyteデータに書き換え、ファイルに書き込む
	byteData, err := encodeNNModel(nnModel)
	if err != nil {
		return err
	}
	return writeModelFile(modelPath, byteData)
}

// ReadNNLayers : ニューラルネットワークの情報をファイルから取得する
func ReadNNLayers(modelPath string) (*neuralNetwork.NeuralNetworkLayers, error) {
	// ファイルからmodelのbyteデータを取得
	byteData, err := readModelFile(modelPath)
	if err != nil {
		return nil, err
	}

	// byteデータからモデル情報を作成
	nnModel, err := decodeNNModel(byteData)
	if err != nil {
		return nil, err
	}

	// モデル情報からレイヤー情報を復元する
	return convertNNLayers(nnModel)
}

func convertNNModel(nnLayers *neuralNetwork.NeuralNetworkLayers) (*NNModel, error) {
	nnModel := NewNNModel()

	// レイヤー情報を取得
	for _, layer := range nnLayers.GetLayers() {
		nnData := NewNNData()

		switch convertLayer := layer.(type) {
		case *neuralNetwork.Affine:
			nnData = convertNNDataFromAffine(convertLayer)
		case *neuralNetwork.Tanh:
			nnData.Type = TanhType
		case *neuralNetwork.Relu:
			nnData.Type = ReluType
		case *neuralNetwork.Sigmoid:
			nnData.Type = SigmoidType
		default:
			return nil, errors.New("意図しないレイヤータイプが指定されています.")
		}

		nnModel.Layers = append(nnModel.Layers, nnData)
	}

	// 最終のレイヤーを設定
	nnData := NewNNData()
	nnData.Type = SoftmaxWithLossType
	nnModel.Layers = append(nnModel.Layers, nnData)

	// Optimizerを設定
	nnData = NewNNData()
	optimizer := nnLayers.GetOptimizer()
	switch optimizer.(type) {
	case *neuralNetwork.SGD:
		nnData.Type = SgdType