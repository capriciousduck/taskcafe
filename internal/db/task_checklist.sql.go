// Code generated by sqlc. DO NOT EDIT.
// source: task_checklist.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTaskChecklist = `-- name: CreateTaskChecklist :one
INSERT INTO task_checklist (task_id, created_at, name, position) VALUES ($1, $2, $3, $4)
  RETURNING task_checklist_id, task_id, created_at, name, position
`

type CreateTaskChecklistParams struct {
	TaskID    uuid.UUID `json:"task_id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Position  float64   `json:"position"`
}

func (q *Queries) CreateTaskChecklist(ctx context.Context, arg CreateTaskChecklistParams) (TaskChecklist, error) {
	row := q.db.QueryRowContext(ctx, createTaskChecklist,
		arg.TaskID,
		arg.CreatedAt,
		arg.Name,
		arg.Position,
	)
	var i TaskChecklist
	err := row.Scan(
		&i.TaskChecklistID,
		&i.TaskID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}

const createTaskChecklistItem = `-- name: CreateTaskChecklistItem :one
INSERT INTO task_checklist_item (task_checklist_id, created_at, name, position, complete, due_date) VALUES ($1, $2, $3, $4, false, null)
  RETURNING task_checklist_item_id, task_checklist_id, created_at, complete, name, position, due_date
`

type CreateTaskChecklistItemParams struct {
	TaskChecklistID uuid.UUID `json:"task_checklist_id"`
	CreatedAt       time.Time `json:"created_at"`
	Name            string    `json:"name"`
	Position        float64   `json:"position"`
}

func (q *Queries) CreateTaskChecklistItem(ctx context.Context, arg CreateTaskChecklistItemParams) (TaskChecklistItem, error) {
	row := q.db.QueryRowContext(ctx, createTaskChecklistItem,
		arg.TaskChecklistID,
		arg.CreatedAt,
		arg.Name,
		arg.Position,
	)
	var i TaskChecklistItem
	err := row.Scan(
		&i.TaskChecklistItemID,
		&i.TaskChecklistID,
		&i.CreatedAt,
		&i.Complete,
		&i.Name,
		&i.Position,
		&i.DueDate,
	)
	return i, err
}

const deleteTaskChecklistByID = `-- name: DeleteTaskChecklistByID :exec
DELETE FROM task_checklist WHERE task_checklist_id = $1
`

func (q *Queries) DeleteTaskChecklistByID(ctx context.Context, taskChecklistID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTaskChecklistByID, taskChecklistID)
	return err
}

const deleteTaskChecklistItem = `-- name: DeleteTaskChecklistItem :exec
DELETE FROM task_checklist_item WHERE task_checklist_item_id = $1
`

func (q *Queries) DeleteTaskChecklistItem(ctx context.Context, taskChecklistItemID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTaskChecklistItem, taskChecklistItemID)
	return err
}

const getTaskChecklistByID = `-- name: GetTaskChecklistByID :one
SELECT task_checklist_id, task_id, created_at, name, position FROM task_checklist WHERE task_checklist_id = $1
`

