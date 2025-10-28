## 🔁 1. Le concept général

En Go, il n’y a **qu’un seul mot-clé pour les boucles** : `for`.

👉 Pas de `while`, pas de `do...while` — **tout passe par `for`**.

---

## 🔹 2. La boucle `for` classique

Structure :

```go
for initialisation; condition; incrément {
    // instructions
}
```

Exemple :

```go
for i := 0; i < 5; i++ {
    fmt.Println("Compteur :", i)
}
```

🧠 *Déroulement :*

1. Initialise `i`
2. Vérifie la condition (`i < 5`)
3. Exécute le bloc
4. Incrémente `i`
5. Retour à l’étape 2

---

## 🔸 3. Boucle type "while"

Tu peux omettre certaines parties du `for` :

```go
i := 0
for i < 5 { // tant que i < 5
    fmt.Println(i)
    i++
}
```

💡 *C’est l’équivalent d’un `while`* dans d’autres langages.

---

## 🔸 4. Boucle infinie

Si tu omets tout, tu obtiens une boucle infinie :

```go
for {
    fmt.Println("Boucle sans fin")
}
```

⚠️ À utiliser avec un `break` ou `return` pour sortir :

```go
for {
    if time.Now().Second()%10 == 0 {
        fmt.Println("Sortie de la boucle")
        break
    }
}
```

---

## 🔸 5. Sortie et saut de boucle

| Mot-clé    | Rôle                         |
| ---------- | ---------------------------- |
| `break`    | interrompt la boucle         |
| `continue` | saute à l’itération suivante |

Exemple :

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        continue // ignore 5
    }
    if i == 8 {
        break // sort de la boucle
    }
    fmt.Println(i)
}
```

---

## 🔹 6. Boucle sur une collection

### 🧮 Sur un slice (ou tableau) :

```go
numbers := []int{10, 20, 30}

for i, value := range numbers {
    fmt.Println("Index :", i, "Valeur :", value)
}
```

* `i` → index
* `value` → valeur à cet index
* Tu peux ignorer l’un des deux avec `_` :

  ```go
  for _, value := range numbers {
      fmt.Println(value)
  }
  ```

### 🗺️ Sur une map :

```go
users := map[string]int{"Alice": 25, "Bob": 30}

for name, age := range users {
    fmt.Println(name, "→", age)
}
```

---

## 🔸 7. Boucle sur une chaîne (`string`)

```go
for i, r := range "GoLang" {
    fmt.Printf("%d: %c\n", i, r)
}
```

💡 `r` est une **rune** (un caractère Unicode).
Go gère bien les caractères multi-octets.

---

## 🔹 8. Boucles imbriquées

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 2; j++ {
        fmt.Println("i:", i, "j:", j)
    }
}
```

---

## ⚠️ 9. Labels (cas rare mais utile)

Tu peux nommer une boucle pour faire un `break` spécifique :

```go
outer:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i*j > 3 {
            break outer // sort de la boucle externe
        }
        fmt.Println(i, j)
    }
}
```

---

## ✅ À retenir

* `for` est **la seule boucle** en Go.
* Trois formes :

  * `for i := 0; i < n; i++` → classique
  * `for condition` → comme un `while`
  * `for {}` → boucle infinie
* `range` simplifie le parcours de slices, maps, strings, etc.
* `break` et `continue` contrôlent le flux.

