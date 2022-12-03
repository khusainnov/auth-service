CREATE TABLE statistics_table (
    stat_id smallserial not null primary key,
    user_numb int,
    file_numb int
);

insert into statistics_table (user_numb, file_numb) VALUES (0, 0);

CREATE OR REPLACE VIEW v_user_numb AS
    SELECT count(ut.username)
        FROM user_table ut;

CREATE OR REPLACE VIEW v_file_numb AS
    SELECT count(ft.file_id)
        FROM file_table ft;

CREATE OR REPLACE FUNCTION fn_add_stat()
RETURNS TRIGGER
LANGUAGE plpgsql
AS $$
    BEGIN
--         INSERT INTO statistics_table (user_numb, file_numb) VALUES ((SELECT * FROM v_user_numb), (SELECT * FROM v_file_numb));
        UPDATE statistics_table SET user_numb = (SELECT * FROM v_user_numb), file_numb = (SELECT * FROM v_file_numb) WHERE stat_id = 1;
        RETURN old;
    END;
$$;


CREATE OR REPLACE TRIGGER tr_upd_stat AFTER INSERT ON user_table
FOR EACH ROW EXECUTE FUNCTION fn_add_stat();

CREATE OR REPLACE TRIGGER tr_upd_stat2 AFTER INSERT ON file_table
    FOR EACH ROW EXECUTE FUNCTION fn_add_stat();