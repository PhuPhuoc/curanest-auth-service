-- +goose Up
-- +goose StatementBegin
CREATE TABLE `roles` (
    `id` varchar(36) NOT NULL,
    `name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_name` (`name`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX `unique_name` ON `roles`;
DELETE FROM `roles`;
DROP TABLE `roles`;
-- +goose StatementEnd
