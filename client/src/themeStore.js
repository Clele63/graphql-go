import { create } from 'zustand'
import { persist } from 'zustand/middleware'

const useTheme = create(
  persist(
    (set) => ({
      isDark: false,
      toggleTheme: () => set((state) => ({ isDark: !state.isDark }))
    }),
    {
      name: 'theme-storage'
    }
  )
)

export default useTheme
