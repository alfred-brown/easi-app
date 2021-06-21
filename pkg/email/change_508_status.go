package email

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"github.com/cmsgov/easi-app/pkg/apperrors"
	"github.com/cmsgov/easi-app/pkg/models"
)

type changeAccessibilityRequestStatus struct {
	RequestName string
	ChangerName string
	OldStatus   string
	NewStatus   string
}

func convertStatusToReadableString(status models.AccessibilityRequestStatus) string {
	switch status {
	case models.AccessibilityRequestStatusOpen:
		return "Open"
	case models.AccessibilityRequestStatusClosed:
		return "Closed"
	case models.AccessibilityRequestStatusInRemediation:
		return "In remediation"
	default:
		return ""
	}
}

func (c Client) changeAccessibilityRequestStatusBody(requestName, removerName string, oldStatus, newStatus models.AccessibilityRequestStatus) (string, error) {
	data := changeAccessibilityRequestStatus{
		ChangerName: removerName,
		RequestName: requestName,
		OldStatus:   convertStatusToReadableString(oldStatus),
		NewStatus:   convertStatusToReadableString(newStatus),
	}
	var b bytes.Buffer
	if c.templates.changeAccessibilityRequestStatus == nil {
		return "", errors.New("email template is nil")
	}
	err := c.templates.changeAccessibilityRequestStatus.Execute(&b, data)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

// SendChangeAccessibilityRequestStatusEmail sends an email for a changed 508 status
func (c Client) SendChangeAccessibilityRequestStatusEmail(
	ctx context.Context,
	requestName,
	changerName string,
	oldStatus,
	newStatus models.AccessibilityRequestStatus,
) error {
	subject := fmt.Sprintf("%s: Status changed to %s", requestName, convertStatusToReadableString(newStatus))
	body, err := c.changeAccessibilityRequestStatusBody(requestName, changerName, oldStatus, newStatus)
	if err != nil {
		return &apperrors.NotificationError{Err: err, DestinationType: apperrors.DestinationTypeEmail}
	}
	err = c.sender.Send(
		ctx,
		c.config.AccessibilityTeamEmail,
		subject,
		body,
	)
	if err != nil {
		return &apperrors.NotificationError{Err: err, DestinationType: apperrors.DestinationTypeEmail}
	}
	return nil
}
