# pgext : jackc/pgx and jackc/pgtype extension

This module provides a simple concrete struct, ``pgext.Puint``, which has the expected pgx encoding and decoding methods. While it is true that postgres does not have a uint64 type, this module asserts that ``pgext.Puint`` will be stored and loaded from a ``numeric(20,0)`` column type. This can be asserted further in Postgres by using the query ``CREATE DOMAIN uint64 AS numeric(20,0);``.
