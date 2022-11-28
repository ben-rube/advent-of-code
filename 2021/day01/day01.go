package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	increments, groupOf3Increments := ProcessesData("./day02_input.txt")
	log.Printf("Total increments %v ", increments)
	log.Printf("Grouped Increments %v", groupOf3Increments)

}

func ProcessesData(fileName string) (int,int) {
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open fil %v", err)
		return 0, 0
	}

	parsedNumbers, err := ReadInts(reader)
	if err != nil {
		log.Fatal("Could not parse strings to ints")
		return 0, 0
	}

	groupOf3Increments := GroupOf3Increments(parsedNumbers)
	increments := LookForIncrements(parsedNumbers)
	return increments, groupOf3Increments

}

func GroupOf3Increments(parsedNumbers []int) int {
	var groupedInts []int
	for i := 0; i < len(parsedNumbers) - 2; i++ {
		groupedInts = append(groupedInts, parsedNumbers[i]+parsedNumbers[i+1]+parsedNumbers[i+2])
	}
	increments := LookForIncrements(groupedInts)
	return increments
}

func LookForIncrements(parsedNumbers []int) (int) {
	previousInt := parsedNumbers[0]
	numIncrements := 0
	for i := 1; i < len(parsedNumbers); i++ {
		currentInt := parsedNumbers[i]
		if currentInt > previousInt {
			numIncrements++
		} else {
			log.Printf("Not incremented cur=%v last=%v", currentInt, previousInt)
		}
		previousInt = currentInt

	}
	return numIncrements
}

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}