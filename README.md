# pgext : pgx and pgtype extension

This module provides a simple concrete struct, ``pgext.Puint``, which has the expected pgx encoding and decoding methods. 

This concrete struct is simply a ``uint64``, however it includes the pgx encoding and decoding methods to be stored in a numeric postgres field.

While it is true that postgres does not have a uint64 type, this module asserts that ``pgext.Puint`` will be stored and loaded from a ``numeric(20,0)`` column type. The module will work with this definition alone, but it is recommended that a new domain is created in the postgres database. For example,

```CREATE DOMAIN uint64 AS numeric(20,0) NOT NULL CHECK(0 <= VALUE AND VALUE <= 18446744073709551615);```