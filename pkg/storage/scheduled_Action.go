package storage

import (
	"context"
	"database/sql"
	"errors"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/models"

	"go.uber.org/zap"
)

//FetchNextScheduledAction Gets the next scheduledAction from the DB if there is one to take action on
func (s *Store) FetchNextScheduledAction(ctx context.Context) (*models.ScheduledAction, error) {

	publicationEntry := models.ScheduledAction{}
	const SQL = ` 

	SELECT * FROM scheduled_action

	WHERE scheduled_action_status in ( 'ready','retry')
	AND scheduled_date < CURRENT_TIMESTAMP
	ORDER BY scheduled_date ASC, updated_at DESC
	
	LIMIT 1
	
	`
	err := s.db.Get(&publicationEntry, SQL)

	//TODO look up this library again... We would expect that there aren't any rows sometimes,

	if err != nil {
		// appcontext.ZLogger(ctx).Error(
		// 	fmt.Sprintf("Error Fetching scheduled action entry %s", err),
		// )
		if errors.Is(err, sql.ErrNoRows) {
			//TODO, it's ok to have null here, verify this

			//appcontext.ZLogger(ctx).Error("Failed to fetch scheduled Actions, because there are no rows", zap.Error(err))

			//return nil, &apperrors.ResourceNotFoundError{Err: err, Resource: models.ScheduledAction{}}
			return nil, nil
		}
		return nil, err
	}
	publicationEntry.Status = "publishing"
	upErr := s.UpdateScheduledAction(ctx, &publicationEntry)
	if upErr != nil {

		appcontext.ZLogger(ctx).Fatal("Error updating scheduled action", zap.Error(upErr))
	}

	return &publicationEntry, nil

}

//UpdateScheduledAction Updates the scheduled action in the DB
func (s *Store) UpdateScheduledAction(ctx context.Context, scheduledAction *models.ScheduledAction) error {

	const UpdateScheduledActionSQL = `
	UPDATE public.scheduled_action
	SET 
		scheduled_action_type= :scheduled_action_type,
		action_data= :action_data,
		created_at= :created_at,
		updated_at= :updated_at,
		recurring= :recurring,
		scheduled_date= :scheduled_date,
		days_to_next_occurence= :days_to_next_occurence,
		last_success_at= :last_success_at,
		scheduled_action_status= :scheduled_action_status,
		current_attempts = :current_attempts,
		max_attempts = :max_attempts,
		note = :note
		
	WHERE id = :id
	
	
	`

	_, err := s.db.NamedExec(UpdateScheduledActionSQL, scheduledAction)
	return err

}

//InsertNewScheduledAction Creates a new scheduled action
func (s *Store) InsertNewScheduledAction(ctx context.Context, scheduledAction *models.ScheduledAction) error {
	const InsertSaSQL = `
	INSERT INTO public.scheduled_action
	( scheduled_action_type,
	 action_data,
	 created_at,
	 updated_at,
	 recurring,
	 scheduled_date,
	 days_to_next_occurence,
	 last_success_at,
	 scheduled_action_status,
	 current_attempts,
	 max_attempts)
	VALUES
	( :scheduled_action_type,
	 :action_data,
	 :DEFAULT,
	 :DEFAULT,
	 :recurring,
	 :scheduled_date,
	 :days_to_next_occurence,
	 :DEFAULT,
	 :DEFAULT,
	 :DEFAULT,
	 :max_attempts)
	RETURNING *; 
	
	
	`

	_, err := s.db.NamedExec(InsertSaSQL, scheduledAction)

	return err

}
