package sqlite

import (
	"database/sql"
	"fmt"
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
	var wantedId int64 = 1

	dbContext := setUp()
	defer dropDatabase(dbContext)

	id, err := dbContext.CreateTask("Create Task", 1, "Category", nil)
	if err != nil {
		t.Fatalf("Failed to insert task into db: %v", err)
	}
	err = testIdValue(wantedId, id)
	if err != nil {
		t.Error(err)
	}
}

func TestGetAllTasks(t *testing.T) {
	tasks := []types.Task{
		{
			Title:    "test task 1",
			Priority: 2,
			Category: "Category 2",
		},
		{
			Title:    "test task 2",
			Priority: 3,
			Category: "Category 3",
		},
		{
			Title:    "test task 3",
			Priority: 4,
			Category: "Category 1",
		},
	}

	dbContext := setUp()
	defer dropDatabase(dbContext)
	insertTaskTestHelper(t, dbContext, tasks...)

	returnedTasks, err := dbContext.GetAllTasks()
	if err != nil {
		t.Errorf("Error getting all tasks from database: %v", err)
	}

	if len(returnedTasks) != 3 {
		t.Errorf("got %d want %d", len(returnedTasks), 3)
	}

	for i := 0; i < len(tasks); i++ {
		got := returnedTasks[i]
		want := tasks[i]
		err := testIdValue(int64(i+1), got.ID)
		if err != nil {
			t.Error(err)
		}
		err = testStringValues(want.Title, got.Title)
		if err != nil {
			t.Error(err)
		}
		err = testPriorityValues(want.Priority, got.Priority)
		if err != nil {
			t.Error(err)
		}
		if got.Completed {
			t.Errorf("expected completed to be false but was true")
		}
		err = testStringValues(string(want.Category), string(got.Category))
		if err != nil {
			t.Error(err)
		}
	}
}

func TestGetTaskById(t *testing.T) {
	tasks := types.Task{
		ID:       1,
		Title:    "test task",
		Priority: 5,
		Category: "By ID",
	}

	dbContext := setUp()
	defer dropDatabase(dbContext)
	insertTaskTestHelper(t, dbContext, tasks)

	got, err := dbContext.GetTaskById(1)
	if err != nil {
		t.Errorf("Error getting task from database: %v", err)
	}
	err = testIdValue(1, got.ID)
	if err != nil {
		t.Error(err)
	}
	err = testStringValues("test task", got.Title)
	if err != nil {
		t.Error(err)
	}
	err = testPriorityValues(types.Priority(5), got.Priority)
	if err != nil {
		t.Error(err)
	}
	if got.Completed {
		t.Errorf("expected completed to be false but was true")
	}
	err = testStringValues(string("By ID"), string(got.Category))
	if err != nil {
		t.Error(err)
	}

}

func TestGetAllTasksByCompletedStatus(t *testing.T) {
	tasks := []types.Task{
		{
			Title:    "some task",
			Priority: 3,
			Category: "Category 2",
		},
		{
			Title:    "next task",
			Priority: 2,
			Category: "Category 3",
		},
		{
			Title:    "final task",
			Priority: 1,
			Category: "Category 1",
		},
	}

	dbContext := setUp()
	defer dropDatabase(dbContext)
	insertTaskTestHelper(t, dbContext, tasks...)

	// Getting all completed tasks from the slice above
	returnedTasks, err := dbContext.GetAllTasksByCompletedStatus(true)
	if err != nil {
		t.Errorf("Error getting all completed tasks from database: %v", err)
	}
	if len(returnedTasks) != 0 {
		t.Errorf("expected=%d, got=%d", len(returnedTasks), 0)
	}

	// Getting all outstanding tasks from the slice above
	returnedTasks, err = dbContext.GetAllTasksByCompletedStatus(false)
	if err != nil {
		t.Errorf("Error getting all completed tasks from database: %v", err)
	}
	if len(returnedTasks) != 3 {
		t.Errorf("expected=%d, got=%d", len(returnedTasks), 0)
	}
}

