package model

import (
	"os"
	"reflect"
	"testing"

	"github.com/goMLLibrary/core/neuralNetwork"
	"github.com/goMLLibrary/core/util"
	. "github.com/smartystreets/goconvey/convey"
	"gonum.org/v1/gonum/mat"
)

func TestModelHandler(t *testing.T) {
	Convey("Given : 2層のニューラルネットワークの情報が与えられた時", t, func() {
		nnLayers := neuralNetwork.NewDefaultNeuralNetworkLayers()

		modelPath := "model.db"
		defer os.Remove(modelPath)

		Convey("AND : 1層目：5*10のAffineレイヤーを作成", nil)
		inputSize := 5
		outputSize := 10
		affine := neuralNetwork.NewAffine(inputSize, outputSize)

		Convey("AND : 1層目：Affineレイヤーのパラメーターを初期化し設定", nil)
		w := mat.NewDense(inputSize, outputSize, util.CreateFloatArrayByStep(inputSize*outputSize, 0, 