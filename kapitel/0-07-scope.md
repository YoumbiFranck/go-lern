## 🧠 1. Qu’est-ce que le *scope* (la portée)

Le **scope** définit **où une variable est accessible** dans ton code.
En Go, la portée est **déterminée par les accolades `{}`** et le **package**.

---

## 🔹 2. Les niveaux de portée

### 🟢 a) **Portée locale (bloc ou fonction)**

Les variables déclarées **dans une fonction** ne sont accessibles **que dans cette fonction**.

```go
func demo() {
    x := 10
    fmt.Println(x) // OK
}
fmt.Println(x) // ❌ Erreur : x n’existe pas ici
```

💡 *Règle simple :* une variable vit **jusqu’à la fin du bloc `{}`** où elle est définie.

---

### 🟡 b) **Portée de package**

Une variable ou fonction déclarée **en dehors de toute fonction** est accessible :

* Dans tout le **fichier**,
* Et dans les **autres fichiers du même package**.

```go
package main

var version = "1.0.0" // visible partout dans le package

func main() {
    fmt.Println(version) // ✅
}
```

---

### 🟠 c) **Portée globale (exportée)**

Si le nom d’une fonction, variable ou struct **commence par une majuscule**,
elle devient **exportée** → accessible depuis **d’autres packages**.

```go
package tools

var Version = "1.0.0" // Exportée
var name = "outil"    // Non exportée
```

Dans un autre fichier :

```go
package main

import "monprojet/tools"

func main() {
    fmt.Println(tools.Version) // ✅
    // fmt.Println(tools.name) // ❌ invisible
}
```

💡 *Convention :*

* **Majuscule = public**
* **Minuscule = privé**

---

## ⚙️ 3. Créer plusieurs fichiers Go

Chaque fichier Go doit :

1. Commencer par le **même package** (pour faire partie du même module)
2. Être dans le **même dossier**

### Exemple :

```
monprojet/
 ├── main.go
 ├── utils.go
```

#### `main.go`

```go
package main

import "fmt"

func main() {
    
}
```

#### `utils.go`

```go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Salut", name)
}
```

# Code

```bash
 go run main.go utils.go
```

➡️ Ici, `main.go` et `utils.go` partagent le même **package main**,
donc `main()` peut appeler `greet()` sans import.

---

## 🧩 4. Organisation par packages

Quand ton projet grandit, tu peux **structurer ton code** avec des **packages**.

### Exemple :

```
monprojet/
 ├── main.go
 └── tools/
     └── mathutils.go
```

#### `tools/mathutils.go`

```go

```

#### `main.go`

```go
package main

import (
    "fmt"
    "monprojet/tools"
)

func main() {
    result := tools.Add(3, 4)
    fmt.Println(result) // 7
}
```

💡 Les packages servent à **organiser ton code** comme des modules.
Le nom du dossier = nom du package (par convention).

---

## 🧩 5. Initialisation automatique

Tu peux créer une fonction spéciale `init()` dans un fichier :
elle sera **appelée automatiquement** avant `main()`.

```go
func init() {
    fmt.Println("Initialisation du package")
}
```

---

## ✅ À retenir

| Concept                | Description                                                      |
| ---------------------- | ---------------------------------------------------------------- |
| **Scope local**        | variable visible uniquement dans la fonction                     |
| **Scope de package**   | variable/fonction visible dans tous les fichiers du même package |
| **Exportation**        | nom en majuscule = visible depuis d’autres packages              |
| **init()**             | exécutée avant `main()` automatiquement                          |
| **Plusieurs fichiers** | même dossier + même `package` = partagent le code                |


