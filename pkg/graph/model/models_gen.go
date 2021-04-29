// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/cmsgov/easi-app/pkg/models"
	"github.com/google/uuid"
)

type AccessibilityRequestDocumentType struct {
	CommonType           models.AccessibilityRequestDocumentCommonType `json:"commonType"`
	OtherTypeDescription *string                                       `json:"otherTypeDescription"`
}

type AccessibilityRequestEdge struct {
	Cursor string                       `json:"cursor"`
	Node   *models.AccessibilityRequest `json:"node"`
}

type AccessibilityRequestsConnection struct {
	Edges      []*AccessibilityRequestEdge `json:"edges"`
	TotalCount int                         `json:"totalCount"`
}

type AddGRTFeedbackInput struct {
	EmailBody string    `json:"emailBody"`
	Feedback  string    `json:"feedback"`
	IntakeID  uuid.UUID `json:"intakeID"`
}

type AddGRTFeedbackPayload struct {
	ID *uuid.UUID `json:"id"`
}

type BasicActionInput struct {
	Feedback string    `json:"feedback"`
	IntakeID uuid.UUID `json:"intakeId"`
}

type BusinessCaseAsIsSolution struct {
	Cons        *string `json:"cons"`
	CostSavings *string `json:"costSavings"`
	Pros        *string `json:"pros"`
	Summary     *string `json:"summary"`
	Title       *string `json:"title"`
}

// A solution proposal within a business case
type BusinessCaseSolution struct {
	AcquisitionApproach     *string `json:"acquisitionApproach"`
	Cons                    *string `json:"cons"`
	CostSavings             *string `json:"costSavings"`
	HasUI                   *string `json:"hasUi"`
	HostingCloudServiceType *string `json:"hostingCloudServiceType"`
	HostingLocation         *string `json:"hostingLocation"`
	HostingType             *string `json:"hostingType"`
	Pros                    *string `json:"pros"`
	SecurityIsApproved      *bool   `json:"securityIsApproved"`
	SecurityIsBeingReviewed *string `json:"securityIsBeingReviewed"`
	Summary                 *string `json:"summary"`
	Title                   *string `json:"title"`
}

type ContractDate struct {
	Day   *string `json:"day"`
	Month *string `json:"month"`
	Year  *string `json:"year"`
}

type CreateAccessibilityRequestDocumentInput struct {
	CommonDocumentType           models.AccessibilityRequestDocumentCommonType `json:"commonDocumentType"`
	MimeType                     string                                        `json:"mimeType"`
	Name                         string                                        `json:"name"`
	OtherDocumentTypeDescription *string                                       `json:"otherDocumentTypeDescription"`
	RequestID                    uuid.UUID                                     `json:"requestID"`
	Size                         int                                           `json:"size"`
	URL                          string                                        `json:"url"`
}

type CreateAccessibilityRequestDocumentPayload struct {
	AccessibilityRequestDocument *models.AccessibilityRequestDocument `json:"accessibilityRequestDocument"`
	UserErrors                   []*UserError                         `json:"userErrors"`
}

type CreateAccessibilityRequestInput struct {
	IntakeID uuid.UUID `json:"intakeID"`
	Name     string    `json:"name"`
}

type CreateAccessibilityRequestPayload struct {
	AccessibilityRequest *models.AccessibilityRequest `json:"accessibilityRequest"`
	UserErrors           []*UserError                 `json:"userErrors"`
}

type CreateSystemIntakeNoteInput struct {
	Content    string    `json:"content"`
	AuthorName string    `json:"authorName"`
	IntakeID   uuid.UUID `json:"intakeId"`
}

type CreateTestDateInput struct {
	Date      time.Time               `json:"date"`
	RequestID uuid.UUID               `json:"requestID"`
	Score     *int                    `json:"score"`
	TestType  models.TestDateTestType `json:"testType"`
}

type CreateTestDatePayload struct {
	TestDate   *models.TestDate `json:"testDate"`
	UserErrors []*UserError     `json:"userErrors"`
}

type DeleteAccessibilityRequestDocumentInput struct {
	ID uuid.UUID `json:"id"`
}

type DeleteAccessibilityRequestDocumentPayload struct {
	ID *uuid.UUID `json:"id"`
}

type DeleteTestDateInput struct {
	ID uuid.UUID `json:"id"`
}

type DeleteTestDatePayload struct {
	TestDate   *models.TestDate `json:"testDate"`
	UserErrors []*UserError     `json:"userErrors"`
}

type GeneratePresignedUploadURLInput struct {
	FileName string `json:"fileName"`
	MimeType string `json:"mimeType"`
	Size     int    `json:"size"`
}

type GeneratePresignedUploadURLPayload struct {
	URL        *string      `json:"url"`
	UserErrors []*UserError `json:"userErrors"`
}

