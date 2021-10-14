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

// SaveGraph : 指定した座標情報で直線グラフを作成し、画像で保存する
func (gc *GraphCreater) SaveLineGraph(parameter GraphParameter, pointsList []GraphPoints) error {
	gc.p.Title.Text = parameter.Title
	gc.p.X.Label.Text = parameter.XLabel
	gc.p.Y.Label.Text = parameter.YLabel

	for _, points := range pointsList {
		if err := plotutil.AddLinePoints(gc.p, points.key, points.convertPlotterXYs()); err != nil {
			return err
		}
	}

	err := gc.p.Save(vg.Length(parameter.Width), vg.Length(parameter.Height), gc.createFilePath(param