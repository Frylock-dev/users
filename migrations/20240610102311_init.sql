-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users
(
    id   SERIAL PRIMARY KEY,
    uuid VARCHAR(125) NOT NULL,
    UNIQUE (uuid)
);

CREATE TABLE IF NOT EXISTS users_info
(
    user_id             INT         NOT NULL UNIQUE REFERENCES users (id),
    phone               VARCHAR(18) NOT NULL UNIQUE,
    email               VARCHAR(55) NOT NULL UNIQUE,
    first_name          VARCHAR(55) NOT NULL,
    second_name         VARCHAR(55) NOT NULL,
    last_name           VARCHAR(55) NOT NULL,
    passport_number     INT,
    passport_code       INT,
    passport_issue_date TIMESTAMP,
    birthday            TIMESTAMP
);

CREATE TABLE IF NOT EXISTS roles
(
    id    SERIAL PRIMARY KEY,
    title VARCHAR(55) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_role
(
    user_id INT NOT NULL,
    role_id INT NOT NULL,
    UNIQUE (user_id, role_id),
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (role_id) REFERENCES roles (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_role;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users_info;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
