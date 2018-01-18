package tsp

import (
	"image/color"
	"log"

	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var Colors = map[string][]uint8{
	"Blue":      []uint8{52, 152, 219, 255},
	"Red":       []uint8{231, 76, 60, 255},
	"Green":     []uint8{22, 160, 133, 255},
	"Turquoise": []uint8{26, 188, 156, 255},
	"Amethist":  []uint8{155, 89, 182, 255},
	"Asphalt":   []uint8{52, 73, 94, 255},
	"Orange":    []uint8{211, 84, 0, 255},
}

func Stats(data []float64) {
	// Calculate standard metrics
	meanVal := stat.Mean(data, nil)
	modeVal, modeCount := stat.Mode(data, nil)
	medianVal, err := stats.Median(data)
	checkErr(err)
	minVal := floats.Min(data)
	maxVal := floats.Max(data)
	rangeVal := maxVal - minVal
	varianceVal := stat.Variance(data, nil)
	stdDevVal := stat.StdDev(data, nil)

	// Get quantiles
	inds := make([]int, len(data))
	floats.Argsort(data, inds)
	quant25 := stat.Quantile(0.25, stat.Empirical, data, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, data, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, data, nil)

	log.Printf("Mean: %0.2f", meanVal)
	log.Printf("Mode val: %0.2f", modeVal)
	log.Printf("Mode cnt: %0.2f", modeCount)
	log.Printf("Median: %0.2f", medianVal)
	log.Printf("Min: %0.2f", minVal)
	log.Printf("Max: %0.2f", maxVal)
	log.Printf("Range: %0.2f", rangeVal)
	log.Printf("Variance: %0.2f", varianceVal)
	log.Printf("Standard Deviation: %0.2f", stdDevVal)
	log.Printf("25%% quantile: %0.2f", quant25)
	log.Printf("50%% quantile: %0.2f", quant50)
	log.Printf("75%% quantile: %0.2f", quant75)
}

func PlotHistogram(v []float64, title string) error {
	log.Printf("Plotting %s.png", title)

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title

	pv := make(plotter.Values, len(v))

	for i, val := range v {
		pv[i] = val
	}

	h, err := plotter.NewHist(pv, 100)
	if err != nil {
		return err
	}

	h.Normalize(1)
	p.Add(h)

	if err := p.Save(15*vg.Centimeter, 15*vg.Centimeter, title+".png"); err != nil {
		return err
	}

	return nil
}

func PlotScatterXY1Var(x1, y1 []float64, c1 []uint8, title string) error {
	log.Printf("Plotting %s.png", title)

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	pv1 := make(plotter.XYs, len(x1))
	for i, val := range x1 {
		pv1[i].X = val
		pv1[i].Y = y1[i]
	}

	s1, err := buildScatter(x1, y1, c1)
	checkErr(err)

	p.Legend.Add("x1, y1", s1)

	p.Add(s1)

	if err := p.Save(30*vg.Centimeter, 15*vg.Centimeter, title+".png"); err != nil {
		return err
	}

	return nil
}

func PlotScatterXY2Var(x1, y1, x2, y2 []float64, c1, c2 []uint8, title string) error {
	log.Printf("Plotting %s.png", title)

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	s1, err := buildScatter(x1, y1, c1)
	checkErr(err)

	s2, err := buildScatter(x2, y2, c2)
	checkErr(err)

	p.Legend.Add("x1, y1", s1)
	p.Legend.Add("x2, y2", s2)

	p.Add(s1, s2)

	if err := p.Save(30*vg.Centimeter, 15*vg.Centimeter, title+".png"); err != nil {
		return err
	}

	return nil
}

func PlotScatterXY3Var(x1, y1, x2, y2, x3, y3 []float64, c1, c2, c3 []uint8, title string) error {
	log.Printf("Plotting %s.png", title)

	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	s1, err := buildScatter(x1, y1, c1)
	checkErr(err)

	s2, err := buildScatter(x2, y2, c2)
	checkErr(err)

	s3, err := buildScatter(x3, y3, c3)
	checkErr(err)

	p.Legend.Add("x1, y1", s1)
	p.Legend.Add("x2, y2", s2)
	p.Legend.Add("x3, y3", s3)

	p.Add(s1, s2, s3)

	if err := p.Save(30*vg.Centimeter, 15*vg.Centimeter, title+".png"); err != nil {
		return err
	}

	return nil
}

func PlotTime1Vals(timestamp, v1 []float64, c1 []uint8, title string) error {
	log.Printf("Plotting %s.png", title)

	// New plot
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	l1, err := buildLine(timestamp, v1, c1)
	checkErr(err)

	p.Add(l1)

	p.Legend.Add("v1", l1)

	if err := p.Save(60*vg.Centimeter, 15*vg.Centimeter, title+".png"); err != nil {
		return err
	}

	return nil
}

func PlotTime2Vals(timestamp, v1, v2 []float64, c1, c2 []uint8, title string) error {
	log.Printf("Plotting %s.png", title)

	// New plot
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	l1, err := buildLine(timestamp, v1, c1)
	checkErr(err)

	l2, err := buildLine(timestamp, v2, c2)
	checkErr(err)

	p.Add(l1, l2)

	p.Legend.Add("v1", l1)
	p.Legend.Add("v2", l2)

	if err := p.Save(60*vg.Centimeter, 15*vg.Centimeter, title+".png"); err != nil {
		return err
	}

	return nil
}

func PlotTime3Vals(timestamp, v1, v2, v3 []float64, c1, c2, c3 []uint8, title string) error {
	log.Printf("Plotting %s.png", title)

	// New plot
	p, err := plot.New()
	if err != nil {
		return err
	}

	p.Title.Text = title
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	l1, err := buildLine(timestamp, v1, c1)
	checkErr(err)

	l2, err := buildLine(timestamp, v2, c2)
	checkErr(err)

	l3, err := buildLine(timestamp, v3, c3)
	checkErr(err)

	p.Add(l1, l2, l3)

	p.Legend.Add("v1", l1)
	p.Legend.Add("v2", l2)
	p.Legend.Add("v3", l3)

	if err := p.Save(60*vg.Centimeter, 15*vg.Centimeter, title+".png"); err != nil {
		return err
	}

	return nil
}

// =================
// = Plot builders =
// =================
func buildScatter(x, y []float64, c []uint8) (*plotter.Scatter, error) {
	pv := make(plotter.XYs, len(x))
	for i, val := range x {
		pv[i].X = val
		pv[i].Y = y[i]
	}

	s, err := plotter.NewScatter(pv)
	if err != nil {
		return s, err
	}

	s.Color = color.RGBA{
		R: c[0],
		G: c[1],
		B: c[2],
		A: c[3],
	}

	return s, nil
}

func buildLine(t, v []float64, c []uint8) (*plotter.Line, error) {
	pv := make(plotter.XYs, len(v))

	for i, val := range v {
		pv[i].X = t[i]
		pv[i].Y = val
	}

	l, err := plotter.NewLine(pv)
	if err != nil {
		return l, err
	}

	l.Color = color.RGBA{
		R: c[0],
		G: c[1],
		B: c[2],
		A: c[3],
	}

	return l, nil
}

// ===========
// = Helpers =
// ===========
// Checks for an error and fatals if error is not nil
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
