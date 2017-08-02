# Pogo

```sh
# connect to the tempo database and write to tempo/
pogo -db postgres://localhost:5432/tempo?sslmode=disable -path tempo
```

TODO:

  - [ ] migrate to pgx@3
  - [ ] make api more like cdp
  - [ ] warn about skipped files, but don't error out
