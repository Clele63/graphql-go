// TODO 3.3 (suite) - Subscription commentaires
// Ce qu'il faut faire :
// - ADD_COMMENT : ajouter un commentaire à une tâche
// - DELETE_COMMENT : supprimer un commentaire par id
// - COMMENT_ADDED : écouter les nouveaux commentaires
// - ME : récupérer l'utilisateur connecté

import React, { useState, useEffect } from 'react'
import { gql, useMutation, useSubscription, useQuery } from '@apollo/client'
import useToast from '../toastStore'
import { COMMENT_FRAGMENT } from '../apollo'

const ADD_COMMENT = gql`
mutation AddComment($taskId: ID!, $content: String!) {
  addComment(taskId: $taskId, content: $content) {
    ...CommentFragment
  }
}
${COMMENT_FRAGMENT}
`

const DELETE_COMMENT = gql`
mutation DeleteComment($id: ID!) {
  deleteComment(id: $id)
}
`

const COMMENT_ADDED = gql`
subscription OnCommentAdded($taskId: ID!) {
  commentAdded(taskId: $taskId) {
    ...CommentFragment
  }
}
${COMMENT_FRAGMENT}
`

const ME = gql`
query Me {
  me {
    id
  }
}
`

export default function CommentSection({ task, onCommentsUpdate }) {
  const [content, setContent] = useState('')
  const [comments, setComments] = useState(task.comments || [])
  const { data: userData } = useQuery(ME)
  const user = userData?.me
  const addToast = useToast(s => s.addToast)

  const [addComment, { loading: adding }] = useMutation(ADD_COMMENT, {
    onCompleted: (data) => {
      setContent('')
      addToast('Commentaire ajouté', 'success')
    },
    onError: (error) => {
      addToast('Erreur: ' + error.message, 'error')
    }
  })

  const [deleteComment] = useMutation(DELETE_COMMENT, {
    refetchQueries: ['TasksForColumn'],
    onCompleted: () => {
      addToast('Commentaire supprimé', 'success')
    },
    onError: (error) => {
      addToast('Erreur: ' + error.message, 'error')
    }
  })

  useSubscription(COMMENT_ADDED, {
    variables: { taskId: task.id },
    onData: ({ data }) => {
      const newComment = data.data.commentAdded
      if (newComment) {
        setComments(prev => {
          if (prev.find(c => c.id === newComment.id)) return prev
          return [...prev, newComment]
        })
      }
    }
  })

  useEffect(() => {
    setComments(task.comments || [])
  }, [task.comments])

  async function handleSubmit(e) {
    e.preventDefault()
    if (!content.trim()) return

    try {
      await addComment({
        variables: {
          taskId: task.id,
          content: content.trim()
        }
      })
    } catch (err) {
      console.error('Erreur ajout commentaire:', err)
    }
  }

  async function handleDelete(commentId) {
    if (!confirm('Supprimer ce commentaire ?')) return

    try {
      await deleteComment({ variables: { id: commentId } })
      setComments(prev => prev.filter(c => c.id !== commentId))
    } catch (err) {
      console.error('Erreur suppression:', err)
    }
  }

  function formatDate(dateString) {
    const date = new Date(dateString)
    const now = new Date()
    const diffMs = now - date
    const diffMins = Math.floor(diffMs / 60000)
    const diffHours = Math.floor(diffMs / 3600000)
    const diffDays = Math.floor(diffMs / 86400000)

    if (diffMins < 1) return 'À l\'instant'
    if (diffMins < 60) return `Il y a ${diffMins} min`
    if (diffHours < 24) return `Il y a ${diffHours}h`
    if (diffDays < 7) return `Il y a ${diffDays}j`
    
    return date.toLocaleDateString('fr-FR', { 
      day: 'numeric', 
      month: 'short', 
      year: date.getFullYear() !== now.getFullYear() ? 'numeric' : undefined 
    })
  }

  return (
    <div className="comment-section">
      <h4 className="comment-section-title">
        💬 Commentaires ({comments.length})
      </h4>

      <div className="comments-list">
        {comments.length === 0 ? (
          <div className="empty-comments">
            <span style={{ fontSize: '2rem', opacity: 0.3 }}>💭</span>
            <p>Aucun commentaire pour le moment</p>
          </div>
        ) : (
          [...comments]
            .sort((a, b) => new Date(a.createdAt) - new Date(b.createdAt))
            .map(comment => (
              <div key={comment.id} className="comment-item">
                <div className="comment-header">
                  <div className="comment-author">
                    <div className="comment-avatar">
                      {comment.author.name.charAt(0).toUpperCase()}
                    </div>
                    <div>
                      <div className="comment-author-name">{comment.author.name}</div>
                      <div className="comment-date">{formatDate(comment.createdAt)}</div>
                    </div>
                  </div>
                  {comment.author.id === user?.id && (
                    <button
                      className="btn-icon btn-delete-comment"
                      onClick={() => handleDelete(comment.id)}
                      title="Supprimer"
                    >
                      🗑️
                    </button>
                  )}
                </div>
                <div className="comment-content">{comment.content}</div>
              </div>
            ))
        )}
      </div>

      <div className="comment-form">
        <textarea
          value={content}
          onChange={e => setContent(e.target.value)}
          placeholder="Ajoutez un commentaire..."
          rows={3}
          disabled={adding}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && (e.metaKey || e.ctrlKey)) {
              e.preventDefault()
              handleSubmit(e)
            }
          }}
        />
        <div className="comment-form-actions">
          <button
            type="button"
            onClick={handleSubmit}
            className="btn btn-primary btn-small"
            disabled={adding || !content.trim()}
          >
            {adding ? (
              <span className="loading-spinner" style={{ width: 14, height: 14 }}></span>
            ) : (
              '💬 Commenter'
            )}
          </button>
        </div>
      </div>
    </div>
  )
}
