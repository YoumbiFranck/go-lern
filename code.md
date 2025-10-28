# Code 
```go
package main

import (
	"fmt"
	"time"
)

// Structure pour démontrer les types personnalisés
type Person struct {
	Name string
	Age  int
}

// Méthode pour la structure Person
func (p Person) Greet() string {
	return fmt.Sprintf("Bonjour, je m'appelle %s et j'ai %d ans!", p.Name, p.Age)
}

func main() {
	fmt.Println("=== Bienvenue dans votre environnement Go! ===")
	fmt.Println()

	// 1. Variables et types de base
	fmt.Println("1. Variables et types:")
	var message string = "Hello, Go!"
	nombre := 42 // Déclaration courte
	estVrai := true
	fmt.Printf("   Message: %s\n", message)
	fmt.Printf("   Nombre: %d\n", nombre)
	fmt.Printf("   Boolean: %v\n", estVrai)
	fmt.Println()

	// 2. Structures
	fmt.Println("2. Structures:")
	personne := Person{Name: "Alice", Age: 30}
	fmt.Printf("   %s\n", personne.Greet())
	fmt.Println()

	// 3. Tableaux et Slices
	fmt.Println("3. Collections:")
	fruits := []string{"pomme", "banane", "orange"}
	fmt.Printf("   Fruits: %v\n", fruits)
	fmt.Printf("   Premier fruit: %s\n", fruits[0])
	fmt.Println()

	// 4. Boucles
	fmt.Println("4. Boucle for:")
	for i, fruit := range fruits {
		fmt.Printf("   %d: %s\n", i+1, fruit)
	}
	fmt.Println()

	// 5. Conditions
	fmt.Println("5. Conditions:")
	age := 25
	if age >= 18 {
		fmt.Println("   Vous êtes majeur")
	} else {
		fmt.Println("   Vous êtes mineur")
	}
	fmt.Println()

	// 6. Fonctions
	fmt.Println("6. Fonctions:")
	resultat := additionner(10, 32)
	fmt.Printf("   10 + 32 = %d\n", resultat)
	fmt.Println()

	// 7. Goroutines (concurrence)
	fmt.Println("7. Goroutines:")
	go afficherAvecDelai("Goroutine 1", 100*time.Millisecond)
	go afficherAvecDelai("Goroutine 2", 50*time.Millisecond)
	
	// Attendre un peu pour voir les goroutines s'exécuter
	time.Sleep(300 * time.Millisecond)
	fmt.Println()

	fmt.Println("=== Fin du programme ===")
}

// Fonction simple d'addition
func additionner(a, b int) int {
	return a + b
}

// Fonction pour démontrer les goroutines
func afficherAvecDelai(message string, delai time.Duration) {
	time.Sleep(delai)
	fmt.Printf("   %s (après %v)\n", message, delai)
}
```