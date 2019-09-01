
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `users` (
	`id` 				BIGINT				NOT NULL, 
	`name` 				VARCHAR(255)				DEFAULT NULL, 
	`created_at` 		DATETIME			NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`id`)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TAPLE "users";