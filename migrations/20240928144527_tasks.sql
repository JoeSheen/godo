-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    priority INTEGER NOT NULL CHECK (priority IN (0, 1, 2, 3, 4, 5)),
    completed BOOLEAN NOT NULL CHECK (completed IN (0, 1)),
    category VARCHAR(255) NOT NULL,
    created_timestamp datetime NOT NULL,
    completed_timestamp datetime,
    deadline datetime,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
