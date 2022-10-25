package main

import (
    "fmt"
    "log"

    "golang.org/x/example/stringutil"

    "example.com/greetings"
)

func main() {
    fmt.Println(stringutil.Reverse("Hello"))

    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    message, err := greetings.Hello("")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
