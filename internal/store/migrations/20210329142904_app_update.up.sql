-- formated https://sqlformat.darold.net/

CREATE TABLE v3.app_update (
    version int PRIMARY KEY,
    url text NOT NULL,
    force boolean NOT NULL,
    checksum text NOT NULL,
    changelog text NOT NULL
);

CREATE OR REPLACE FUNCTION v3.get_updates ()
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

CREATE OR REPLACE FUNCTION v3.create_update (_version int, _url text, _force boolean, _checksum text, _changelog text)
    RETURNS int
    AS $$
    INSERT INTO app_update (version, url, force, checksum, changelog)
        VALUES (_version, _url, _force, _checksum, _changelog)
    RETURNING
        version;

$$
LANGUAGE SQL;

