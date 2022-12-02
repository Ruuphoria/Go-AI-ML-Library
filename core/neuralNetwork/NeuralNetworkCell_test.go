package neuralNetwork

import (
	"testing"

	"github.com/goMLLibrary/core/util"
	. "github.com/smartystreets/goconvey/convey"
	"gonum.org/v1/gonum/mat"
)

func TestAffine(t *testing.T) {
	Convey("Given : アフィン変換のレイヤーが一つ与えられた時", t, func() {
		Convey("AND : 重みが3*2行列(in = 3, out = 2), 初期値は1-6とする", nil)
		w := mat.NewDense(3, 2, util.CreateFloatArrayByStep(6, 1, 1))
		Convey("AND : バイアスが2次元, 初期値は-2,-1とする", nil)
		b := mat.NewVecDense(2, []float64{-2, -1})
		aff := newAffine(w, b)
		Convey("When : 入力xを2*3行列とし、値を5-10とする", func() {
			x := mat.NewDense(2, 3, 