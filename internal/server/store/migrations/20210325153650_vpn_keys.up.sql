-- formated https://sqlformat.darold.net/

CREATE OR REPLACE FUNCTION v3.create_vpn_key (_key text, _validity integer, _used boolean, _user_id uuid, _txhash text, _with_pro boolean)
    RETURNS text
    AS $$
    INSERT INTO vpn_keys (key, validity, used, user_id, txhash, with_pro)
        VALUES (_key, _validity, _used, _user_id, _txhash, _with_pro)
    RETURNING key;

$$
LANGUAGE SQL;
