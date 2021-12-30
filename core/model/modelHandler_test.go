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

		modelPath := "model.