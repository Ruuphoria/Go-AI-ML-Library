package neuralNetwork

import (
	"gonum.org/v1/gonum/mat"
)

// Optimizer : パラメーターと勾配情報からパラメーターの最適化を行うIF
type Optimizer interface {
	// Update : パラメーターを勾配情報を元に最適化(更新)する
	Update(params map[string]mat.Matrix, grads map[string]m