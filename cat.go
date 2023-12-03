
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("Usage: cat [file1 file2 ...]")
        os.Exit(1)
    }
    for _, fname := range args {
        file, err := os.Open(fname)
        if err != nil {
            fmt.Println("Error opening file %s: %v", fname, err);
            os.Exit(1)
        }
        defer file.Close()
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            fmt.Println(scanner.Text())
        }
    }
    os.Exit(0)
}
