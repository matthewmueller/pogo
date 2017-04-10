package tempo

// GENERATED BY POGO. DO NOT EDIT.

import (
 "errors"
 "strings"
)

// Update the Task by the Primary Key
func (t *Tasks) Update(id *string, tt *Task) (task Task, err error) {
 fields := t.getFields(tt)

 // first check if we have the primary key
 if id == nil {
  return task, errors.New(`primary key "id" must be non-nil`)
 }

 // don't update the primary key
 delete(fields, "id")

 // prepare the slices
 c, i, v := querySlices(fields, 1)

 // sql query
 sqlstr := `UPDATE public.tasks SET (` +
  strings.Join(c, ", ") + `) = (` +
  strings.Join(i, ", ") + `)
		WHERE id = $1
		RETURNING id, key, target, rate, offset, timezone, rate_type, rate_options, active, target_type, target_options, user, meta, refreshed_at, created_at, updated_at`

 // run query
 values := append([]interface{}{tt.ID}, v...)
 DBLog(sqlstr, values...)

 row := t.DB.QueryRow(sqlstr, values...)
 err = row.Scan(&task.ID, &task.Key, &task.Target, &task.Rate, &task.Offset, &task.Timezone, &task.RateType, &task.RateOptions, &task.Active, &task.TargetType, &task.TargetOptions, &task.User, &task.Meta, &task.RefreshedAt, &task.CreatedAt, &task.UpdatedAt)
 if err != nil {
  return task, err
 }

 return task, nil
}
