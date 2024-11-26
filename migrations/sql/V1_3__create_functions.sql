CREATE FUNCTION find_id(table_name TEXT, field TEXT, want TEXT)
RETURNS INT
AS $$
DECLARE res_id INT;
BEGIN
    EXECUTE format('SELECT id FROM %I WHERE %I = $1', table_name, field)
    INTO res_id
    USING want;

    RETURN res_id;
END;
$$ LANGUAGE plpgsql;


CREATE FUNCTION get_time_control_id(want_minutes INT, want_increment INT)
RETURNS INT
LANGUAGE sql
AS $$
    SELECT id
    FROM time_control
    WHERE time_minutes = want_minutes AND increment = want_increment;
$$;