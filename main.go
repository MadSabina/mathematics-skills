package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"sort"
	"math"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error! Please use: go run . data.txt")
		return
	}
	data, err := ReadData(os.Args)
	if err != nil {
		log.Fatalf("Sorry! Unable to read data file: %v", err)
		return
	}
	if data == nil {
		return
	}
	average := FindAverage(data)
	FindMedian(data)
	variance := FindVariance(data, average)
	FindStandardDeviation(variance)
}

func ReadData(datafile []string) ([]float64, error) {
	file := datafile[1]
	numbers, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	n := strings.Split(string(numbers), "\n")
	var nums []float64
	for _, data := range n {
		if data == "" || data == " " {
			continue
		}
		nbrs, err := strconv.Atoi(string(data))
		if err != nil {
			return nil, err
		}
		nums = append(nums, float64(nbrs))
	}
	return nums, nil
}

func FindAverage(data []float64) (average float64) {
	n := float64(len(data))
	var result float64
	for _, numbers := range data {
		result += numbers
	}
	a := result / n
	average = math.Round(a)
	fmt.Println("Average:", average)
	return
}

func FindVariance(data []float64, average float64) (variance float64) {
	numbers := float64(len(data))
	var value float64
	for i := 0; i < len(data); i++ {
		value += (average-data[i])*(average-data[i])
	}
	variance = math.Round(value/numbers)
	fmt.Println("Variance:", int(variance))
	return
}

func FindMedian(data []float64) {
	sort.Float64s(data)
	center := (len(data)-1)/2
	var value float64
	if len(data)%2 == 0 {
		value = (data[center] + data[center+1]) / 2
	} else if len(data)%2 != 0 {
		value = (data[center])
	}
	median := math.Round(value)
	fmt.Println("Median:", median)
}

func FindStandardDeviation(data float64) {
	variance := math.Round(math.Sqrt(data))
	fmt.Println("Standard Deviation:", variance)
}
