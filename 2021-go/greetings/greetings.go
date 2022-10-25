package greetings

import (
    "errors"
    "fmt"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("no name given")
    }

    message := fmt.Sprintf("Hi %v, Welcome!", name)
    return message, nil
}
