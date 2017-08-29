{{/*************************************************************************/}}
{{/* Variables */}}
{{/*************************************************************************/}}



{{/*************************************************************************/}}
{{/* Our Package */}}
{{/*************************************************************************/}}

package {{ .Settings.Package }}

{{/*************************************************************************/}}
{{/* Imports */}}
{{/*************************************************************************/}}

import uuid "github.com/satori/go.uuid"

{{/*************************************************************************/}}
{{/* Pogo marker */}}
{{/*************************************************************************/}}

// GENERATED BY POGO. DO NOT EDIT.

{{/*************************************************************************/}}
{{/* satori.UUID */}}
{{/*************************************************************************/}}

// DecodeUUID UUID to string
func DecodeUUID(id uuid.UUID) *string {
	uid := id.String()
	return &uid
}

// EncodeUUID string to UUID
func EncodeUUID(id *string) *uuid.UUID {
	if id == nil {
		return nil
	}

	uid, err := uuid.FromString(*id)
	if err != nil {
		return nil
	}

	return &uid
}