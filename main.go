package main 

import "fmt"

func main() {


day := "mardi"

switch day {
case "lundi":
    fmt.Println("Début de semaine")
case "vendredi":
    fmt.Println("Presque le week-end")
default:
    fmt.Println("Jour ordinaire")
}


}