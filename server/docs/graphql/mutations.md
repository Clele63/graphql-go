# Mutations

### About mutations



### addComment



#### Input fields

- taskId ([ID!](scalars.md#id))

- content ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| author ([User!](objects.md#user)) |  |
| content ([String!](scalars.md#string)) |  |
| createdAt ([Date!](scalars.md#date)) |  |
| id ([ID!](scalars.md#id)) |  |
| task ([Task!](objects.md#task)) |  |

---

### createTask



#### Input fields

- input ([CreateTaskInput!](input_objects.md#createtaskinput))
 

#### Returns

| Name | Description |
|------|-------------|
| assignees ([[User!]!](objects.md#user)) |  |
| column ([Column!](objects.md#column)) |  |
| comments ([[Comment!]!](objects.md#comment)) |  |
| createdAt ([Date!](scalars.md#date)) |  |
| description ([String](scalars.md#string)) |  |
| id ([ID!](scalars.md#id)) |  |
| status ([String!](scalars.md#string)) |  |
| title ([String!](scalars.md#string)) |  |

---

### createUser



#### Input fields

- input ([CreateUserInput!](input_objects.md#createuserinput))
 

#### Returns

| Name | Description |
|------|-------------|
| avatar ([String!](scalars.md#string)) |  |
| createdAt ([Date!](scalars.md#date)) |  |
| email ([String!](scalars.md#string)) |  |
| id ([ID!](scalars.md#id)) |  |
| name ([String!](scalars.md#string)) |  |

---

### deleteComment



#### Input fields

- id ([ID!](scalars.md#id))
 

---

### deleteTask



#### Input fields

- id ([ID!](scalars.md#id))
 

---

### deleteUser



#### Input fields

- id ([ID!](scalars.md#id))
 

---

### login



#### Input fields

- email ([String!](scalars.md#string))

- password ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| expired_at ([Date!](scalars.md#date)) |  |
| token ([String!](scalars.md#string)) |  |

---

### moveTask



#### Input fields

- id ([ID!](scalars.md#id))

- toColumnId ([ID!](scalars.md#id))
 

#### Returns

| Name | Description |
|------|-------------|
| assignees ([[User!]!](objects.md#user)) |  |
| column ([Column!](objects.md#column)) |  |
| comments ([[Comment!]!](objects.md#comment)) |  |
| createdAt ([Date!](scalars.md#date)) |  |
| description ([String](scalars.md#string)) |  |
| id ([ID!](scalars.md#id)) |  |
| status ([String!](scalars.md#string)) |  |
| title ([String!](scalars.md#string)) |  |

---

### updateTask



#### Input fields

- input ([UpdateTaskInput!](input_objects.md#updatetaskinput))
 

#### Returns

| Name | Description |
|------|-------------|
| assignees ([[User!]!](objects.md#user)) |  |
| column ([Column!](objects.md#column)) |  |
| comments ([[Comment!]!](objects.md#comment)) |  |
| createdAt ([Date!](scalars.md#date)) |  |
| description ([String](scalars.md#string)) |  |
| id ([ID!](scalars.md#id)) |  |
| status ([String!](scalars.md#string)) |  |
| title ([String!](scalars.md#string)) |  |

---

### updateUser



#### Input fields

- input ([UpdateUserInput!](input_objects.md#updateuserinput))
 

#### Returns

| Name | Description |
|------|-------------|
| avatar ([String!](scalars.md#string)) |  |
| createdAt ([Date!](scalars.md#date)) |  |
| email ([String!](scalars.md#string)) |  |
| id ([ID!](scalars.md#id)) |  |
| name ([String!](scalars.md#string)) |  |

---