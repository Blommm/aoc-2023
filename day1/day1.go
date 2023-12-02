package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("C:\\Dev\\adventofcode-2023\\day1\\input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Printf("(Part 1) Total: %d", part1(file))
	log.Printf("(Part 2) Total: %d", part2(file))
}

func part1(file *os.File) int {
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += extractNumbers(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func part2(file *os.File) int {
	sum := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		cleanedString := convertNumberWordsToNumbers(scanner.Text())
		sum += extractNumbers(cleanedString)
		log.Printf("%s (%s) = %d", cleanedString, scanner.Text(), extractNumbers(cleanedString))
	}

	return sum
}

func extractNumbers(line string) int {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(line, -1)

	s1 := numbers[0]
	s2 := numbers[len(numbers)-1]

	first := s1[0:1]
	last := s2[len(s2)-1:]

	ret, err := strconv.Atoi(first + last)

	if err != nil {
		log.Fatal(err)
	}

	return ret
}

func convertNumberWordsToNumbers(line string) string {
	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9")

	return replacer.Replace(line)
}
