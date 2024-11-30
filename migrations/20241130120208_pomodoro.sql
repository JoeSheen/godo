-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS pomodoro(
    id INTEGER NOT NULL,
    name VARCHAR(255),
    active BOOLEAN NOT NULL CHECK (active IN (0, 1)),
    start_time datetime NOT NULL,
    end_time datetime NOT NULL,
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pomodoro;
-- +goose StatementEnd
