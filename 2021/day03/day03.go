package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {

	content, _ := ioutil.ReadFile("./day03_data.txt")

	oxygenStrings := strings.Split(string(content), "\n")
	co2Strings := strings.Split(string(content), "\n")
	oxygen := filter(oxygenStrings, func(zeros, ones int) bool {
		return zeros > ones
	})
	co2 := filter(co2Strings, func(zeros, ones int) bool {
		return zeros < ones || zeros == ones
	})

	c, _ := strconv.ParseInt(co2, 2, 64)
	o, _ := strconv.ParseInt(oxygen, 2, 64)
	fmt.Println(c * o)
}

func filter(list []string, pred func(zeros, ones int) bool) string {
	pos := 0
	for len(list) != 1 {
		zeros := 0
		ones := 0

		for _, b := range list {
			if b[pos] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		var bit byte
		if pred(zeros, ones) {
			bit = '1'
		} else {
			bit = '0'
		}
		for i := 0; i < len(list); i++ {
			if list[i][pos] == bit {
				list = append(list[:i], list[i+1:]...)
				i--
			}
		}
		pos++
	}
	return list[0]
}
