package neuralNetwork

import (
	"reflect"
	"testing"

	"github.com/goMLLibrary/core/util"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewImage(t *testing.T) {
	Convey("Given : 3 * 4の数値が配列で与えられた時", t, func() {
		w := 3
		h := 4
		input := util.CreateFloatArrayByStep(w*h, 0, 1)
		Convey("When : Imageを作成する", func() {
			image := NewImage(input, w, h)
			Conv