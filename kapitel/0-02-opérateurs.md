## ⚙️ 1. Introduction

Un **opérateur** sert à effectuer des calculs ou des comparaisons sur des valeurs.
Go propose plusieurs **familles d’opérateurs** : arithmétiques, de comparaison, logiques, d’affectation, et sur les bits.

---

## ➕ 2. Opérateurs arithmétiques

| Opérateur | Signification  | Exemple  | Résultat                            |
| --------- | -------------- | -------- | ----------------------------------- |
| `+`       | addition       | `10 + 3` | `13`                                |
| `-`       | soustraction   | `10 - 3` | `7`                                 |
| `*`       | multiplication | `10 * 3` | `30`                                |
| `/`       | division       | `10 / 3` | `3` (car division entière si `int`) |
| `%`       | reste (modulo) | `10 % 3` | `1`                                 |

```go
a, b := 10, 3
fmt.Println(a + b) // 13
fmt.Println(a / b) // 3  (division entière)
fmt.Println(float64(a) / float64(b)) // 3.333...
```

💡 *Attention :* la division entre entiers donne un entier.

---

## ⚖️ 3. Opérateurs de comparaison

Renvoient toujours un **booléen** (`true` ou `false`).

| Opérateur | Signification       | Exemple  |
| --------- | ------------------- | -------- |
| `==`      | égal à              | `a == b` |
| `!=`      | différent de        | `a != b` |
| `>`       | supérieur à         | `a > b`  |
| `<`       | inférieur à         | `a < b`  |
| `>=`      | supérieur ou égal à | `a >= b` |
| `<=`      | inférieur ou égal à | `a <= b` |

```go
x, y := 5, 8
fmt.Println(x == y) // false
fmt.Println(x < y)  // true
```

---

## 🔐 4. Opérateurs logiques

Utilisés avec des booléens (`true` / `false`).
(||) = OU logique

| Opérateur | Signification | Exemple         | Résultat   |       |   |        |        |
| --------- | ------------- | --------------- | ---------- | ----- | - | ------ | ------ |
| `&&`      | ET logique    | `true && false` | `false`    |       |   |        |        |
| `         |               | `               | OU logique | `true |   | false` | `true` |
| `!`       | NON logique   | `!true`         | `false`    |       |   |        |        |

```go
isAdmin := true
isLogged := false

if isAdmin && isLogged {
    fmt.Println("Accès autorisé")
} else {
    fmt.Println("Accès refusé")
}
```

💡 *Astuce :* Go évalue de gauche à droite et **court-circuite** :
si `false && ...`, la suite n’est pas évaluée.

---

## 🧮 5. Opérateurs d’affectation

| Opérateur | Signification                | Exemple  |
| --------- | ---------------------------- | -------- |
| `=`       | affectation simple           | `a = 10` |
| `+=`      | addition + affectation       | `a += 5` |
| `-=`      | soustraction + affectation   | `a -= 2` |
| `*=`      | multiplication + affectation | `a *= 3` |
| `/=`      | division + affectation       | `a /= 2` |
| `%=`      | modulo + affectation         | `a %= 2` |

```go
a := 10
a += 5 // équivaut à a = a + 5
fmt.Println(a) // 15
```

---

## ⚡ 6. Opérateurs sur les bits (bitwise)

Utilisés pour manipuler les bits d’entiers.

| Opérateur | Signification     | Exemple      | Résultat binaire       |    |                          |
| --------- | ----------------- | ------------ | ---------------------- | -- | ------------------------ |
| `&`       | ET bit à bit      | `a & b`      | 1 si les deux bits = 1 |    |                          |
| `         | `                 | OU bit à bit | `a                     | b` | 1 si au moins un bit = 1 |
| `^`       | XOR (OU exclusif) | `a ^ b`      | 1 si bits différents   |    |                          |
| `&^`      | AND NOT           | `a &^ b`     | 1 si bit de `b` = 0    |    |                          |
| `<<`      | décalage à gauche | `a << 1`     | multiplie par 2        |    |                          |
| `>>`      | décalage à droite | `a >> 1`     | divise par 2           |    |                          |

```go
a, b := 6, 3 // 6 = 110, 3 = 011
fmt.Println(a & b)  // 2 (010)
fmt.Println(a | b)  // 7 (111)
fmt.Println(a ^ b)  // 5 (101)
fmt.Println(a << 1) // 12 (1100)
```

---

## 🔍 7. Opérateurs divers

| Opérateur | Signification   | Exemple |
| --------- | --------------- | ------- |
| `++`      | incrémente de 1 | `i++`   |
| `--`      | décrémente de 1 | `i--`   |

⚠️ Ces opérateurs **ne sont pas des expressions** (contrairement à C/Java).
Tu **ne peux pas** faire :

```go
fmt.Println(i++) // ❌ interdit
```

Tu dois écrire :

```go
i++
fmt.Println(i)
```

---

## ✅ À retenir

* Go supporte les opérateurs classiques (`+`, `-`, `==`, `&&`, etc.)
* Division entière et float sont distinctes.
* `++` et `--` ne s’utilisent que seuls, pas dans des expressions.
* Opérateurs bit à bit = manipulation bas-niveau, utile pour flags, permissions, etc.


