package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
    "strings"
)

const RedMax   = 12
const GreenMax = 13
const BlueMax  = 14

var redRegexp   = regexp.MustCompile("(\\d+) red")
var blueRegexp  = regexp.MustCompile("(\\d+) blue")
var greenRegexp = regexp.MustCompile("(\\d+) green")

var gameRegexp  = regexp.MustCompile("Game (\\d+)")

type GameResult struct {
    red   int
    blue  int
    green int
}

func newGameResult(description string) *GameResult {

    var redVal int
    var blueVal int
    var greenVal int

    redMatches := redRegexp.FindStringSubmatch(description)
    if len(redMatches) == 0 {
        redVal = 0
    } else {
        redVal, _ = strconv.Atoi(redMatches[1])
    }

    blueMatches := blueRegexp.FindStringSubmatch(description)
    if len(blueMatches) == 0 {
        blueVal = 0
    } else {
        blueVal, _ = strconv.Atoi(blueMatches[1])
    }

    greenMatches := greenRegexp.FindStringSubmatch(description)
    if len(greenMatches) == 0 {
        greenVal = 0
    } else {
        greenVal, _ = strconv.Atoi(greenMatches[1])
    }

    result := GameResult{red: redVal, blue: blueVal, green: greenVal}
    return &result
}

func main() {
    var fileName string
    args := os.Args

    if len(args) == 1 {
        fileName = "day_02/inputPartOne"
    } else if len(args) == 2 {
        fileName = args[1]
    } else {
        fmt.Println("Usage: day_02 [filename, defaults to day_02/inputPartOne")
        os.Exit(1)
    }

    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("Error opening file:", err)
        os.Exit(2)
    }

    scanner := bufio.NewScanner(file)

    gamesToValidity := make(map[string]bool)
    gamesToResults := make(map[string]*[]GameResult)

    for scanner.Scan() {
        curLine := scanner.Text()
        split := strings.Split(curLine, ":")

        if len(split) != 2 {
            fmt.Println("Malformed line:", curLine)
            os.Exit(3)
        }

        gameSection := split[0]
        gameMatch:= gameRegexp.FindStringSubmatch(gameSection)

        if len(gameMatch) == 0 {
            fmt.Println("The line was malformed, expecting it to start with \"Game #\":", curLine)
            os.Exit(3)
        }

        gameNumber := gameMatch[1]
        gameDescriptions := strings.Split(split[1], ";")
        gamesToResults[gameNumber] = readGamesFromLine(gameDescriptions)

        // check if game is valid
        gamesToValidity[gameNumber] = isGameValid(gamesToResults[gameNumber])

    }

    sum := getSumOfGameIDs(&gamesToValidity)
    fmt.Println(sum)

    sumOfPowerSets := getSumOfPowerSets(&gamesToResults)
    fmt.Println(sumOfPowerSets)

    os.Exit(0)
}

func readGamesFromLine(gameDescriptions []string) *[]GameResult {
    var results []GameResult
    for _, gd := range gameDescriptions {
        res := newGameResult(gd)
        results = append(results, *res)
    }

    return &results
}

func isGameValid(rounds *[]GameResult) bool {
    for _, gameResult := range *rounds {
        if gameResult.red > RedMax {
            return false
        }
        if gameResult.blue > BlueMax {
            return false
        }
        if gameResult.green > GreenMax {
            return false
        }
    }
    return true
}

func getSumOfGameIDs(gamesToValidity *map[string]bool) int {
    sum := 0
    for k, v := range *gamesToValidity {
        if v {
            intK, _ := strconv.Atoi(k)
            sum += intK
        }
    }

    return sum
}

func getSumOfPowerSets(gamesToResults *map[string]*[]GameResult) int {
    sum := 0
    for _, curResults := range *gamesToResults {
        minRed, minBlue, minGreen := getMinSet(curResults)
        pow := minRed * minBlue * minGreen
        sum += pow
    }
    return sum
}

func getMinSet(results *[]GameResult) (minRed, minBlue, minGreen int) {
    minRed, minBlue, minGreen = 0, 0, 0

    for _, result := range *results {
        if result.red > minRed {
            minRed = result.red
        }
        if result.blue > minBlue {
            minBlue = result.blue
        }
        if result.green > minGreen {
            minGreen = result.green
        }
    }
    return
}
