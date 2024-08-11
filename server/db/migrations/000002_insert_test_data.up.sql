-- Insert test data into users table
INSERT INTO users (username, email, password) VALUES
('testuser1', 'testuser1@example.com', 'password1'),
('testuser2', 'testuser2@example.com', 'password2'),
('testuser3', 'testuser3@example.com', 'password3');

-- Insert test data into todos table
INSERT INTO todos (user_id, text, status) VALUES
(1, 'Test TODO 1 for user 1', 'pending'),
(1, 'Test TODO 2 for user 1', 'completed'),
(2, 'Test TODO 1 for user 2', 'pending'),
(3, 'Test TODO 1 for user 3', 'pending');