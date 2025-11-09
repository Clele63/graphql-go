import React, { useEffect } from 'react'
import { ApolloProvider } from '@apollo/client'
import { client } from './apollo'
import Login from './pages/Login'
import Board from './pages/Board'
import ToastContainer from './components/ToastContainer'
import useAuth from './store'
import useTheme from './themeStore'

export default function App() {
  const token = useAuth(s => s.token)
  const isDark = useTheme(s => s.isDark)

  useEffect(() => {
    if (isDark) {
      document.documentElement.classList.add('dark')
    } else {
      document.documentElement.classList.remove('dark')
    }
  }, [isDark])

  return (
    <ApolloProvider client={client}>
      {!token ? <Login /> : <Board />}
      <ToastContainer />
    </ApolloProvider>
  )
}
