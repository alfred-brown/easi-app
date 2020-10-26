package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/cmsgov/easi-app/pkg/apperrors"
	"github.com/cmsgov/easi-app/pkg/models"
)

type createSystemIntakeAction func(context.Context, *models.Action) error

// NewSystemIntakeActionHandler is a constructor for SystemIntakeActionHandler
func NewSystemIntakeActionHandler(
	base HandlerBase,
	create createSystemIntakeAction,
) SystemIntakeActionHandler {
	return SystemIntakeActionHandler{
		HandlerBase:              base,
		CreateSystemIntakeAction: create,
	}
}

// SystemIntakeActionHandler is the handler for CRUD operations on a system intake action
type SystemIntakeActionHandler struct {
	HandlerBase
	CreateSystemIntakeAction createSystemIntakeAction
}

// Handle handles a request for the system intake action
func (h SystemIntakeActionHandler) Handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			id := mux.Vars(r)["intake_id"]
			valErr := apperrors.NewValidationError(
				errors.New("system intake failed validation"),
				models.SystemIntake{},
				"",
			)
			if id == "" {
				valErr.WithValidation("path.intakeID", "is required")
				h.WriteErrorResponse(r.Context(), w, &valErr)
				return
			}
			intakeID, err := uuid.Parse(id)
			if err != nil {
				valErr.WithValidation("path.intakeID", "must be UUID")
				h.WriteErrorResponse(r.Context(), w, &valErr)
				return
			}
			actionType := models.ActionType(mux.Vars(r)["action_type"])
			if actionType == "" {
				valErr.WithValidation("path.actionType", "is required")
				h.WriteErrorResponse(r.Context(), w, &valErr)
				return
			}
			defer r.Body.Close()
			decoder := json.NewDecoder(r.Body)
			action := models.Action{}
			err = decoder.Decode(&action)
			if err != nil && err != io.EOF {
				h.WriteErrorResponse(r.Context(), w, &apperrors.BadRequestError{Err: err})
				return
			}
			action.ActionType = actionType
			action.IntakeID = &intakeID
			err = h.CreateSystemIntakeAction(r.Context(), &action)
			if err != nil {
				h.WriteErrorResponse(r.Context(), w, err)
				return
			}

			w.WriteHeader(http.StatusCreated)
			return
		default:
			h.WriteErrorResponse(r.Context(), w, &apperrors.MethodNotAllowedError{Method: r.Method})
			return
		}
	}
}