type IssueLifecycleIDInput struct {
	ExpiresAt time.Time `json:"expiresAt"`
	Feedback  string    `json:"feedback"`
	IntakeID  uuid.UUID `json:"intakeId"`
	Lcid      *string   `json:"lcid"`
	NextSteps *string   `json:"nextSteps"`
	Scope     string    `json:"scope"`
}

type RejectIntakeInput struct {
	Feedback  string    `json:"feedback"`
	IntakeID  uuid.UUID `json:"intakeId"`
	NextSteps *string   `json:"nextSteps"`
	Reason    string    `json:"reason"`
}

type SystemConnection struct {
	Edges      []*SystemEdge `json:"edges"`
	TotalCount int           `json:"totalCount"`
}

type SystemEdge struct {
	Cursor string         `json:"cursor"`
	Node   *models.System `json:"node"`
}

// An action taken on a system intake, often resulting in a change in status.
type SystemIntakeAction struct {
	ID           uuid.UUID                `json:"id"`
	SystemIntake *models.SystemIntake     `json:"systemIntake"`
	Type         SystemIntakeActionType   `json:"type"`
	Actor        *SystemIntakeActionActor `json:"actor"`
	Feedback     *string                  `json:"feedback"`
	CreatedAt    time.Time                `json:"createdAt"`
}

type SystemIntakeActionActor struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type SystemIntakeBusinessOwner struct {
	Component *string `json:"component"`
	Name      *string `json:"name"`
}

type SystemIntakeCollaborator struct {
	Acronym      *string `json:"acronym"`
	Collaborator *string `json:"collaborator"`
	Key          *string `json:"key"`
	Label        *string `json:"label"`
	Name         *string `json:"name"`
}

type SystemIntakeContract struct {
	Contractor  *string       `json:"contractor"`
	EndDate     *ContractDate `json:"endDate"`
	HasContract *string       `json:"hasContract"`
	StartDate   *ContractDate `json:"startDate"`
	Vehicle     *string       `json:"vehicle"`
}

type SystemIntakeCosts struct {
	ExpectedIncreaseAmount *string `json:"expectedIncreaseAmount"`
	IsExpectingIncrease    *string `json:"isExpectingIncrease"`
}

type SystemIntakeFundingSource struct {
	FundingNumber *string `json:"fundingNumber"`
	IsFunded      *bool   `json:"isFunded"`
	Source        *string `json:"source"`
}

type SystemIntakeGovernanceTeam struct {
	IsPresent *bool                       `json:"isPresent"`
	Teams     []*SystemIntakeCollaborator `json:"teams"`
}

type SystemIntakeIsso struct {
	IsPresent *bool   `json:"isPresent"`
	Name      *string `json:"name"`
}

type SystemIntakeNote struct {
	Author    *SystemIntakeNoteAuthor `json:"author"`
	Content   string                  `json:"content"`
	CreatedAt time.Time               `json:"createdAt"`
	ID        uuid.UUID               `json:"id"`
}

type SystemIntakeNoteAuthor struct {
	Eua  string `json:"eua"`
	Name string `json:"name"`
}

type SystemIntakeProductManager struct {
	Component *string `json:"component"`
	Name      *string `json:"name"`
}

type SystemIntakeRequester struct {
	Component *string `json:"component"`
	Email     *string `json:"email"`
	Name      string  `json:"name"`
}

type UpdateSystemIntakeAdminLeadInput struct {
	AdminLead string    `json:"adminLead"`
	ID        uuid.UUID `json:"id"`
}

type UpdateSystemIntakePayload struct {
	SystemIntake *models.SystemIntake `json:"systemIntake"`
	UserErrors   []*UserError         `json:"userErrors"`
}

type UpdateSystemIntakeReviewDatesInput struct {
	GrbDate *time.Time `json:"grbDate"`
	GrtDate *time.Time `json:"grtDate"`
	ID      uuid.UUID  `json:"id"`
}

type UpdateTestDateInput struct {
	Date     time.Time               `json:"date"`
	ID       uuid.UUID               `json:"id"`
	Score    *int                    `json:"score"`
	TestType models.TestDateTestType `json:"testType"`
}

type UpdateTestDatePayload struct {
	TestDate   *models.TestDate `json:"testDate"`
	UserErrors []*UserError     `json:"userErrors"`
}

// UserError represents application-level errors that are the result of
// either user or application developer error.
type UserError struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

// A user role associated with a job code
type Role string

const (
	// A 508 Tester
	RoleEasi508Tester Role = "EASI_508_TESTER"
	// A 508 request program team member
	RoleEasi508User Role = "EASI_508_USER"
	// A 508 request program team member or tester
	RoleEasi508TesterOrUser Role = "EASI_508_TESTER_OR_USER"
	// A member of the GRT
	RoleEasiGovteam Role = "EASI_GOVTEAM"
	// A generic EASi user
	RoleEasiUser Role = "EASI_USER"
)

