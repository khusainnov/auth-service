CREATE OR REPLACE FUNCTION fn_delete_files ()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $$
   BEGIN
       DELETE FROM file_table WHERE username = old.username;

       RETURN old;
   END;
$$;

CREATE OR REPLACE TRIGGER tr_delete_files BEFORE DELETE ON user_table
    FOR EACH ROW EXECUTE FUNCTION fn_delete_files();