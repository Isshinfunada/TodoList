-- Drop trigger for todos table
DROP TRIGGER IF EXISTS update_todos_updated_at ON todos;

-- Drop trigger function
DROP FUNCTION IF EXISTS update_updated_at_column;

-- Drop todos table
DROP TABLE IF EXISTS todos;

-- Drop users table
DROP TABLE IF EXISTS users;