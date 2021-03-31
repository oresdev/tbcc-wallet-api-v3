-- formated https://sqlformat.darold.net/

CREATE OR REPLACE FUNCTION v3.get_users ()
    RETURNS json
    AS $$
    SELECT
        array_to_json(array_agg(row_to_json(t))) AS arrayToJsonarrayAggrowToJsont
    FROM (
        SELECT
            id,
            useraddress,
            accounttype,
            smartcard
        FROM
            users) t;

$$
LANGUAGE SQL;

CREATE OR REPLACE FUNCTION v3.get_user_by_id (_user_id uuid)
    RETURNS json
    AS $$
    SELECT
        row_to_json(t) AS rowToJsont
    FROM (
        SELECT
            id,
            useraddress,
            accounttype,
            smartcard
        FROM
            users
        WHERE
            users.id = _user_id) t;

$$
LANGUAGE SQL;

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
            COALESCE(jsonb_agg(vpn_keys) FILTER (WHERE vpn_keys.id IS NOT NULL), '[]') AS vpn_keys,
            COALESCE(jsonb_agg(app_config) FILTER (WHERE app_config.config_group IS NOT NULL), '{}') AS app_config
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

-- Migrate user data from depricated database (public scheme)
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
            COALESCE(jsonb_agg(vpn_keys) FILTER (WHERE vpn_keys.id IS NOT NULL), '[]') AS vpn_keys,
            COALESCE(jsonb_agg(app_config) FILTER (WHERE app_config.config_group IS NOT NULL), '{}') AS app_config
        FROM
            users
        LEFT JOIN vpn_keys ON vpn_keys.user_id = users.id
        LEFT JOIN app_config ON 1 = 1
    WHERE
        EXISTS (
            SELECT
                1
            FROM
                users u
            WHERE
                u.useraddress && ARRAY[_addresses])
        GROUP BY
            users.id) t;

$$
LANGUAGE SQL;

CREATE OR REPLACE FUNCTION v3.update_user (_uuid uuid, _address text)
    RETURNS json
    AS $$
    UPDATE
        users
    SET
        useraddress = array_append(useraddress, _address)
    WHERE
        id = _uuid
    RETURNING
        row_to_json(users.*);

$$
LANGUAGE SQL;

CREATE OR REPLACE FUNCTION v3.create_user (_useraddress text[], _accounttype text, _smartcard boolean)
    RETURNS json
    AS $$
    INSERT INTO users (useraddress, accounttype, smartcard)
        VALUES (_useraddress, _accounttype, _smartcard)
    RETURNING
        row_to_json(users.*);

$$
LANGUAGE SQL;
