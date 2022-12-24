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
	