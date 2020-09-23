import React, { useMemo } from 'react';
import { Link } from 'react-router-dom';
import { useTable } from 'react-table';
import { Table } from '@trussworks/react-uswds';
import { DateTime } from 'luxon';

import Footer from 'components/Footer';
import Header from 'components/Header';
import MainContent from 'components/MainContent';
import PageWrapper from 'components/PageWrapper';

const AllRequests = () => {
  const columns: any = useMemo(
    () => [
      {
        Header: 'Submission date',
        accessor: 'submittedAt',
        Cell: ({ value }: any) => {
          return DateTime.fromISO(value).toLocaleString(DateTime.DATE_FULL);
        }
      },
      {
        Header: 'Request name',
        accessor: 'requestName',
        Cell: ({ row, value }: any) => {
          return (
            <Link to={`/governance-review-team/${row.original.id}`}>
              {value}
            </Link>
          );
        }
      },
      {
        Header: 'Component',
        accessor: 'requester.component'
      },
      { Header: 'Status', accessor: 'status' }
    ],
    []
  );

  const data = [
    {
      id: 'addaa218-34d3-4dd8-a12f-38f6ff33b22d',
      submittedAt: new Date().toISOString(),
      requestName: 'Easy Access to System Information',
      requester: {
        name: 'Christopher Hui',
        component: 'Division of Pop Corners'
      },
      status: 'Submitted'
    },
    {
      id: '229f9b64-18fc-4ee1-95c4-9d4b143d215c',
      submittedAt: new Date().toISOString(),
      requestName: 'Hard Access to System Information',
      requester: {
        name: 'George Baukerton',
        component: 'Office of Information Technology'
      },
      status: 'Submitted'
    }
  ];

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    rows,
    prepareRow
  } = useTable({
    columns,
    data
  });

  /* eslint-disable react/jsx-props-no-spreading */
  return (
    <PageWrapper>
      <Header />
      <MainContent className="grid-container">
        <Table bordered={false} {...getTableProps()} fullWidth>
          <caption className="usa-sr-only">
            Table of open requests currently managed by the admin team
          </caption>
          {/* TODO for closed tab: Table of closed requests */}
          <thead>
            {headerGroups.map(headerGroup => (
              <tr {...headerGroup.getHeaderGroupProps()}>
                {headerGroup.headers.map(column => (
                  <th {...column.getHeaderProps()}>
                    {column.render('Header')}
                  </th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody {...getTableBodyProps()}>
            {rows.map(row => {
              prepareRow(row);
              return (
                <tr {...row.getRowProps()}>
                  {row.cells.map((cell, i) => {
                    if (i === 0) {
                      return (
                        <th {...cell.getCellProps()} scope="row">
                          {cell.render('Cell')}
                        </th>
                      );
                    }
                    return (
                      <td {...cell.getCellProps()}>{cell.render('Cell')}</td>
                    );
                  })}
                </tr>
              );
            })}
          </tbody>
        </Table>
      </MainContent>
      <Footer />
    </PageWrapper>
  );
};

export default AllRequests;
