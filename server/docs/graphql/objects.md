# Objects

### About objects

[Objects](https://graphql.github.io/graphql-spec/June2018/#sec-Objects) in GraphQL represent the resources you can access. An object can contain a list of fields, which are specifically typed.

### Board

 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>columns</strong> (<a href="objects.md#column">[Column!]!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### Column

 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>order</strong> (<a href="scalars.md#int">Int!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>tasks</strong> (<a href="objects.md#taskconnection">TaskConnection!</a>)</td> 
    <td>
      <p></p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
        <tr>
          <td>
            <p>first (<a href="scalars.md#int">Int</a>)</p>
            <p></p>
          </td>
        </tr>
        <tr>
          <td>
            <p>after (<a href="scalars.md#string">String</a>)</p>
            <p></p>
          </td>
        </tr>
      </table>
    </td>
  </tr>
</table>

---

### Comment

 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>author</strong> (<a href="objects.md#user">User!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>content</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#date">Date!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>task</strong> (<a href="objects.md#task">Task!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### PageInfo

  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>endCursor</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>hasNextPage</strong> (<a href="scalars.md#boolean">Boolean!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### Task

 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>assignees</strong> (<a href="objects.md#user">[User!]!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>column</strong> (<a href="objects.md#column">Column!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>comments</strong> (<a href="objects.md#comment">[Comment!]!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#date">Date!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>description</strong> (<a href="scalars.md#string">String</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>status</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>title</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### TaskConnection

  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>edges</strong> (<a href="objects.md#taskedge">[TaskEdge!]!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>pageInfo</strong> (<a href="objects.md#pageinfo">PageInfo!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### TaskEdge

  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>cursor</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>node</strong> (<a href="objects.md#task">Task!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### Token

  

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>expired_at</strong> (<a href="scalars.md#date">Date!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>token</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
</table>

---

### User

 

#### Implements


- [Node](interfaces.md#node) 

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
  <tr>
    <td><strong>avatar</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>createdAt</strong> (<a href="scalars.md#date">Date!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>email</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>id</strong> (<a href="scalars.md#id">ID!</a>)</td> 
    <td></td>
  </tr>
  <tr>
    <td><strong>name</strong> (<a href="scalars.md#string">String!</a>)</td> 
    <td></td>
  </tr>
</table>

---