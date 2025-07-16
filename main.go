package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    var unusedVar int // unused variable

    file, _ := os.Open("nonexistent.txt") // error is ignored

    fmt.Println("Hello, Linter!") // no newline at end of file

    // Shadowed variable
    x := 10
    if true {
        x := 20 // shadows outer x
        fmt.Println(x)
    }

    // Inefficient string concatenation in loop
    result := ""
    for i := 0; i < 10; i++ {
        result += fmt.Sprintf("%d,", i)
    }
    fmt.Println(result)

    // Cyclomatic complexity: multiple branches
    val := 5
    if val == 1 {
        fmt.Println("One")
    } else if val == 2 {
        fmt.Println("Two")
    } else if val == 3 {
        fmt.Println("Three")
    } else if val == 4 {
        fmt.Println("Four")
    } else if val == 5 {
        fmt.Println("Five")
    } else {
        fmt.Println("Other")
    }

    // Deprecated function usage (example: os.Exit is not deprecated but used here for demonstration)
    os.Exit(1)
}
