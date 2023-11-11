'use client';

import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from '@nextui-org/table';
import { Pagination } from '@nextui-org/pagination';
import { useState, useMemo, useEffect } from 'react';
import { notifyError } from '../components/toast/notifications';
import { GET } from '../libs/axios/axios';
import HistoryStyleCell from './style-cell';
import useAuth from '../hooks/useAuth';

const rowsPerPage = 10;

const columns = [
  {
    key: 'UpdatedAt',
    label: 'COMPLETED AT',
  },
  {
    key: 'title',
    label: 'QUESTION',
  },
  {
    key: 'language',
    label: 'LANGUAGE',
  },
  {
    key: 'actions',
    label: 'VIEW SOLUTION',
  },
];

interface History {}

export default function History() {
  const { authId } = useAuth();
  const [page, setPage] = useState(1);
  const [history, setHistory] = useState([]);

  const noOfPages =
    history && Math.ceil(history.length / rowsPerPage)
      ? Math.ceil(history.length / rowsPerPage)
      : 1;

  const fetchHistory = async () => {
    try {
      const response = await GET(`histories/${authId}`);
      if (response.data != null) {
        setHistory(response.data.histories || []);
      }
    } catch (error) {
      notifyError(error.message.data.message);
    }
  };

  useEffect(() => {
    fetchHistory();
  }, []);

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    const paginatedHistory = history.slice(start, end);
    const paginatedHistoryArr = [...paginatedHistory];
    return paginatedHistoryArr.map((history, i) => {
      return {
        ...history,
        listId: i + 1 + start,
      };
    });
  }, [page, history]);

  return (
    <div className="mx-auto max-w-7xl px-6 h-4/5 my-10">
      <div className="flex justify-between items-center mb-5">
        <span className="text-3xl">History</span>
      </div>
      <Table
        aria-label="History Table"
        isStriped
        bottomContent={
          <div className="flex w-full justify-center">
            <Pagination
              isCompact
              showControls
              showShadow
              color="secondary"
              page={page}
              total={noOfPages}
              onChange={page => setPage(page)}
            />
          </div>
        }
      >
        <TableHeader columns={columns}>
          {column => (
            <TableColumn key={column.key} align="center">
              {column.label}
            </TableColumn>
          )}
        </TableHeader>
        <TableBody items={items} emptyContent={'No rows to display.'}>
          {item => (
            <TableRow key={item.ID}>
              {columnKey => (
                <TableCell>
                  <HistoryStyleCell item={item} columnKey={columnKey} />
                </TableCell>
              )}
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
