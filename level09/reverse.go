package main

import (
    "fmt"
    "os"
)

func main() {
    argv := []byte(os.Args[1])
    for i, c := range argv {
        fmt.Printf("%c", c-byte(i));
    }
    fmt.Printf("\n")
}
