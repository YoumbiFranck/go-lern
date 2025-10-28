## 🧱 1. Les types de données en Go

En Go, tout est **fortement typé** : chaque variable a un type précis, connu à la compilation.
Cela rend ton code **fiable et performant**.

---

### 🔹 Les grands groupes de types

| Catégorie                         | Exemples                          | Description rapide                      |
| --------------------------------- | --------------------------------- | --------------------------------------- |
| **Numériques**                    | `int`, `float64`, `complex64`     | Nombres entiers, flottants ou complexes |
| **Booléens**                      | `bool`                            | Valeurs logiques : `true` ou `false`    |
| **Chaînes de caractères**         | `string`                          | Texte, séquence de caractères UTF-8     |
| **Tableaux & Slices**             | `[5]int`, `[]int`                 | Collections d’éléments du même type     |
| **Maps**                          | `map[string]int`                  | Dictionnaire clé-valeur                 |
| **Structs**                       | `struct { name string; age int }` | Regroupement de données                 |
| **Interfaces / Pointeurs / etc.** | `interface{}`, `*int`             | Concepts avancés (plus tard)            |

---

## ⚙️ 2. Déclaration de variables

Deux syntaxes possibles :

```go
// Syntaxe complète
var age int = 30

// Go peut inférer le type automatiquement
name := "Alice" // type = string

// Déclaration multiple
var x, y = 10, 20
```

💡 *Astuce :* utilise `:=` pour aller vite, sauf dans les blocs globaux (où seul `var` est permis).

---

## 🔢 3. Types numériques

```go
var a int = 10        // entier
var b float64 = 3.14  // réel
var c uint = 255      // entier non signé
```

🧮 **Remarque :** Go ne fait pas de conversions implicites entre types numériques :

```go
var x int = 10
var y float64 = 3.5
// var z float64 = x + y ❌ Erreur
var z float64 = float64(x) + y // ✅
```

---

## 🧠 4. Booléens

```go
isAdmin := true
if isAdmin {
    fmt.Println("Accès autorisé")
}
```

---

## 🧵 5. Chaînes de caractères (`string`)

```go
message := "Bonjour"
fmt.Println(len(message)) // longueur = 7
fmt.Println(message[0])   // premier octet : 66 ('B')
```

🔸 Les strings sont **immuables** (on ne peut pas modifier un caractère directement).
🔸 Utilise **les backticks** pour des chaînes multilignes :

```go
text := `Ligne 1
Ligne 2`
```

---

## 🧩 6. Tableaux et slices

```go
var arr [3]int = [3]int{1, 2, 3} // tableau fixe
slice := []int{1, 2, 3, 4}       // slice dynamique

slice = append(slice, 5) // on peut ajouter
fmt.Println(slice)       // [1 2 3 4 5]
```

💡 *Analogie :* un **slice** = une vue flexible sur un tableau sous-jacent (comme une “liste” en Python).

---

## 🗺️ 7. Maps (dictionnaires)

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
fmt.Println(ages["Alice"]) // 25

ages["Charlie"] = 22       // ajout
```

---

## 🧱 8. Structs (types personnalisés)

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name) // "Alice"
```

💬 *Analogie :* un `struct` = un “objet sans méthode” (comme une simple classe de données).

---

## ✅ À retenir

* Go est **fortement typé** → pas de conversions automatiques.
* Types de base : `int`, `float`, `bool`, `string`.
* Collections : `array`, `slice`, `map`.
* Structures : `struct` pour regrouper des champs.
* Utilise `:=` pour des déclarations rapides.


