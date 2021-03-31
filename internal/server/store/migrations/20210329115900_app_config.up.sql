-- formated https://sqlformat.darold.net/

CREATE OR REPLACE FUNCTION v3.get_config ()
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

CREATE OR REPLACE FUNCTION v3.create_config (_config_group text, _value json)
    RETURNS text
    AS $$
    INSERT INTO app_config (config_group, value)
        VALUES (_config_group, _value)
    RETURNING
        config_group;

$$
LANGUAGE SQL;

