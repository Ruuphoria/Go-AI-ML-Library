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
		dw := mat.NewDense(4, 3,