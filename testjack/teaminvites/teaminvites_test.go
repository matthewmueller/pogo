package teaminvites_test

import (
	"testing"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/testjack"
	"github.com/matthewmueller/pogo/testjack/teaminvites"
	"github.com/stretchr/testify/assert"
)

func DB(t *testing.T) (testjack.DB, func()) {
	config, err := pgx.ParseURI("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	db, err := pgx.Connect(config)
	if err != nil {
		t.Fatal(err)
	}

	return db, func() {
		if e := db.Close(); e != nil {
			t.Fatal(e)
		}
	}
}

func TestArrayInsert(t *testing.T) {
	db, close := DB(t)
	defer close()

	ti := teaminvites.New().Emails([]string{"m@gmail.com", "f@gmail.com", "g@gmail.com"})

	tis, err := teaminvites.Insert(db, ti)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "m@gmail.com", (*tis.GetEmails())[0])
	assert.Equal(t, "f@gmail.com", (*tis.GetEmails())[1])
	assert.Equal(t, "g@gmail.com", (*tis.GetEmails())[2])
}
