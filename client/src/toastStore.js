import { create } from 'zustand'

const useToast = create(set => ({
  toasts: [],
  addToast: (message, type = 'info') => {
    const id = Date.now()
    set(state => ({ 
      toasts: [...state.toasts, { id, message, type }] 
    }))
    setTimeout(() => {
      set(state => ({ 
        toasts: state.toasts.filter(t => t.id !== id) 
      }))
    }, 3000)
  },
  removeToast: (id) => {
    set(state => ({ 
      toasts: state.toasts.filter(t => t.id !== id) 
    }))
  }
}))

export default useToast
