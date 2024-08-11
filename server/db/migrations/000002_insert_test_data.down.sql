-- Remove test data from todos table
DELETE FROM todos WHERE text LIKE 'Test TODO%';

-- Remove test data from users table
DELETE FROM users WHERE email LIKE 'testuser%@example.com';