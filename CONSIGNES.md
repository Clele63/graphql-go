# Évaluation GraphQL - Kanban Board

## 🎯 Objectifs de l'évaluation

L'objectif de cette évaluation est de mettre en pratique les compétences acquises durant le cours GraphQL en créant une application Kanban complète avec :

- Configuration d'un serveur GraphQL avec Apollo Server
- Définition d'un schéma GraphQL complet
- Implémentation de resolvers (queries, mutations, subscriptions)
- Configuration d'un client Apollo avec React
- Exploitation des données GraphQL côté frontend

## 📚 Barème

- **Partie 1 - Configuration du serveur (30 points)**
  - Configuration Apollo Server : 10 points
  - Schéma GraphQL complet : 20 points

- **Partie 2 - Resolvers (40 points)**
  - Queries : 10 points
  - Mutations : 15 points
  - Subscriptions : 15 points

- **Partie 3 - Client Apollo (30 points)**
  - Configuration Apollo Client : 10 points
  - Queries et Mutations : 10 points
  - Subscriptions temps réel : 10 points

## 🚀 Partie 1 - Configuration du serveur (30 points)

> **Note** : Deux implémentations sont disponibles selon votre préférence :
> - Node.js avec Apollo Server → dossier `server-apollo/`
> - Python avec Graphene → dossier `server-python/`
> Si vous voulez utiliser une autre technologie, veuillez le préciser. La seule exigence est que le serveur supporte les fonctionnalités demandées (queries, mutations, subscriptions, authentification, etc.). Et assurez-vous de fournir des instructions claires pour l'installation et le démarrage du serveur dans ce cas.

### 1.1 Configuration du serveur (10 points)

**Apollo Server (Node.js)** - 📁 Fichier : `server-apollo/src/index.js`

**TODO 1.1** : Configurer Apollo Server v4 avec :
- Support des subscriptions (graphql-ws)
- Context d'authentification (JWT)
- Serveur HTTP Express

```javascript
// Indice : Utiliser ApolloServer, expressMiddleware, createServer
// Penser à configurer le WebSocket pour les subscriptions
```

**Graphene (Python)** - 📁 Fichier : `server-python/app.py`

**TODO 1.1** : Configurer Flask avec Graphene :
- Support des subscriptions WebSocket
- Context d'authentification (JWT)
- Endpoint GraphQL avec Flask

```python
# Indice : Utiliser Flask, GraphQLView, flask-cors
# Penser à configurer le middleware d'authentification
```

### 1.2 Schéma GraphQL (20 points)

**Apollo Server** - 📁 Fichier : `server-apollo/src/schema.js`
**Graphene** - 📁 Fichier : `server-python/schema.py`

**TODO 1.2** : Définir le schéma GraphQL complet avec les types suivants :

#### Types principaux (10 points)
- `User` : id, name, email, avatar
- `Board` : id, name, columns
- `Column` : id, name, tasks (avec pagination)
- `Task` : id, title, description, status, assignees (array), comments, createdAt
- `Comment` : id, content, authorId, author, createdAt

#### Types de pagination (5 points)
- `TaskEdge` : node, cursor
- `PageInfo` : hasNextPage, endCursor
- `TaskConnection` : edges, pageInfo

#### Operations (5 points)
- **Query** : me, users, board
- **Mutation** : login, createTask, updateTask, moveTask, deleteTask, addComment, deleteComment
- **Subscription** : taskCreated, taskUpdated, taskMoved, commentAdded

```graphql
# Indice : Utiliser le pattern Connection pour la pagination
# Les assignees doivent être un array [User!]
# Les subscriptions doivent retourner les objets modifiés
```

## 🔧 Partie 2 - Resolvers (40 points)

**Apollo Server** - 📁 Fichier : `server-apollo/src/resolvers.js`
**Graphene** - 📁 Fichier : `server-python/schema.py` (classes Query, Mutation, Subscription)

### 2.1 Queries (10 points)

**TODO 2.1** : Implémenter les resolvers de queries :

- `me` : Retourner l'utilisateur connecté à partir du context
- `users` : Retourner la liste de tous les utilisateurs
- `board` : Retourner le board avec ses colonnes

**TODO 2.2** : Implémenter le resolver de pagination :

- `Column.tasks` : Implémenter la pagination cursor-based
  - Paramètres : `first` (limite), `after` (cursor)
  - Retourner un `TaskConnection` avec edges et pageInfo

```javascript
// Indice : Utiliser .slice() pour paginer
// Calculer hasNextPage et endCursor pour PageInfo
```

### 2.2 Mutations (15 points)

**TODO 2.3** : Implémenter les mutations de tâches :

