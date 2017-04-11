package postgres_test

import (
	"testing"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/postgres"
	"github.com/stretchr/testify/assert"
)

var db *pgx.Conn

type environment struct {
	PostgresURL string `env:"POSTGRES_URL,required"`
}

func TestSetup(t *testing.T) {
	var ev environment
	err := env.Parse(&ev)
	if err != nil {
		t.Fatal("unable to parse env variables")
	}

	config, err := pgx.ParseURI(ev.PostgresURL)
	if err != nil {
		t.Fatal("postgres uri invalid")
	}

	db, err = pgx.Connect(config)
	if err != nil {
		t.Fatal("unable to connect to postgres")
	}
}

func TestJackTables(t *testing.T) {
	tables, err := postgres.Tables(db, "jack")
	if err != nil {
		t.Fatal(err)
	}

	var tableTests = []struct {
		expected string
	}{
		{"reports"},
		{"standups"},
		{"standups_teammates"},
		{"teammates"},
		{"teams"},
	}

	for i, tt := range tableTests {
		actual := tables[i].TableName
		assert.Equal(t, tt.expected, actual)
	}
}

func TestJackColumns(t *testing.T) {
	columns, err := postgres.Columns(db, "jack", "teams")
	if err != nil {
		t.Fatal(err)
	}

	t1 := "gen_random_uuid()"
	t7 := "'{}'::text[]"
	t10 := "true"
	t11 := "4"
	t12 := "1"
	t13 := "(timezone('utc'::text, now()) + '14 days'::interval)"
	t14 := "timezone('utc'::text, now())"
	t15 := "timezone('utc'::text, now())"

	var columnTests = []*postgres.Column{
		{1, "id", "uuid", true, &t1, true},
		{2, "slack_team_id", "text", true, nil, false},
		{3, "slack_team_access_token", "text", true, nil, false},
		{4, "slack_bot_access_token", "text", true, nil, false},
		{5, "slack_bot_id", "text", true, nil, false},
		{6, "team_name", "text", true, nil, false},
		{7, "scope", "text[]", true, &t7, false},
		{8, "email", "text", false, nil, false},
		{9, "stripe_id", "text", false, nil, false},
		{10, "active", "boolean", true, &t10, false},
		{11, "free_teammates", "integer", true, &t11, false},
		{12, "cost_per_user", "integer", true, &t12, false},
		{13, "trial_ends", "timestamp with time zone", true, &t13, false},
		{14, "created_at", "timestamp with time zone", false, &t14, false},
		{15, "updated_at", "timestamp with time zone", false, &t15, false},
	}

	// for _, column := range columns {
	// 	fmt.Println("%+v", column)
	// }

	for i, tt := range columnTests {
		actual := columns[i]
		assert.Equal(t, tt, actual)
	}
}

func TestJackForeignKeys(t *testing.T) {
	fks, err := postgres.ForeignKeys(db, "jack", "reports")
	if err != nil {
		t.Fatal(err)
	}

	var fkTests = []*postgres.ForeignKey{
		// ForeignKeyName ColumnName RefIndexName RefTableName RefColumnName KeyID SeqNo OnUpdate OnDelete Match
		{"reports_standup_id_fkey", "standup_id", "standups_pkey", "standups", "id", 0, 0, "", "", ""},
		{"reports_user_id_fkey", "user_id", "teammates_pkey", "teammates", "id", 0, 0, "", "", ""},
	}

	// for _, fk := range fks {
	// 	fmt.Println("%+v", fk)
	// }

	for i, tt := range fkTests {
		actual := fks[i]
		assert.Equal(t, tt, actual)
	}
}

func TestJackIndexes(t *testing.T) {
	indexes, err := postgres.Indexes(db, "jack", "standups_teammates")
	if err != nil {
		t.Fatal(err)
	}

	var indexTests = []*postgres.Index{
		// IndexName IsUnique IsPrimary SeqNo Origin IsPartial
		{"standups_teammates_teammate_id_standup_id_key", true, false, 0, "", false},
	}

	for i, tt := range indexTests {
		actual := indexes[i]
		assert.Equal(t, tt, actual)
	}
}

func TestJackIndexColumns(t *testing.T) {
	columns, err := postgres.IndexColumns(db, "jack", "standups_teammates", "standups_teammates_teammate_id_standup_id_key")
	if err != nil {
		t.Fatal(err)
	}

	// for _, column := range columns {
	// 	fmt.Println("%+v", column)
	// }

	// i think the cid or seqno refers to column position
	// so unique order is correct, despite 2 being after 1
	var indexTests = []*postgres.IndexColumn{
		{2, 2, "teammate_id", "uuid"},
		{1, 1, "standup_id", "uuid"},
	}

	for i, tt := range indexTests {
		actual := columns[i]
		assert.Equal(t, tt, actual)
	}
}
