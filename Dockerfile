# Image de base Go
FROM golang:1.23-alpine

# Installation des outils utiles pour le développement
RUN apk add --no-cache git

# Définir le répertoire de travail
WORKDIR /app

# Configuration pour éviter les problèmes de permissions avec WSL
RUN addgroup -g 1000 devuser && \
    adduser -D -u 1000 -G devuser devuser && \
    chown -R devuser:devuser /app

# Copier go.mod et go.sum si ils existent (pour la gestion des dépendances)
COPY go.* ./

# Télécharger les dépendances (cette couche sera cachée si go.mod ne change pas)
RUN go mod download || true

# Copier le reste du code
COPY . .

# Changer l'utilisateur pour éviter d'exécuter en root
USER devuser

# Commande par défaut
CMD ["go", "run", "main.go"]
