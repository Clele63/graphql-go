# Documentation : Package `utils`

Ce package fournit des fonctions utilitaires de bas niveau, réutilisables dans toute l'application.

## `utils/password.go`

Gère le hachage et la vérification des mots de passe en utilisant `bcrypt`.

* **`HashPassword(password string) (string, error)`**
    Prend un mot de passe en clair et le hache en utilisant `bcrypt.GenerateFromPassword` avec un coût de 10.
* **`ComparePassword(password string, hash string) bool`**
    Compare un mot de passe en clair avec un hachage `bcrypt` existant.
    *Note : L'implémentation actuelle (`return err != nil`) est incorrecte. Elle devrait renvoyer `err == nil` pour indiquer que la comparaison a réussi.*

## `utils/jwt.go`

Gère la création et le décodage des JSON Web Tokens (JWT).

* **`var issuer = []byte("github/thealmarques")`**
    Définit une clé secrète (issuer) codée en dur pour signer et valider les tokens.
* **`DecodeJwt(token string) (*jwt.Token, error)`**
    Prend une chaîne de token JWT, la parse et la valide en utilisant la clé `issuer`. Il s'attend à ce que le token contienne des `model.UserClaims`.
* **`GenerateJwt(userID string, expiredAt int64) string`**
    Crée un nouveau token JWT signé (HS256). Il inclut l'ID de l'utilisateur (`UserID`) et la date d'expiration (`ExpiresAt`) dans les "claims" du token.