package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"golang.org/x/exp/slices"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
    
    part2 := p2(file)
    // part1 := p1(file)

    // log.Printf("(Part1) Sum: %d", part1)
    log.Printf("(Part2) Sum: %d", part2)
}

type Card struct{
    Nummer int
    AmountWon int
    CardNumbers []int
    CardWinNumbers []int
    CardCopies []Card
}

func p1(file *os.File) int{
    sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        card := parseCards(scanner.Text())
        if card.AmountWon > 0 {
            sum += int(math.Pow(2, float64(card.AmountWon-1)))
        }
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    return sum
}

func timer(name string) func() {
    start := time.Now()
    return func() {
        log.Printf("%s took %v\n", name, time.Since(start))
    }
}

func p2(file *os.File) int{
    defer timer("p2")()
    sum := 0
    var cards []Card

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
        card := parseCards(scanner.Text())
        cards = append(cards, card)
	}

    handledCards := handleCopies(cards)

    for _, hcard := range handledCards {
        // log.Println(hcard.Nummer, len(hcard.CardCopies)+1)
        sum += len(hcard.CardCopies)+1
    }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

    return sum
}

func parseCards(line string) Card {
    var card Card

    cardSplit := strings.Split(line, ":")

    cnumber := 0
    fmt.Sscanf(cardSplit[0], "Card %d", &cnumber)
    
    card.Nummer = cnumber
    card.AmountWon = 0

    scratches := strings.Split(cardSplit[1], "|")
    wnumbers := strings.Split(scratches[0], " ");
    for _, wnum := range wnumbers {
        num := 0
        fmt.Sscanf(wnum, "%d", &num)
        if (num == 0) {
            continue;
        }
        card.CardWinNumbers = append(card.CardWinNumbers, num)
    }

    cnumbers := strings.Split(scratches[1], " ");
    for _, cnum := range cnumbers {
        num := 0
        fmt.Sscanf(cnum, "%d", &num)
        if (num == 0) {
            continue;
        }
        card.CardNumbers = append(card.CardNumbers, num)
    }
    

    for _, cardNumber := range card.CardNumbers {
        if slices.Contains(card.CardWinNumbers, cardNumber) {
            card.AmountWon += 1
        }
    }

    return card;
}

func handleCopies(cards []Card) []Card {
    for _, card := range cards {
        if card.AmountWon > 0 {
            for i := card.Nummer+1; i <= card.Nummer + card.AmountWon; i++ {
                cards[i-1].CardCopies = append(cards[i-1].CardCopies, cards[i-1])
            }
        }

        if len(card.CardCopies) > 0 {
            for _, CCopy := range card.CardCopies {
                if CCopy.AmountWon > 0 {
                    for i := card.Nummer+1; i <= card.Nummer + CCopy.AmountWon; i++ {
                        cards[i-1].CardCopies = append(cards[i-1].CardCopies, cards[i-1])
                    }
                }
            }
        }
    }

    return cards;
}
