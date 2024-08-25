package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Printf("Enter a parameter\n")
        return 
    }
    argv := []byte(os.Args[1])
    for i, c := range argv {
        fmt.Printf("%c", c-byte(i));
    }
    fmt.Printf("\n")
}
