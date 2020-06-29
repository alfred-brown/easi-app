import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { SystemIntakeReview } from 'components/SystemIntakeReview';
import MainContent from 'components/MainContent';
import Header from 'components/Header';
import { AppState } from 'reducers/rootReducer';
import { fetchSystemIntake } from 'types/routines';
import { Field, Form, Formik, FormikProps } from 'formik';
import { RadioField } from 'components/shared/RadioField';
import FieldGroup from 'components/shared/FieldGroup';

type GRTSystemIntakeReviewResponse = {
  decision:
    | 'ISSUE_ID'
    | 'REVIEW_PROCESS_NEEDED'
    | 'GOVERNANCE_NOT_NEEDED'
    | null;
  emailText: string;
};

export const GRTSystemIntakeReview = () => {
  const { systemId } = useParams();
  const dispatch = useDispatch();

  const initialValues: GRTSystemIntakeReviewResponse = {
    decision: null,
    emailText: ''
  };

  useEffect(() => {
    dispatch(fetchSystemIntake(systemId));
  }, [dispatch, systemId]);
  const systemIntake = useSelector(
    (state: AppState) => state.systemIntake.systemIntake
  );
  return (
    <div>
      <Header />
      <MainContent className="grid-container">
        <div className="system-intake__review margin-bottom-7">
          <h1 className="font-heading-xl margin-top-4">CMS System Request</h1>
          {!systemIntake && (
            <h2 className="font-heading-xl">
              System intake with ID: {systemId} was not found
            </h2>
          )}
          {systemIntake && <SystemIntakeReview systemIntake={systemIntake} />}
        </div>
      </MainContent>
      <div className="bg-gray-5">
        <MainContent className="grid-container">
          <Formik
            initialValues={initialValues}
            onSubmit={values => console.log(values)}
          >
            {(formikProps: FormikProps<GRTSystemIntakeReviewResponse>) => {
              const { values, setFieldValue } = formikProps;
              return (
                <Form>
                  <h1 className="font-heading-xl">Next steps</h1>
                  <div>
                    <div className="grid-col-6">
                      <FieldGroup>
                        <fieldset className="usa-fieldset margin-top-4">
                          <legend className="usa-label margin-bottom-1">
                            How to proceed?
                          </legend>
                          <Field
                            as={RadioField}
                            checked={values.decision === 'ISSUE_ID'}
                            id="GrtIntakeReviewForm-DecisionIssueId"
                            name="decision"
                            label="Issue Lifecycle ID with no further governance"
                            onChange={() => {
                              setFieldValue('decision', 'ISSUE_ID');
                            }}
                            value
                          />

                          <Field
                            as={RadioField}
                            checked={
                              values.decision === 'REVIEW_PROCESS_NEEDED'
                            }
                            id="GrtIntakeReviewForm-DecisionReviewProcessNeeded"
                            name="decision"
                            label="Governance review process needed"
                            onChange={() => {
                              setFieldValue(
                                'decision',
                                'REVIEW_PROCESS_NEEDED'
                              );
                            }}
                            value
                          />

                          <Field
                            as={RadioField}
                            checked={
                              values.decision === 'GOVERNANCE_NOT_NEEDED'
                            }
                            id="GrtIntakeReviewForm-DecisionGovernanceNotNeeded"
                            name="decision"
                            label="Governance not needed"
                            helpText="Close this request"
                            onChange={() => {
                              setFieldValue(
                                'decision',
                                'GOVERNANCE_NOT_NEEDED'
                              );
                            }}
                            value
                          />
                        </fieldset>
                      </FieldGroup>
                    </div>
                  </div>
                </Form>
              );
            }}
          </Formik>
        </MainContent>
      </div>
    </div>
  );
};

export default GRTSystemIntakeReview;
