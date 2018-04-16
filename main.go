package main

import (
    "fmt"
    "os"
    "os/user"
    "orange/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    fmt.Printf("Hello %s! This is the Orange interpreter!\n", user.Username)
    fmt.Printf("Feel free to try some commands\n")
    repl.Start(os.Stdin, os.Stdout)
}
