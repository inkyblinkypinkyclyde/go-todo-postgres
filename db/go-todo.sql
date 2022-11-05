DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS projects;
DROP TABLE IF EXISTS users;


CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE projects(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    completed BOOLEAN,
    user_id INT REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    description VARCHAR(255),
    completed BOOLEAN,
    project_id INT REFERENCES projects(id) ON DELETE CASCADE
);

