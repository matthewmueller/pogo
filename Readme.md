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
  - [x] Upsert
  - [x] Where condition
  - [x] Finish testjack tests
    - [x] teams
    - [x] standups_teammates

LATER:

  - [ ] FindOrInsert (upsert with DO NOTHING)
  - [ ] Handle empties. e.g. `INSERT INTO jack.teammates default values`
  - [ ] InsertMany as Insert(db, team...)
  - [ ] UpdateManyByID as Update(db, team...)
  - [ ] DeleteManyByID as Delete(db, team...)
  - [ ] UNIQUE(a, b) multi-column index tests (email + teamname?)
  - [ ] Many-to-Many feature parity with regular models (UpdateBy, Upsert)
  - [ ] One-to-Many operations using fks: teams.DeleteManyByStandupID(...)
  - [ ] Investigate type aliases to simplify accessor coercion  
  - [ ] Implement fake data generator
