package neuralNetwork

import (
	"github.com/goMLLibrary/core/util"
	"gonum.org/v1/gonum/mat"
)

// NeuralNetworkBaseLayer : ニューラルネットワークの素子に関する基本的なIF
type NeuralNetworkBaseLayer interface {
	// Forward : 順方