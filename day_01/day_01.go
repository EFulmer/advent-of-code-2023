package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const DefaultFileName = "day_01/input"
const Digits = "0123456789"

func main() {
    os.Exit(run())
}

func run() int {
    fileName, rc := getFileName()
    if rc != 0 {
        return 0
    }

    result, err := readFileAndComputeSum(fileName)
    if err != 0 {
        return err
    }
    fmt.Println("Answer =", result)
    return 0
}

func getFileName() (string, int) {
    args := os.Args[1:]
    if len(args) == 0 {
        return DefaultFileName, 0
    } else if len(args) == 1 {
        return args[0], 0
    } else {
        fmt.Println("Usage: day_01.go [filename, defaults to input]")
        return "", 1
    }
}

func readFileAndComputeSum(fileName string) (int, int) {
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println("error opening file: ", err)
        return -1, 1
    }

    scanner := bufio.NewScanner(file)
    sum := 0
    for scanner.Scan() { // defaults to lines, which I want
        curLine := scanner.Text()
        firstIndex := strings.IndexAny(curLine, Digits)
        lastIndex := strings.LastIndexAny(curLine, Digits)
        firstChar := string(curLine[firstIndex])
        lastChar := string(curLine[lastIndex])
        numberAsString := firstChar + lastChar
        number, err := strconv.Atoi(numberAsString)
        if err != nil {
            fmt.Println("Error converting string to number: ", err)
            return -1, 1
        }
        sum += number
    }
    return sum, 0
}
