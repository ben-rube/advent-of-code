package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var startingData []int
var total int
var nextDay = map[int]int{}
var fishes = map[int]int{}

func main() {
	open, _ := os.Open("./input.txt")
	startingData, _ = GetData(open)
	fmt.Println(startingData)
	//ProcessDays(256)
	partTwo(256)
	fmt.Println(total)

}

func partOne(days int) {
	for i := 0; i < days; i++ {
		fmt.Println(i)
		var fishToAdd int
		for i, fish := range startingData {
			if fish == 0 {
				startingData[i] = 6
				fishToAdd++
			} else {
				startingData[i]--
			}
		}
		for i := 0; i < fishToAdd; i++ {
			startingData = append(startingData, 8)
		}
	}
	return
}

func partTwo(days int) {
	for _, fish := range startingData {
		fishes[fish]++
	}

	for i := 0; i < days; i++ {

		for j := 0; j <= 8; j++ {
			nextDay[j] += fishes[(j+1)%9]
		}
		nextDay[6] += fishes[0]
		fishes = nextDay
	}

	for _, fish := range fishes {
		total += fish
	}

	return
}

func GetData(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []string
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	var intResult []int
	for _, s := range result {
		split := strings.Split(s, ",")
		for _, s2 := range split {
			atoi, _ := strconv.Atoi(s2)
			intResult = append(intResult, atoi)
		}
	}

	return intResult, scanner.Err()
}
