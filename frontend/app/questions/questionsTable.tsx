'use client';

import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from '@nextui-org/table';
import { Pagination } from '@nextui-org/pagination';
import { useState, useMemo, useCallback, useEffect } from 'react';
import { Question } from '../types/question';
import { getQuestions } from '../questionRequests';
import { ToastContainer, toast } from 'react-toastify';
import StyleCell from './style-cell';
import 'react-toastify/dist/ReactToastify.css';

const QuestionsTable = ({update, setUpdate}) => {
  const [questions, setQuestions] = useState([]);
  const [page, setPage] = useState(1);
  const rowsPerPage = 10;
  const notifyError = (err : string) => toast.warn(err, {
    theme: "dark"
  });

  useEffect(() => {
    const fetchData = () => {
      getQuestions().then(response  => {
        setQuestions(response.data == null ? [] : response.data);
        setUpdate(false);
      })
      .catch(error => {
        notifyError("An error occurred while retrieving questions");
      }) 
    }

    if (update) {
      fetchData();
    }
  }, [update]);

  const noOfPages = Math.ceil(questions.length / rowsPerPage)
    ? Math.ceil(questions.length / rowsPerPage)
    : 1;

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    const paginatedQuestions = questions.slice(start, end);
    return paginatedQuestions.map((question, i) => {
      return {
        ...(question as Question),
        listId: i + 1 + start,
      }
    });

  }, [page, questions]);

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
    <>
      <ToastContainer />
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
              {columnKey => <TableCell>{renderCell({ item, columnKey, setUpdate })}</TableCell>}
            </TableRow>
          )}
        </TableBody>
      </Table>
    </>
  );
};

export default QuestionsTable;