func TestToggleTaskCompleted(t *testing.T) {
	task := types.Task{
		ID:       1,
		Title:    "toggle task",
		Priority: 2,
		Category: "Category 2",
	}

	dbContext := setUp()
	defer dropDatabase(dbContext)
	insertTaskTestHelper(t, dbContext, task)

	err := dbContext.ToggleTaskCompleted(1)
	if err != nil {
		t.Errorf("failed to toggle task: %v", err)
	}

	got, err := dbContext.GetTaskById(1)
	if err != nil {
		t.Errorf("Error getting task from database: %v", err)
	}
	err = testIdValue(1, got.ID)
	if err != nil {
		t.Error(err)
	}
	err = testStringValues("toggle task", got.Title)
	if err != nil {
		t.Error(err)
	}
	err = testPriorityValues(types.Priority(2), got.Priority)
	if err != nil {
		t.Error(err)
	}
	err = testStringValues(string("Category 2"), string(got.Category))
	if err != nil {
		t.Error(err)
	}
	if got.Completed != true {
		t.Error("expected completed flag to be true")
	}
	if got.CompletedTimestamp == nil {
		t.Error("expected completed timestamp to be set")
	}
}

func TestDeleteTaskById(t *testing.T) {
	tasks := types.Task{
		ID:       1,
		Title:    "test task",
		Priority: 5,
		Category: "By ID",
	}

	dbContext := setUp()
	defer dropDatabase(dbContext)

	// check table is empty at start
	returnedTasks, err := dbContext.GetAllTasks()
	if err != nil {
		t.Errorf("Error getting all tasks from database: %v", err)
	}
	if len(returnedTasks) != 0 {
		t.Errorf("got %d want %d", len(returnedTasks), 0)
	}

	// inserts a task to be deleted by ID
	insertTaskTestHelper(t, dbContext, tasks)

	// checks the task has been added
	returnedTasks, err = dbContext.GetAllTasks()
	if err != nil {
		t.Errorf("Error getting all tasks from database: %v", err)
	}
	if len(returnedTasks) != 1 {
		t.Errorf("got %d want %d", len(returnedTasks), 1)
	}

	// deletes the task by ID
	err = dbContext.DeleteTaskById(1)
	if err != nil {
		t.Error(err)
	}

	// requests all tasks to ensure it is deleted
	returnedTasks, err = dbContext.GetAllTasks()
	if err != nil {
		t.Errorf("Error getting all tasks from database: %v", err)
	}
	if len(returnedTasks) != 0 {
		t.Errorf("got %d want %d", len(returnedTasks), 0)
	}
}

func TestDeleteAllTasks(t *testing.T) {
	tasks := []types.Task{
		{
			Title:    "test task 1",
			Priority: 2,
			Category: "Category 2",
		},
		{
			Title:    "test task 2",
			Priority: 3,
			Category: "Category 3",
		},
		{
			Title:    "test task 3",
			Priority: 4,
			Category: "Category 1",
		},
	}

	dbContext := setUp()
	defer dropDatabase(dbContext)
	insertTaskTestHelper(t, dbContext, tasks...)

	// asserts all three tasks have been added
	returnedTasks, err := dbContext.GetAllTasks()
	if err != nil {
		t.Errorf("Error getting all tasks from database: %v", err)
	}
	if len(returnedTasks) != 3 {
		t.Errorf("got %d want %d", len(returnedTasks), 3)
	}

	// performs delete all
	dbContext.DeleteAllTasks()

	// checks that the DB is now empty
	returnedTasks, err = dbContext.GetAllTasks()
	if err != nil {
		t.Errorf("Error getting all tasks from database: %v", err)
	}
	if len(returnedTasks) != 0 {
		t.Errorf("got %d want %d", len(returnedTasks), 0)
	}
}

func insertTaskTestHelper(t testing.TB, dc *DatabaseContext, tasks ...types.Task) {
	t.Helper()
	for _, task := range tasks {
		_, err := dc.CreateTask(task.Title, int(task.Priority), string(task.Category), nil)
		if err != nil {
			t.Fatalf("Failed to insert task into db: %v", err)
		}
	}
}

func testIdValue(want, got int64) error {
	if want != got {
		return fmt.Errorf("expected=%d, got=%d", want, got)
	}
	return nil
}

func testStringValues(want, got string) error {
	if want != got {
		return fmt.Errorf("expected=%s, got=%s", want, got)
	}
	return nil
}

func testPriorityValues(want, got types.Priority) error {
	if want != got {
		return fmt.Errorf("expected=%d, got=%d", want, got)
	}
	return nil
}
