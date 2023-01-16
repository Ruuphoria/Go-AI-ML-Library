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
			Convey("Then : 3*4の2次元配列データが出来ていること", func() {
				So(h, ShouldEqual, len(image))
				for i := 0; i < h; i++ {
					So(w, ShouldEqual, len(image[i]))
					So(reflect.DeepEqual(image[i], util.CreateFloatArrayByStep(w, float64(i*w), 1)), ShouldBeTrue)
				}
			})
		})
	})
}

func TestNewImageWithChannel(t *testing.T) {
	Convey("Given : 幅3, 高さ4, チャネル数3の数値が配列で与えられた時", t, func() {
		w := 3
		h := 4
		c := 3
		input := util.CreateFloatArrayByStep(w*h*c, 0, 1)
		Convey("When : ImageWithChannelを作成する", func() {
			imageWithChannel := NewImag