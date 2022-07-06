CREATE DOMAIN uint64 AS numeric(20,0) NOT NULL CHECK(0 <= VALUE AND VALUE <= 18446744073709551615);

CREATE TABLE uint_table (
    id          BIGSERIAL       PRIMARY KEY,
    num         uint64
);