package graph

import (
	"os"
	"path/filepath"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type GraphCreater struct {
	p      *plot.Plot
	outDir string
}

func NewGraphCreater(outputDir string) (*GraphCreater, error) {
	gc := GraphCreater{}
	p, err := plot.New()
	if err != nil {
		return nil, err
	}
	gc.p = p

	// グラフ情報の保存先フォルダの確認（無ければ作成）
	gc.outDir = outputDir
	if _, err := os.Stat(outputDir); err != nil {
		os.Mkdir(outputDir, 0777)
	}

	return &gc, nil
}

// SaveGraph : 指定した座標