package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	a := []float64{}
	for s.Scan() {
		f, err := strconv.ParseFloat(s.Text(), 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing float: %v\n", err)
			continue
		}
		a = append(a, f)
	}

	fmt.Fprintf(os.Stdout, "n %d\n", len(a))
	fmt.Fprintf(os.Stdout, "min %.4f\n", min(a))
	fmt.Fprintf(os.Stdout, "max %.4f\n", max(a))
	fmt.Fprintf(os.Stdout, "sum %.4f\n", sum(a))
	fmt.Fprintf(os.Stdout, "mean %.4f\n", mean(a))
	fmt.Fprintf(os.Stdout, "median %.4f\n", median(a))
	fmt.Fprintf(os.Stdout, "modes %v\n", modes(a))
	fmt.Fprintf(os.Stdout, "stdev %.4f\n", stdev(a))
}

// https://github.com/ae6rt/golang-examples/blob/master/goeg/src/statistics_ans/statistics.go

func min(a []float64) float64 {
	min := math.MaxFloat64
	for _, f := range a {
		if f < min {
			min = f
		}
	}
	return min
}

func max(a []float64) float64 {
	var max float64
	for i, f := range a {
		if i == 0 || f > max {
			max = f
		}
	}
	return max
}

func sum(a []float64) float64 {
	var sum float64
	for _, f := range a {
		sum += f
	}
	return sum
}

func mean(a []float64) float64 {
	return sum(a) / float64(len(a))
}

func median(a []float64) float64 {
	sort.Float64s(a)
	return a[len(a)/2]
}

func modes(a []float64) []float64 {
	frequencies := make(map[float64]int, len(a))
	highestFrequency := 0
	for _, x := range a {
		frequencies[x]++
		if frequencies[x] > highestFrequency {
			highestFrequency = frequencies[x]
		}
	}
	modes := []float64{}
	for x, frequency := range frequencies {
		if frequency == highestFrequency {
			modes = append(modes, x)
		}
	}
	if highestFrequency == 1 || len(modes) == len(a) {
		modes = []float64{}
	}
	sort.Float64s(modes)
	return modes
}

func stdev(a []float64) float64 {
	mean := mean(a)
	total := 0.0
	for _, f := range a {
		total += math.Pow(f-mean, 2)
	}
	variance := total / float64(len(a)-1)
	return math.Sqrt(variance)
}
