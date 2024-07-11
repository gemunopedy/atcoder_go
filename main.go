package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    first_input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    second_input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading input:", err)
        return
    }

    first_input = strings.TrimSpace(first_input)
    second_input = strings.TrimSpace(second_input)

    first_inputs := strings.Split(first_input, " ")
    second_inputs := strings.Split(second_input, " ")

    fmt.Println(first_inputs)
    fmt.Println(second_inputs)

}
