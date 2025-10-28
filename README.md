# Environnement Go avec Docker pour WSL

Configuration Docker complète pour apprendre Go sous Windows avec WSL.

## 📋 Prérequis

- Docker Desktop installé et configuré pour WSL 2
- WSL 2 activé sur Windows
- Un terminal WSL (Ubuntu, Debian, etc.)

## 🚀 Démarrage rapide

### 1. Premier lancement (mode détaché)

```bash
# Construire et lancer le container en arrière-plan
docker-compose up -d --build
```

Cette commande va :
- Construire l'image Docker avec Go
- Lancer le container en arrière-plan (mode détaché)
- Garder le container actif pour exécuter vos commandes

### 2. Exécuter votre code Go

```bash
# Exécuter main.go dans le container actif
docker-compose exec golang-dev go run main.go

# Ou utiliser run (qui crée un nouveau container temporaire)
docker-compose run --rm golang-dev go run main.go
```

### 3. Arrêter le container

```bash
# Arrêter et supprimer le container
docker-compose down
```

## 💻 Commandes utiles

### Différence entre `exec` et `run`

**`docker-compose exec`** : Exécute une commande dans un container **déjà actif**
```bash
# D'abord, s'assurer que le container tourne
docker-compose up -d

# Puis exécuter des commandes dans ce container
docker-compose exec golang-dev go run main.go
docker-compose exec golang-dev go version
```

**`docker-compose run`** : Crée un **nouveau container temporaire**
```bash
# Pas besoin que le container soit déjà lancé
docker-compose run --rm golang-dev go run main.go
```

### Exécuter un script spécifique

```bash
# Avec exec (container déjà lancé)
docker-compose exec golang-dev go run votre-fichier.go

# Avec run (crée un nouveau container)
docker-compose run --rm golang-dev go run votre-fichier.go
```

### Mode interactif (bash dans le container)

```bash
# Ouvrir un shell dans le container (doit être lancé avec up -d d'abord)
docker-compose exec golang-dev sh

# Ou créer un nouveau container interactif
docker-compose run --rm -it golang-dev sh

# Une fois dans le container, vous pouvez :
go run main.go
go build main.go
./main
exit  # Pour quitter le shell
```

### Compiler un binaire

```bash
# Avec le container actif (exec)
docker-compose exec golang-dev go build -o app main.go

# Ou avec run
docker-compose run --rm golang-dev go build -o app main.go

# Le binaire "app" sera créé dans votre dossier local
```

### Exécuter les tests

```bash
# Créer un fichier de test (ex: main_test.go)
docker-compose exec golang-dev go test ./...
# Ou
docker-compose run --rm golang-dev go test ./...
```

### Formater le code

```bash
docker-compose exec golang-dev go fmt ./...
# Ou
docker-compose run --rm golang-dev go fmt ./...
```

### Ajouter des dépendances

```bash
# Par exemple, pour ajouter une librairie
docker-compose exec golang-dev go get github.com/gorilla/mux
# Ou
docker-compose run --rm golang-dev go get github.com/gorilla/mux

# Les dépendances seront ajoutées à go.mod
```

### Voir les logs du container

```bash
# Voir les logs en temps réel
docker-compose logs -f golang-dev

# Voir les derniers logs
docker-compose logs --tail=50 golang-dev
```

### Vérifier le statut du container

```bash
# Voir si le container tourne
docker-compose ps

# Voir les détails
docker ps
```

## 📁 Structure du projet

```
.
├── Dockerfile              # Configuration de l'image Docker
├── docker-compose.yml      # Configuration Docker Compose
├── go.mod                  # Fichier de module Go
├── main.go                 # Votre script Go principal
└── README.md              # Ce fichier
```

## 🎯 Créer de nouveaux scripts

Créez simplement un nouveau fichier `.go` dans le répertoire :

```bash
# Sous WSL
touch mon-script.go
```

Puis exécutez-le :

```bash
# Méthode 1 : Avec exec (container déjà lancé)
docker-compose up -d  # Si pas encore lancé
docker-compose exec golang-dev go run mon-script.go

# Méthode 2 : Avec run (plus simple, mais crée un nouveau container)
docker-compose run --rm golang-dev go run mon-script.go
```

## 🔄 Workflow recommandé

**Option A : Mode détaché (recommandé pour le développement actif)**
```bash
# 1. Lancer le container une fois
docker-compose up -d --build

# 2. Travailler sur votre code, puis exécuter
docker-compose exec golang-dev go run main.go

# 3. Mode interactif quand nécessaire
docker-compose exec golang-dev sh

# 4. Arrêter quand vous avez terminé
docker-compose down
```

**Option B : Mode ponctuel (pour des tests rapides)**
```bash
# Exécuter directement sans lancer le container
docker-compose run --rm golang-dev go run main.go
```

## 🔧 Personnalisation

### Changer le fichier Go par défaut

Éditez `docker-compose.yml` et modifiez la ligne `command`:

```yaml
command: go run votre-fichier.go
```

### Ajouter des variables d'environnement

Dans `docker-compose.yml`, ajoutez des variables sous `environment`:

```yaml
environment:
  - MA_VARIABLE=valeur
  - API_KEY=votre_clé
```

## 📚 Ressources pour apprendre Go

- [Tour de Go](https://go.dev/tour/) - Tutoriel interactif officiel
- [Go by Example](https://gobyexample.com/) - Exemples pratiques
- [Documentation officielle](https://go.dev/doc/)

## ❓ Problèmes courants

### Le container s'arrête immédiatement avec `up -d`

C'est normal ! Le container est configuré pour rester actif en arrière-plan. Vérifiez qu'il tourne :
```bash
docker-compose ps
```

Si le container n'apparaît pas, vérifiez les logs :
```bash
docker-compose logs golang-dev
```

### Le container ne démarre pas

Vérifiez que Docker Desktop est lancé :
```bash
docker --version
docker-compose --version
```

### Problèmes de permissions

Sous WSL, assurez-vous que vos fichiers ont les bonnes permissions :
```bash
chmod +x main.go
```

### Changements non pris en compte

Reconstruisez l'image :
```bash
docker-compose up --build
```

## 🎉 Bon apprentissage !

Modifiez `main.go` pour expérimenter avec Go. Vos changements sont automatiquement synchronisés avec le container grâce aux volumes Docker.
