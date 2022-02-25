package actionpublisher

import (
	"context"
	"fmt"

	"time"

	"github.com/cmsgov/easi-app/pkg/appcontext"
	"github.com/cmsgov/easi-app/pkg/email"

	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/cmsgov/easi-app/pkg/storage"

	"go.uber.org/zap"
)

//ActionPublisher is the task that checks for timed actions to be taken, and take them when appropriate
type ActionPublisher struct {
	store           *storage.Store
	context         context.Context
	emailClient     email.Client //TODO should this be a reference?
	exitChan        chan bool
	scheduledAction *models.ScheduledAction
}

const jobCheckTime = 1
const errSleepTime = 1

//NewActionPublisher creates an instance of an ActionPublisher with reference to a store
func NewActionPublisher(context context.Context, store *storage.Store, emailClient email.Client) (ActionPublisher, error) {
	ap := ActionPublisher{}
	ap.store = store
	ap.context = context
	ap.emailClient = emailClient
	ap.exitChan = make(chan bool)
	appcontext.ZLogger(context).Info("actionPublisherCreated")
	//ap.context

	return ap, nil
}

//Start initiates the ActionPublisers main loop
func (ap *ActionPublisher) Start() {
	//return

	go ap.apLoop()

}

//Stop cancels the ActionPublishers loop
func (ap *ActionPublisher) Stop() {
	ap.exitChan <- true
}

func (ap *ActionPublisher) apLoop() {

	for {
		select {
		case <-ap.exitChan:
			return
		default:
			jobErr := ap.checkForJob()
			if jobErr != nil {
				appcontext.ZLogger(ap.context).Error("issue finding job", zap.Error(jobErr))
			}

		}

	}
}

func (ap *ActionPublisher) checkForJob() error {
	sA, err := ap.store.FetchNextScheduledAction(ap.context)

	// if there is an error and the action was retrieved, we will want to update the status of hte scheduledaCtion
	if err != nil {
		time.Sleep(errSleepTime * time.Minute)
		return err
	}
	if sA != nil {
		ap.scheduledAction = sA
		sA.CurrentAttempts++
		//todo, do something with event data here, either generic dictionary structure, or let each event know the format. Cannot change if like this.
		err = ap.publishAction()
		return err
	}
	appcontext.ZLogger(ap.context).Info(fmt.Sprintf("There are no scheduled actions. Sleeping for %d minutes", jobCheckTime))

	time.Sleep(jobCheckTime * time.Minute)
	return nil

}

func (ap *ActionPublisher) publishAction() error {

	//todo get succes of action or not here, set to retry if
	var err error

	switch ap.scheduledAction.ActionType {
	case "LCID_Expiring":
		appcontext.ZLogger(ap.context).Info(fmt.Sprintf("Publishing new ScheduledAction: %s, %s", ap.scheduledAction.ActionType, ap.scheduledAction.ActionData))
		err = ap.publishLCIDExpiringAction()

	default:
		ap.scheduledAction.Note = "no action defined for this actionType"
		err = ap.setActionToFailed()

	}

	if err != nil {
		_ = ap.setActionToRetry()
	}

	return nil
}

func (ap *ActionPublisher) publishLCIDExpiringAction() error {

	// emailAddress := "steven.wade@oddball.io"
	// emailAddress := "success@simulator.amazonses.com"
	emailAddress := ap.scheduledAction.ActionData["emailTo"].(string)
	//TODO, is there a better way to get / repreent the data in the dictionary

	// emailAddress := fmt.Sprintf("%v", ap.scheduledAction.ActionData["emailTo"])
	emailSubject := "LCID EXPIRING"

	// emailBody := ap.scheduledAction.ActionData
	emailBody := " A bunch of good information"
	appcontext.ZLogger(ap.context).Info(fmt.Sprintf("Looking at action data. Found email %s %s, %s", emailAddress, ap.scheduledAction.ActionType, ap.scheduledAction.ActionData))

	emailErr := ap.emailClient.SendLCIDExpiringEMAIL(ap.context, emailAddress, emailSubject, emailBody)

	if emailErr != nil {
		return emailErr
	}

	//TODO
	/*
		1. Take action based on amount of time to come
		2. Send email?
		3. Set Status to success / retry. Increment the attempts
	*/
	//ap.store.

	//_,err :=
	ap.scheduledAction.Note = "action completed"
	err := ap.setActionToComplete()
	return err
}
func (ap *ActionPublisher) setActionToComplete() error {

	ap.scheduledAction.Status = "complete"
	successTime := time.Now()
	ap.scheduledAction.LastSuccessTime = &successTime
	// ap.scheduledAction.LastSuccessTime = *time.Now()
	//ap.store.UpdateScheduledAction(context.Background(), ap.scheduledAction)

	if ap.scheduledAction.Recurring {
		//TODO polish this. We could either just at this point insert another record and consider this is done? Or just update it.
		scheduledT := time.Now().AddDate(0, 0, int(ap.scheduledAction.DaysToNext))
		ap.scheduledAction.ScheduledTime = &scheduledT
		ap.scheduledAction.CurrentAttempts = 0

	}
	err := ap.store.UpdateScheduledAction(context.Background(), ap.scheduledAction)
	return err
}

func (ap *ActionPublisher) setActionToFailed() error {

	ap.scheduledAction.Status = "failed"
	//ap.scheduledAction.LastSuccessTime = time.Now()
	//ap.store.UpdateScheduledAction(context.Background(), ap.scheduledAction)

	err := ap.store.UpdateScheduledAction(context.Background(), ap.scheduledAction)
	return err
}

func (ap *ActionPublisher) setActionToRetry() error {
	var err error
	ap.scheduledAction.Status = "retry"
	if ap.scheduledAction.CurrentAttempts == ap.scheduledAction.MaxAttempts {
		err = ap.setActionToFailed()
		return err
	}

	err = ap.store.UpdateScheduledAction(context.Background(), ap.scheduledAction)
	return err
}
