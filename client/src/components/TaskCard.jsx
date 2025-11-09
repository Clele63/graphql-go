import React from 'react'

export default function TaskCard({ task, users, columns, onMove, isMoving, onEdit }) {
  const handleDragStart = (e) => {
    e.dataTransfer.effectAllowed = 'move'
    e.dataTransfer.setData('taskId', task.id)
    e.dataTransfer.setData('sourceColumnId', task.column.id)
    e.currentTarget.style.opacity = '0.5'
  }

  const handleDragEnd = (e) => {
    e.currentTarget.style.opacity = '1'
  }

  const handleClick = (e) => {
    // Ne pas ouvrir la modal si on clique sur le drag handle
    if (e.target.classList.contains('drag-handle')) return
    onEdit(task)
  }

  return (
    <div 
      className="task-card"
      draggable={!isMoving}
      onDragStart={handleDragStart}
      onDragEnd={handleDragEnd}
      onClick={handleClick}
    >
      <div className="task-card-header">
        <div className="task-card-title">{task.title}</div>
        <div className="drag-handle" title="Glisser pour déplacer">⋮⋮</div>
      </div>
      
      {task.description && (
        <div className="task-card-description">{task.description}</div>
      )}
      
      <div className="task-meta" style={{ marginTop: '0.75rem' }}>
        <div style={{ display: 'flex', gap: '0.25rem', flexWrap: 'wrap', flex: 1 }}>
          {task.assignees && task.assignees.length > 0 ? (
            task.assignees.map(assignee => (
              <div key={assignee.id} className="assignee-badge">
                <div className="assignee-avatar">
                  {assignee.name.charAt(0).toUpperCase()}
                </div>
                {assignee.name}
              </div>
            ))
          ) : (
            <div className="assignee-badge" style={{ opacity: 0.5 }}>
              Non assigné
            </div>
          )}
        </div>
        
        {task.comments && task.comments.length > 0 && (
          <span style={{ fontSize: '0.75rem', color: 'var(--text-tertiary)' }}>
            💬 {task.comments.length}
          </span>
        )}
        
        <button 
          className="btn-icon btn-edit" 
          onClick={(e) => {
            e.stopPropagation()
            onEdit(task)
          }}
          title="Modifier"
        >
          ✏️
        </button>
      </div>
    </div>
  )
}
