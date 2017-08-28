package team

import uuid "github.com/satori/go.uuid"

type Team struct {
	ID          *uuid.UUID
	SlackTeamID *string
}


