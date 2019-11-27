package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/Arafatk/glot"
)

//N number of elements
const N = 100000

func main() {
	var xMax, xMin, yMax, yMin float64
	xCords := make([]float64, N)
	yCords := make([]float64, N)
	dat, _ := ioutil.ReadFile("./input-luke-22.txt")
	for i, line := range strings.Split(string(dat), "\n") {
		crds := strings.Split(line, ",")
		if len(crds) != 2 || i == N {
			break
		}
		x, _ := strconv.ParseFloat(crds[0], 64)
		y, _ := strconv.ParseFloat(crds[1], 64)
		xCords[i] = x
		yCords[i] = y
		if x > xMax {
			xMax = x
		} else if x < xMin {
			xMin = x
		} else if y > yMax {
			yMax = y
		} else if y < yMin {
			yMin = y
		}
	}
	plot, _ := glot.NewPlot(2, false, false)
	plot.AddPointGroup("cords", "points", [][]float64{xCords, yCords})
	plot.SavePlot("circle.png")
	fmt.Println(int(2 * math.Pi * math.Sqrt((math.Pow((xMax-xMin)/2.0, 2)+math.Pow((yMax-yMin)/2.0, 2))/2.0)))
}
