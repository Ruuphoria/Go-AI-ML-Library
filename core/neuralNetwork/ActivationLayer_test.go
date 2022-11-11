
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
		})
	})
}

func TestRelu(t *testing.T) {
	Convey("Given : Reluレイヤーが一つ与えられた時", t, func() {
		relu := NewRelu()
		Convey("AND : 行列サイズを3*2とする", nil)
		r := 3
		c := 2
		Convey("When : 入力行列xが与えられた時", func() {
			x := mat.NewDense(r, c, util.CreateFloatArrayByStep(r*c, 1, 1))
			out := relu.Forward(x)
			Convey("Then : Forward処理を実施", func() {
				act_r, act_c := out.Dims()
				So(act_r, ShouldEqual, r)
				So(act_c, ShouldEqual, c)
				for i := 0; i < r; i++ {
					for j := 0; j < c; j++ {
						So(out.At(i, j), ShouldEqual, relu_forward(x.At(i, j)))
					}
				}
			})

			Convey("AND : 誤差doutが与えられた時", nil)
			dout := mat.NewDense(r, c, util.CreateFloatArrayByStep(r*c, 0.5, 0.5))
			Convey("Then : Backward処理を実施", func() {
				out := relu.Backward(dout)
				act_r, act_c := out.Dims()
				So(act_r, ShouldEqual, r)
				So(act_c, ShouldEqual, c)
				for i := 0; i < r; i++ {
					for j := 0; j < c; j++ {
						So(out.At(i, j), ShouldEqual, relu_backward(x.At(i, j), dout.At(i, j)))
					}
				}
			})
		})
	})
}

func TestTanh(t *testing.T) {
	Convey("Given : Tanhレイヤーが一つ与えられた時", t, func() {
		tanh := NewTanh()
		Convey("AND : 行列サイズを3*2とする", nil)
		r := 3
		c := 2
		Convey("When : 入力行列xが与えられた時", func() {
			x := mat.NewDense(r, c, util.CreateFloatArrayByStep(r*c, 1, 1))
			out := tanh.Forward(x)
			Convey("Then : Forward処理を実施", func() {
				act_r, act_c := out.Dims()
				So(act_r, ShouldEqual, r)
				So(act_c, ShouldEqual, c)
				for i := 0; i < r; i++ {
					for j := 0; j < c; j++ {
						So(out.At(i, j), ShouldEqual, tanh_forward(x.At(i, j)))
					}
				}
			})

			Convey("AND : 誤差doutが与えられた時", nil)
			dout := mat.NewDense(r, c, util.CreateFloatArrayByStep(r*c, 0.5, 0.5))
			Convey("Then : Backward処理を実施", func() {
				out := tanh.Backward(dout)
				act_r, act_c := out.Dims()
				So(act_r, ShouldEqual, r)
				So(act_c, ShouldEqual, c)
				for i := 0; i < r; i++ {
					for j := 0; j < c; j++ {
						checkValue(out.At(i, j), tanh_backward(x.At(i, j), dout.At(i, j)), math.Pow10(-8))
					}
				}
			})
		})
	})
}

func TestSoftmaxCrossEntropy(t *testing.T) {
	Convey("Given : SoftmaxCrossEntropyレイヤーが一つ与えられた時", t, func() {
		sce := NewSoftmaxWithLoss()
		Convey("AND : 3次元ベクトルを2個入力するとする（2*3の行列）", nil)
		r := 2 // データ数（バッチ数）
		c := 3 // データの次元
		Convey("When : 入力行列xと正解データが与えられた時", func() {
			// [1,2,3]
			// [4,5,7]
			x := mat.NewDense(r, c, []float64{1, 2, 3, 4, 5, 7})
			// [0,0,1]
			// [1,0,0]
			t := mat.NewDense(r, c, []float64{0, 0, 1, 1, 0, 0})
			loss, _ := sce.Forward(x, t)
			Convey("Then : Forward処理を実施", func() {
				loss_expected := softmaxCrossEntropy_forward(x, t)
				So(loss, ShouldEqual, loss_expected)
			})
			Convey("Then : Backward処理を実施", func() {
				douts := sce.Backward()
				softmax_values := softmax_batch(x)
				douts_expected := softmaxCrossEntropy_backward(softmax_values, t)
				So(mat.Equal(douts, douts_expected), ShouldBeTrue)
			})
		})
	})
}

func checkValue(act, exp, diff float64) {
	So(math.Abs(act-exp), ShouldBeLessThan, diff)