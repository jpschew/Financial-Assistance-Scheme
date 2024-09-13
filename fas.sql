CREATE DATABASE `fas`;

CREATE TABLE IF NOT EXISTS `fas`.`admins` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(128) NOT NULL,
    `username` varchar(128) NOT NULL,
    `password` varchar(128) NOT NULL,
    `create_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    `update_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    PRIMARY KEY (`id`),
    UNIQUE (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO fas.schemes (name, username, password, create_time, update_time)
VALUES ("ADMIN", "admin", "$2a$10$2lyAlVsbrnJi4Y1z0p7psujSR1HlOMGadqnv8PgBPeYHxQVH6vlZu", UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

CREATE TABLE IF NOT EXISTS `fas`.`schemes` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(128) NOT NULL,
    `description` varchar(128) NOT NULL,
    `employment_status` int NOT NULL,
    `martial_status` int NOT NULL,
    `children_status` int NOT NULL,
    `benefits` JSON NOT NULL,
    `create_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    `update_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    PRIMARY KEY (`id`),
    UNIQUE (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO fas.schemes (name, description, employment_status, martial_status, children_status, benefits, create_time, update_time)
VALUES ("Retrenchment assistance scheme", "Financial assistance for retrenched workers", 2, 0, 0, '[{"name": "SkillsFuture credits", "amount": 500.00}]', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO fas.schemes (name, description, employment_status, martial_status, children_status, benefits, create_time, update_time)
VALUES ("Retrenchment assistance scheme (family)", "Financial assistance for schooling children of retrenched workers", 2, 0, 1, '[{"name": "Daily school meal voucher", "amount": 10.00}]', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO fas.schemes (name, description, employment_status, martial_status, children_status, benefits, create_time, update_time)
VALUES ("Financial Assistance for Single Parent Scheme", "Financial assistance for single parent with children", 0, 1, 1, '[{"name": "Monthly Financial Assistance", "amount": 1000.00}]', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

CREATE TABLE IF NOT EXISTS `fas`.`applicants` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `first_name` varchar(128) NOT NULL,
    `last_name` varchar(128) NOT NULL,
    `nric` varchar(16) NOT NULL,
    `employment_status` int NOT NULL,
    `martial_status` int NOT NULL,
    `sex` int NOT NULL,
    `date_of_birth` varchar(16) NOT NULL,
    `household` JSON NOT NULL,
    `create_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    `update_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    PRIMARY KEY (`id`),
    UNIQUE (`nric`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO fas.applicants (first_name, last_name, nric, employment_status, martial_status, sex, date_of_birth, household, create_time, update_time)
VALUES ("Catherine", "Lim", "T0082722Z", 2, 1, 1, "2000-05-19", '[]', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO fas.applicants (first_name, last_name, nric, employment_status, martial_status, sex, date_of_birth, household, create_time, update_time)
VALUES ("Vivian", "Seow", "S8812097C", 1, 2, 1, "1988-03-22",
        '[{"sex": 0, "nric": "T2082722F", "relation": "son", "last_name": "Lee", "first_name": "Vincent", "date_of_birth": "2020-10-09"}]', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());
INSERT INTO fas.applicants (first_name, last_name, nric, employment_status, martial_status, sex, date_of_birth, household, create_time, update_time)
VALUES ("Joshnson", "Tan", "S7638590D", 2, 3, 0, "1988-03-22",
        '[{"sex": 0, "nric": "T1637581G", "relation": "son", "last_name": "Tan", "first_name": "Johnny", "date_of_birth": "2016-09-12"},{"sex": 1, "nric": "T0498374I", "relation": "daughter", "last_name": "Tan", "first_name": "Jenny", "date_of_birth": "2004-04-28"}]', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());

CREATE TABLE IF NOT EXISTS `fas`.`applications` (
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `applicant_id` int unsigned NOT NULL,
    `scheme_id` int unsigned NOT NULL,
    `status` int unsigned NOT NULL,
    `create_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    `update_time` int unsigned NOT NULL DEFAULT (UNIX_TIMESTAMP()),
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO fas.applications (applicant_id, scheme_id, status, create_time, update_time)
VALUES (1, 1, 0, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());