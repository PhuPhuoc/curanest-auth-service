-- +goose Up
-- +goose StatementBegin
CREATE TABLE `accounts` (
    `id` varchar(36) NOT NULL,
    `role_id` varchar(36) NOT NULL,
    `full_name` varchar(50) DEFAULT NULL,
    `email` varchar(100) NOT NULL,
    `phone_number` varchar(12) NOT NULL,
    `password` varchar(100) NOT NULL,
    `salt` varchar(80) DEFAULT NULL,
    `status` enum('activated','banned') DEFAULT 'activated',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime,
    PRIMARY KEY (`id`),
    UNIQUE KEY `unique_email_phone` (`email`, `phone_number`),
    KEY `idx_role_id` (`role_id`),
    CONSTRAINT `fk_accounts_roles` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `accounts` DROP FOREIGN KEY `fk_accounts_roles`;
DROP INDEX `idx_role_id` ON `accounts`;
DROP INDEX `unique_email_phone` ON `accounts`;
DELETE FROM `accounts`;
DROP TABLE `accounts`;
-- +goose StatementEnd
