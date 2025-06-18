CREATE TABLE request_body_tasks (
    id SERIAL PRIMARY KEY,
    task VARCHAR(255) NOT NULL,
    Accomplishment BOOLEAN DEFAULT FALSE
);