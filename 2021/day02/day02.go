package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Submarine struct {
	horizontal int
	depth      int
	aim        int
}

func main() {
	file, err := os.Open("./day02_input.txt")
	if err != nil {
		log.Fatalf("could not open file %v", err)
	}
	positions := &Submarine{}
	err = positions.ProcessInput(file)
	if err != nil {
		return
	}

	fmt.Printf("Sum=%v\n", positions.horizontal*positions.depth)
}

func (p *Submarine) ProcessInput(r io.Reader) (err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var index int
	var indexToWriteTo string
	for scanner.Scan() {
		if index%2 == 0 {
			indexToWriteTo = scanner.Text()
			index++
			continue
		}
		switch indexToWriteTo {
		case "forward":
			intForward, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			p.horizontal += intForward
			p.depth += intForward * p.aim
		case "up":
			intUp, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			p.aim -= intUp
		case "down":
			intDown, err := strconv.Atoi(scanner.Text())
			if err != nil {
				return err
			}
			p.aim += intDown
		}
		index++
	}
	if scanner.Err() != nil {
		return err
	}

	return
}
