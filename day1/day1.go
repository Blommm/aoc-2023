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
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//log.Printf("(Part 1) Total: %d", part1(file))
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
    words := map[string]int {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
    }

    indexedMap := map[string]int {
        "one": 0,
        "two": 0,
        "three": 0,
        "four": 0,
        "five": 0,
        "six": 0,
        "seven": 0,
        "eight": 0,
        "nine": 0,
    }

rerun:
    for word, value := range words {
        idx := strings.Index(line, word)

        log.Printf("%d (%d)", idx, value)
        
        if idx != -1 {
            if (idx+1 > len(line)) {
                line = line[:idx] + strconv.Itoa(value) + line[idx:]
            } else {
                line = line[:idx+1] + strconv.Itoa(value) + line[idx+1:];
            }
        }

        indexedMap[word] = idx;

        if (indexedMap[word] != -1) {
            goto rerun;      
        }
    }
    return line;
}
