# GraphQL App — React (TypeScript) + Go (gqlgen + BoltDB)

* **Frontend :** React + TypeScript + Apollo Client
* **Backend :** Go + gqlgen (GraphQL server)
* **Base de données :** BoltDB (base embarquée, sans serveur)

---

## 🧱 Architecture générale

<!-- TODO -->
```
client/            → Application React (TypeScript)
server/            → Serveur Go GraphQL
  ├── graph/       → Schéma, resolvers et modèles gqlgen
  ├── db/          → Gestion de BoltDB (ouverture, buckets, utils)
  ├── middlewares/ → Gestion des middlewares (jwt)
  ├── utils/       → Gestion de ressources utiles (jwt, password)
  ├── go.mod       → Déclaration du module
  └── server.go    → Entrée principale du serveur
```

Le serveur GraphQL expose un **unique endpoint** :

```
http://localhost:6000/query
```

et un **GraphQL Playground** est disponible sur :

```
http://localhost:6000/
```

---

## 🚀 Stack technique

### Frontend

* **React + Vite + TypeScript**
* **Apollo Client** pour la communication GraphQL
* **TailwindCSS** (optionnel, pour le style rapide)
* **React Router** pour la navigation

### Backend

* **Go 1.25.0**
* **gqlgen** pour la génération automatique du serveur GraphQL
* **bbolt** (fork officiel de BoltDB) pour la base de données embarqué
* **uuid** pour les identifiants uniques
* **net/http** pour l’exposition du serveur
* **chi** pour le routage avec authentification jwt

---

## ⚙️ Installation et lancement

### 1. Cloner le projet

```bash
git clone https://github.com/Clele63/graphql-go.git
cd graphql-go
```

### 2. Lancer le conteneur monolithe

```bash
docker compose up --build -d
```

\> Le serveur écoute sur `http://localhost:6000` ou les ports suivants.

\> Le client écoute sur `http://localhost:3000` ou les ports suivants.

#### Structure minimale du backend

<!-- TODO -->
```
server/
│
├── graph/
│   ├── schema.graphqls     # Définition du schéma GraphQL
│   ├── model/              # Types générés
│   ├── resolver.go         # Injection des dépendances
│   ├── mutation_resolver.go
│   └── query_resolver.go
│
├── db/
│   └── db.go               # Initialisation de BoltDB
│
└── server.go                 # Serveur principal
```

#### Structure minimale du frontend

<!-- TODO -->
```
frontend/
│
├── src/
│   ├── apollo/
│   │   └── client.ts       # Configuration Apollo Client
│   ├── components/
│   │   └── UserList.tsx    # Exemple de requête GraphQL
│   ├── pages/
│   │   ├── Home.tsx
│   │   └── CreateUser.tsx
│   ├── App.tsx
│   └── main.tsx
│
├── index.html
└── package.json
```

---

<!-- TODO -->
<!-- ## Exemple de schéma GraphQL

```graphql
# backend/graph/schema.graphqls

type User {
  id: ID!
  name: String!
}

input NewUser {
  name: String!
}

type Query {
  users: [User!]!
}

type Mutation {
  createUser(input: NewUser!): User!
}
```

---

## Exemple de stockage avec BoltDB

Chaque utilisateur est stocké dans un bucket `users` sous forme JSON.

```go
// backend/db/db.go
package db

import (
  "encoding/json"
  "go.etcd.io/bbolt"
  "log"
)

func InitDB(path string) *bbolt.DB {
  db, err := bbolt.Open(path, 0666, nil)
  if err != nil {
    log.Fatal(err)
  }

  db.Update(func(tx *bbolt.Tx) error {
    _, err := tx.CreateBucketIfNotExists([]byte("users"))
    return err
  })

  return db
}
```

---

## 🔀 Exemple de resolver

```go
// backend/graph/query_resolver.go
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
  var users []*model.User
  err := r.DB.View(func(tx *bbolt.Tx) error {
    b := tx.Bucket([]byte("users"))
    return b.ForEach(func(_, v []byte) error {
      var u model.User
      if err := json.Unmarshal(v, &u); err != nil {
        return err
      }
      users = append(users, &u)
      return nil
    })
  })
  return users, err
}
```

---

## 🔄 Exemple de mutation

```go
// backend/graph/mutation_resolver.go
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
  user := &model.User{
    ID:   uuid.NewString(),
    Name: input.Name,
  }

  err := r.DB.Update(func(tx *bbolt.Tx) error {
    b := tx.Bucket([]byte("users"))
    data, _ := json.Marshal(user)
    return b.Put([]byte(user.ID), data)
  })
  if err != nil {
    return nil, err
  }
  return user, nil
}
```

---

## 🌐 Exemple d’appel côté frontend

```tsx
// frontend/src/components/UserList.tsx
import { gql, useQuery } from "@apollo/client";

const GET_USERS = gql`
  query GetUsers {
    users {
      id
      name
    }
  }
`;

export default function UserList() {
  const { data, loading } = useQuery(GET_USERS);

  if (loading) return <p>Loading...</p>;

  return (
    <ul>
      {data.users.map((u: any) => (
        <li key={u.id}>{u.name}</li>
      ))}
    </ul>
  );
}
``` -->

<!-- --- -->
