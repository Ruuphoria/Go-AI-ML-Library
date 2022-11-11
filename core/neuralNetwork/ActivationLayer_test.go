
package neuralNetwork

import (
	"math"
	"testing"

	"github.com/goMLLibrary/core/util"
	. "github.com/smartystreets/goconvey/convey"
	"gonum.org/v1/gonum/mat"
)

func TestSigmoid(t *testing.T) {
	Convey("Given : sigmoidレイヤーが一つ与えられた時", t, func() {
		s := NewSigmoid()
		Convey("AND : 行列サイズを2*3とする", nil)
		r := 2
		c := 3
		Convey("When : 入力行列xが与えられた時", func() {
			x := mat.NewDense(r, c, util.CreateFloatArrayByStep(r*c, 1, 1))
			out := s.Forward(x)
			Convey("Then : Forward処理を実施", func() {
				act_r, act_c := out.Dims()
				So(act_r, ShouldEqual, r)
				So(act_c, ShouldEqual, c)
				for i := 0; i < r; i++ {
					for j := 0; j < c; j++ {
						So(out.At(i, j), ShouldEqual, sigmoid_forward(x.At(i, j)))
					}
				}
			})

			Convey("AND : 誤差doutが与えられた時", nil)
			dout := mat.NewDense(r, c, util.CreateFloatArrayByStep(r*c, 0.5, 0.5))
			Convey("Then : Backward処理を実施", func() {
				out := s.Backward(dout)
				act_r, act_c := out.Dims()
				So(act_r, ShouldEqual, r)
				So(act_c, ShouldEqual, c)
				for i := 0; i < r; i++ {
					for j := 0; j < c; j++ {
						So(out.At(i, j), ShouldEqual, sigmoid_backward(x.At(i, j), dout.At(i, j)))
					}
				}
			})