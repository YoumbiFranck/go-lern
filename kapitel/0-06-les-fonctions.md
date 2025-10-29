## ⚙️ 1. Qu’est-ce qu’une fonction ?

Une **fonction** est un bloc de code **nommé**, qui peut :

* recevoir des **paramètres** en entrée,
* exécuter du code,
* et (optionnellement) **retourner une valeur**.

> 💬 *Analogie :* une fonction est comme une machine : tu lui donnes des entrées → elle te rend un résultat.

---

## 🔹 2. Déclaration de base

```go
func nom(paramètres) typeRetour {
    // instructions
    return valeur
}
```

Exemple simple :

```go
func sayHello() {
    fmt.Println("Bonjour !")
}
```

Appel :

```go
sayHello()
```

---

## 🔸 3. Fonction avec paramètres

```go
func greet(name string) {
    fmt.Println("Salut", name)
}

func main() {

greet("Alice") // → Salut Alice


} 
```

💡 *Rappel :* le type vient **après** le nom du paramètre (contrairement à d’autres langages).

---

## 🔸 4. Plusieurs paramètres

```go
func add(a int, b int) {
    fmt.Println(a + b)
}
```

Tu peux simplifier si les types sont identiques :

```go
func add(a, b int) {
    fmt.Println(a + b)
}
```

---

## 🔹 5. Fonction avec valeur de retour

```go
func add(a, b int) int {
    return a + b
}

result := add(3, 5)
fmt.Println(result) // 8
```

💡 La valeur après les parenthèses indique **le type de retour**.

---

## 🔸 6. Retour multiple

Une fonction peut retourner **plusieurs valeurs** (spécificité de Go !).

```go
func divide(a, b int) (int, int) {
    quotient := a / b
    reste := a % b
    return quotient, reste
}

q, r := divide(10, 3)
fmt.Println("Quotient:", q, "Reste:", r)
```

Tu peux **ignorer** une valeur avec `_` :

```go
q, _ := divide(10, 3)
```

---

## 🔹 7. Noms de retours (optionnel)

Tu peux **nommer les valeurs de retour** pour rendre le code plus lisible.

```go
func rectangle(a, b int) (area int, perimeter int) {
    area = a * b
    perimeter = 2 * (a + b)
    return // retourne area, perimeter automatiquement
}
```

---

## 🔸 8. Fonctions anonymes (lambdas)

Tu peux définir une fonction **sans nom** :

```go
func() {
    fmt.Println("Fonction anonyme")
}()
```

Ou la stocker dans une variable :

```go
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(2, 3)) // 5
```

---

## 🔹 9. Passage par valeur

Tous les **paramètres sont passés par valeur** (copiés).
Modifier un argument **n’affecte pas la variable d’origine**.

```go
func increment(x int) {
    x++
}

n := 5
increment(n)
fmt.Println(n) // 5 (pas 6)
```

💡 Pour modifier la valeur d’origine, il faut passer **un pointeur** (on verra plus tard).

---

## 🧠 10. Portée des variables

Une variable définie dans une fonction **n’existe que dans cette fonction**.

```go
func show() {
    x := 10
    fmt.Println(x)
}
// fmt.Println(x) ❌ → erreur, x est local
```

---

## ✅ À retenir

| Concept            | Description                              |
| ------------------ | ---------------------------------------- |
| `func`             | mot-clé pour déclarer une fonction       |
| Paramètres         | nom + type, séparés par des virgules     |
| Retour             | après les parenthèses                    |
| Plusieurs retours  | `(val1, val2 type)`                      |
| Passage par valeur | les arguments sont copiés                |
| Fonction anonyme   | peut être stockée ou appelée directement |


