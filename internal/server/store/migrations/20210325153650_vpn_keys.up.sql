-- Returns json data from vpn_keys
CREATE OR REPLACE FUNCTION v3.vpn_keys_update_by_uuid (_uuid uuid, _txhash text)
    RETURNS json
    AS $$
    UPDATE
        vpn_keys
    SET
        user_id = _uuid,
        used = TRUE,
        timestamp = CURRENT_TIMESTAMP AT TIME ZONE 'UTC',
        txhash = _txhash
    WHERE
        id = (
            SELECT
                id
            FROM
                vpn_keys
            WHERE
                used IS NULL
            LIMIT 1)
RETURNING
    row_to_json(vpn_keys.*);

$$
LANGUAGE SQL;