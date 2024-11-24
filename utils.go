package main

import (
    "fmt"
    "strings"
)

func log(msgs ...string) {
    msg := strings.Join(msgs, " ")
    fmt.Printf("\x1b[34mLOG:\x1b[m\t%s\n", msg)
}

func logError(msgs ...string) {
    msg := strings.Join(msgs, " ")
    fmt.Printf("\x1b[31mERROR:\x1b[m\t%s\n", msg)
}

func logWarning(msgs ...string) {
    msg := strings.Join(msgs, " ")
    fmt.Printf("\x1b[33mWARN:\x1b[m\t%s\n", msg)
}
