import React, { useState } from 'react'
import { gql, useMutation } from '@apollo/client'
import useAuth from '../store'

const LOGIN = gql`
mutation Login($email: String!, $password: String!) {
  login(email: $email, password: $password) {
    token
  }
}
`

export default function Login() {
  const [email, setEmail] = useState('alice@example.com')
  const [password, setPassword] = useState('password')
  const [login, { loading, error }] = useMutation(LOGIN)
  const setToken = useAuth(s => s.setToken)

  async function onSubmit(e) {
    e.preventDefault()
    try {
      const { data } = await login({ variables: { email, password } })
      if (data?.login?.token) {
        setToken(data.login.token)
      }
    } catch (err) {
      // L'erreur est déjà gérée par Apollo
    }
  }

  return (
    <div className="login-container">
      <form onSubmit={onSubmit} className="login-card">
        <h1>🎯 Kanban Board</h1>
        <p className="subtitle">Connectez-vous pour gérer vos tâches</p>
        
        <div className="form-group">
          <label htmlFor="email">Email</label>
          <input 
            id="email"
            type="email"
            value={email} 
            onChange={e => setEmail(e.target.value)}
            placeholder="alice@example.com"
            required
          />
        </div>

        <div className="form-group">
          <label htmlFor="password">Mot de passe</label>
          <input 
            id="password"
            type="password" 
            value={password} 
            onChange={e => setPassword(e.target.value)}
            placeholder="••••••••"
            required
          />
        </div>

        <button className="btn btn-primary" disabled={loading} type="submit">
          {loading ? (
            <>
              <span className="loading-spinner"></span>
              Connexion...
            </>
          ) : (
            'Se connecter'
          )}
        </button>

        {error && (
          <div className="error-message">
            ❌ {error.message}
          </div>
        )}

        <div className="info-text">
          <strong>Comptes de test :</strong><br />
          alice@example.com / bob@example.com / carol@example.com<br />
          <em>Mot de passe : password</em>
        </div>
      </form>
    </div>
  )
}
