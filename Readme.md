# Pogo

```sh
# connect to the tempo database and write to tempo/
pogo -db postgres://localhost:5432/tempo?sslmode=disable -path tempo
```

TODO:

  - [x] migrate to pgx@3
  - [x] warn about skipped files, but don't error out
  - ~~[x] make api more like cdp~~
    - Need to use this more to figure out if there are any API improvements to make
  - [x] Test out API. Probably return non-pointers.
  - [x] Finish generated tests using a pogo sql script (not jack)
  - [x] Create many-to-many generated tests

  - [ ] Upsert => InsertOrUpdate(...) & InsertOrNothing(...)
  - [ ] Where condition
  - [ ] Finish testjack tests
    - [ ] teams
    - [ ] standups_teammates
    - [ ] UNIQUE(a, b) multi-column index tests

LATER:

  - [ ] Investigate type aliases to simplify accessor coercion  
  - [ ] Implement fake data generator
