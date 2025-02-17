-- Create the tasks table
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Insert sample data
INSERT INTO tasks (title, description, status) 
VALUES 
    ('Task 1', 'Description for Task 1', 'new'),
    ('Task 2', 'Description for Task 2', 'in_progress'),
    ('Task 3', 'Description for Task 3', 'done');