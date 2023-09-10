'use client';

import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from '@nextui-org/table';
import { Pagination } from '@nextui-org/pagination';
import { useState, useMemo, useCallback } from 'react';
// import { rows, columns } from './data'
import StyleCell from './style-cell';
import { AppState } from '../redux/store';
import { useSelector } from 'react-redux';

const QuestionsTable = () => {
  const { questionBank } = useSelector((state: AppState) => state.questionBank);
  const [page, setPage] = useState(1);
  const questionBankLength = questionBank.length;
  const rowsPerPage = 10;

  const noOfPages = Math.ceil(questionBankLength / rowsPerPage)
    ? Math.ceil(questionBankLength / rowsPerPage)
    : 1;

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;

    const questions = questionBank.slice(start, end);
    // return questionBank.slice(start, end);
    return questions.map((question, i) => {
      return {
        ...question,
        id: i + 1 + start,
      };
    });
  }, [page, questionBank]);

  const renderCell = useCallback(StyleCell, []);

  const columns = useMemo(() => {
    return [
      { key: 'id', label: 'ID' },
      { key: 'title', label: 'TITLE' },
      { key: 'category', label: 'CATEGORY' },
      { key: 'complexity', label: 'COMPLEXITY' },
      { key: 'actions', label: 'ACTIONS' },
    ];
  }, []);

  return (
    <Table
      aria-label="Questions Table"
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
          <TableRow key={item.id}>
            {columnKey => <TableCell>{renderCell({ item, columnKey })}</TableCell>}
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
};

export default QuestionsTable;
