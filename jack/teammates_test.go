package jack_test

import (
	"testing"
	"time"

	"github.com/jackc/pgx"
	"github.com/matthewmueller/pogo/jack"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

// GENERATED BY POGO. DO NOT EDIT.

func teammatesDB(t *testing.T) jack.DB {
	config, err := pgx.ParseURI("postgres://localhost:5432/pogo?sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}

	db, err := pgx.Connect(config)
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestTeammatesInsert(t *testing.T) {
	// setup the model
	model := jack.New(teammatesDB(t))

	// random values
	_id := uuid.NewV4()
	_slackid := "4316073c-3a58-434a-8b1e-4ffaa18b735b"
	_username := "4c2e07cb-329d-4ff6-ac5e-97c458f5c984"
	_firstname := "b0ff5295-5cd2-40bd-ab05-f1d5d8610deb"
	_lastname := "81724b03-ae52-4ad1-a24e-0d55247ef4de"
	_email := "72c87ec9-65d0-4e5c-9abc-d0025680db67"
	_avatar := "cdde01c2-43e9-4cfe-b772-4356c9414bfa"
	_timezone := "f2fe3e27-9f2d-4f18-95fd-445440975715"
	_createdat := time.Now()
	_updatedat := time.Now()

	// struct
	teammate1 := jack.Teammate{
		ID:        &_id,
		SlackID:   &_slackid,
		Username:  &_username,
		FirstName: &_firstname,
		LastName:  &_lastname,
		Email:     &_email,
		Avatar:    &_avatar,
		Timezone:  &_timezone,
		CreatedAt: &_createdat,
		UpdatedAt: &_updatedat,
	}

	teammate2, err := model.Teammate.Insert(teammate1)
	if err != nil {
		t.Fatal(err)
	}

	// assertions
	assert.Equal(t, _id, *teammate2.ID)
	assert.Equal(t, _slackid, *teammate2.SlackID)
	assert.Equal(t, _username, *teammate2.Username)
	assert.Equal(t, _firstname, *teammate2.FirstName)
	assert.Equal(t, _lastname, *teammate2.LastName)
	assert.Equal(t, _email, *teammate2.Email)
	assert.Equal(t, _avatar, *teammate2.Avatar)
	assert.Equal(t, _timezone, *teammate2.Timezone)
	assert.Equal(t, _createdat, *teammate2.CreatedAt)
	assert.Equal(t, _updatedat, *teammate2.UpdatedAt)

	// cleanup
	if e := model.Teammate.Delete(&_id); e != nil {
		t.Fatal(e)
	}
}

func TestTeammatesUpdate(t *testing.T) {
	// setup the model
	model := jack.New(teammatesDB(t))

	// random values
	_id := uuid.NewV4()
	_slackid := "d553fbda-9902-43de-903b-e900970a6938"
	_username := "574dcef6-e4ce-4512-aeb5-d3f2ead21256"
	_firstname := "e3e859d3-cfe1-4876-8c2d-6395ce53cbd4"
	_lastname := "b4105298-9aa5-4d8e-b354-990ff3098aba"
	_email := "7b5b9842-9742-4f3f-b530-7869fd62ee27"
	_avatar := "e9f0f2d7-6cd0-4eb3-a445-b62deb3c5644"
	_timezone := "c52c5c21-312e-4389-9acb-cbc3acc5c353"
	_createdat := time.Now()
	_updatedat := time.Now()

	// struct
	teammate1 := jack.Teammate{
		ID:        &_id,
		SlackID:   &_slackid,
		Username:  &_username,
		FirstName: &_firstname,
		LastName:  &_lastname,
		Email:     &_email,
		Avatar:    &_avatar,
		Timezone:  &_timezone,
		CreatedAt: &_createdat,
		UpdatedAt: &_updatedat,
	}

	teammate2, err := model.Teammate.Insert(teammate1)
	if err != nil {
		t.Fatal(err)
	}

	// random values
	_id2 := uuid.NewV4()
	_slackid2 := "4634b984-e054-4ead-8ca3-ea8984b85241"
	_username2 := "ca934639-d608-4c28-a01b-6930730ac594"
	_firstname2 := "6613b42e-5715-4075-b009-9b4082f9be29"
	_lastname2 := "3091540b-dc26-4d04-944a-6c58bc71ffff"
	_email2 := "765ddfc2-ea2b-4b2f-9feb-e1b4102b0b61"
	_avatar2 := "8a07740e-b66f-46be-8f3d-c524738c53ac"
	_timezone2 := "9dd84ba0-7848-432e-a2d5-01eca7ad96a7"
	_createdat2 := time.Now()
	_updatedat2 := time.Now()

	// random values
	teammate2.ID = &_id2
	teammate2.SlackID = &_slackid2
	teammate2.Username = &_username2
	teammate2.FirstName = &_firstname2
	teammate2.LastName = &_lastname2
	teammate2.Email = &_email2
	teammate2.Avatar = &_avatar2
	teammate2.Timezone = &_timezone2
	teammate2.CreatedAt = &_createdat2
	teammate2.UpdatedAt = &_updatedat2

	teammate3, err := model.Teammate.Update(*teammate2, &_id)
	if err != nil {
		t.Fatal(err)
	}

	// assertions
	assert.Equal(t, _id2, *teammate3.ID)
	assert.Equal(t, _slackid2, *teammate3.SlackID)
	assert.Equal(t, _username2, *teammate3.Username)
	assert.Equal(t, _firstname2, *teammate3.FirstName)
	assert.Equal(t, _lastname2, *teammate3.LastName)
	assert.Equal(t, _email2, *teammate3.Email)
	assert.Equal(t, _avatar2, *teammate3.Avatar)
	assert.Equal(t, _timezone2, *teammate3.Timezone)
	assert.Equal(t, _createdat2, *teammate3.CreatedAt)
	assert.Equal(t, _updatedat2, *teammate3.UpdatedAt)

	// cleanup
	if e := model.Teammate.Delete(&_id); e != nil {
		t.Fatal(e)
	}
}