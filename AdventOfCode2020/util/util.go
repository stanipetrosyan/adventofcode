package util

import (
	"bufio"
	"os"
	"strconv"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ConvertToIntLines(path string) ([]int, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []int

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, n)
	}

	return lines, nil
}

func FindMax(numbers []int) int {
	var max int = 0
	for _, item := range numbers {
		if item > max {
			max = item
		}
	}
	return max
}

func FindMin(numbers []int) int {
	var max int = numbers[0]
	for _, item := range numbers {
		if item < max {
			max = item
		}
	}
	return max
}
