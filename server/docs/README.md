# Documentation Principale du Projet

Bienvenue sur la documentation technique du projet.

Cette documentation est divisée en deux sections principales :

1.  **Backend Go** : Documentation manuelle du code Go (logique métier, BDD, middlewares).
2.  **API GraphQL** : Documentation de l'API (généralement générée par un outil).

---

## 1. Documentation Backend (Go)

Cette section détaille le fonctionnement interne du serveur Go, hors logique GraphQL générée.

* **[Point d'entrée (`server.go`)](./server.md)**
    * *Démarrage, routage `chi`, initialisation des middlewares et de `gqlgen`.*

* **[Utilitaires (`utils/`)](./utils.md)**
    * *Gestion des mots de passe (`bcrypt`) et des tokens (`JWT`).*

* **[Middlewares (`middlewares/`)](./middleware.md)**
    * *Logique d'authentification et d'injection de contexte `JWT`.*

* **[Connexion BDD (`graph/connect/`)](./connect.md)**
    * *Initialisation de la connexion `MySQL`, pooling, et exécution du `schema.sql`.*

* **[Accès aux Données (`queries/` & `sqlc`)](./queries.md)**
    * *Configuration `sqlc.yml`, requêtes `SQL` sources, et logique du `wrapper` pour le logging et la transformation des données.*

---

## 2. Documentation API (GraphQL)

Cette section documente le schéma GraphQL, y compris les types, les requêtes, les mutations et les scalaires personnalisés.

* **[Documentation du Schéma GraphQL](./graphql/schema.md)**
    * *Documentation des schémas GraphQL*