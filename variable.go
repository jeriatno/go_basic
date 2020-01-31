package main

import "fmt"

/* first step */
var first_name = "Iwanna"
var last_name  = "Jeriatno"
var fullname   = first_name+" "+last_name

/* second step */
// var first_name, last_name = "Iwanna", "Jeriatno"
// var fullname   = first_name+" "+last_name

/* string null */
// var fullname = ""
/* or */
// var fullname string


func main() {
    fmt.Printf("hello "+fullname+" !!!")
}
