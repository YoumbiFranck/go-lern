package main

import "fmt"

func main() {
    users := map[string]string{
        "admin": "Alice",
        "user":  "Bob",
    }

    // Lecture
    fmt.Println(users["admin"])

    // Ajout
    users["guest"] = "Eve"

    // Vérification
    if val, ok := users["user"]; ok {
        fmt.Println("Utilisateur :", val)
    }

    // Suppression
    delete(users, "guest")

    // Parcours
    for role, name := range users {
        fmt.Println(role, "=>", name)
    }

    fmt.Println("Taille :", len(users))
}


