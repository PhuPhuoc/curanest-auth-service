-- +goose Up
-- +goose StatementBegin
CREATE TABLE `roles` (
    `id` varchar(36) NOT NULL,
    `name` varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM `roles`;
DROP TABLE `roles`;
-- +goose StatementEnd
