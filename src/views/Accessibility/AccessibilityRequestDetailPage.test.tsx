import React from 'react';
import { Provider } from 'react-redux';
import { MemoryRouter, Route } from 'react-router-dom';
import { MockedProvider } from '@apollo/client/testing';
import {
  render,
  screen,
  waitForElementToBeRemoved,
  within
} from '@testing-library/react';
import GetAccessibilityRequestAccessibilityTeamOnlyQuery from 'queries/GetAccessibilityRequestAccessibilityTeamOnlyQuery';
import GetAccessibilityRequestQuery from 'queries/GetAccessibilityRequestQuery';
import configureMockStore from 'redux-mock-store';

import { MessageProvider } from 'hooks/useMessage';

import { ACCESSIBILITY_TESTER_DEV } from '../../constants/jobCodes';

import AccessibilityRequestDetailPage from './AccessibilityRequestDetailPage';

describe('AccessibilityRequestDetailPage', () => {
  const mockStore = configureMockStore();
  const store = mockStore({
    auth: {
      euaId: 'AAAA'
    }
  });

  const default508RequestQuery = {
    request: {
      query: GetAccessibilityRequestQuery,
      variables: {
        id: 'e0a4de2f-a2c2-457d-ac08-bbd011104855'
      }
    },
    result: {
      data: {
        accessibilityRequest: {
          id: 'e0a4de2f-a2c2-457d-ac08-bbd011104855',
          euaUserId: 'ABCD',
          submittedAt: new Date().toISOString(),
          name: 'My Special Request',
          system: {
            name: 'TACO',
            lcid: '123456',
            businessOwner: { name: 'Clark Kent', component: 'OIT' }
          },
          documents: [],
          testDates: [],
          statusRecord: {
            status: 'OPEN'
          }
        }
      }
    }
  };

  it('renders without crashing', async () => {
    render(
      <MemoryRouter
        initialEntries={['/508/requests/e0a4de2f-a2c2-457d-ac08-bbd011104855']}
      >
        <MockedProvider mocks={[default508RequestQuery]} addTypename={false}>
          <MessageProvider>
            <Provider store={store}>
              <Route path="/508/requests/:accessibilityRequestId">
                <AccessibilityRequestDetailPage />
              </Route>
            </Provider>
          </MessageProvider>
        </MockedProvider>
      </MemoryRouter>
    );

    await waitForElementToBeRemoved(() => screen.getByTestId('page-loading'));

    expect(
      screen.getByTestId('accessibility-request-detail-page')
    ).toBeInTheDocument();
  });

  describe('for a business owner', () => {
    // ✅ Documents
    const withDocsQuery = {
      request: {
        query: GetAccessibilityRequestQuery,
        variables: {
          id: 'a11yRequest123'
        }
      },
      result: {
        data: {
          accessibilityRequest: {
            id: 'a11yRequest123',
            euaUserId: 'AAAA',
            submittedAt: new Date().toISOString(),
            name: 'MY Request',
            system: {
              name: 'TACO',
              lcid: '0000',
              businessOwner: { name: 'Clark Kent', component: 'OIT' }
            },
            documents: [
              {
                id: 'doc1',
                url: 'myurl',
                uploadedAt: 'time',
                status: 'PENDING',
                documentType: {
                  commonType: 'TEST_PLAN',
                  otherTypeDescription: ''
                }
              }
            ],
            testDates: [],
            statusRecord: {
              status: 'OPEN'
            },
            notes: [
              {
                id: 'noteID',
                authorName: 'Common Human',
                note: 'This is very well done'
              }
            ]
          }
        }
      }
    };

    // ❌ Documents
    const withoutDocsQuery = {
      request: {
        query: GetAccessibilityRequestQuery,
        variables: {
          id: 'a11yRequest123'
        }
      },
      result: {
        data: {
          accessibilityRequest: {
            id: 'a11yRequest123',
            euaUserId: 'AAAA',
            submittedAt: new Date().toISOString(),
            name: 'MY Request',
            system: {
              name: 'TACO',
              lcid: '0000',
              businessOwner: { name: 'Clark Kent', component: 'OIT' }
            },
            documents: [],
            testDates: [],
            statusRecord: {
              status: 'OPEN'
            }
          }
        }
      }
    };
    it('renders Next step if no documents', async () => {
      render(
        <MemoryRouter initialEntries={['/508/requests/a11yRequest123']}>
          <MockedProvider mocks={[withoutDocsQuery]} addTypename={false}>
            <Provider store={store}>
              <MessageProvider>
                <Route path="/508/requests/:accessibilityRequestId">
                  <AccessibilityRequestDetailPage />
                </Route>
              </MessageProvider>
            </Provider>
          </MockedProvider>
        </MemoryRouter>
      );

      await waitForElementToBeRemoved(() => screen.getByTestId('page-loading'));

      expect(
        screen.getByRole('heading', {
          level: 2,
          name: /Next step/i
        })
      ).toBeInTheDocument();
    });

    it('renders the AccessibilityDocumentList when documents exist', async () => {
      render(
        <MemoryRouter initialEntries={['/508/requests/a11yRequest123']}>
          <MockedProvider mocks={[withDocsQuery]} addTypename={false}>
            <Provider store={store}>
              <MessageProvider>
                <Route path="/508/requests/:accessibilityRequestId">
                  <AccessibilityRequestDetailPage />
                </Route>
              </MessageProvider>
            </Provider>
          </MockedProvider>
        </MemoryRouter>
      );

      await waitForElementToBeRemoved(() => screen.getByTestId('page-loading'));

      expect(screen.getByTestId('body-with-doc-table')).toBeInTheDocument();
    });
  });

  describe('for a 508 user or 508 tester', () => {
    const defaultQuery = {
      request: {
        query: GetAccessibilityRequestAccessibilityTeamOnlyQuery,
        variables: {
          id: 'a11yRequest123'
        }
      },
      result: {
        data: {
          accessibilityRequest: {
            id: 'a11yRequest123',
            euaUserId: 'AAAA',
            submittedAt: new Date().toISOString(),
            name: 'MY Request',
            system: {
              name: 'TACO',
              lcid: '0000',
              businessOwner: { name: 'Clark Kent', component: 'OIT' }
            },
            documents: [],
            testDates: [],
            statusRecord: {
              status: 'OPEN'
            },
            notes: []
          }
        }
      }
    };

    // ✅ Notes
    // ❌ Documents
    const withNotesQuery = {
      request: {
        query: GetAccessibilityRequestAccessibilityTeamOnlyQuery,
        variables: {
          id: 'a11yRequest123'
        }
      },
      result: {
        data: {
          accessibilityRequest: {
            id: 'a11yRequest123',
            euaUserId: 'AAAA',
            submittedAt: new Date().toISOString(),
            name: 'MY Request',
            system: {
              name: 'TACO',
              lcid: '0000',
              businessOwner: { name: 'Clark Kent', component: 'OIT' }
            },
            documents: [],
            testDates: [],
            statusRecord: {
              status: 'OPEN'
            },
            notes: [
              {
                id: 'noteID',
                createdAt: new Date().toISOString(),
                authorName: 'Common Human',
                note: 'This is very well done'
              },
              {
                id: 'noteID2',
                createdAt: new Date().toISOString(),
                authorName: 'Common Human',
                note: 'This is okay'
              }
            ]
          }
        }
      }
    };

    // ✅ Documents
    // ❌ Notes
    const withDocsQuery = {
      request: {
        query: GetAccessibilityRequestAccessibilityTeamOnlyQuery,
        variables: {
          id: 'a11yRequest123'
        }
      },
      result: {
        data: {
          accessibilityRequest: {
            id: 'a11yRequest123',
            euaUserId: 'AAAA',
            submittedAt: new Date(),
            name: 'MY Request',
            system: {
              name: 'TACO',
              lcid: '0000',
              businessOwner: { name: 'Clark Kent', component: 'OIT' }
            },
            documents: [
              {
                id: 'doc1',
                url: 'myurl',
                uploadedAt: 'time',
                status: 'PENDING',
                documentType: {
                  commonType: 'TEST_PLAN',
                  otherTypeDescription: ''
                }
              }
            ],
            testDates: [],
            statusRecord: {
              status: 'OPEN'
            },
            notes: [
              {
                id: 'noteID',
                authorName: 'Common Human',
                note: 'This is very well done'
              }
            ]
          }
        }
      }
    };

    const testerStore = mockStore({
      auth: { groups: [ACCESSIBILITY_TESTER_DEV], isUserSet: true }
    });

    const defaultRender = () =>
      render(
        <MemoryRouter initialEntries={['/508/requests/a11yRequest123']}>
          <MockedProvider mocks={[defaultQuery]} addTypename={false}>
            <Provider store={testerStore}>
              <MessageProvider>
                <Route path="/508/requests/:accessibilityRequestId">
                  <AccessibilityRequestDetailPage />
                </Route>
              </MessageProvider>
            </Provider>
          </MockedProvider>
        </MemoryRouter>
      );

    it("doesn't render table if there are no documents", async () => {
      defaultRender();

      await waitForElementToBeRemoved(() => screen.getByTestId('page-loading'));

      expect(
        screen.queryByRole('table', { name: /Docuemnts uploaded/i })
      ).not.toBeInTheDocument();
    });

    it('renders table if there are documents', async () => {
      render(
        <MemoryRouter initialEntries={['/508/requests/a11yRequest123']}>
          <MockedProvider mocks={[withDocsQuery]} addTypename={false}>
            <Provider store={testerStore}>
              <MessageProvider>
                <Route path="/508/requests/:accessibilityRequestId">
                  <AccessibilityRequestDetailPage />
                </Route>
              </MessageProvider>
            </Provider>
          </MockedProvider>
        </MemoryRouter>
      );

      await waitForElementToBeRemoved(() => screen.getByTestId('page-loading'));

      expect(
        screen.getByRole('table', { name: /Documents upload/i })
      ).toBeInTheDocument();
    });

    describe('notes tab', () => {
      it('can view existing notes', async () => {
        render(
          <MemoryRouter initialEntries={['/508/requests/a11yRequest123']}>
            <MockedProvider mocks={[withNotesQuery]} addTypename={false}>
              <Provider store={testerStore}>
                <MessageProvider>
                  <Route path="/508/requests/:accessibilityRequestId">
                    <AccessibilityRequestDetailPage />
                  </Route>
                </MessageProvider>
              </Provider>
            </MockedProvider>
          </MemoryRouter>
        );

        await waitForElementToBeRemoved(() =>
          screen.getByTestId('page-loading')
        );

        const notesList = screen.getByRole('list', {
          name: /existing notes/i
        });

        expect(within(notesList).getAllByRole('listitem').length).toEqual(2);
      });

      it('displays validation errors', async () => {
        defaultRender();

        await waitForElementToBeRemoved(() =>
          screen.getByTestId('page-loading')
        );

        screen.getByRole('button', { name: /Add note/i }).click();

        expect(await screen.findByRole('button', { name: /Enter a note/i }));
      });
    });
  });
});
