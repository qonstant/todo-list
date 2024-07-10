package db

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/mock"
// )

// type MockQueries struct {
// 	mock.Mock
// }

// func (m *MockQueries) ExecTx(ctx context.Context, fn func(*Queries) error) error {
// 	args := m.Called(ctx, fn)
// 	return args.Error(0)
// }

// func TestSQLStore_ExecTx_Success(t *testing.T) {
// 	mockQueries := new(MockQueries)
// 	mockQueries.On("ExecTx", mock.Anything, mock.Anything).Return(nil)

// 	mockDB := new(sql.DB)
// 	store := NewStore(mockDB)
// 	sqlStore := store.(*SQLStore)

// 	err := sqlStore.ExecTx(context.Background(), func(q *Queries) error {
// 		return mockQueries.ExecTx(context.Background(), nil)
// 	})

// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}

// 	mockQueries.AssertCalled(t, "ExecTx", mock.Anything, mock.Anything)
// }

// func TestSQLStore_ExecTx_Error(t *testing.T) {
// 	mockQueries := new(MockQueries)
// 	mockQueries.On("ExecTx", mock.Anything, mock.Anything).Return(fmt.Errorf("mock error"))

// 	mockDB := new(sql.DB)
// 	store := NewStore(mockDB)
// 	sqlStore := store.(*SQLStore)

// 	err := sqlStore.ExecTx(context.Background(), func(q *Queries) error {
// 		return mockQueries.ExecTx(context.Background(), nil)
// 	})

// 	if err == nil {
// 		t.Fatal("expected an error but got nil")
// 	}

// 	mockQueries.AssertCalled(t, "ExecTx", mock.Anything, mock.Anything)
// }
