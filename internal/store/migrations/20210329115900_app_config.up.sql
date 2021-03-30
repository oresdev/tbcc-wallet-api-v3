-- formated https://sqlformat.darold.net/

CREATE TABLE v3.app_config (
    config_group text NOT NULL PRIMARY KEY,
    value json NOT NULL
);

CREATE OR REPLACE FUNCTION v3.get_user_ext (_user_id uuid)
    RETURNS json
    AS $$
    SELECT
        row_to_json(t) AS rowToJsont
    FROM (
        SELECT
            users.id,
            users.useraddress,
            users.accounttype,
            users.smartcard,
            jsonb_agg(to_jsonb (vpn_keys)) AS vpn_keys,
            jsonb_agg(to_jsonb (app_config)) AS app_config
        FROM
            users
        LEFT JOIN vpn_keys ON vpn_keys.user_id = users.id
        LEFT JOIN app_config ON 1 = 1
    WHERE
        users.id = _user_id
    GROUP BY
        users.id) t;

$$
LANGUAGE SQL;

CREATE OR REPLACE FUNCTION v3.get_user_ext_by_address (_address text)
    RETURNS json
    AS $$
    SELECT
        row_to_json(t) AS rowToJsont
    FROM (
        SELECT
            users.id,
            users.useraddress,
            users.accounttype,
            users.smartcard,
            jsonb_agg(to_jsonb (vpn_keys)) AS vpn_keys,
            jsonb_agg(to_jsonb (app_config)) AS app_config
        FROM
            users
        LEFT JOIN vpn_keys ON vpn_keys.user_id = users.id
        LEFT JOIN app_config ON 1 = 1
    WHERE
        _address = ANY (users.useraddress)
    GROUP BY
        users.id) t;

$$
LANGUAGE SQL;

-- migrate users

CREATE OR REPLACE FUNCTION check_exists_user_and_return_ext (_addresses text[])
    RETURNS json
    AS $$
        SELECT
        row_to_json(t) AS rowToJsont
    FROM (
        SELECT
            users.id,
            users.useraddress,
            users.accounttype,
            users.smartcard,
            jsonb_agg(to_jsonb (vpn_keys)) AS vpn_keys,
            jsonb_agg(to_jsonb (app_config)) AS app_config
        FROM
            users
        LEFT JOIN vpn_keys ON vpn_keys.user_id = users.id
        LEFT JOIN app_config ON 1 = 1
    WHERE
        EXISTS (
        SELECT 1
        FROM
            users u
        WHERE
            u.useraddress && ARRAY[_addresses])
    GROUP BY
        users.id) t;

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