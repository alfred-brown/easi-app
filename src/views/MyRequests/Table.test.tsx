import React from 'react';
import { act } from 'react-dom/test-utils';
import { MemoryRouter } from 'react-router-dom';
import { MockedProvider } from '@apollo/client/testing';
import { mount, ReactWrapper } from 'enzyme';
import GetRequestsQuery from 'queries/GetRequestsQuery';

import Table from './Table';

describe('My Requests Table', () => {
  describe('when there are no requests', () => {
    it('displays an empty message', async () => {
      const mocks = [
        {
          request: {
            query: GetRequestsQuery,
            variables: { first: 20 }
          },
          result: {}
        }
      ];

      let component: ReactWrapper;
      await act(async () => {
        component = mount(
          <MockedProvider mocks={mocks}>
            <Table />
          </MockedProvider>
        );
        await new Promise(resolve => setTimeout(resolve, 0));
        component.update();
      });

      expect(component.find('p').exists()).toBeTruthy();
      expect(component.find('table').exists()).toBeFalsy();
    });
  });

  describe('when there are requests', () => {
    const mocks = [
      {
        request: {
          query: GetRequestsQuery,
          variables: { first: 20 }
        },
        result: {
          data: {
            requests: {
              edges: [
                {
                  node: {
                    id: '123',
                    name: '508 Test 1',
                    submittedAt: '2021-05-25T19:22:40Z',
                    type: 'ACCESSIBILITY_REQUEST'
                  }
                },
                {
                  node: {
                    id: '456',
                    name: 'Intake 1',
                    submittedAt: '2021-05-22T19:22:40Z',
                    type: 'GOVERNANCE_REQUEST'
                  }
                }
              ]
            }
          }
        }
      }
    ];

    const renderComponent = async () => {
      let component: ReactWrapper;
      await act(async () => {
        component = mount(
          <MemoryRouter>
            <MockedProvider mocks={mocks} addTypename={false}>
              <Table />
            </MockedProvider>
          </MemoryRouter>
        );
        await new Promise(resolve => setTimeout(resolve, 0));
        component.update();
      });
      return component;
    };

    it('displays a table', async () => {
      const component = await renderComponent();
      expect(component.find('p').exists()).toBeFalsy();
      expect(component.find('table').exists()).toBeTruthy();
    });

    it('displays headers', async () => {
      const component = await renderComponent();
      const headers = component.find('thead').find('th');
      expect(headers.length).toEqual(3);
      expect(headers.at(0).text()).toEqual('Request name');
      expect(headers.at(1).text()).toEqual('Governance');
      expect(headers.at(2).text()).toEqual('Submission date');
    });

    it('displays rows of data', async () => {
      const component = await renderComponent();
      const rows = component.find('tbody').find('tr');
      expect(rows.length).toEqual(2);

      const rowOne = rows.at(0);
      expect(rowOne.find('th').find('a').html()).toEqual(
        '<a class="usa-link" href="/508/requests/123">508 Test 1</a>'
      );
      expect(rowOne.find('td').at(0).text()).toEqual('Section 508');
      expect(rowOne.find('td').at(1).text()).toEqual('May 25 2021');

      const rowTwo = rows.at(1);
      expect(rowTwo.find('th').find('a').html()).toEqual(
        '<a class="usa-link" href="/governance-task-list/456">Intake 1</a>'
      );
      expect(rowTwo.find('td').at(0).text()).toEqual('IT Governance');
      expect(rowTwo.find('td').at(1).text()).toEqual('May 22 2021');
    });
  });
});