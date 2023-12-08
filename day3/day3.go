package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func main() {
    content, err := ioutil.ReadFile("./input")
    if err != nil {
        log.Fatal(err)
    }

    log.SetFlags(0)

    part1 := p1(string(content))

    log.Println(part1)

}

type Block struct{
    Number int
    Top string
    Middle string
    Bottom string
}

func p1(content string) int {
    lines := strings.Split(content, "\n");
    blocks := parseBlocks(lines)

    sum := 0

    for _, block := range blocks {
        topHasPart := checkForPart(block.Top)
        middleHasPart := checkForPart(block.Middle)
        bottomHasPart := checkForPart(block.Bottom)

        log.Printf("\n%s\n%s\n%s\ntop: %t, middle: %t, bottom: %t\n", block.Top, block.Middle, block.Bottom, topHasPart, middleHasPart, bottomHasPart)
        if (topHasPart || middleHasPart || bottomHasPart) {
            log.Printf("Added nr %d", block.Number)
            sum += block.Number
        }
    }

    return sum
}

func checkForPart(string string) bool {
    hasPart := false
    for _, char := range string {
        if (char == '.' || unicode.IsDigit(char)) {
            continue;
        }
        hasPart = true;
    }
    return hasPart
}

func parseBlocks(lines []string) []Block {
    var Blocks []Block
    
    for i, line := range lines {
        startNumber, endNumber := -1, -1
        for pos, character := range line {
            endNumber = pos
            if (unicode.IsDigit(character) && startNumber == -1) {
                startNumber = pos
                continue;
            } else if (unicode.IsDigit(character) && len(line)-1 != pos) {
                continue;
            } else if (len(line)-1 == pos && character != '.') {
                endNumber = pos+1
            }
            if (startNumber == -1) {
                continue;
            }

            start := getPossibleStart(startNumber)
            end := getPossibleEnd(len(line), endNumber)

            var block Block

            number, err := strconv.Atoi(line[startNumber:endNumber])
            if err != nil {
                log.Panic(err)
            }

            block.Number = number

            if (i-1 != -1) {
                block.Top = lines[i-1][start:end]
            }

            block.Middle = line[start:end]

            if (i+1 < len(lines)-1) {
                block.Bottom = lines[i+1][start:end]
            }

            Blocks = append(Blocks, block)

            startNumber = -1
            endNumber = -1
        }
    }

    return Blocks
}

func getPossibleStart(startNumber int) int {
    if (startNumber-1 != -1) {
        return startNumber-1
    }
    return startNumber
}

func getPossibleEnd(lengthOfString int, endNumber int) int {
    if (lengthOfString >= endNumber+1) {
        return endNumber+1
    }
    return endNumber
}
