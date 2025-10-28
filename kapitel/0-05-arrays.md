## 🧱 1. Qu’est-ce qu’un array en Go ?

Un **array** est une **collection d’éléments du même type**, **de taille fixe**.
Chaque élément est accessible par un **index** (commençant à 0).

> 📘 *Analogie :* imagine un casier avec un nombre fixe de cases.
> Une fois le nombre de cases défini, tu ne peux plus le changer.

---

## 🔹 2. Déclaration d’un array

### Syntaxe générale :

```go
var nom [taille]Type
```

Exemples :

```go
var nombres [3]int              // [0 0 0]
var noms [2]string              // ["", ""]
nombres = [3]int{10, 20, 30}    // affectation directe
```

Tu peux aussi **laisser Go déduire la taille** :

```go
nombres := [...]int{10, 20, 30} // "..." = taille auto (3 ici)
```

---

## 🔸 3. Accès et modification

```go
var fruits [3]string
fruits[0] = "Pomme"
fruits[1] = "Banane"
fruits[2] = "Orange"

fmt.Println(fruits[1]) // "Banane"
```

💡 *Indice de départ = 0*, donc :

* `fruits[0]` → 1er élément
* `fruits[2]` → dernier élément (si taille = 3)

---

## 🔍 4. Parcourir un array

### Méthode classique :

```go
for i := 0; i < len(fruits); i++ {
    fmt.Println(fruits[i])
}
```

### Avec `range` :

```go
for i, fruit := range fruits {
    fmt.Println(i, fruit)
}
```

Ignorer l’index :

```go
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

---

## ⚙️ 5. Taille fixe = limitation

```go
var a [3]int = [3]int{1, 2, 3}
// a = [4]int{1,2,3,4} ❌ Erreur : tailles différentes
```

Go considère `[3]int` et `[4]int` comme **deux types différents**.
👉 Pour plus de flexibilité, on utilise les **slices** (prochain chapitre).

---

## 🧮 6. Comparaison d’arrays

Les tableaux sont **comparables** si :

* Ils ont la **même taille**
* Leurs éléments sont **comparables** entre eux

```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{1, 2, 4}

fmt.Println(a == b) // true
fmt.Println(a == c) // false
```

---

## 📦 7. Copie par valeur

Quand tu assignes un tableau à un autre, **Go copie le contenu** (pas une référence).

```go
a := [3]int{1, 2, 3}
b := a // copie indépendante

b[0] = 100
fmt.Println(a) // [1 2 3]
fmt.Println(b) // [100 2 3]
```

💡 Cela veut dire que modifier `b` **ne change pas `a`**.

---

## 🧠 8. Array multidimensionnel

Tu peux créer des tableaux à plusieurs dimensions (comme des matrices) :

```go
matrix := [2][3]int{
    {1, 2, 3},
    {4, 5, 6},
}

fmt.Println(matrix[1][2]) // 6
```

---

## ✅ À retenir

| Point clé         | Détail                     |
| ----------------- | -------------------------- |
| Taille fixe       | définie à la déclaration   |
| Type strict       | `[3]int` ≠ `[4]int`        |
| Copie par valeur  | pas de référence implicite |
| Parcours          | via `for` ou `range`       |
| Multidimensionnel | possible (`[2][3]int`)     |


