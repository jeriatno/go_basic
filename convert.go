package main

import (
    "fmt"
    "strconv"
    )


func main() {
    const nominal = 8000000
    // convert int to string
    salary := strconv.Itoa(nominal);
    // convert string to int
    var bonus, _  = strconv.Atoi("1000000")

    salarybonus := nominal + bonus;

    fmt.Printf("Salary : "+ salary +"\n")
    fmt.Printf("Salary + Bonus : ", salarybonus)
}
