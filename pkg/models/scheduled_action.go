package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

//import "github.com/google/uuid"

//ScheduledAction represents an action to be taken in the future, persisted in the DB
type ScheduledAction struct {
	ID         int64  `json:"id" db:"id"`
	Status     string `json:"status" db:"scheduled_action_status"`
	ActionType string `json:"actionType" db:"scheduled_action_type"`
	// ActionData      string     `json:"actionData" db:"action_data"`
	ActionData      Dictionary `json:"actionData" db:"action_data"`
	CreatedTime     *time.Time `json:"createdTime" db:"created_at"`
	UpdatedTime     *time.Time `json:"updatedTime" db:"updated_at"`
	Recurring       bool       `json:"recurring" db:"recurring"`
	ScheduledTime   *time.Time `json:"scheduledTime" db:"scheduled_date"`
	DaysToNext      int16      `json:"daysToNext" db:"days_to_next_occurence"`
	LastSuccessTime *time.Time `json:"lastSuccessTime" db:"last_success_at"`
	CurrentAttempts int16      `json:"currentAttempts" db:"current_attempts"`
	MaxAttempts     int16      `json:"maxAttempts" db:"max_attempts"`
	Note            string     `json:"note" db:"note"`
	//Note            string     `json:"note db:"note"`
}

//Dictionary is a type to handle variable input of data from the DB. It is like a hash table.
type Dictionary map[string]interface{}

//Value let's the SQL driver transform the data to the dictionary type
func (d Dictionary) Value() (driver.Value, error) {

	j, err := json.Marshal(d)
	return j, err
}

//Scan is used by sql.scan to read the values from the DB
func (d *Dictionary) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion .([]byte) failed")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*d, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("type assertion .(map[string]interface{}) failed")
	}

	return nil
}
