## ⚙️ 1. Introduction

Les **conditions** permettent d’exécuter du code **seulement si** une expression est vraie.
En Go, la syntaxe est **simple et sans parenthèses**.

---

## 🔹 2. Structure de base : `if`

```go
if condition {
    // bloc exécuté si condition == true
}
```

Exemple :

```go
age := 20
if age >= 18 {
    fmt.Println("Majeur")
}
```

💡 Pas besoin de `()` autour de la condition, mais **les accolades `{}` sont obligatoires**, même pour une seule ligne.

---

## 🔸 3. `if` ... `else`

```go
if age >= 18 {
    fmt.Println("Majeur")
} else {
    fmt.Println("Mineur")
}
```

---

## 🔸 4. `if` ... `else if` ... `else`

```go
score := 75

if score >= 90 {
    fmt.Println("Excellent")
} else if score >= 70 {
    fmt.Println("Bien")
} else {
    fmt.Println("Peut mieux faire")
}
```

🔹 Les conditions sont évaluées **dans l’ordre**.
Dès qu’une condition est vraie, les suivantes sont **ignorées**.

---

## 🧠 5. Déclaration dans un `if`

Go permet de **déclarer une variable directement dans la condition**.
Sa portée est limitée au bloc `if`.

```go
if n := len("Bonjour"); n > 5 {
    fmt.Println("Mot long")
} else {
    fmt.Println("Mot court")
}
// n n'existe plus ici ❌
```

💡 Pratique pour calculer une valeur à tester **sans polluer le reste du code**.

---

## ⚖️ 6. Conditions multiples avec opérateurs logiques

Tu peux combiner plusieurs conditions :

```go
age := 25
isMember := true

if age >= 18 && isMember {
    fmt.Println("Accès autorisé")
}
```

| Opérateur | Signification |   |            |
| --------- | ------------- | - | ---------- |
| `&&`      | ET logique    |   |            |
| `         |               | ` | OU logique |
| `!`       | NON logique   |   |            |

---

## 🔄 7. `switch` – alternative à plusieurs `if`

Go propose un `switch` **lisible et puissant**.

### Exemple simple :

```go
day := "mardi"

switch day {
case "lundi":
    fmt.Println("Début de semaine")
case "vendredi":
    fmt.Println("Presque le week-end")
default:
    fmt.Println("Jour ordinaire")
}
```

💡 Pas besoin d’écrire `break` : Go **sort automatiquement** du `case` après exécution.
(Si tu veux continuer, tu peux ajouter `fallthrough`, mais c’est rare.)

---

### Exemple avec conditions personnalisées :

```go
age := 20

switch {
case age < 12:
    fmt.Println("Enfant")
case age < 18:
    fmt.Println("Adolescent")
default:
    fmt.Println("Adulte")
}
```

🔸 Ici, le `switch` fonctionne comme une série de `if/else if`.

---

## ⚠️ 8. Particularités Go

* Pas de parenthèses dans les conditions (`if x > 5` ✅, `if (x > 5)` ❌).
* Toujours des accolades `{}` même pour une ligne.
* `switch` sans expression = suite de tests booléens.

---

## ✅ À retenir

* `if`, `else if`, `else` → logique conditionnelle classique.
* `switch` → plus clair pour plusieurs cas.
* Tu peux **déclarer une variable dans un if**.
* Pas de `break` dans les `case` (Go le fait pour toi).