var AllRole = []Role{
	RoleEasi508Tester,
	RoleEasi508User,
	RoleEasi508TesterOrUser,
	RoleEasiGovteam,
	RoleEasiUser,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleEasi508Tester, RoleEasi508User, RoleEasi508TesterOrUser, RoleEasiGovteam, RoleEasiUser:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SystemIntakeActionType string

const (
	SystemIntakeActionTypeBizCaseNeedsChanges            SystemIntakeActionType = "BIZ_CASE_NEEDS_CHANGES"
	SystemIntakeActionTypeCreateBizCase                  SystemIntakeActionType = "CREATE_BIZ_CASE"
	SystemIntakeActionTypeGUIDEReceivedClose             SystemIntakeActionType = "GUIDE_RECEIVED_CLOSE"
	SystemIntakeActionTypeIssueLcid                      SystemIntakeActionType = "ISSUE_LCID"
	SystemIntakeActionTypeNeedBizCase                    SystemIntakeActionType = "NEED_BIZ_CASE"
	SystemIntakeActionTypeNoGovernanceNeeded             SystemIntakeActionType = "NO_GOVERNANCE_NEEDED"
	SystemIntakeActionTypeNotItRequest                   SystemIntakeActionType = "NOT_IT_REQUEST"
	SystemIntakeActionTypeNotRespondingClose             SystemIntakeActionType = "NOT_RESPONDING_CLOSE"
	SystemIntakeActionTypeProvideFeedbackNeedBizCase     SystemIntakeActionType = "PROVIDE_FEEDBACK_NEED_BIZ_CASE"
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseDraft SystemIntakeActionType = "PROVIDE_GRT_FEEDBACK_BIZ_CASE_DRAFT"
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseFinal SystemIntakeActionType = "PROVIDE_GRT_FEEDBACK_BIZ_CASE_FINAL"
	SystemIntakeActionTypeReadyForGrb                    SystemIntakeActionType = "READY_FOR_GRB"
	SystemIntakeActionTypeReadyForGrt                    SystemIntakeActionType = "READY_FOR_GRT"
	SystemIntakeActionTypeReject                         SystemIntakeActionType = "REJECT"
	SystemIntakeActionTypeSendEmail                      SystemIntakeActionType = "SEND_EMAIL"
	SystemIntakeActionTypeSubmitBizCase                  SystemIntakeActionType = "SUBMIT_BIZ_CASE"
	SystemIntakeActionTypeSubmitFinalBizCase             SystemIntakeActionType = "SUBMIT_FINAL_BIZ_CASE"
	SystemIntakeActionTypeSubmitIntake                   SystemIntakeActionType = "SUBMIT_INTAKE"
)

var AllSystemIntakeActionType = []SystemIntakeActionType{
	SystemIntakeActionTypeBizCaseNeedsChanges,
	SystemIntakeActionTypeCreateBizCase,
	SystemIntakeActionTypeGUIDEReceivedClose,
	SystemIntakeActionTypeIssueLcid,
	SystemIntakeActionTypeNeedBizCase,
	SystemIntakeActionTypeNoGovernanceNeeded,
	SystemIntakeActionTypeNotItRequest,
	SystemIntakeActionTypeNotRespondingClose,
	SystemIntakeActionTypeProvideFeedbackNeedBizCase,
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseDraft,
	SystemIntakeActionTypeProvideGrtFeedbackBizCaseFinal,
	SystemIntakeActionTypeReadyForGrb,
	SystemIntakeActionTypeReadyForGrt,
	SystemIntakeActionTypeReject,
	SystemIntakeActionTypeSendEmail,
	SystemIntakeActionTypeSubmitBizCase,
	SystemIntakeActionTypeSubmitFinalBizCase,
	SystemIntakeActionTypeSubmitIntake,
}

func (e SystemIntakeActionType) IsValid() bool {
	switch e {
	case SystemIntakeActionTypeBizCaseNeedsChanges, SystemIntakeActionTypeCreateBizCase, SystemIntakeActionTypeGUIDEReceivedClose, SystemIntakeActionTypeIssueLcid, SystemIntakeActionTypeNeedBizCase, SystemIntakeActionTypeNoGovernanceNeeded, SystemIntakeActionTypeNotItRequest, SystemIntakeActionTypeNotRespondingClose, SystemIntakeActionTypeProvideFeedbackNeedBizCase, SystemIntakeActionTypeProvideGrtFeedbackBizCaseDraft, SystemIntakeActionTypeProvideGrtFeedbackBizCaseFinal, SystemIntakeActionTypeReadyForGrb, SystemIntakeActionTypeReadyForGrt, SystemIntakeActionTypeReject, SystemIntakeActionTypeSendEmail, SystemIntakeActionTypeSubmitBizCase, SystemIntakeActionTypeSubmitFinalBizCase, SystemIntakeActionTypeSubmitIntake:
		return true
	}
	return false
}

func (e SystemIntakeActionType) String() string {
	return string(e)
}

func (e *SystemIntakeActionType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SystemIntakeActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SystemIntakeActionType", str)
	}
	return nil
}

func (e SystemIntakeActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
