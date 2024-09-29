package sqlite

import (
	"database/sql"
	"log"
	"testing"

	"github.com/JoeSheen/godo/internal/types"
)

func setUp() *DatabaseContext {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`CREATE TABLE "tasks"("id" INTEGER NOT NULL, 
	"title" VARCHAR(255) NOT NULL, "priority" INTEGER NOT NULL CHECK (priority IN (0, 1, 2, 3, 4, 5)),
	"completed" BOOLEAN NOT NULL CHECK (completed IN (0, 1)), "category" VARCHAR(255) NOT NULL, 
	"created_timestamp" datetime NOT NULL, "completed_timestamp" datetime, "deadline" datetime, 
	PRIMARY KEY(id))`)

	return &DatabaseContext{db: db}
}

func dropDatabase(dbContext *DatabaseContext) {
	dbContext.db.Close()
}

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name       string
		want       types.Task
		expectedId int64
	}{
		{
			name: "Test Create Task",
			want: types.Task{
				Title:     "Test Title",
				Category:  "Category 1",
				Completed: false,
				Priority:  5,
				Deadline:  nil,
			},
			expectedId: 1,
		},
	}

	for _, test := range tests {
		dbContext := setUp()
		defer dropDatabase(dbContext)

		id, err := dbContext.CreateTask(test.want.Title, int(test.want.Priority), string(test.want.Category), nil)
		if err != nil {
			t.Fatalf("Failed to insert task into db: %v", err)
		}
		t.Run(test.name, func(t *testing.T) {
			if id != test.expectedId {
				t.Errorf("got %d want %d", id, test.expectedId)
			}
		})
	}
}

func TestGetAllTasks(t *testing.T) {
}
