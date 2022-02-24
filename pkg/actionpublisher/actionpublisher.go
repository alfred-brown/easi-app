package actionpublisher

import (
	"context"

	"time"

	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/cmsgov/easi-app/pkg/storage"
)

//ActionPublisher is the task that checks for timed actions to be taken, and take them when appropriate
type ActionPublisher struct {
	store           *storage.Store
	exitChan        chan bool
	scheduledAction *models.ScheduledAction
}

const jobCheckTime = 15
const errSleepTime = 1

//NewActionPublisher creates an instance of an ActionPublisher with reference to a store
func NewActionPublisher(store *storage.Store) ActionPublisher {
	ap := ActionPublisher{}
	ap.store = store
	ap.exitChan = make(chan bool)
	return ap
}

//Start initiates the ActionPublisers main loop
func (ap *ActionPublisher) Start() {

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
			ap.checkForJob()

		}

	}
}

func (ap *ActionPublisher) checkForJob() {
	sA, err := ap.store.FetchNextScheduledAction(context.Background())

	// if there is an error and the action was retrieved, we will want to update the status of hte scheduledaCtion
	if err != nil {
		time.Sleep(errSleepTime * time.Minute)
		return
	}
	if sA != nil {
		ap.scheduledAction = sA
		sA.CurrentAttempts++
		//todo, do something with event data here, either generic dictionary structure, or let each event know the format. Cannot change if like this.
		_ = ap.publishAction()
		return
	}

	time.Sleep(jobCheckTime * time.Minute)

}

func (ap *ActionPublisher) publishAction() error {

	//todo get succes of action or not here, set to retry if
	var err error

	switch ap.scheduledAction.ActionType {
	case "LCID_Expiring":
		err = ap.publishLCIDExpiringAction()

	default:
		err = ap.setActionToFailed()

	}

	if err != nil {
		_ = ap.setActionToRetry()
	}

	return nil
}

func (ap *ActionPublisher) publishLCIDExpiringAction() error {
	//TODO
	/*
		1. Take action based on amount of time to come
		2. Send email?
		3. Set Status to success / retry. Increment the attempts
	*/
	//ap.store.

	//_,err :=
	err := ap.setActionToComplete()
	return err
}
func (ap *ActionPublisher) setActionToComplete() error {

	ap.scheduledAction.Status = "complete"
	ap.scheduledAction.LastSuccessTime = time.Now()
	//ap.store.UpdateScheduledAction(context.Background(), ap.scheduledAction)

	if ap.scheduledAction.Recurring {
		//TODO polish this
		ap.scheduledAction.ScheduledTime = time.Now().AddDate(0, 0, int(ap.scheduledAction.DaysToNext))
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
