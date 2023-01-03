package neuralNetwork

// Image : ニューラルネットワークでの画像データ（1チャンネル分）を格納する配列データ
type Image [][]float64

// NewImage : 単一チャネルの画像データを格納する配列データを作成
// input : 画像の元データを格納した配列データ
// w : 幅
// h : 高さ
func NewImage(input []float