# Documentation : Package `queries` (Logique sqlc)

C'est le cœur de la couche d'accès aux données, géré par `sqlc`. Cette section couvre la configuration, les requêtes SQL sources et le wrapper manuel.

## `server/sqlc.yml` (Fichier de Configuration)

Définit comment `sqlc` doit générer le code Go.

* **`engine: "mysql"`** : Utilise le moteur MySQL.
* **`schema: "graph/connect/data/schema.sql"`** : Lit la structure de la BDD depuis ce fichier.
* **`queries: "queries/sql/"`** : Lit les requêtes SQL à partir de ce dossier.
* **`out: "queries/generated"`** : Écrit le code Go généré dans ce dossier.
* **`emit_interface: true`** : Génère l'interface `Querier` (cruciale pour le wrapping et les tests).
* **`overrides`** :
    * Indique à `sqlc` de ne pas utiliser `time.Time` pour la colonne `users.creation_date`.
    * À la place, il doit utiliser le type personnalisé `scalar.Date` (défini dans `graph/resolver/scalar/date.go`), assurant la compatibilité directe avec `gqlgen`.

## `queries/sql/` (Fichiers SQL sources)

Ce sont les requêtes SQL brutes que `sqlc` parse pour générer du code Go typesafe.

### `mutationsUsers.sql`
* **`-- name: CreateUser :exec`**
    Insère un nouvel utilisateur (nom, mot de passe haché, email, date).
* **`-- name: GetCreatedUser :one`**
    Récupère l'utilisateur qui vient d'être créé en utilisant `LAST_INSERT_ID()` (spécifique à MySQL).
* **`-- name: UpdateUser :exec`**
    Met à jour un utilisateur. Utilise `COALESCE` pour permettre les mises à jour partielles (les champs non fournis ne sont pas écrasés).
* **`-- name: GetUpdatedUser :one`**
    Récupère un utilisateur par son `id` (typiquement après une mise à jour).
* **`-- name: DeleteUser :exec`**
    Supprime un utilisateur par son `id`.

### `users.sql`
* **`-- name: ListUsers :many`**
    Récupère la liste de tous les utilisateurs (id, nom, email, date).
* **`-- name: GetUserByID :one`**
    Récupère un utilisateur spécifique par `id`.
* **`-- name: SearchUsersByName :many`**
    Recherche des utilisateurs dont le nom correspond (insensible à la casse) au terme de recherche.
* **`-- name: GetUserAuthByName :one`**
    Récupère un utilisateur par son nom, *y compris son mot de passe*, spécifiquement pour le processus de login/authentification.

## `queries/wrapper/` (Wrapper Manuel)

Ce répertoire contient du code manuel qui "enveloppe" (wrap) le code généré par `sqlc`. C'est une bonne pratique pour ajouter de la logique (comme le logging) et pour transformer les types de données.

* **`wrapper/Wrapper.go`**
    * **`struct WrappedQueries`** : Contient le `Querier` généré (`inner`) et un `logger`.
    * **`NewWrappedQueries(...)`** : Le constructeur qui initialise la structure.
* **`wrapper/UsersWrapper.go`**
    Implémente l'interface `Querier` générée par `sqlc`.
    * **Logique de "pass-through"** : Pour les mutations simples (`CreateUser`, `DeleteUser`, `UpdateUser`), il se contente d'appeler la fonction `inner` correspondante.
    * **Logique de "transformation"** : Pour les lectures (`ListUsers`, `SearchUsersByName`, `GetUserByID`, `GetUserAuthByName`), il fait plus :
        1.  Il ajoute du **logging** avant et après l'appel à la BDD.
        2.  Il **transforme les types de retour**. `sqlc` génère des types de retour spécifiques pour chaque requête (ex: `generated.ListUsersRow`, `generated.GetUserByIDRow`). Ce wrapper convertit ces types spécifiques en un type commun (`generated.User`), ce qui simplifie grandement le code dans les resolvers GraphQL.

## `queries/Queries.go`

* **`InitWrappedQueries(db generated.Querier) *wrapper.WrappedQueries`**
    Une simple fonction "factory" qui crée et renvoie une nouvelle instance de `wrapper.WrappedQueries`, en lui injectant le `Querier` généré par `sqlc` et un logger par défaut.