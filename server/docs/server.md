# Documentation : Point d'entrée (server.go)

Ce fichier assemble tous les composants du backend Go.

## `func main()`

La fonction `main` est le point d'entrée de l'application et orchestre les étapes suivantes :

1.  **Configuration** : Récupère le port (`PORT`) et l'environnement (`APP_ENV`) depuis les variables d'environnement.
2.  **Connexion BDD** : Appelle **`connect.InitDB()`** pour initialiser la connexion à la base de données MySQL et préparer le `Querier` `sqlc`.
3.  **Initialisation GraphQL** : Crée le `resolver.Resolver` de GraphQL en lui injectant le `Querier` (`Queries: connect.GetQueries()`).
4.  **Routage (Chi)** : Crée un nouveau routeur `chi`.
5.  **Middlewares (Chi)** :
    * Applique le **`middlewares.JwtMiddleware()`** à toutes les routes pour gérer l'authentification.
    * Configure le middleware `cors` pour autoriser les requêtes cross-domain.
6.  **Configuration GraphQL (gqlgen)** :
    * Configure le serveur `gqlgen` (`handler.New(exec.NewExecutableSchema(...))`) en liant le schéma aux resolvers.
    * Ajoute les "transports" (WebSocket, GET, POST) et la configuration du cache (LRU).
7.  **Routes** :
    * `/playground` : Sert l'interface web GraphQL Playground.
    * `/query` : Sert le point d'entrée principal de l'API GraphQL.
8.  **Service de Fichiers Statiques (Frontend)** :
    * Si `APP_ENV` est `"prod"`, le serveur `chi` est configuré pour servir l'application frontend (probablement React/Node) buildée, qui se trouve dans `client/build`.
    * Il gère également la redirection "fallback" vers `index.html` pour les routeurs côté client (Single Page Application).
9.  **Démarrage** : Lance le serveur HTTP (`http.ListenAndServe`) sur le port configuré.