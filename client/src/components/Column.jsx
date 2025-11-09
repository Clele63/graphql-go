// TODO 3.2 (suite) - Query paginée
// Ce qu'il faut faire :
// - TASKS_FOR_COLUMN : board avec columns.tasks (pagination)
// - MOVE_TASK : déplacer une tâche vers une autre colonne
// - Gérer le drag & drop et le bouton "Charger plus"

import React, { useState } from 'react'
import { gql, useQuery, useMutation } from '@apollo/client'
import TaskCard from './TaskCard'
import TaskModal from './TaskModal'
import useToast from '../toastStore'
import { TASK_FRAGMENT } from '../apollo'

const TASKS_FOR_COLUMN = gql`
query TasksForColumn($columnId: ID!, $after: String) {
  column(id: $columnId) {
    id
    name
    tasks(first: 5, after: $after) {
      edges {
        cursor
        node {
          ...TaskFragment
        }
      }
      pageInfo {
        endCursor
        hasNextPage
      }
    }
  }
}
${TASK_FRAGMENT}
`

const MOVE_TASK = gql`
mutation MoveTask($id: ID!, $to: ID!) {
  moveTask(id: $id, toColumnId: $to) { id column { id name } }
}
`

const COLUMN_ICONS = {
  'Todo': '📋',
  'Doing': '⚡',
  'Done': '✅'
}

export default function Column({ column, users, allColumns }) {
  const { data, loading, error, fetchMore } = useQuery(TASKS_FOR_COLUMN, {
    variables: { columnId: column.id, after: null },
    fetchPolicy: 'cache-and-network'
  })
  const addToast = useToast(s => s.addToast)
  
  const [moveTask, { loading: moving }] = useMutation(MOVE_TASK, { 
    refetchQueries: ['TasksForColumn', 'Board'],
    onCompleted: (data) => {
      const task = data?.moveTask
      if (task) {
        addToast(`Déplacé vers ${task.column.name}`, 'success')
      }
    }
  })

  const [isDragOver, setIsDragOver] = useState(false)
  const [editingTask, setEditingTask] = useState(null)

  if (loading) {
    return (
      <div className="column">
        <div className="column-header">
          <h3>{COLUMN_ICONS[column.name] || '📌'} {column.name}</h3>
          <div className="column-badge">...</div>
        </div>
        <div className="tasks-list">
          <div className="skeleton" style={{ height: 100 }}></div>
          <div className="skeleton" style={{ height: 120 }}></div>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="column">
        <div className="column-header">
          <h3>{column.name}</h3>
        </div>
        <div className="error-message">❌ {error.message}</div>
      </div>
    )
  }

  const col = data?.column
  const tasks = col?.tasks?.edges ?? []
  const pageInfo = col?.tasks?.pageInfo ?? { hasNextPage: false }

  async function move(id, toColumnId) {
    try {
      await moveTask({ variables: { id, to: toColumnId } })
    } catch (err) {
      console.error('Erreur déplacement:', err)
      addToast('Erreur lors du déplacement', 'error')
    }
  }

  const handleDragOver = (e) => {
    e.preventDefault()
    e.dataTransfer.dropEffect = 'move'
    setIsDragOver(true)
  }

  const handleDragLeave = (e) => {
    e.preventDefault()
    setIsDragOver(false)
  }

  const handleDrop = async (e) => {
    e.preventDefault()
    setIsDragOver(false)
    
    const taskId = e.dataTransfer.getData('taskId')
    const sourceColumnId = e.dataTransfer.getData('sourceColumnId')
    
    if (taskId && sourceColumnId !== column.id) {
      await move(taskId, column.id)
    }
  }

  return (
    <>
      <div 
        className={`column ${isDragOver ? 'column-drag-over' : ''}`}
        onDragOver={handleDragOver}
        onDragLeave={handleDragLeave}
        onDrop={handleDrop}
      >
        <div className="column-header">
          <h3>{COLUMN_ICONS[column.name] || '📌'} {column.name}</h3>
          <div className="column-badge">{tasks.length}</div>
        </div>
        
        <div className="tasks-list">
          {tasks.length === 0 ? (
            <div className="empty-state">
              <div className="empty-state-icon">📭</div>
              <div className="empty-state-text">
                {isDragOver ? 'Déposer ici' : 'Aucune tâche'}
              </div>
            </div>
          ) : (
            tasks.map(({ node }) => (
              <TaskCard 
                key={node.id} 
                task={node} 
                users={users} 
                columns={allColumns} 
                onMove={move}
                onEdit={setEditingTask}
                isMoving={moving}
              />
            ))
          )}
        </div>
        
        {pageInfo.hasNextPage && (
          <button
            className="load-more-btn"
            onClick={() => 
              fetchMore({ 
                variables: { 
                  columnId: column.id, 
                  after: col.tasks.pageInfo.endCursor 
                },
                updateQuery: (prev, { fetchMoreResult }) => {
                  if (!fetchMoreResult) return prev;
                  
                  const prevTasks = prev.column.tasks;
                  const newTasks = fetchMoreResult.column.tasks;
                  
                  return {
                    column: {
                      ...prev.column,
                      tasks: {
                        ...newTasks,
                        edges: [
                          ...prevTasks.edges,
                          ...newTasks.edges
                        ]
                      }
                    }
                  };
                }
              })
            }
          >
            ⬇️ Charger plus de tâches
          </button>
        )}
      </div>

      {editingTask && (
        <TaskModal
          task={editingTask}
          users={users}
          onClose={() => setEditingTask(null)}
        />
      )}
    </>
  )
}
