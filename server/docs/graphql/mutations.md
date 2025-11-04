# Mutations

### About mutations



### createUser



#### Input fields

- input ([CreateUserInput!](input_objects.md#createuserinput))
 

#### Returns

| Name | Description |
|------|-------------|
| creationDate ([Date!](scalars.md#date)) |  |
| email ([String!](scalars.md#string)) |  |
| id ([ID!](scalars.md#id)) |  |
| name ([String!](scalars.md#string)) |  |

---

### deleteUser



#### Input fields

- id ([ID!](scalars.md#id))
 

---

### login



#### Input fields

- name ([String!](scalars.md#string))

- password ([String!](scalars.md#string))
 

#### Returns

| Name | Description |
|------|-------------|
| expired_at ([Date!](scalars.md#date)) |  |
| token ([String!](scalars.md#string)) |  |

---

### updateUser



#### Input fields

- input ([UpdateUserInput!](input_objects.md#updateuserinput))
 

#### Returns

| Name | Description |
|------|-------------|
| creationDate ([Date!](scalars.md#date)) |  |
| email ([String!](scalars.md#string)) |  |
| id ([ID!](scalars.md#id)) |  |
| name ([String!](scalars.md#string)) |  |

---