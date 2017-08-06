# Pogo

```sh
# connect to the tempo database and write to tempo/
pogo -db postgres://localhost:5432/tempo?sslmode=disable -path tempo
```

TODO:

  - [x] migrate to pgx@3
  - [x] warn about skipped files, but don't error out
  - ~~[ ] make api more like cdp~~
    - Need to use this more to figure out if there are any API improvements to make
  - [ ] Test out API. Probably return non-pointers.
  - [ ] Finish generated tests using a pogo sql script (not jack)
  - [ ] Create many-to-many generated tests
  - [ ] Implement fake data generator
