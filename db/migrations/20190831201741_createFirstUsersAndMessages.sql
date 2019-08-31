
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO users
VALUES
(1, 'テストユーザ1')

INSERT INTO messages
VALUES
(1, 'テストメッセージ1')

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

