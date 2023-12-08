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
var SeedRanges map[int]int = make(map[int]int)

type MapTo struct{
    Start int
    DestinationStart int
    Range int
}

var SeedToSoil []MapTo
var SoilToFertilizer []MapTo
var FertilizerToWater []MapTo
var WaterToLight []MapTo
var LightToTemperature []MapTo
var TemperatureToHumidity []MapTo
var HumidityToLocation []MapTo

func main() {
    content, err := ioutil.ReadFile("./input")
    if err != nil {
        log.Fatal(err)
    }

    // part1 := p1(string(content))
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

    for start, end := range SeedRanges {
        log.Println("new range")
        lowestLocation := -1
        for i := start; i < end; i++ {
            soil := checkRange(SeedToSoil, i)
            fertilizer := checkRange(SoilToFertilizer, soil)
            water := checkRange(FertilizerToWater, fertilizer)
            light := checkRange(WaterToLight, water)
            temperature := checkRange(LightToTemperature, light)
            humidity := checkRange(TemperatureToHumidity, temperature)
            location := checkRange(HumidityToLocation, humidity)
            
            if (location < lowestLocation || lowestLocation == -1) {
                lowestLocation = location
            }
        }
        Locations = append(Locations, lowestLocation)
    }

    sort.Ints(Locations)
    return Locations[0]
}

func checkRange(mapToCheck []MapTo, checkValue int) int {
    ret := -1
    for _, d := range mapToCheck {
        if (d.Start <= checkValue && d.Start + d.Range > checkValue) {
            ret = d.DestinationStart + (checkValue - d.Start);   
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

        seedStart, seedRange := 0, 0

        fmt.Sscanf(s, "%d", &seedStart)
        fmt.Sscanf(seedArray[i+1], "%d", &seedRange)
        
        SeedRanges[seedStart] = seedStart+seedRange
    }

    log.Println(SeedRanges)

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
                SeedToSoil = append(SeedToSoil, MapTo{
                    Start: sourceRangeStart,
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                })
                break;
            case "soil-to-fertilizer":
                log.Println("soil-to-fertilizer ", rangeStart, sourceRangeStart, rangeLength)
                SoilToFertilizer = append(SoilToFertilizer, MapTo{
                    Start: sourceRangeStart,
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                })
            case "fertilizer-to-water":
                log.Println("fertilizer-to-water ", rangeStart, sourceRangeStart, rangeLength)
                FertilizerToWater = append(FertilizerToWater, MapTo{
                    Start: sourceRangeStart,
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                })
                break;
            case "water-to-light":
                log.Println("water-to-light ", rangeStart, sourceRangeStart, rangeLength)
                WaterToLight= append(WaterToLight, MapTo{
                    Start: sourceRangeStart,
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                })
                break;
            case "light-to-temperature":
                log.Println("light-to-temperature ", rangeStart, sourceRangeStart, rangeLength)
                LightToTemperature = append(LightToTemperature, MapTo{
                    Start: sourceRangeStart,
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                })
                break;
            case "temperature-to-humidity":
                log.Println("temperature-to-humidity ", rangeStart, sourceRangeStart, rangeLength)
                TemperatureToHumidity = append(TemperatureToHumidity, MapTo{
                    Start: sourceRangeStart,
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                })
                break;
            case "humidity-to-location":
                log.Println("humidity-to-location ", rangeStart, sourceRangeStart, rangeLength)
                HumidityToLocation = append(HumidityToLocation, MapTo{
                    Start: sourceRangeStart,
                    DestinationStart: rangeStart,
                    Range: rangeLength,
                })
                break;
            default: {
                log.Println("Default")
            }
        }

    }

}
