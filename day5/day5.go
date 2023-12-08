package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

var Seeds []int
var SeedRanges []string

type MapTo struct{
    DestinationStart int
    Range int
}

var SeedToSoil map[string]MapTo = make(map[string]MapTo)
var SoilToFertilizer map[string]MapTo = make(map[string]MapTo)
var FertilizerToWater map[string]MapTo = make(map[string]MapTo)
var WaterToLight map[string]MapTo = make(map[string]MapTo)
var LightToTemperature map[string]MapTo = make(map[string]MapTo)
var TemperatureToHumidity map[string]MapTo = make(map[string]MapTo)
var HumidityToLocation map[string]MapTo = make(map[string]MapTo)

func main() {
    content, err := ioutil.ReadFile("./input")
    if err != nil {
        log.Fatal(err)
    }

    // part1 := p1(string(content))
    // Should work just takes a year
    part2 := p2(string(content))
    log.Println(part2)

    // log.Println(part1)
}

func p1(content string) int {
    parseMaps(content)
    parseSeeds(content)

    var Locations []int

    for _, seed := range Seeds {
        soil := checkRange(SeedToSoil, seed)
        fertilizer := checkRange(SoilToFertilizer, soil)
        water := checkRange(FertilizerToWater, fertilizer)
        light := checkRange(WaterToLight, water)
        temperature := checkRange(LightToTemperature, light)
        humidity := checkRange(TemperatureToHumidity, temperature)
        location := checkRange(HumidityToLocation, humidity)

        Locations = append(Locations, location)
    }

    sort.Ints(Locations)
    return Locations[0]
}

func timer(name string) func() {
    start := time.Now()
    return func() {
        log.Printf("%s took %v\n", name, time.Since(start))
    }
}

func p2(content string) int {
    defer timer("p2")()
    parseMaps(content)
    parseSeedRange(content)

    var Locations []int

    log.Println(SeedRanges)
    for _, seedRange := range SeedRanges {
        start, rangeLen := 0, 0

        fmt.Sscanf(seedRange, "%d-%d", &start, &rangeLen)
        end := start+rangeLen
        log.Printf("Next range (%d - %d)", start, end)

        for i := start; i < end; i++ {
            soil := checkRange(SeedToSoil, i)
            fertilizer := checkRange(SoilToFertilizer, soil)
            water := checkRange(FertilizerToWater, fertilizer)
            light := checkRange(WaterToLight, water)
            temperature := checkRange(LightToTemperature, light)
            humidity := checkRange(TemperatureToHumidity, temperature)
            location := checkRange(HumidityToLocation, humidity)

            Locations = append(Locations, location)
        }
    }

    sort.Ints(Locations)
    log.Println(Locations)
    return Locations[0]
}

func checkRange(mapToCheck map[string]MapTo, checkValue int) int {
    ret := -1
    for inRange, Dest := range mapToCheck {
        start, end := 0, 0;
        fmt.Sscanf(inRange, "%d-%d", &start, &end)
        if (start <= checkValue && end > checkValue) {
            ret = Dest.DestinationStart + (checkValue - start);   
        }
    }
    if (ret == -1) {
        return checkValue;
    }
    return ret;
}
func parseSeedRange(fileContent string) {
    seedList := strings.Split(fileContent, "\n")[0]
    seedArray := strings.Split(seedList[7:], " ");
    
    for i, s := range seedArray {
        if (i % 2 != 0) {
            continue;
        }

        seedStart, seedEnd := 0, 0

        fmt.Sscanf(s, "%d", &seedStart)
        fmt.Sscanf(seedArray[i+1], "%d", &seedEnd)

        SeedRanges = append(SeedRanges, fmt.Sprintf("%d-%d", seedStart, seedEnd));
    }

}

func parseSeeds(fileContent string) {
    seedList := strings.Split(fileContent, "\n")[0]
    seedArray := strings.Split(seedList[7:], " ");

    for _, s := range seedArray {
        appendSeed := 0
        fmt.Sscanf(s, "%d", &appendSeed)
        Seeds = append(Seeds, appendSeed);
    }
}

func parseMaps(fileContent string) {
    lines := strings.Split(fileContent, "\n")

    currentParse := ""
    for _, line := range lines {
        if (line == "") {
            log.Print("\n")
            currentParse = ""
            continue;
        }

        if (currentParse == "") {
            fmt.Sscanf(line, "%s\tmap:", &currentParse)
            continue;
        }

        rangeStart, sourceRangeStart, rangeLength := 0, 0, 0
        fmt.Sscanf(line, "%d %d %d", &rangeStart, &sourceRangeStart, &rangeLength)

        switch(currentParse) {
            case "seed-to-soil":
                log.Println("seed-to-soil ", rangeStart, sourceRangeStart, rangeLength)
                SeedToSoil[fmt.Sprintf("%d-%d", sourceRangeStart, sourceRangeStart + rangeLength)] = MapTo{
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                }
                break;
            case "soil-to-fertilizer":
                log.Println("soil-to-fertilizer ", rangeStart, sourceRangeStart, rangeLength)
                SoilToFertilizer[fmt.Sprintf("%d-%d", sourceRangeStart, sourceRangeStart + rangeLength)] = MapTo{
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                }
                break;
            case "fertilizer-to-water":
                log.Println("fertilizer-to-water ", rangeStart, sourceRangeStart, rangeLength)
                FertilizerToWater[fmt.Sprintf("%d-%d", sourceRangeStart, sourceRangeStart + rangeLength)] = MapTo{
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                }
                break;
            case "water-to-light":
                log.Println("water-to-light ", rangeStart, sourceRangeStart, rangeLength)
                WaterToLight[fmt.Sprintf("%d-%d", sourceRangeStart, sourceRangeStart + rangeLength)] = MapTo{
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                }
                break;
            case "light-to-temperature":
                log.Println("light-to-temperature ", rangeStart, sourceRangeStart, rangeLength)
                LightToTemperature[fmt.Sprintf("%d-%d", sourceRangeStart, sourceRangeStart + rangeLength)] = MapTo{
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                }
                break;
            case "temperature-to-humidity":
                log.Println("temperature-to-humidity ", rangeStart, sourceRangeStart, rangeLength)
                TemperatureToHumidity[fmt.Sprintf("%d-%d", sourceRangeStart, sourceRangeStart + rangeLength)] = MapTo{
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                }
                break;
            case "humidity-to-location":
                log.Println("humidity-to-location ", rangeStart, sourceRangeStart, rangeLength)
                HumidityToLocation[fmt.Sprintf("%d-%d", sourceRangeStart, sourceRangeStart + rangeLength)] = MapTo{
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                }
                break;
            default: {
                log.Println("Default")
            }
        }

    }

}
