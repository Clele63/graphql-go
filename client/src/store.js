import { create } from 'zustand'

const useAuth = create(set => ({
  token: null,
  setToken: (token) => set({ token }),
  logout: () => set({ token: null })
}))

export default useAuth
