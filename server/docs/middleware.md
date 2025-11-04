# Documentation : Package `middlewares`

Ce package contient les middlewares HTTP utilisés par le routeur `chi` (défini dans `server.go`).

## `middlewares/jwt.go`

Fournit un middleware pour l'authentification JWT et l'injection de contexte.

* **`struct UserAuth`**
    Une structure simple pour stocker les informations d'authentification (`UserID`, `IPAddress`, `Token`) qui seront injectées dans le contexte de la requête.

* **`var userCtxKey = &contextKey{name: "user"}`**
    Définit une clé de contexte privée pour éviter les collisions lors de l'injection/extraction de `UserAuth` dans le contexte.

* **`JwtMiddleware() func(http.Handler) http.Handler`**
    La fonction de middleware principale. Pour chaque requête entrante :
    1.  Appelle `tokenFromHTTPRequestgo` pour extraire la chaîne du token de l'en-tête `Authorization: Bearer ...`.
    2.  Appelle `userIDFromHTTPRequestgo` pour valider le token (en utilisant `utils.DecodeJwt`) et en extraire le `UserID`.
    3.  Si le token est invalide ou absent, le `UserID` est défini sur la chaîne `"0"`.
    4.  Crée un objet `UserAuth` avec l'ID trouvé et l'IP de la requête.
    5.  Injecte cet objet `UserAuth` dans le contexte de la requête (`r.WithContext`).

* **`tokenFromHTTPRequestgo(r *http.Request) string`**
    Fonction d'aide pour parser l'en-tête `Authorization` et extraire le token.

* **`userIDFromHTTPRequestgo(tokenString string) string`**
    Fonction d'aide qui utilise `utils.DecodeJwt` pour valider le token et retourner le `UserID` ou `"0"` en cas d'échec.

* **`GetAuthFromContext(ctx context.Context) *UserAuth`**
    Une fonction d'aide (utilisée par les resolvers GraphQL) pour récupérer l'objet `UserAuth` à partir du contexte.