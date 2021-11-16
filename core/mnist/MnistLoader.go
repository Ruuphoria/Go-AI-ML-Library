
package mnist

import (
	"fmt"
	"image"
	"image/color"
	"path"

	"github.com/goMLLibrary/core/util"
	"gonum.org/v1/gonum/mat"
	"github.com/petar/GoMNIST"
)

var downloadMaps = map[string]string {
	"train-images-idx3-ubyte.gz": "http://yann.lecun.com/exdb/mnist/train-images-idx3-ubyte.gz",
	"train-labels-idx1-ubyte.gz": "http://yann.lecun.com/exdb/mnist/train-labels-idx1-ubyte.gz",
	"t10k-images-idx3-ubyte.gz": "http://yann.lecun.com/exdb/mnist/t10k-images-idx3-ubyte.gz",
	"t10k-labels-idx1-ubyte.gz": "http://yann.lecun.com/exdb/mnist/t10k-labels-idx1-ubyte.gz",
}

// LoadData : Mnistのデータセットを取得
func LoadData(rootPath string) (trainSet *MnistDataSet, testSet *MnistDataSet, err error) {
	// mnistデータがなければ、ダウンロードする
	downloadMnistDataIfNeeded(rootPath)

	// train : Mnistの学習用データ
	// test : Mnistのテスト用データ
	train, test, err := GoMNIST.Load(rootPath)
	if err != nil {
		return nil, nil, err
	}

	trainDataSet := newMnistDataSet(train)
	testDataSet := newMnistDataSet(test)
	return trainDataSet, testDataSet, nil
}

type MnistDataSet struct {
	dataSet []MnistData
	nCol    int
	nRow    int
}

func newMnistDataSet(set *GoMNIST.Set) *MnistDataSet {
	dataSet := MnistDataSet{nCol: set.NCol, nRow: set.NRow}
	dataSet.dataSet = make([]MnistData, 0, set.Count())
	for i, rawData := range set.Images {