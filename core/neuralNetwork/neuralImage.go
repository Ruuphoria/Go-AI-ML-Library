package neuralNetwork

// Image : ニューラルネットワークでの画像データ（1チャンネル分）を格納する配列データ
type Image [][]float64

// NewImage : 単一チャネルの画像データを格納する配列データを作成
// input : 画像の元データを格納した配列データ
// w : 幅
// h : 高さ
func NewImage(input []float64, w int, h int) Image {
	if len(input) != w*h {
		panic("入力された画像データと指定した幅・高さがマッチしてません")
	}
	image := make([][]float64, 0, h)

	for i := 0; i < h; i++ {
		row := input[i*w : (i+1)*w]
		image = append(image, row)
	}
	return image
}

// ImageWithChannel : 複数チャネル（RGBなど）を持つ画像データを格納する配列データ
type ImageWithChannel []Image

// NewImageWithChannel : 複数チャネルを持つ画像データを作成
// input : 画像の元データを格納した配列データ
// w : 幅
// h : 高さ
// c : チャネル数
func NewImageWithChannel(input []float64, w int, h int, c int) ImageWithChannel {
	if len(input) != w*h*c {
		panic("入力された画像データと指定した幅・高さ・チャネル数がマッチしてません")
	}
	iwc := make([]Image, 0, c)

	for i := 0; i < c; i++ {
		image := NewImage(input[i*w*h:(i+1)*w*h], w, h)
		iwc = append(iwc, image)
	}
	return iwc
}

// ImagesWithChannel : 複数チャネルを持つ画像データを複数格納した配列データ
type ImagesWithChannel []ImageWithChannel

// NewImagesWithChannel : 複数チャネルを持つ画像データを作成
// in