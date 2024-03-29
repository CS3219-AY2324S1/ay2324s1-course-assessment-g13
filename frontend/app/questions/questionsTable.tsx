'use client';

import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from '@nextui-org/table';
import { Pagination } from '@nextui-org/pagination';
import { useState, useMemo, useEffect } from 'react';
import { Question } from '../types/question';
import StyleCell from './style-cell';

interface QuestionProps {
  questions: Question[];
  fetchQuestions: () => void;
}

const QuestionsTable = ({ questions, fetchQuestions }: QuestionProps) => {
  const [page, setPage] = useState(1);
  const rowsPerPage = 10;

  useEffect(() => {
    fetchQuestions();
  }, []);

  const noOfPages =
    questions && Math.ceil(questions.length / rowsPerPage)
      ? Math.ceil(questions.length / rowsPerPage)
      : 1;

  useEffect(() => {
    setPage(1);
  }, [noOfPages]);

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    const paginatedQuestions = questions.slice(start, end);
    const paginatedQuestionsArr = [...paginatedQuestions];
    return paginatedQuestionsArr.map((question: Question, i: number) => {
      return {
        ...(question as Question),
        listId: i + 1 + start,
      };
    });
  }, [page, questions]);

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
          <TableRow key={item.id}>
            {columnKey => (
              <TableCell>
                <StyleCell item={item} columnKey={columnKey} fetchQuestions={fetchQuestions} />
              </TableCell>
            )}
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
};

export default QuestionsTable;
