drop view if exists v_user_numb;
drop view if exists v_file_numb;
drop table statistics_table;
drop function if exists fn_add_stat();
drop trigger if exists tr_upd_stat on user_table;
drop trigger if exists tr_upd_stat2 on file_table;