# Projet GraphQL Go (gqlgen + sqlc + MySQL)

Ce projet est un serveur backend en Go qui expose une API GraphQL. Il utilise **gqlgen** pour la génération du serveur GraphQL et **sqlc** pour générer un code Go typesafe à partir de requêtes SQL brutes, le tout connecté à une base de données **MySQL**.

## 📚 Documentation

Une documentation technique détaillée du code backend et de l'API est disponible dans le dossier [`/server/docs/README.md`](./server/docs/README.md).

## 🚀 Stack Technique

* **Backend** : Go
* **API** : GraphQL via `gqlgen`
* **Base de Données** : MySQL
* **Accès BDD** : `sqlc` (pour la génération de code à partir de SQL)
* **Routage** : `chi`
* **Authentification** : JWT (tokens) via un middleware personnalisé
* **Déploiement** : Docker & Docker Compose

## ✨ Fonctionnalités

* **Gestion des Utilisateurs (CRUD)** : Créer, mettre à jour et supprimer des utilisateurs.
* **Authentification** : Une mutation `login` qui vérifie un utilisateur et retourne un JSON Web Token (JWT).
* **Recherche** : Rechercher des utilisateurs par nom (protégé par authentification).
* **Listing** : Lister tous les utilisateurs (protégé par authentification).

## 🏗️ Architecture du Projet

Voici une arborescence de projet corrigée et complétée, basée sur l'ensemble des fichiers fournis :

```
./
├── .devcontainer/        # Configuration pour le développement distant
│   ├── .dockerignore
│   ├── Dockerfile
│   └── devcontainer.json
├── .vscode/              # Paramètres de l'éditeur VSCode
│   ├── launch.json
│   └── settings.json
├── client/               # Application Frontend (inféré du Dockerfile)
│   ├── build/            # Output du build, servi par server.go
│   └── ...               # Autres fichiers: package.json, src/, ...
├── server/               # Application Backend Go
│   ├── docs/                 # Documentation technique
│   │   ├── server.md
│   │   ├── utils.md
│   │   ├── middleware.md
│   │   ├── connect.md
│   │   ├── queries.md
│   │   └── README.md         # Table des matières de la documentation
│   ├── graph/
│   │   ├── connect/      # Logique de connexion BDD et chargement schema
│   │   │   ├── data/
│   │   │   │   ├── mock.sql
│   │   │   │   └── schema.sql
│   │   │   ├── dbConnect.go
│   │   │   ├── execSql.go
│   │   │   └── ...
│   │   ├── exec/         # Code auto-généré par gqlgen
│   │   ├── model/        # Modèles Go (générés et manuels)
│   │   ├── resolver/
│   │   │   ├── scalar/
│   │   │   │   └── date.go
│   │   │   ├── mutations.resolvers.go
│   │   │   ├── query.resolvers.go
│   │   │   ├── resolver.go
│   │   │   └── ...
│   │   └── schema/       # Schémas GraphQL (.graphqls)
│   ├── middlewares/      # Middlewares HTTP (ex: JWT)
│   │   └── jwt.go
│   ├── queries/
│   │   ├── generated/    # Code Go typesafe généré par sqlc
│   │   ├── sql/          # Requêtes SQL sources pour sqlc
│   │   └── wrapper/      # Wrapper pour les requêtes (logging, etc.)
│   ├── tools/
│   │   └── tools.go
│   ├── utils/            # Fonctions utilitaires (JWT, mot de passe)
│   │   ├── jwt.go
│   │   └── password.go
│   ├── go.mod
│   ├── go.sum
│   ├── gqlgen.yml
│   ├── server.go         # Point d'entrée du serveur (routes chi)
│   └── sqlc.yml
├── .gitignore
├── Dockerfile            # Dockerfile multi-stage (build client + build serveur)
├── docker-compose.yml    # Orchestre l'application et la BDD MySQL
└── README.md             # Ce fichier
```

## ⚙️ Installation et Lancement

Ce projet est entièrement conteneurisé.

**Prérequis** : Docker et Docker Compose.

1.  Clonez le dépôt.
2.  À la racine du projet, exécutez :

    ```bash
    docker compose up --build
    ```

3.  Le `docker-compose.yml` va :
    * Construire et démarrer le service `graphql-go-database` (MySQL).
    * Construire l'image `graphql-go-app` en utilisant le `Dockerfile` (qui build le client Node.js *puis* le serveur Go).
    * Démarrer le service `graphql-go-app` (le serveur Go).
    * Le serveur Go attendra que la base de données soit saine avant de démarrer.

## 🛠️ Utilisation

* **GraphQL Playground** : `http://localhost:6060/playground`
* **Endpoint API** : `http://localhost:6060/query`
* **Client Frontend** : `http://localhost:6060/` (servi par Go en mode `prod`)

### Exemples de Requêtes

#### 1. Créer un utilisateur

```graphql
mutation CreateUser {
  createUser(
    input: {
      name: "TestUser"
      password: "password123"
      email: "test@example.com"
      creation_date: "2024-01-01T00:00:00Z" # Format RFC3339
    }
  ) {
    id
    name
    email
  }
}
```

```graphql
mutation Login {
  login(name: "TestUser", password: "password123") {
    token
  }
}
```

```json
{
  "Authorization": "Bearer VOTRE_TOKEN_JWT_ICI"
}
```

```graphql
query GetUsers {
  users {
    id
    name
    creationDate
  }
}
```