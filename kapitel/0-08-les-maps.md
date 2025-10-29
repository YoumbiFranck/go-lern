## 🧭 1. Qu’est-ce qu’une map ?

Une **map** est une collection **clé → valeur**.
Chaque clé doit être **unique**, et le type des clés et des valeurs est **fixe**.

> 💬 *Analogie :* imagine un répertoire téléphonique —
> le **nom** est la clé, et le **numéro** est la valeur.

---

## 🔹 2. Déclaration d’une map

### Méthode 1 : déclaration vide avec `make()`

```go
ages := make(map[string]int)
```

* `string` → type des clés
* `int` → type des valeurs

### Méthode 2 : initialisation directe

```go
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
```

---

## 🔸 3. Ajout et modification d’éléments

```go
ages := make(map[string]int)
ages["Alice"] = 25
ages["Bob"] = 30
ages["Bob"] = 31 // modification
```

💡 Si la clé existe déjà → la valeur est **mise à jour**.

---

## 🔍 4. Lecture d’une valeur

```go
fmt.Println(ages["Alice"]) // 25
fmt.Println(ages["Eve"])   // 0 (clé absente)
```

⚠️ Si la clé n’existe pas, Go retourne la **valeur zéro du type** (ici `0` pour `int`).

---

## 🧠 5. Vérifier si une clé existe

Go fournit une syntaxe spéciale pour tester la présence d’une clé :

```go
age, exists := ages["Eve"]

if exists {
    fmt.Println("Âge trouvé :", age)
} else {
    fmt.Println("Clé inconnue")
}
```

💡 `exists` est un booléen (`true` ou `false`).

---

## 🧮 6. Supprimer un élément

```go
delete(ages, "Bob")
```

* Si la clé n’existe pas, **aucune erreur** n’est levée.

---

## 🔁 7. Parcourir une map avec `range`

```go
for name, age := range ages {
    fmt.Println(name, "→", age)
}
```

L’ordre de parcours est **aléatoire** (non garanti) → Go ne trie pas les maps.

---

## ⚙️ 8. Taille d’une map

```go
fmt.Println(len(ages)) // nombre d’éléments
```

---

## 🧩 9. Types autorisés comme clés

Les **types de clé** doivent être **comparables** (car Go utilise une table de hachage).
✅ Valides : `string`, `int`, `bool`, `float`, `struct` comparable
❌ Invalides : `slice`, `map`, `fonction`

---

## 🧱 10. Exemple complet

```go
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
```

Sortie possible :

```
Alice
Utilisateur : Bob
admin => Alice
user => Bob
Taille : 2
```

---

## ✅ À retenir

| Concept                 | Description                                 |
| ----------------------- | ------------------------------------------- |
| `map[keyType]valueType` | structure clé-valeur                        |
| `make(map[string]int)`  | crée une map vide                           |
| `delete(map, key)`      | supprime une entrée                         |
| `len(map)`              | donne la taille                             |
| `range`                 | permet de parcourir                         |
| Ordre aléatoire         | pas garanti                                 |
| Majuscule = exportable  | utile si tu déclares la map dans un package |


