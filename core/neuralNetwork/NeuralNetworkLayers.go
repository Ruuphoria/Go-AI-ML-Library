
package neuralNetwork

import "gonum.org/v1/gonum/mat"

// NeuralNetworkLayers : ニューラルネットワークの素子を複数持つ多重層
type NeuralNetworkLayers struct {
	layers              []NeuralNetworkBaseLayer
	lastActivationLayer *SoftmaxWithLoss
	optimizer           Optimizer
}

// NewDefaultNeuralNetworkLayers : NeuralNetworkLayersのインスタンスを作成
func NewDefaultNeuralNetworkLayers() *NeuralNetworkLayers {
	nnl := NeuralNetworkLayers{}
	nnl.layers = make([]NeuralNetworkBaseLayer, 0)
	nnl.lastActivationLayer = NewSoftmaxWithLoss()
	nnl.optimizer = NewSGD()
	return &nnl
}

// Add : ニューラルネットワークの素子を追加
func (nnl *NeuralNetworkLayers) Add(layer NeuralNetworkBaseLayer) {
	nnl.layers = append(nnl.layers, layer)
}

// SetOptimizer : optimizerの設定
func (nnl *NeuralNetworkLayers) SetOptimizer(optimizer Optimizer) {
	nnl.optimizer = optimizer
}

// SetLastActivationLayer : ニューラルネットワークの最終層を設定
func (nnl *NeuralNetworkLayers) SetLastActivationLayer(layer *SoftmaxWithLoss) {
	nnl.lastActivationLayer = layer
}

// Forward : 順伝搬処理の実施
func (nnl *NeuralNetworkLayers) Forward(x mat.Matrix, t mat.Matrix) (loss float64, accuracy float64) {
	var input mat.Matrix = mat.DenseCopyOf(x)
	for _, layer := range nnl.layers {
		input = layer.Forward(input)
	}
	return nnl.lastActivationLayer.Forward(input, t)
}

// Backward : 逆伝搬処理の実施
func (nnl *NeuralNetworkLayers) Backward() {
	var dout mat.Matrix = nnl.lastActivationLayer.Backward()
	for i := len(nnl.layers) - 1; i >= 0; i-- {
		dout = nnl.layers[i].Backward(dout)
	}
}

// Update : 各レイヤーのパラメーターを勾配情報を元に更新
func (nnl *NeuralNetworkLayers) Update() {
	for i, layer := range nnl.layers {
		neuralNetworkLayer, ok := layer.(NeuralNetworkLayer)
		if !ok {
			// 重みをもっているレイヤーではないためスキップする
			//tmpLayers = append(tmpLayers, layer)
			continue
		}

		// 各パラメーターの更新処理
		params := neuralNetworkLayer.GetParams()
		grads := neuralNetworkLayer.GetGradients()
		nnl.optimizer.Update(params, grads)
		neuralNetworkLayer.UpdateParams(params)
		nnl.layers[i] = neuralNetworkLayer
	}
}

// GetLayers : レイヤー情報を取得
func (nnl *NeuralNetworkLayers) GetLayers() []NeuralNetworkBaseLayer {
	return nnl.layers
}

// GetOptimizer : optimizer情報を取得
func (nnl *NeuralNetworkLayers) GetOptimizer() Optimizer {
	return nnl.optimizer
}

// GetLastActivationLayer : 最終的な活性化レイヤーを取得
func (nnl *NeuralNetworkLayers) GetLastActivationLayer() *SoftmaxWithLoss {
	return nnl.lastActivationLayer
}