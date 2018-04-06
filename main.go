package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
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
	fmt.Fprintf(os.Stdout, "%s", histogram(a))
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

func histogram(a []float64) string {
	// Alloc the buckets.
	type bucket struct {
		min, max float64
		count    int
	}
	var (
		n       = 10
		buckets = make([]bucket, n)
		delta   = (max(a) - min(a)) / float64(n)
	)
	for i := 0; i < n; i++ {
		buckets[i].min = min(a) + (float64(i+0) * delta)
		buckets[i].max = min(a) + (float64(i+1) * delta)
	}

	// Count into the buckets.
	for _, v := range a {
		for i := range buckets {
			if buckets[i].min <= v && v <= buckets[i].max {
				buckets[i].count++
			}
		}
	}

	// Draw the buckets.
	draw := func(v, total int) string {
		n := (float64(v) / float64(total)) * 100
		return strings.Repeat("*", int(n))
	}
	var (
		buf = &bytes.Buffer{}
		tw  = tabwriter.NewWriter(buf, 0, 1, 1, ' ', 0)
	)
	for _, b := range buckets {
		fmt.Fprintf(tw, "%.4f-%.4f\t%d\t%s\n", b.min, b.max, b.count, draw(b.count, len(a)))
	}
	tw.Flush()
	return buf.String()
}
