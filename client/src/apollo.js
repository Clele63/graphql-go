// TODO 3.1 - Configuration Apollo Client (10 points)
// Ce qu'il faut faire :
// 1. Créer HttpLink pour les queries/mutations (http://localhost:4000/graphql)
// 2. Créer AuthLink pour ajouter le token JWT dans les headers
// 3. Créer WebSocketLink pour les subscriptions (ws://localhost:4000/graphql)
// 4. Utiliser split() pour router selon le type d'opération
// 5. Créer ApolloClient avec le link et le cache InMemoryCache
// 6. Configurer typePolicies pour merger les edges de pagination

import { ApolloClient, InMemoryCache, createHttpLink, split, gql } from '@apollo/client'
import { setContext } from '@apollo/client/link/context'
import { GraphQLWsLink } from '@apollo/client/link/subscriptions'
import { createClient } from 'graphql-ws'
import { getMainDefinition } from '@apollo/client/utilities'
import useAuth from './store'

export const COMMENT_FRAGMENT = gql`
  fragment CommentFragment on Comment {
    id
    content
    createdAt
    author {
      id
      name
    }
  }
`

export const TASK_FRAGMENT = gql`
  fragment TaskFragment on Task {
    id
    title
    description
    column {
      id
    }
    assignees {
      id
      name
    }
    comments {
      ...CommentFragment
    }
  }
  ${COMMENT_FRAGMENT}
`

const httpLink = createHttpLink({
  uri: 'http://localhost:6060/query',
})

const authLink = setContext((_, { headers }) => {
  const token = useAuth.getState().token
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
})

const wsLink = new GraphQLWsLink(createClient({
  url: 'ws://localhost:6060/query',
  connectionParams: () => {
    const token = useAuth.getState().token
    return {
      Authorization: token ? `Bearer ${token}` : '',
    }
  }
}))

const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query)
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    )
  },
  wsLink,
  authLink.concat(httpLink),
)

export const client = new ApolloClient({
  link: splitLink,
  cache: new InMemoryCache({
    typePolicies: {
      Column: {
        fields: {
          tasks: {
            keyArgs: false,
            merge(existing = { edges: [] }, incoming, { args }) {
              if (!args || !args.after) {
                return incoming;
              }
              const merged = [...existing.edges, ...incoming.edges];
              const seen = new Set();
              const deduped = [];
              for (const edge of merged) {
                if (!seen.has(edge.cursor)) {
                  seen.add(edge.cursor);
                  deduped.push(edge);
                }
              }              
              return {
                ...incoming,
                edges: deduped,
              };
            },
          },
        },
      },
    },
  })
})