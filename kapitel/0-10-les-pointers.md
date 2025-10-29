# 🧭 L’idée générale

Un **pointeur** est une **adresse mémoire** vers une valeur d’un certain type.

* Obtenir l’adresse : **`&x`**
* Déréférencer (accéder à la valeur pointée) : **`*p`**
* Type d’un pointeur vers `T` : **`*T`**
* Valeur zéro d’un pointeur : **`nil`** (pointe vers rien)

> Analogie : un pointeur est **l’adresse** d’un appartement ; `*p` te fait entrer **dans** l’appartement.

---

# ⚙️ 1) Base syntaxique

```go
x := 10
p := &x        // p a le type *int, contient l’adresse de x
fmt.Println(*p) // 10

*p = 42        // modifie x via le pointeur
fmt.Println(x) // 42
```

* Comparaison : deux pointeurs peuvent être comparés (`p == q`) → compare les **adresses**.

---

# 🧱 2) Pointeur et `struct`

Très courant pour **modifier** un objet ou éviter la **copie**.

```go
type User struct {
    Name string
    Age  int
}

func Birthday(u *User) { u.Age++ } // modifie l'original

u := User{Name: "Alice", Age: 29}
Birthday(&u)
fmt.Println(u.Age) // 30
```

Accès pratique aux champs : Go autorise `p.Name` au lieu de `(*p).Name`.

---

# 🆚 3) `new` vs `make` vs littéraux adressés

* `new(T)` → alloue `T`, initialise à zéro, renvoie `*T`.
* `make(...)` → **uniquement** pour `slice`, `map`, `chan` (retourne une **valeur** initialisée, pas un `*T`).
* `&T{...}` → idiome courant pour struct :

```go
u := new(User)        // *User avec champs zéro
v := &User{Name: "Bob"} // *User initialisé

s := make([]int, 0, 8)  // slice vide, pas un pointeur
```

---

# 🧠 4) Pointeurs et **pass-by-value**

Go est **pass-by-value** : un pointeur **lui-même** est **copié** quand on le passe à une fonction,
mais **l’adresse** qu’il contient référence la **même** donnée — donc on peut la modifier.

```go
func Set100(p *int) { *p = 100 }

n := 1
Set100(&n)
fmt.Println(n) // 100
```

---

# 🧩 5) Pointeurs avec **array**, **slice**, **map**, **string**

* **Array `[N]T`** : lourd à copier ; si tu veux modifier l’original, passe `*[N]T`.
* **Slice `[]T`** : déjà un **descripteur** (ptr/len/cap).

  * Modifier un **élément** via le paramètre → visible par l’appelant.
  * **`append`** peut réallouer → retourne toujours le slice mis à jour.
* **Map** : descripteur vers une table ; modifier les **entrées** est visible (pas besoin de `*map[...]...`).
* **String** : immuable ; un pointeur vers string est rarement utile.

```go
func TouchFirst(s []int) { s[0] = 9 } // visible
func AppendOne(s []int) []int { return append(s, 1) }

nums := []int{0,2,3}
TouchFirst(nums)          // nums -> [9 2 3]
nums = AppendOne(nums)    // accepte la nouvelle vue
```

---

# 🧯 6) Nil, panics et vérifications

* Déréférencer un **pointeur nil** → **panic**.
* Vérifie avant usage si nécessaire :

```go
var p *User
if p == nil {
    p = &User{Name: "Init"}
}
```

---

# 🚫 7) Ce que Go **n’a pas** avec les pointeurs

* **Pas d’arithmétique** des pointeurs (`p++` interdit).
* **Impossible de prendre l’adresse** d’un **élément de map** : `&m["k"]` ❌ (les entrées peuvent bouger).

  * Contourner : lire dans une variable, modifier, puis réaffecter.

```go
v := m["k"]
v++
m["k"] = v
```

---

# 🧵 8) Méthodes : receveur **valeur** vs **pointeur**

* `(t T)` : méthode travaille sur une **copie**.
* `(t *T)` : méthode travaille sur **l’original** (et évite la copie).

```go
type Counter struct{ N int }

func (c Counter) AddV()  { c.N++ }  // ne modifie pas l'original
func (c *Counter) AddP() { c.N++ }  // modifie

c := Counter{}
c.AddV()
fmt.Println(c.N) // 0
c.AddP()
fmt.Println(c.N) // 1
```

**Règle pratique :** receveur pointeur si la méthode **modifie** l’état ou si le type est **lourd**.

---

# 🐛 9) Pièges fréquents

### a) Adresse de variable de boucle (`range`)

```go
users := []User{{"A",1},{"B",2}}
ptrs := []*User{}
for _, u := range users {
    ptrs = append(ptrs, &u) // ❌ même adresse à chaque tour (variable u réutilisée)
}
```

**Correct :**

```go
for i := range users {
    ptrs = append(ptrs, &users[i]) // ✅ adresse de l’élément réel
}
```

### b) Re-allocation silencieuse avec `append`

* Un pointeur vers un élément d’un slice peut devenir **dangling** si le slice réalloue.
  → Évite de stocker des pointeurs vers des éléments si tu continues à `append`.

### c) Interfaces et méthodes pointeur

* Une interface demandant une méthode avec receveur **pointeur** n’est pas satisfaite par `T` (il faut `*T`).
  (Rappel rapide : *method set* de `T` ≠ de `*T`.)

---

# ⚖️ 10) Performance et allocation (aperçu)

* Passer un **pointeur** peut éviter de **copier** de gros objets.
* Mais un pointeur peut forcer une **allocation heap** (escape analysis) → coût GC potentiel.
* Souvent, **les petits structs** passent **mieux par valeur** (stack-friendly) et sont plus simples.

> Règle pragmatique : **par valeur par défaut**, **pointeur** si tu **modifies** ou si c’est **gros**.

---

# 🧪 11) Exemples ciblés

### Pointeur vers array

```go
func Fill(a *[3]int) { a[0], a[1], a[2] = 7,8,9 }

arr := [3]int{1,2,3}
Fill(&arr)
fmt.Println(arr) // [7 8 9]
```

### Construction idiomatique d’un struct pointeur

```go
type Conf struct{ Port int }
cfg := &Conf{Port: 8080} // idiome
```

### Option “setter” avec pointeur

```go
type Server struct{ Port int }

func (s *Server) WithPort(p int) *Server { s.Port = p; return s }

srv := (&Server{}).WithPort(8080)
```

---

# 📝 Cheatsheet (mémo rapide)

* `&x` → adresse de `x` ; `*p` → valeur pointée.
* `nil` est la valeur zéro d’un pointeur ; déréférencer `nil` → panic.
* `new(T)` → `*T` zéro ; `&T{...}` → `*T` initialisé ; `make` ≠ pointeur.
* `slice/map/chan` : passer par valeur **suffit** pour modifier leur **contenu**.
* Pas d’arithmétique des pointeurs ; pas `&m["k"]`.
* Méthodes modifiantes → receveur `*T`.
* Valeur par défaut : **par valeur**, pointeur si **modification** ou **gros type**.

---

# 🎯 Mini-exercice (optionnel)

1. Que va afficher ce programme, et pourquoi ?

```go
func push(s []int) {
    s = append(s, 4)
    s[0] = 99
}
func main() {
    s := []int{1,2,3}
    push(s)
    fmt.Println(s)
}
```

2. Corrige-le pour que `s` devienne `[99 2 3 4]`.

3. Repère l’erreur :

```go
for _, u := range users {
    list = append(list, &u)
}
```

Explique le correctif.


