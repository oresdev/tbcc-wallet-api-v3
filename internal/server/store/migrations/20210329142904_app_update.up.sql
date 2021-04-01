-- Returns app_update json data
CREATE OR REPLACE FUNCTION v3.app_update_get_rows ()
    RETURNS json
    AS $$
    SELECT
        row_to_json(t) AS rowToJsont
    FROM (
        SELECT
            version,
            url,
            force,
            checksum,
            changelog
        FROM
            app_update) t;

$$
LANGUAGE SQL;

-- Returns version integer row
CREATE OR REPLACE FUNCTION v3.app_update_create_row (_version integer, _url text, _force boolean, _checksum text, _changelog text)
    RETURNS integer
    AS $$
    INSERT INTO app_update (version, url, force, checksum, changelog)
        VALUES (_version, _url, _force, _checksum, _changelog)
    RETURNING
        version;

$$
LANGUAGE SQL;