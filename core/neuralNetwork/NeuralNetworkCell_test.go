package neuralNetwork

import (
	"testing"

	"github.com/goMLLibrary/core/util"
	. "github.com/smartystreets/goconvey/convey"
	"gonum.org/v1/gonum/mat"
)

func TestAffine(t *testing.T) {
	Convey("Given : アフィン変換のレイヤーが一つ与えられた時", t, func() {
		Conv