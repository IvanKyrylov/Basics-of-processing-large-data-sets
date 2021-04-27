package main

import (
	"sort"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// type csv struct {
// 	person          string
// 	reg_addr_koatuu int64
// 	oper_code       int
// 	oper_name       string
// 	d_reg           string
// 	dep_code        int
// 	dep             string
// 	brand           string
// 	model           string
// 	make_year       int
// 	color           string
// 	kind            string
// 	body            string
// 	purpose         string
// 	fuel            string
// 	capacity        int
// 	own_weight      int
// 	total_weight    int
// 	n_reg_new       string
// }

func main() {
	// wg := sync.WaitGroup{}
	// wg.Add(1)
	plotColor()
	plotMakeYear()
	// go plotBrand(&wg)
	// wg.Wait()
}

func plotColor() {
	var colorAuto2013 []Count
	var colorAuto2014 []Count
	var colorAuto2015 []Count
	var colorAuto2016 []Count
	var colorAuto2017 []Count
	var colorAuto2018 []Count
	var colorAuto2019 []Count
	var colorAuto2020 []Count

	// wg := sync.WaitGroup{}
	// wg.Add(2)
	// go func() {
	colorAuto2013 = <-MapReduce("./csv/tz_opendata_z01012013_po31122013.csv", 9)
	// 	wg.Done()
	// }()

	// go func() {
	colorAuto2014 = <-MapReduce("./csv/tz_opendata_z01012014_po31122014.csv", 9)
	// 	wg.Done()
	// }()
	// wg.Wait()
	// wg.Add(2)

	// go func() {
	colorAuto2015 = <-MapReduce("./csv/tz_opendata_z01012015_po31122015.csv", 9)
	// 	wg.Done()
	// }()

	// go func() {
	colorAuto2016 = <-MapReduce("./csv/tz_opendata_z01012016_po31122016.csv", 9)
	// 	wg.Done()
	// }()

	// wg.Wait()
	// wg.Add(2)

	// go func() {
	colorAuto2017 = <-MapReduce("./csv/tz_opendata_z01012017_po31122017.csv", 9)
	// 	wg.Done()
	// }()

	// go func() {
	colorAuto2018 = <-MapReduce("./csv/tz_opendata_z01012018_po01012019.csv", 9)
	// 	wg.Done()
	// }()

	// wg.Wait()
	// wg.Add(2)

	// go func() {
	colorAuto2019 = <-MapReduce("./csv/tz_opendata_z01012019_po01012020.csv", 9)
	// 	wg.Done()
	// }()

	// go func() {
	colorAuto2020 = <-MapReduce("./csv/tz_opendata_z01012020_po01012021.csv", 9)
	// 	wg.Done()
	// }()

	// wg.Wait()

	sort.Slice(colorAuto2013, func(i, j int) bool {
		return colorAuto2013[i].Key > colorAuto2013[j].Key
	})
	sort.Slice(colorAuto2014, func(i, j int) bool {
		return colorAuto2014[i].Key > colorAuto2014[j].Key
	})
	sort.Slice(colorAuto2015, func(i, j int) bool {
		return colorAuto2015[i].Key > colorAuto2015[j].Key
	})
	sort.Slice(colorAuto2015, func(i, j int) bool {
		return colorAuto2015[i].Key > colorAuto2015[j].Key
	})
	sort.Slice(colorAuto2016, func(i, j int) bool {
		return colorAuto2016[i].Key > colorAuto2016[j].Key
	})
	sort.Slice(colorAuto2017, func(i, j int) bool {
		return colorAuto2017[i].Key > colorAuto2017[j].Key
	})
	sort.Slice(colorAuto2018, func(i, j int) bool {
		return colorAuto2018[i].Key > colorAuto2018[j].Key
	})
	sort.Slice(colorAuto2019, func(i, j int) bool {
		return colorAuto2019[i].Key > colorAuto2019[j].Key
	})
	sort.Slice(colorAuto2020, func(i, j int) bool {
		return colorAuto2020[i].Key > colorAuto2020[j].Key
	})

	groupCount2013 := make(plotter.Values, 0, len(colorAuto2013))
	groupCount2014 := make(plotter.Values, 0, len(colorAuto2014))
	groupCount2015 := make(plotter.Values, 0, len(colorAuto2015))
	groupCount2016 := make(plotter.Values, 0, len(colorAuto2016))
	groupCount2017 := make(plotter.Values, 0, len(colorAuto2017))
	groupCount2018 := make(plotter.Values, 0, len(colorAuto2018))
	groupCount2019 := make(plotter.Values, 0, len(colorAuto2019))
	groupCount2020 := make(plotter.Values, 0, len(colorAuto2020))

	groupName := []string{}

	for i, v := range colorAuto2013 {
		if i == 6 {
			groupCount2013 = append(groupCount2013, float64(0))
			continue
		}
		groupCount2013 = append(groupCount2013, float64(v.Count))
	}
	for _, v := range colorAuto2014 {
		groupCount2014 = append(groupCount2014, float64(v.Count))
	}
	for _, v := range colorAuto2015 {
		groupCount2015 = append(groupCount2015, float64(v.Count))
	}
	for _, v := range colorAuto2016 {
		groupCount2016 = append(groupCount2016, float64(v.Count))
	}
	for _, v := range colorAuto2017 {
		groupCount2017 = append(groupCount2017, float64(v.Count))
	}
	for _, v := range colorAuto2018 {
		groupCount2018 = append(groupCount2018, float64(v.Count))
	}
	for _, v := range colorAuto2019 {
		groupCount2019 = append(groupCount2019, float64(v.Count))
	}
	for _, v := range colorAuto2020 {
		groupCount2020 = append(groupCount2020, float64(v.Count))
		groupName = append(groupName, v.Key)
	}

	p := plot.New()

	p.Title.Text = "Color stat"
	p.Y.Label.Text = "Count"
	w := vg.Points(12)

	bars2013, err := plotter.NewBarChart(groupCount2013, w)
	if err != nil {
		panic(err)
	}
	bars2013.LineStyle.Width = vg.Length(0)
	bars2013.Color = plotutil.Color(0)
	bars2013.Offset = -w * 3

	bars2014, err := plotter.NewBarChart(groupCount2014, w)
	if err != nil {
		panic(err)
	}
	bars2014.LineStyle.Width = vg.Length(0)
	bars2014.Color = plotutil.Color(1)
	bars2014.Offset = -w * 2

	bars2015, err := plotter.NewBarChart(groupCount2015, w)
	if err != nil {
		panic(err)
	}
	bars2015.LineStyle.Width = vg.Length(0)
	bars2015.Color = plotutil.Color(2)
	bars2015.Offset = -w

	bars2016, err := plotter.NewBarChart(groupCount2016, w)
	if err != nil {
		panic(err)
	}
	bars2016.LineStyle.Width = vg.Length(0)
	bars2016.Color = plotutil.Color(3)
	// bars2016.Offset = w

	bars2017, err := plotter.NewBarChart(groupCount2017, w)
	if err != nil {
		panic(err)
	}
	bars2017.LineStyle.Width = vg.Length(0)
	bars2017.Color = plotutil.Color(4)
	bars2017.Offset = w

	bars2018, err := plotter.NewBarChart(groupCount2018, w)
	if err != nil {
		panic(err)
	}
	bars2018.LineStyle.Width = vg.Length(0)
	bars2018.Color = plotutil.Color(5)
	bars2018.Offset = w * 2

	bars2019, err := plotter.NewBarChart(groupCount2019, w)
	if err != nil {
		panic(err)
	}
	bars2019.LineStyle.Width = vg.Length(0)
	bars2019.Color = plotutil.Color(6)
	bars2019.Offset = w * 3

	bars2020, err := plotter.NewBarChart(groupCount2020, w)
	if err != nil {
		panic(err)
	}
	bars2020.LineStyle.Width = vg.Length(0)
	bars2020.Color = plotutil.Color(7)
	bars2020.Offset = w * 4

	p.Add(bars2013, bars2014, bars2015, bars2016, bars2017, bars2018, bars2019, bars2020)

	p.Legend.Add("Group 2013", bars2013)
	p.Legend.Add("Group 2014", bars2014)
	p.Legend.Add("Group 2015", bars2015)
	p.Legend.Add("Group 2016", bars2016)
	p.Legend.Add("Group 2017", bars2017)
	p.Legend.Add("Group 2018", bars2018)
	p.Legend.Add("Group 2019", bars2019)
	p.Legend.Add("Group 2020", bars2020)

	p.Legend.Top = true

	p.NominalX(groupName...)

	if err := p.Save(20*vg.Inch, 22*vg.Inch, "./plot/colorPlot.png"); err != nil {
		panic(err)
	}
}
func plotMakeYear() {
	var makeYearAuto2013 []Count
	var makeYearAuto2014 []Count
	var makeYearAuto2015 []Count
	var makeYearAuto2016 []Count
	var makeYearAuto2017 []Count
	var makeYearAuto2018 []Count
	var makeYearAuto2019 []Count
	var makeYearAuto2020 []Count

	// wg := sync.WaitGroup{}
	// wg.Add(8)
	// go func() {
	makeYearAuto2013 = (<-MapReduce("./csv/tz_opendata_z01012013_po31122013.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// go func() {
	makeYearAuto2014 = (<-MapReduce("./csv/tz_opendata_z01012014_po31122014.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// go func() {
	makeYearAuto2015 = (<-MapReduce("./csv/tz_opendata_z01012015_po31122015.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// go func() {
	makeYearAuto2016 = (<-MapReduce("./csv/tz_opendata_z01012016_po31122016.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// go func() {
	makeYearAuto2017 = (<-MapReduce("./csv/tz_opendata_z01012017_po31122017.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// go func() {
	makeYearAuto2018 = (<-MapReduce("./csv/tz_opendata_z01012018_po01012019.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// go func() {
	makeYearAuto2019 = (<-MapReduce("./csv/tz_opendata_z01012019_po01012020.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// go func() {
	makeYearAuto2020 = (<-MapReduce("./csv/tz_opendata_z01012020_po01012021.csv", 9))[:10]
	// 	wg.Done()
	// }()

	// wg.Wait()

	sort.Slice(makeYearAuto2013, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2013[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2013[j].Key)
		return iC > jC
	})
	sort.Slice(makeYearAuto2014, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2014[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2014[j].Key)
		return iC > jC
	})
	sort.Slice(makeYearAuto2015, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2015[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2015[j].Key)
		return iC > jC
	})
	sort.Slice(makeYearAuto2016, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2016[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2016[j].Key)
		return iC > jC
	})
	sort.Slice(makeYearAuto2017, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2017[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2017[j].Key)
		return iC > jC
	})
	sort.Slice(makeYearAuto2018, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2018[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2018[j].Key)
		return iC > jC
	})
	sort.Slice(makeYearAuto2019, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2019[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2019[j].Key)
		return iC > jC
	})
	sort.Slice(makeYearAuto2020, func(i, j int) bool {
		iC, _ := strconv.Atoi(makeYearAuto2020[i].Key)
		jC, _ := strconv.Atoi(makeYearAuto2020[j].Key)
		return iC > jC
	})

	groupCount2013 := make(plotter.Values, 0, len(makeYearAuto2013))
	groupCount2014 := make(plotter.Values, 0, len(makeYearAuto2014))
	groupCount2015 := make(plotter.Values, 0, len(makeYearAuto2015))
	groupCount2016 := make(plotter.Values, 0, len(makeYearAuto2016))
	groupCount2017 := make(plotter.Values, 0, len(makeYearAuto2017))
	groupCount2018 := make(plotter.Values, 0, len(makeYearAuto2018))
	groupCount2019 := make(plotter.Values, 0, len(makeYearAuto2019))
	groupCount2020 := make(plotter.Values, 0, len(makeYearAuto2020))

	groupName := []string{}

	for _, v := range makeYearAuto2013[:10] {
		groupCount2013 = append(groupCount2013, float64(v.Count))
	}
	for _, v := range makeYearAuto2014[:10] {
		groupCount2014 = append(groupCount2014, float64(v.Count))
	}
	for _, v := range makeYearAuto2015[:10] {
		groupCount2015 = append(groupCount2015, float64(v.Count))
	}
	for _, v := range makeYearAuto2016[:10] {
		groupCount2016 = append(groupCount2016, float64(v.Count))
	}
	for _, v := range makeYearAuto2017[:10] {
		groupCount2017 = append(groupCount2017, float64(v.Count))
	}
	for _, v := range makeYearAuto2018[:10] {
		groupCount2018 = append(groupCount2018, float64(v.Count))
	}
	for _, v := range makeYearAuto2019[:10] {
		groupCount2019 = append(groupCount2019, float64(v.Count))
	}
	for _, v := range makeYearAuto2020[:10] {
		groupCount2020 = append(groupCount2020, float64(v.Count))
		groupName = append(groupName, v.Key)
	}

	p := plot.New()

	p.Title.Text = "Make year stat"
	p.Y.Label.Text = "Count"
	w := vg.Points(12)

	bars2013, err := plotter.NewBarChart(groupCount2013, w)
	if err != nil {
		panic(err)
	}
	bars2013.LineStyle.Width = vg.Length(0)
	bars2013.Color = plotutil.Color(0)
	bars2013.Offset = -w * 3

	bars2014, err := plotter.NewBarChart(groupCount2014, w)
	if err != nil {
		panic(err)
	}
	bars2014.LineStyle.Width = vg.Length(0)
	bars2014.Color = plotutil.Color(1)
	bars2014.Offset = -w * 2

	bars2015, err := plotter.NewBarChart(groupCount2015, w)
	if err != nil {
		panic(err)
	}
	bars2015.LineStyle.Width = vg.Length(0)
	bars2015.Color = plotutil.Color(2)
	bars2015.Offset = -w

	bars2016, err := plotter.NewBarChart(groupCount2016, w)
	if err != nil {
		panic(err)
	}
	bars2016.LineStyle.Width = vg.Length(0)
	bars2016.Color = plotutil.Color(3)
	// bars2016.Offset = w

	bars2017, err := plotter.NewBarChart(groupCount2017, w)
	if err != nil {
		panic(err)
	}
	bars2017.LineStyle.Width = vg.Length(0)
	bars2017.Color = plotutil.Color(4)
	bars2017.Offset = w

	bars2018, err := plotter.NewBarChart(groupCount2018, w)
	if err != nil {
		panic(err)
	}
	bars2018.LineStyle.Width = vg.Length(0)
	bars2018.Color = plotutil.Color(5)
	bars2018.Offset = w * 2

	bars2019, err := plotter.NewBarChart(groupCount2019, w)
	if err != nil {
		panic(err)
	}
	bars2019.LineStyle.Width = vg.Length(0)
	bars2019.Color = plotutil.Color(6)
	bars2019.Offset = w * 3

	bars2020, err := plotter.NewBarChart(groupCount2020, w)
	if err != nil {
		panic(err)
	}
	bars2020.LineStyle.Width = vg.Length(0)
	bars2020.Color = plotutil.Color(7)
	bars2020.Offset = w * 4

	p.Add(bars2013, bars2014, bars2015, bars2016, bars2017, bars2018, bars2019, bars2020)

	p.Legend.Add("Group 2013", bars2013)
	p.Legend.Add("Group 2014", bars2014)
	p.Legend.Add("Group 2015", bars2015)
	p.Legend.Add("Group 2016", bars2016)
	p.Legend.Add("Group 2017", bars2017)
	p.Legend.Add("Group 2018", bars2018)
	p.Legend.Add("Group 2019", bars2019)
	p.Legend.Add("Group 2020", bars2020)

	p.Legend.Top = true

	p.NominalX(groupName...)
	if err := p.Save(20*vg.Inch, 22*vg.Inch, "./plot/makeYearPlot.png"); err != nil {
		panic(err)
	}
}

// func plotBrand() {
// 	brandAuto2013 := MapReduce("./csv/tz_opendata_z01012013_po31122013.csv", 7)
// 	brandAuto2014 := MapReduce("./csv/tz_opendata_z01012014_po31122014.csv", 7)
// 	brandAuto2015 := MapReduce("./csv/tz_opendata_z01012015_po31122015.csv", 7)
// 	brandAuto2016 := MapReduce("./csv/tz_opendata_z01012016_po31122016.csv", 7)
// 	brandAuto2017 := MapReduce("./csv/tz_opendata_z01012017_po31122017.csv", 7)
// 	brandAuto2018 := MapReduce("./csv/tz_opendata_z01012018_po01012019.csv", 7)
// 	brandAuto2019 := MapReduce("./csv/tz_opendata_z01012019_po01012020.csv", 7)
// 	brandAuto2020 := MapReduce("./csv/tz_opendata_z01012020_po01012021.csv", 7)

// 	sort.Slice(brandAuto2013, func(i, j int) bool {
// 		return brandAuto2013[i].Key > brandAuto2013[j].Key
// 	})
// 	sort.Slice(brandAuto2014, func(i, j int) bool {
// 		return brandAuto2014[i].Key > brandAuto2014[j].Key
// 	})
// 	sort.Slice(brandAuto2015, func(i, j int) bool {
// 		return brandAuto2015[i].Key > brandAuto2015[j].Key
// 	})
// 	sort.Slice(brandAuto2015, func(i, j int) bool {
// 		return brandAuto2015[i].Key > brandAuto2015[j].Key
// 	})
// 	sort.Slice(brandAuto2016, func(i, j int) bool {
// 		return brandAuto2016[i].Key > brandAuto2016[j].Key
// 	})
// 	sort.Slice(brandAuto2017, func(i, j int) bool {
// 		return brandAuto2017[i].Key > brandAuto2017[j].Key
// 	})
// 	sort.Slice(brandAuto2018, func(i, j int) bool {
// 		return brandAuto2018[i].Key > brandAuto2018[j].Key
// 	})
// 	sort.Slice(brandAuto2019, func(i, j int) bool {
// 		return brandAuto2019[i].Key > brandAuto2019[j].Key
// 	})
// 	sort.Slice(brandAuto2020, func(i, j int) bool {
// 		return brandAuto2020[i].Key > brandAuto2020[j].Key
// 	})

// 	groupCount2013 := make(plotter.Values, 0, len(brandAuto2013))
// 	groupCount2014 := make(plotter.Values, 0, len(brandAuto2014))
// 	groupCount2015 := make(plotter.Values, 0, len(brandAuto2015))
// 	groupCount2016 := make(plotter.Values, 0, len(brandAuto2016))
// 	groupCount2017 := make(plotter.Values, 0, len(brandAuto2017))
// 	groupCount2018 := make(plotter.Values, 0, len(brandAuto2018))
// 	groupCount2019 := make(plotter.Values, 0, len(brandAuto2019))
// 	groupCount2020 := make(plotter.Values, 0, len(brandAuto2020))

// 	groupName := []string{}

// 	for _, v := range brandAuto2013 {
// 		groupCount2013 = append(groupCount2013, float64(v.Count))
// 	}
// 	for _, v := range brandAuto2014 {
// 		groupCount2014 = append(groupCount2014, float64(v.Count))
// 	}
// 	for _, v := range brandAuto2015 {
// 		groupCount2015 = append(groupCount2015, float64(v.Count))
// 	}
// 	for _, v := range brandAuto2016 {
// 		groupCount2016 = append(groupCount2016, float64(v.Count))
// 	}
// 	for _, v := range brandAuto2017 {
// 		groupCount2017 = append(groupCount2017, float64(v.Count))
// 	}
// 	for _, v := range brandAuto2018 {
// 		groupCount2018 = append(groupCount2018, float64(v.Count))
// 	}
// 	for _, v := range brandAuto2019 {
// 		groupCount2019 = append(groupCount2019, float64(v.Count))
// 	}
// 	for _, v := range brandAuto2020 {
// 		groupCount2020 = append(groupCount2020, float64(v.Count))
// 		groupName = append(groupName, v.Key)
// 	}

// 	p := plot.New()

// 	p.Title.Text = "brand stat"
// 	p.Y.Label.Text = "Count"
// 	w := vg.Points(12)

// 	bars2013, err := plotter.NewBarChart(groupCount2013, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2013.LineStyle.Width = vg.Length(0)
// 	bars2013.Color = plotutil.Color(0)
// 	bars2013.Offset = -w * 3

// 	bars2014, err := plotter.NewBarChart(groupCount2014, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2014.LineStyle.Width = vg.Length(0)
// 	bars2014.Color = plotutil.Color(1)
// 	bars2014.Offset = -w * 2

// 	bars2015, err := plotter.NewBarChart(groupCount2015, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2015.LineStyle.Width = vg.Length(0)
// 	bars2015.Color = plotutil.Color(2)
// 	bars2015.Offset = -w

// 	bars2016, err := plotter.NewBarChart(groupCount2016, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2016.LineStyle.Width = vg.Length(0)
// 	bars2016.Color = plotutil.Color(3)
// 	// bars2016.Offset = w

// 	bars2017, err := plotter.NewBarChart(groupCount2017, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2017.LineStyle.Width = vg.Length(0)
// 	bars2017.Color = plotutil.Color(4)
// 	bars2017.Offset = w

// 	bars2018, err := plotter.NewBarChart(groupCount2018, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2018.LineStyle.Width = vg.Length(0)
// 	bars2018.Color = plotutil.Color(5)
// 	bars2018.Offset = w * 2

// 	bars2019, err := plotter.NewBarChart(groupCount2019, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2019.LineStyle.Width = vg.Length(0)
// 	bars2019.Color = plotutil.Color(6)
// 	bars2019.Offset = w * 3

// 	bars2020, err := plotter.NewBarChart(groupCount2020, w)
// 	if err != nil {
// 		panic(err)
// 	}
// 	bars2020.LineStyle.Width = vg.Length(0)
// 	bars2020.Color = plotutil.Color(7)
// 	bars2020.Offset = w * 4

// 	p.Add(bars2013, bars2014, bars2015, bars2016, bars2017, bars2018, bars2019, bars2020)

// 	p.Legend.Add("Group 2013", bars2013)
// 	p.Legend.Add("Group 2014", bars2014)
// 	p.Legend.Add("Group 2015", bars2015)
// 	p.Legend.Add("Group 2016", bars2016)
// 	p.Legend.Add("Group 2017", bars2017)
// 	p.Legend.Add("Group 2018", bars2018)
// 	p.Legend.Add("Group 2019", bars2019)
// 	p.Legend.Add("Group 2020", bars2020)

// 	p.Legend.Top = true

// 	p.NominalX(groupName...)
// 	if err := p.Save(20*vg.Inch, 22*vg.Inch, "./plot/brandPlot.png"); err != nil {
// 		panic(err)
// 	}

// }
