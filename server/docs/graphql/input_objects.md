# Input objects

### About input objects

[Input objects](https://graphql.github.io/graphql-spec/June2018/#sec-Input-Objects) can be described as "composable objects" because they include a set of input fields that define the object.

### CreateUserInput




#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>creation_date</strong> (<a href="scalars.md#date">Date!</a>)</td>
    <td></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td></td>
  </tr>
  <tr>
    <td><strong>password</strong> (<a href="scalars.md#string">String!</a>)</td>
    <td></td>
  </tr>
</table>

---

### UpdateUserInput




#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String</a>)</td>
    <td></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td>
    <td></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String</a>)</td>
    <td></td>
  </tr>
</table>

---