- `createTask` : Créer une nouvelle tâche avec assigneeIds
- `updateTask` : Modifier une tâche existante
- `moveTask` : Déplacer une tâche vers une autre colonne
- `deleteTask` : Supprimer une tâche

**TODO 2.4** : Implémenter les mutations de commentaires :

- `addComment` : Ajouter un commentaire à une tâche
- `deleteComment` : Supprimer un commentaire (vérifier l'auteur)

```javascript
// Indice : Utiliser nanoid() pour générer les IDs
// Penser à publier les événements pour les subscriptions
```

### 2.3 Subscriptions (15 points)

**TODO 2.5** : Implémenter les subscriptions avec PubSub :

- `taskCreated` : Notifier la création d'une tâche
- `taskUpdated` : Notifier la modification d'une tâche
- `taskMoved` : Notifier le déplacement d'une tâche
- `commentAdded` : Notifier l'ajout d'un commentaire

```javascript
// Indice : Utiliser pubsub.publish() dans les mutations
// Utiliser pubsub.asyncIterator() dans les subscriptions
// Exemple : TOPICS.TASK_CREATED(columnId)
```

## 💻 Partie 3 - Client Apollo (30 points)

### 3.1 Configuration Apollo Client (10 points)

📁 Fichier : `client/src/apollo.js`

**TODO 3.1** : Configurer Apollo Client avec :

- HttpLink pour les queries/mutations
- WebSocketLink pour les subscriptions
- Split link pour router selon le type d'opération
- AuthLink pour ajouter le token JWT
- Cache InMemoryCache

```javascript
// Indice : Utiliser split() pour séparer HTTP et WebSocket
// getMainDefinition() permet de détecter les subscriptions
// Passer le token dans connectionParams pour WebSocket
```

### 3.2 Queries et Mutations (10 points)

📁 Fichiers : `client/src/pages/Board.jsx`, `client/src/components/Column.jsx`, `client/src/components/TaskModal.jsx`

**TODO 3.2** : Implémenter les opérations GraphQL :

#### Board.jsx
- Query `BOARD` : Récupérer le board avec colonnes et utilisateurs
- Mutation `CREATE_TASK` : Créer une tâche avec assigneeIds

#### Column.jsx
- Query `TASKS_FOR_COLUMN` : Récupérer les tâches paginées
- Mutation `MOVE_TASK` : Déplacer une tâche

#### TaskModal.jsx
- Mutation `UPDATE_TASK` : Modifier une tâche
- Mutation `DELETE_TASK` : Supprimer une tâche

```javascript
// Indice : Utiliser gql pour définir les queries
// useQuery pour les requêtes, useMutation pour les modifications
// Penser à refetchQueries après les mutations
```

### 3.3 Subscriptions temps réel (10 points)

📁 Fichiers : `client/src/pages/Board.jsx`, `client/src/components/CommentSection.jsx`

**TODO 3.3** : Implémenter les subscriptions :

#### Board.jsx
- `TASK_CREATED` : Écouter les nouvelles tâches
- `TASK_UPDATED` : Écouter les modifications
- `TASK_MOVED` : Écouter les déplacements

#### CommentSection.jsx
- `COMMENT_ADDED` : Écouter les nouveaux commentaires

```javascript
// Indice : Utiliser useSubscription hook
// Mettre à jour le state local dans onData
// Afficher un toast de notification
```

## 📝 Conseils

1. **Suivez l'ordre des parties** : Configuration → Schéma → Resolvers → Client
2. **Testez régulièrement** : Utilisez Apollo Sandbox pour tester vos queries
3. **Gérez les erreurs** : Ajoutez des try/catch et messages d'erreur clairs
4. **Authentification** : Le token JWT est géré par le store Zustand (`useAuth`)
5. **Pagination** : Implémentez le pattern Cursor-based comme vu en cours
6. **Subscriptions** : N'oubliez pas de publier les événements dans les mutations

## 🛠️ Démarrage

### Serveur

**Option 1 : Apollo Server (Node.js)**
```bash
cd server-apollo
npm install
npm start
```

**Option 2 : Graphene (Python)**
```bash
cd server-python
python3 -m venv venv
source venv/bin/activate  # Sur Windows: venv\Scripts\activate
pip install -r requirements.txt
python app.py
```

Le serveur démarre sur `http://localhost:4000`

### Client
```bash
cd client
npm install
npm run dev
```

Le client démarre sur `http://localhost:5173`

## 📦 Livrables

- Code source complet avec tous les TODO implémentés
- Application fonctionnelle (serveur + client)
- Pas besoin de documentation supplémentaire => Prenez soin de commenter votre code là où c'est nécessaire

## ⏱️ Durée

**3 heures**

Bon courage ! 🚀
