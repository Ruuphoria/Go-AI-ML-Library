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
		w := mat.NewDense(3, 2, util.CreateFloatA