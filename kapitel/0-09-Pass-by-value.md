En **Go, tout est *pass by value*** : quand tu appelles une fonction, **chaque argument est copié**.
👉 Mais ce qui est copié peut être soit la **donnée elle-même** (ex : `int`, `array`), soit un **descripteur** qui **pointe** vers des données (ex : `slice`, `map`, `string`, `chan`, `interface`, `func`).
C’est là que naissent les confusions.

---

# 1) Règle d’or

> **Toujours une copie**, jamais de passage “par référence” implicite.
> La question est : **que copie-t-on ?**

* **Types valeur** (ex. `int`, `bool`, `struct`, `array`) → on copie **tous les bits** de la valeur.
* **Types “référence”** (ex. `slice`, `map`, `string`, `chan`, `func`, `interface`) → on copie **un petit en-tête** (pointeurs + métadonnées) qui **référence** des données ailleurs en mémoire.

Conséquence :

* Copier un **array** ou un **struct** = **deux objets indépendants**.
* Copier un **slice** ou une **map** = **deux en-têtes distincts**, mais **la même mémoire sous-jacente** (donc les modifications portent sur la même “vraie” donnée).

---

# 2) Par type : ce qui se passe vraiment

## a) Scalaires (int, float, bool)

Copie simple. La fonction ne peut pas modifier la variable d’origine.

```go
func inc(x int) { x++ }

n := 5
inc(n)
fmt.Println(n) // 5 (inchangé)
```

## b) Arrays (taille fixe)

**Copie complète** du tableau.

```go
func touch(a [3]int) { a[0] = 99 }

arr := [3]int{1,2,3}
touch(arr)
fmt.Println(arr) // [1 2 3] (indépendant)
```

➡️ Pour modifier un array en place : passer un **pointeur** `*[3]int`.

## c) Structs

Même logique que les arrays : **copie champ par champ**.

```go
type User struct{ Name string; Age int }

func birthday(u User) { u.Age++ }

alice := User{"Alice", 30}
birthday(alice)
fmt.Println(alice.Age) // 30 (inchangé)
```

➡️ Modifier l’original : prendre `*User`.

## d) Slices

Un slice = **en-tête (ptr, len, cap)**. On **copie l’en-tête**, **pas** les éléments.
Deux slices peuvent donc **voir les mêmes éléments**.

```go
func setFirst(s []int) { s[0] = 42 }

nums := []int{1,2,3}
setFirst(nums)
fmt.Println(nums) // [42 2 3] (modifié)
```

⚠️ `append` peut **réallouer** : si la capacité est dépassée, un **nouveau tableau** est alloué → les modifications suivantes peuvent **ne plus** toucher l’ancien slice.

```go
func push(s []int) {
    s = append(s, 9) // peut réallouer
    s[0] = 100
}
nums := []int{1,2,3} // cap=3
push(nums)
fmt.Println(nums) // [1 2 3] (souvent inchangé ici)
```

➡️ Si tu veux garantir la modification de la **composition** (taille) du slice vu par l’appelant, passe un **pointeur vers slice** : `*[]int` (rare) ou retourne le nouveau slice.

## e) Maps

Une map est un **pointeur interne** vers une table de hachage.
Copier une map = **nouvel en-tête** pointant vers **la même table**.

```go
func put(m map[string]int) { m["x"] = 1 }

m := map[string]int{}
put(m)
fmt.Println(m["x"]) // 1 (modifié)
```

➡️ Pas besoin de pointeur pour modifier le contenu.
(En revanche, **affecter une nouvelle map** à l’argument ne changera pas la variable de l’appelant.)

## f) Strings

Un string = (ptr, len) vers des **octets immuables**.
Copie = **nouvel en-tête** vers **les mêmes octets**. Impossible de modifier le contenu.

```go
func touch(s string) { /* rien à faire : immuable */ }
```

## g) Channels, functions, interfaces

Ce sont aussi des **descripteurs copiés**. La copie référence le **même canal / même valeur sous-jacente**.

---

# 3) Pointeurs : quand et pourquoi

Passer un **pointeur** (`*T`) signifie “je veux modifier l’original” ou éviter une **grosse copie**.

```go
func incPtr(x *int) { *x++ }

n := 5
incPtr(&n)
fmt.Println(n) // 6 (modifié)
```

Autre intérêt : **performance** (éviter la copie d’un gros struct/array).

---

# 4) Receveurs de méthode : valeur vs pointeur

* **Receveur valeur** `(u User)` : méthode travaille sur une **copie**.
* **Receveur pointeur** `(u *User)` : méthode travaille sur **l’original**.

```go
type Counter struct { N int }

func (c Counter) AddV()   { c.N++ }   // copie
func (c *Counter) AddP()  { c.N++ }   // original

c := Counter{N: 0}
c.AddV()
fmt.Println(c.N) // 0
c.AddP()
fmt.Println(c.N) // 1
```

Règle pratique :

* Si la méthode **modifie** le receveur ou si le type est **lourd**, utilise le **pointeur**.

---

# 5) Petits pièges fréquents (et parades)

1. **Je modifie un slice mais l’appelant ne voit rien après append**

   * Parce que `append` a réalloué → retourne toujours le slice :

   ```go
   func push(s []int, v int) []int { return append(s, v) }
   s = push(s, 9)
   ```

2. **Je pensais qu’une map copiée serait indépendante**

   * Non : même table sous-jacente. Pour cloner :

   ```go
   clone := make(map[string]int, len(m))
   for k, v := range m { clone[k] = v }
   ```

3. **Je passe un gros struct et c’est lent**

   * Prends `*MonGrosStruct`.

4. **Je veux modifier un array**

   * Passe `*[N]T` ou, mieux, utilise un **slice**.

---

# 6) Cheatsheet (mémo)

| Type              | Copie passée à la fonction    | Modifs visibles par l’appelant ? |
| ----------------- | ----------------------------- | -------------------------------- |
| `int/float/bool`  | valeur entière                | ❌                                |
| `array`           | tous les éléments             | ❌                                |
| `struct`          | tous les champs               | ❌                                |
| `*T` (pointeur)   | adresse                       | ✅ via `*ptr`                     |
| `slice`           | en-tête (ptr,len,cap)         | ✅ sur éléments; `append` : ⚠️    |
| `map`             | en-tête vers la table         | ✅ sur les entrées                |
| `string`          | en-tête vers octets immuables | ❌ (contenu immuable)             |
| `chan/func/iface` | descripteur                   | ✅ (même ressource)               |

---

# 7) Exemples rapides

### Modifier un struct proprement

```go
type User struct{ Name string; Age int }
func Birthday(u *User) { u.Age++ }

u := User{"Bob", 29}
Birthday(&u)
fmt.Println(u.Age) // 30
```

### API slice sûre vis-à-vis d’append

```go
func AppendSafe(s []int, vals ...int) []int {
    return append(s, vals...)
}

s := []int{1,2,3}
s = AppendSafe(s, 4,5)
```

### Clonage de slice (copie des éléments)

```go
src := []int{1,2,3}
dst := append([]int(nil), src...) // ou: make+copy
```

---

# 8) Quand choisir valeur vs pointeur (règle pratique)

* **Passe par valeur** si :

  * petit type (ex. petits structs), lecture seule, plus simple et sûr.
* **Passe par pointeur** si :

  * tu dois **modifier** l’original,
  * le type est **lourd** à copier,
  * tu veux partager l’état (en conscience des risques de concurrence).


