-- formated https://sqlformat.darold.net/

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
    useraddress text[] NOT NULL,
    accounttype text NOT NULL,
    smartcard boolean NOT NULL
);

CREATE OR REPLACE FUNCTION get_users ()
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

CREATE OR REPLACE FUNCTION get_user_by_id (_user_id uuid)
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

CREATE OR REPLACE FUNCTION get_user_by_address (_address text)
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
            _address = ANY (users.useraddress)) t;

$$
LANGUAGE SQL;

CREATE OR REPLACE FUNCTION create_user (_useraddress text[], _accounttype text, _smartcard boolean)
    RETURNS uuid
    AS $$
    INSERT INTO users (useraddress, accounttype, smartcard)
        VALUES (_useraddress, _accounttype, _smartcard)
    RETURNING
        id;

$$
LANGUAGE SQL;
