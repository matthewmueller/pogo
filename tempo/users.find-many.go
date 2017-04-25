package tempo

// GENERATED BY POGO. DO NOT EDIT.

// FindMany find many Users by a condition
func (u *Users) FindMany(condition string, params ...interface{}) (users []User, err error) {
	// sql select query, primary key provided by sequence
	sqlstr := `
    SELECT "id", "role", "email", "stripe_id", "active", "github_access_token", "scope", "free_events", "cost_per_event", "total_events", "paid_events", "token", "created_at", "updated_at", "free_tasks"
    FROM public.users
    WHERE ` + condition

	DBLog(sqlstr, params...)
	rows, err := u.DB.Query(sqlstr, params...)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Role, &user.Email, &user.StripeID, &user.Active, &user.GithubAccessToken, &user.Scope, &user.FreeEvents, &user.CostPerEvent, &user.TotalEvents, &user.PaidEvents, &user.Token, &user.CreatedAt, &user.UpdatedAt, &user.FreeTasks)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return users, rows.Err()
	}

	// ensure we return an empty array
	// rather than nil when we marshal
	if len(users) == 0 {
		return make([]User, 0), nil
	}

	return users, nil
}