func (q *Queries) GetTaskChecklistByID(ctx context.Context, taskChecklistID uuid.UUID) (TaskChecklist, error) {
	row := q.db.QueryRowContext(ctx, getTaskChecklistByID, taskChecklistID)
	var i TaskChecklist
	err := row.Scan(
		&i.TaskChecklistID,
		&i.TaskID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}

const getTaskChecklistItemByID = `-- name: GetTaskChecklistItemByID :one
SELECT task_checklist_item_id, task_checklist_id, created_at, complete, name, position, due_date FROM task_checklist_item WHERE task_checklist_item_id = $1
`

func (q *Queries) GetTaskChecklistItemByID(ctx context.Context, taskChecklistItemID uuid.UUID) (TaskChecklistItem, error) {
	row := q.db.QueryRowContext(ctx, getTaskChecklistItemByID, taskChecklistItemID)
	var i TaskChecklistItem
	err := row.Scan(
		&i.TaskChecklistItemID,
		&i.TaskChecklistID,
		&i.CreatedAt,
		&i.Complete,
		&i.Name,
		&i.Position,
		&i.DueDate,
	)
	return i, err
}

const getTaskChecklistItemsForTaskChecklist = `-- name: GetTaskChecklistItemsForTaskChecklist :many
SELECT task_checklist_item_id, task_checklist_id, created_at, complete, name, position, due_date FROM task_checklist_item WHERE task_checklist_id = $1
`

func (q *Queries) GetTaskChecklistItemsForTaskChecklist(ctx context.Context, taskChecklistID uuid.UUID) ([]TaskChecklistItem, error) {
	rows, err := q.db.QueryContext(ctx, getTaskChecklistItemsForTaskChecklist, taskChecklistID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskChecklistItem
	for rows.Next() {
		var i TaskChecklistItem
		if err := rows.Scan(
			&i.TaskChecklistItemID,
			&i.TaskChecklistID,
			&i.CreatedAt,
			&i.Complete,
			&i.Name,
			&i.Position,
			&i.DueDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTaskChecklistsForTask = `-- name: GetTaskChecklistsForTask :many
SELECT task_checklist_id, task_id, created_at, name, position FROM task_checklist WHERE task_id = $1
`

func (q *Queries) GetTaskChecklistsForTask(ctx context.Context, taskID uuid.UUID) ([]TaskChecklist, error) {
	rows, err := q.db.QueryContext(ctx, getTaskChecklistsForTask, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TaskChecklist
	for rows.Next() {
		var i TaskChecklist
		if err := rows.Scan(
			&i.TaskChecklistID,
			&i.TaskID,
			&i.CreatedAt,
			&i.Name,
			&i.Position,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setTaskChecklistItemComplete = `-- name: SetTaskChecklistItemComplete :one
UPDATE task_checklist_item SET complete = $2 WHERE task_checklist_item_id = $1
  RETURNING task_checklist_item_id, task_checklist_id, created_at, complete, name, position, due_date
`

type SetTaskChecklistItemCompleteParams struct {
	TaskChecklistItemID uuid.UUID `json:"task_checklist_item_id"`
	Complete            bool      `json:"complete"`
}

func (q *Queries) SetTaskChecklistItemComplete(ctx context.Context, arg SetTaskChecklistItemCompleteParams) (TaskChecklistItem, error) {
	row := q.db.QueryRowContext(ctx, setTaskChecklistItemComplete, arg.TaskChecklistItemID, arg.Complete)
	var i TaskChecklistItem
	err := row.Scan(
		&i.TaskChecklistItemID,
		&i.TaskChecklistID,
		&i.CreatedAt,
		&i.Complete,
		&i.Name,
		&i.Position,
		&i.DueDate,
	)
	return i, err
}

const updateTaskChecklistItemLocation = `-- name: UpdateTaskChecklistItemLocation :one
UPDATE task_checklist_item SET position = $2, task_checklist_id = $3 WHERE task_checklist_item_id = $1 RETURNING task_checklist_item_id, task_checklist_id, created_at, complete, name, position, due_date
`

type UpdateTaskChecklistItemLocationParams struct {
	TaskChecklistItemID uuid.UUID `json:"task_checklist_item_id"`
	Position            float64   `json:"position"`
	TaskChecklistID     uuid.UUID `json:"task_checklist_id"`
}

func (q *Queries) UpdateTaskChecklistItemLocation(ctx context.Context, arg UpdateTaskChecklistItemLocationParams) (TaskChecklistItem, error) {
	row := q.db.QueryRowContext(ctx, updateTaskChecklistItemLocation, arg.TaskChecklistItemID, arg.Position, arg.TaskChecklistID)
	var i TaskChecklistItem
	err := row.Scan(
		&i.TaskChecklistItemID,
		&i.TaskChecklistID,
		&i.CreatedAt,
		&i.Complete,
		&i.Name,
		&i.Position,
		&i.DueDate,
	)
	return i, err
}

const updateTaskChecklistItemName = `-- name: UpdateTaskChecklistItemName :one
UPDATE task_checklist_item SET name = $2 WHERE task_checklist_item_id = $1
  RETURNING task_checklist_item_id, task_checklist_id, created_at, complete, name, position, due_date
`

type UpdateTaskChecklistItemNameParams struct {
	TaskChecklistItemID uuid.UUID `json:"task_checklist_item_id"`
	Name                string    `json:"name"`
}

func (q *Queries) UpdateTaskChecklistItemName(ctx context.Context, arg UpdateTaskChecklistItemNameParams) (TaskChecklistItem, error) {
	row := q.db.QueryRowContext(ctx, updateTaskChecklistItemName, arg.TaskChecklistItemID, arg.Name)
	var i TaskChecklistItem
	err := row.Scan(
		&i.TaskChecklistItemID,
		&i.TaskChecklistID,
		&i.CreatedAt,
		&i.Complete,
		&i.Name,
		&i.Position,
		&i.DueDate,
	)
	return i, err
}

const updateTaskChecklistName = `-- name: UpdateTaskChecklistName :one
UPDATE task_checklist SET name = $2 WHERE task_checklist_id = $1
  RETURNING task_checklist_id, task_id, created_at, name, position
`

type UpdateTaskChecklistNameParams struct {
	TaskChecklistID uuid.UUID `json:"task_checklist_id"`
	Name            string    `json:"name"`
}

func (q *Queries) UpdateTaskChecklistName(ctx context.Context, arg UpdateTaskChecklistNameParams) (TaskChecklist, error) {
	row := q.db.QueryRowContext(ctx, updateTaskChecklistName, arg.TaskChecklistID, arg.Name)
	var i TaskChecklist
	err := row.Scan(
		&i.TaskChecklistID,
		&i.TaskID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}

const updateTaskChecklistPosition = `-- name: UpdateTaskChecklistPosition :one
UPDATE task_checklist SET position = $2 WHERE task_checklist_id = $1 RETURNING task_checklist_id, task_id, created_at, name, position
`

type UpdateTaskChecklistPositionParams struct {
	TaskChecklistID uuid.UUID `json:"task_checklist_id"`
	Position        float64   `json:"position"`
}

func (q *Queries) UpdateTaskChecklistPosition(ctx context.Context, arg UpdateTaskChecklistPositionParams) (TaskChecklist, error) {
	row := q.db.QueryRowContext(ctx, updateTaskChecklistPosition, arg.TaskChecklistID, arg.Position)
	var i TaskChecklist
	err := row.Scan(
		&i.TaskChecklistID,
		&i.TaskID,
		&i.CreatedAt,
		&i.Name,
		&i.Position,
	)
	return i, err
}
