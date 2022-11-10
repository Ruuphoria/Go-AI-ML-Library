
package neuralNetwork

import (
	"math"

	"github.com/goMLLibrary/core/util"
	"gonum.org/v1/gonum/mat"
)

const (
	delta = 0.0000001 // logの中身が0にならないように対応
)

// Sigmoid : シグモイド関数
type Sigmoid struct {
	out mat.Matrix
}

// NewSigmoid : シグモイド関数の素子を取得
func NewSigmoid() *Sigmoid {
	sigmoid := Sigmoid{}
	return &sigmoid
}

func (sigmoid *Sigmoid) Forward(x mat.Matrix) mat.Matrix {
	r, c := x.Dims()