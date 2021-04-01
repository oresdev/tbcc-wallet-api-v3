-- Returns app_config json data
CREATE OR REPLACE FUNCTION v3.app_config_get_rows ()
    RETURNS json
    AS $$
    SELECT
        row_to_json(t) AS rowToJsont
    FROM (
        SELECT
            config_group,
            value
        FROM
            app_config) t;

$$
LANGUAGE SQL;

-- Returns config_group string
CREATE OR REPLACE FUNCTION v3.app_config_create_row (_config_group text, _value json)
    RETURNS text
    AS $$
    INSERT INTO app_config (config_group, value)
        VALUES (_config_group, _value)
    RETURNING
        config_group;

$$
LANGUAGE SQL;