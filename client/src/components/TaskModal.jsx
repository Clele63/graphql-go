// TODO 3.2 (suite) - Mutations UPDATE et DELETE
// Ce qu'il faut faire :
// - UPDATE_TASK : modifier une tâche avec UpdateTaskInput
// - DELETE_TASK : supprimer une tâche par id
// - Utiliser refetchQueries et afficher des toasts

import React, { useState, useEffect } from 'react'
import { gql, useMutation } from '@apollo/client'
import useToast from '../toastStore'
import CommentSection from './CommentSection'
import { TASK_FRAGMENT } from "../apollo";

const UPDATE_TASK = gql`
mutation UpdateTask($input: UpdateTaskInput!) {
  updateTask(input: $input) {
    ...TaskFragment
  }
}
${TASK_FRAGMENT}
`

const DELETE_TASK = gql`
mutation DeleteTask($id: ID!) {
  deleteTask(id: $id)
}
`

export default function TaskModal({ task, users, onClose }) {
  const [title, setTitle] = useState(task.title)
  const [description, setDescription] = useState(task.description || '')
  const [assigneeIds, setAssigneeIds] = useState(task.assignees?.map(a => a.id) || [])
  const addToast = useToast(s => s.addToast)
  
  const [updateTask, { loading: updating }] = useMutation(UPDATE_TASK, {
    refetchQueries: ['TasksForColumn'],
    onCompleted: () => {
      addToast('Tâche mise à jour', 'success')
    }
  })
  
  const [deleteTask, { loading: deleting }] = useMutation(DELETE_TASK, {
    refetchQueries: ['TasksForColumn', 'Board'],
    onCompleted: () => {
      addToast('Tâche supprimée', 'success')
    }
  })

  useEffect(() => {
    const handleEscape = (e) => {
      if (e.key === 'Escape') onClose()
    }
    document.addEventListener('keydown', handleEscape)
    return () => document.removeEventListener('keydown', handleEscape)
  }, [onClose])

  const handleSubmit = async (e) => {
    e.preventDefault()
    if (!title.trim()) return

    try {
      await updateTask({
        variables: {
          input: {
            id: task.id,
            title: title.trim(),
            description: description.trim() || null,
            assigneeIds: assigneeIds.length > 0 ? assigneeIds : null
          }
        }
      })
      onClose()
    } catch (err) {
      console.error('Erreur mise à jour:', err)
      addToast('Erreur: ' + err.message, 'error')
    }
  }

  function toggleAssignee(userId) {
    setAssigneeIds(prev => 
      prev.includes(userId) 
        ? prev.filter(id => id !== userId)
        : [...prev, userId]
    )
  }

  const handleDelete = async () => {
    if (!confirm('Êtes-vous sûr de vouloir supprimer cette tâche ?')) return

    try {
      await deleteTask({ variables: { id: task.id } })
      onClose()
    } catch (err) {
      console.error('Erreur suppression:', err)
      addToast('Erreur: ' + err.message, 'error')
    }
  }

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={e => e.stopPropagation()}>
        <div className="modal-header">
          <h2>Modifier la tâche</h2>
          <button className="btn-icon" onClick={onClose} title="Fermer (Esc)">
            ✕
          </button>
        </div>

        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="edit-title">Titre *</label>
            <input
              id="edit-title"
              value={title}
              onChange={e => setTitle(e.target.value)}
              placeholder="Titre de la tâche"
              required
              autoFocus
            />
          </div>

          <div className="form-group">
            <label htmlFor="edit-description">Description</label>
            <textarea
              id="edit-description"
              value={description}
              onChange={e => setDescription(e.target.value)}
              placeholder="Ajoutez une description..."
              rows={4}
            />
          </div>

          <div className="form-group">
            <label>Assigner à (sélection multiple)</label>
            <div style={{ display: 'flex', gap: '0.5rem', flexWrap: 'wrap', marginTop: '0.5rem' }}>
              {users.map(u => (
                <button
                  key={u.id}
                  type="button"
                  className={assigneeIds.includes(u.id) ? 'btn btn-primary btn-small' : 'btn btn-secondary btn-small'}
                  onClick={() => toggleAssignee(u.id)}
                  style={{ minWidth: 'auto' }}
                >
                  {assigneeIds.includes(u.id) ? '✓ ' : ''}{u.name}
                </button>
              ))}
            </div>
          </div>

          {/* Section des commentaires */}
          <CommentSection task={task} />

          <div className="modal-actions">
            <button
              type="button"
              className="btn btn-danger btn-small"
              onClick={handleDelete}
              disabled={deleting}
            >
              {deleting ? '...' : '🗑️ Supprimer'}
            </button>
            <div style={{ marginLeft: 'auto', display: 'flex', gap: '0.5rem' }}>
              <button
                type="button"
                className="btn btn-secondary btn-small"
                onClick={onClose}
              >
                Annuler
              </button>
              <button
                type="submit"
                className="btn btn-primary btn-small"
                disabled={updating}
              >
                {updating ? '...' : '✓ Enregistrer'}
              </button>
            </div>
          </div>
        </form>
      </div>
    </div>
  )
}
