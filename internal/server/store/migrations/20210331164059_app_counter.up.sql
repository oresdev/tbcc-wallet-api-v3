-- Returns void
CREATE OR REPLACE FUNCTION v3.app_counter_update_row (_version integer)
    RETURNS VOID
    AS $$
    UPDATE
        app_counter
    SET
        count = count + 1
    WHERE
        version = _version;

$$
LANGUAGE SQL;