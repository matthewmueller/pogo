package testutil

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/jackc/pgx"
	gen "github.com/matthewmueller/go-gen"
	text "github.com/matthewmueller/go-text"
	"github.com/matthewmueller/pogo"
)

// Connect fn
func Connect(t testing.TB, url string) (*pgx.Conn, func()) {
	cfg, err := pgx.ParseConnectionString(url)
	if err != nil {
		t.Fatal(err)
	}

	conn, err := pgx.Connect(cfg)
	if err != nil {
		t.Fatal(err)
	}

	return conn, func() {
		conn.Close()
	}
}

// Exec in a transaction
func Exec(t testing.TB, conn *pgx.Conn, sql string) {
	tx, err := conn.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	if _, err := tx.Exec(sql); err != nil {
		t.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}

// Build pogo
func Build(t testing.TB, url, schema, dir string) func() {
	p := pogo.Pogo{
		URL:    url,
		Schema: schema,
		Output: dir,
	}

	if err := p.Run(); err != nil {
		t.Fatal(err)
	}

	return func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatal(err)
		}
	}
}

// Run fn
func Run(t testing.TB, name, main string) (string, string, func()) {
	tmpdir := filepath.Join("_tmp", text.Snake(name))
	path := filepath.Join(tmpdir, "main.go")

	code, err := gen.Format(main)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.MkdirAll(tmpdir, 0755); err != nil {
		t.Fatal(err)
	}

	if err := ioutil.WriteFile(path, []byte(code), 0644); err != nil {
		t.Fatal(err)
	}

	gobin, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}

	// go run
	cmd := exec.Command(gobin, "run", path)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	// run the command
	if err := cmd.Run(); err != nil {
		t.Fatal(stderr.String())
	}

	return stdout.String(), stderr.String(), func() {
		if err := os.RemoveAll(tmpdir); err != nil {
			t.Fatal(err)
		}
	}
}
