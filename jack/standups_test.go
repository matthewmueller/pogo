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

func standupsDB(t *testing.T) jack.DB {
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

func TestStandupsInsert(t *testing.T) {
	// setup the model
	model := jack.New(standupsDB(t))

	// random values
	_id := uuid.NewV4()
	_name := "9571c8bd-8d85-4aa1-b923-b380f72f30d9"
	_slackchannelid := "8948d5fc-8f28-4b68-a253-5ae4041bffa5"
	_time := "493299ce-c061-4889-bada-b0a46fc5fa85"
	_timezone := "97dd2bd2-e865-4633-ae74-ac6365626c5c"
	_questions := map[string]interface{}{}
	_teamid := uuid.NewV4()
	_createdat := time.Now()
	_updatedat := time.Now()

	// struct
	standup1 := jack.Standup{
		ID:             &_id,
		Name:           &_name,
		SlackChannelID: &_slackchannelid,
		Time:           &_time,
		Timezone:       &_timezone,
		Questions:      &_questions,
		TeamID:         &_teamid,
		CreatedAt:      &_createdat,
		UpdatedAt:      &_updatedat,
	}

	standup2, err := model.Standup.Insert(standup1)
	if err != nil {
		t.Fatal(err)
	}

	// assertions
	assert.Equal(t, _id, *standup2.ID)
	assert.Equal(t, _name, *standup2.Name)
	assert.Equal(t, _slackchannelid, *standup2.SlackChannelID)
	assert.Equal(t, _time, *standup2.Time)
	assert.Equal(t, _timezone, *standup2.Timezone)
	assert.Equal(t, _questions, *standup2.Questions)
	assert.Equal(t, _teamid, *standup2.TeamID)
	assert.Equal(t, _createdat, *standup2.CreatedAt)
	assert.Equal(t, _updatedat, *standup2.UpdatedAt)

	// cleanup
	if e := model.Standup.Delete(&_id); e != nil {
		t.Fatal(e)
	}
}

func TestStandupsUpdate(t *testing.T) {
	// setup the model
	model := jack.New(standupsDB(t))

	// random values
	_id := uuid.NewV4()
	_name := "75443fec-654b-40b7-ad8c-148ac7a25364"
	_slackchannelid := "a421f102-13c9-4fae-a806-b65f4745d216"
	_time := "9ec06052-853f-4f47-8af0-537eada30fb8"
	_timezone := "73687cb2-ceba-41ae-9be5-f95ebd796433"
	_questions := map[string]interface{}{}
	_teamid := uuid.NewV4()
	_createdat := time.Now()
	_updatedat := time.Now()

	// struct
	standup1 := jack.Standup{
		ID:             &_id,
		Name:           &_name,
		SlackChannelID: &_slackchannelid,
		Time:           &_time,
		Timezone:       &_timezone,
		Questions:      &_questions,
		TeamID:         &_teamid,
		CreatedAt:      &_createdat,
		UpdatedAt:      &_updatedat,
	}

	standup2, err := model.Standup.Insert(standup1)
	if err != nil {
		t.Fatal(err)
	}

	// random values
	_id2 := uuid.NewV4()
	_name2 := "297d5d08-bde6-4243-b5ad-ececd52885be"
	_slackchannelid2 := "ebc254f7-b2f8-44f1-a427-fe1e59ed79a9"
	_time2 := "59196dc1-936f-404c-acfa-39c12d9d140b"
	_timezone2 := "9cdd2115-d593-4dc6-bd6f-60f227c93efa"
	_questions2 := map[string]interface{}{}
	_teamid2 := uuid.NewV4()
	_createdat2 := time.Now()
	_updatedat2 := time.Now()

	// random values
	standup2.ID = &_id2
	standup2.Name = &_name2
	standup2.SlackChannelID = &_slackchannelid2
	standup2.Time = &_time2
	standup2.Timezone = &_timezone2
	standup2.Questions = &_questions2
	standup2.TeamID = &_teamid2
	standup2.CreatedAt = &_createdat2
	standup2.UpdatedAt = &_updatedat2

	standup3, err := model.Standup.Update(*standup2, &_id)
	if err != nil {
		t.Fatal(err)
	}

	// assertions
	assert.Equal(t, _id2, *standup3.ID)
	assert.Equal(t, _name2, *standup3.Name)
	assert.Equal(t, _slackchannelid2, *standup3.SlackChannelID)
	assert.Equal(t, _time2, *standup3.Time)
	assert.Equal(t, _timezone2, *standup3.Timezone)
	assert.Equal(t, _questions2, *standup3.Questions)
	assert.Equal(t, _teamid2, *standup3.TeamID)
	assert.Equal(t, _createdat2, *standup3.CreatedAt)
	assert.Equal(t, _updatedat2, *standup3.UpdatedAt)

	// cleanup
	if e := model.Standup.Delete(&_id); e != nil {
		t.Fatal(e)
	}
}
