package jack

// GENERATED BY POGO. DO NOT EDIT.

// FindMany find many Conversations by a condition
func (c *Conversations) FindMany(condition string, params ...interface{}) (conversations []Conversation, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "user_id", "topic", "context", "state", "created", "updated"
    FROM public.conversations
    WHERE ` + condition

	DBLog(sqlstr, params...)
	rows, err := c.DB.Query(sqlstr, params...)
	if err != nil {
		return conversations, err
	}
	defer rows.Close()

	for rows.Next() {
		conversation := Conversation{}
		err = rows.Scan(&conversation.ID, &conversation.UserID, &conversation.Topic, &conversation.Context, &conversation.State, &conversation.Created, &conversation.Updated)
		if err != nil {
			return conversations, err
		}
		conversations = append(conversations, conversation)
	}

	if rows.Err() != nil {
		return conversations, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(conversations) == 0 {
		return make([]Conversation, 0), nil
	}

	return conversations, nil
}