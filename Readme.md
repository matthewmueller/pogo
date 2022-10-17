# Pogo

Type-safe Database Client for Go. Supports PostgreSQL and SQLite.

The code is stable and has been in production for many years now. This repository won't change much in the future.

This code along with future updates will eventually be folded into https://github.com/livebud/bud.

## Usage

### Introspect your database

```sh
# connect to a postgres database and build the client in pogo/
pogo --db $(POSTGRES_URL) --schema public --dir ./pogo
```

### Using the Generated Client

```go
pgconfig, err := pgx.ParseURI(env.DatabaseURL)
if err != nil {
  return err
}
db, err := pgx.Connect(pgconfig)
if err != nil {
  return err
}

users, err := user.FindMany(db,
  user.NewFilter().Email("alice@livebud.com"),
  user.NewOrder().CreatedAt(user.DESC),
)
```

Check out [the tests](https://github.com/matthewmueller/pogo/blob/master/internal/postgres/generate_test.go) for more usage examples.

### License

MIT
