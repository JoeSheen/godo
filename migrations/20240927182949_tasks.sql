-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    priority INTEGER CHECK (priority IN (0, 5)),
    completed BOOLEAN NOT NULL CHECK (completed IN (0, 1)),
    created_timestamp datetime not null,
    completed_timestamp datetime
    ceadline datetime
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
