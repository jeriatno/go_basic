package main

import "fmt"

func main() {
    traffic := "silver"

    switch {
    case traffic == "red":
            fmt.Printf("stop \n")
        case traffic == "green":
            fmt.Printf("start \n")
        case traffic == "yellow":
            fmt.Printf("ready \n")
        default:
            fmt.Printf("error \n")
    }
}
