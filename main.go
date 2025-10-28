package main 

import "fmt"

func main() {


for i := 0; i < 10; i++ {
    if i == 5 {
		fmt.Println("ignore 5")
        continue // ignore 5
		
    }
    if i == 8 {
        break // sort de la boucle
    }
    fmt.Println(i)
}



}