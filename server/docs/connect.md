# Documentation : Package `graph/connect`

Ce package est responsable de l'initialisation de la connexion à la base de données MySQL et de l'exécution du schéma SQL.

## `connect/dbConnect.go`

Contient la logique de connexion principale.

* **Variables Globales**
    * `var DB *sql.DB` : Pool de connexion à la base de données.
    * `var Queries *wrapper.WrappedQueries` : Instance du querier `sqlc` (enveloppé) accessible globalement.

* **`InitDB()`**
    Fonction centrale qui orchestre la connexion :
    1.  **Configuration** : Lit les variables d'environnement (`DB_USERNAME`, `DB_PASSWORD`, `DB_HOST`, etc.).
    2.  **Connexion en deux temps** :
        * Ouvre une première connexion *sans* sélectionner de base de données (`dsnNoDB`) pour exécuter la commande `CREATE DATABASE IF NOT EXISTS ...`.
        * Ferme cette connexion temporaire.
    3.  **Connexion Principale** : Ouvre la connexion "réelle" à la base de données spécifiée (`dsn`).
    4.  **Pool de Connexions** : Configure le pool (`SetConnMaxLifetime`, `SetMaxOpenConns`, etc.) pour la robustesse.
    5.  **Exécution Schéma** : Appelle `ExecSchema(DB)` pour créer les tables.
    6.  **Initialisation `sqlc`** :
        * Crée le `Querier` de base : `baseQueries := generated.New(DB)`.
        * L'enveloppe dans le wrapper (pour le logging) : `Queries = queries.InitWrappedQueries(baseQueries)`.

* **`GetQueries() *wrapper.WrappedQueries`**
    Un "getter" public qui permet au reste de l'application (principalement le `resolver.Resolver`) d'accéder à l'objet `Queries` initialisé.

## `connect/execSql.go`

Gère l'exécution des fichiers SQL de configuration.

* **`ExecSchema(DB *sql.DB)`**
    Appelle `loadSchema()` et exécute le SQL retourné sur la BDD.
* **`ExecMock(DB *sql.DB)`**
    Appelle `loadMock()` et exécute le SQL retourné (pour les données de test).

## `connect/loadSql.go` (Tag: `!embed`)

Version de développement pour charger le SQL.

* **`loadSchema()` et `loadMock()`**
    Lisent les fichiers `schema.sql` et `mock.sql` directement depuis le disque (`os.ReadFile`).

## `connect/loadSqlEmbed.go` (Tag: `embed`)

Version de production pour charger le SQL.

* **`//go:embed data/schema.sql`**
    Utilise la directive `embed` de Go pour inclure le contenu des fichiers SQL (`schema.sql` et `mock.sql`) directement dans le binaire compilé. C'est cette version qui est utilisée lors du build Docker (`-tags embed`).