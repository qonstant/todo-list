package db

import (
	"context"
	"testing"
	"time"

	"todo-list/util"

	"github.com/stretchr/testify/require"
)

func createRandomTask(t *testing.T) Task {
	arg := CreateTaskParams{
		Title:    util.RandomString(10),
		ActiveAt: util.RandomTime(),
	}

	task, err := testQueryies.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.Title, task.Title)
	require.WithinDuration(t, arg.ActiveAt.UTC(), task.ActiveAt.UTC(), time.Second)

	require.NotZero(t, task.ID)
	require.NotZero(t, task.CreatedAt)
	require.NotZero(t, task.UpdatedAt)

	return task
}

func TestCreateTask(t *testing.T) {
	createRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := createRandomTask(t)
	task2, err := testQueryies.GetTask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, task1.Title, task2.Title)
	require.WithinDuration(t, task1.ActiveAt.UTC(), task2.ActiveAt.UTC(), time.Second)
	require.Equal(t, task1.Done, task2.Done)
	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second)
	require.WithinDuration(t, task1.UpdatedAt, task2.UpdatedAt, time.Second)
}

func TestUpdateTask(t *testing.T) {
	task1 := createRandomTask(t)

	arg := UpdateTaskParams{
		ID:       task1.ID,
		Title:    util.RandomString(10),
		ActiveAt: util.RandomTime(),
	}

	task2, err := testQueryies.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, arg.Title, task2.Title)
	require.WithinDuration(t, arg.ActiveAt.UTC(), task2.ActiveAt.UTC(), time.Second)
	require.Equal(t, task1.Done, task2.Done)
	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second)
	require.True(t, task2.UpdatedAt.After(task1.UpdatedAt))
}

func TestDeleteTask(t *testing.T) {
	task1 := createRandomTask(t)
	err := testQueryies.DeleteTask(context.Background(), task1.ID)
	require.NoError(t, err)

	task2, err := testQueryies.GetTask(context.Background(), task1.ID)
	require.Error(t, err)
	require.Empty(t, task2)
}

func TestListTasks(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTask(t)
	}

	tasks, err := testQueryies.ListTasks(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, tasks)

	for _, task := range tasks {
		require.NotEmpty(t, task)
	}
}
