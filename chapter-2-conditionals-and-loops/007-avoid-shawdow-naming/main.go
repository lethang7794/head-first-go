package main

import (
        "fmt"
)

func main() {
        var count int = 12
        var suffix string = "minutes of bonus footage"
        var format string = "DVD"

        var languages = append([]string{}, "Espanol")
        fmt.Println(count, suffix, "on", format, languages)
}
