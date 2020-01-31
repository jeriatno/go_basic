package main

import "fmt"

// string datatype
var first_name = "Iwanna"
var last_name  = "Jeriatno"
// int datatype
var weight = 70
var height = 170

func main() {
    var body int

    // string datatype
    fullname    := first_name+" "+last_name
    // int datatype
    body        = weight + height

    fmt.Printf("My name is " +fullname+ " and body is ", body)
}
