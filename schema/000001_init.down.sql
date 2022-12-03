DROP TABLE if exists role_table CASCADE;
DROP TABLE if exists user_table CASCADE;
DROP TABLE if exists student_table CASCADE;
DROP TABLE if exists admin_table CASCADE;
DROP TABLE if exists file_table CASCADE;
DROP FUNCTION IF EXISTS insert_user_by_role();
DROP TRIGGER IF EXISTS tr_sort_by_id ON user_table;
