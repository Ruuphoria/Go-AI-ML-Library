package neuralNetwork

import (
	"testing"

	"github.com/goMLLibrary/core/util"
	. "github.com/smartystreets/goconvey/convey"
	"gonum.org/v1/gonum/mat"
)

func TestSGD(t *testing.T) {
	Convey("Given : 1つの重み、1つのバイアスが与えられた時", t, func() {
		params := make(map[string]mat.Matrix)
		grads := make(map[string]mat.Matrix)
		Convey("AND 重み行列は4*3行列で各値は1-12とし、勾配は0-5.5(0.5刻み）とする", nil)
		w := mat.NewDense(4, 3, util.CreateFloatArrayByStep(12, 1.0, 1.0))
		dw := mat.NewDense(4, 3, util.CreateFloatArrayByStep(12, 0, 0.5))
		params["w"] = w
		grads["w"] = dw
		Convey("AND バイアスは3次元で各値は0-2とし、勾配は-2-0とする", nil)
		b := mat.NewVecDense(3, util.CreateFloatArrayByStep(3, 0, 1.0))
		db := mat.NewVecDense(3, util.CreateFloatArrayByStep(3, -2, 1.0))
		params["b"] = b
		grads["b"] = db
		Convey("When : SGDの学習率0.1で初期化", func() {
			sgd := NewSGD(WithSGDLearningRate(0.1))
			Convey("Then : Optimizerでu