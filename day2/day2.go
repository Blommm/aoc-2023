package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    sum := 0
    sum2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        sum += part1(scanner.Text())
        sum2 += part2(scanner.Text())
	}

    log.Println(sum)
    log.Println(sum2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

var limits = map[string]int {
    "red": 12,
    "green": 13,
    "blue": 14,
}

func part1(line string) int {
    log.Println(line)
    s := strings.Split(line, ":")
    gid := 0;
    fmt.Sscanf(s[0], "Game %d", &gid)
    sets := strings.Split(s[1], ";")

    for _, set := range sets {
        colors := strings.Split(set, ",")
        for _, cube := range colors {
            log.Println(cube)
            color := ""
            value := 0

            fmt.Sscanf(cube, "%d %s", &value, &color)

            if (value > limits[color]) {
                gid = 0;
            }
        }
    }

    return gid
}

func part2(line string) int {
    s := strings.Split(line, ":")
    gid := 0;
    fmt.Sscanf(s[0], "Game %d", &gid)
    sets := strings.Split(s[1], ";")

    minCubes := map[string]int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }

    for _, set := range sets {
        colors := strings.Split(set, ",")
        for _, cube := range colors {
            color := ""
            value := 0

            fmt.Sscanf(cube, "%d %s", &value, &color)
            
            if value > minCubes[color] {
                minCubes[color] = value;
            }
        }
    }

    return minCubes["red"] * minCubes["green"] * minCubes["blue"]
}
