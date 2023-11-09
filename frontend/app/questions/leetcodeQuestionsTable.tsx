'use client';

import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from '@nextui-org/table';
import { Pagination } from '@nextui-org/pagination';
import { useState, useMemo, useEffect } from 'react';
import { Question } from '../types/question';
import StyleCell from './style-cell';
import axios from 'axios';

interface ApiResponse {
    total: number;
    problems: Question[];
}

const leetcodeQuestionsURL = `https://asia-southeast1-peer-preps-assignment6.cloudfunctions.net/GetProblems`

const LeetCodeQuestionsTable = () => {
  const [page, setPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const [questions, setQuestions] = useState<Question[]>([]);
  const rowsPerPage = 10;

  const fetchLeetcodeQuestions = async () => {
    try {
        const queryParams = {
          'offset': 0,
          'page-size': 100,
        }
        const response = await axios.get<ApiResponse>(leetcodeQuestionsURL, {
        params: queryParams
        });
        const {problems} = response.data;
        setTotalPages(Math.ceil(100/rowsPerPage));
        setQuestions(problems);
    } catch (error) {
        console.error("Unable to get leetcode questions")
    }
  }

  useEffect(() => {
    fetchLeetcodeQuestions();
  }, []);

  const handlePageChange = (newPage: number) => {
    setPage(newPage);
  }

  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;
    const paginatedQuestions = questions.slice(start, end);
    const paginatedQuestionsArr = [...paginatedQuestions]
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
      bottomContent={
        <div className="flex w-full justify-center">
          <Pagination
            isCompact
            showControls
            showShadow
            color="secondary"
            page={page}
            total={totalPages}
            onChange={handlePageChange}
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
            {columnKey => 
              <TableCell>
                <StyleCell item={item} columnKey={columnKey} isLeetCode={true} />
              </TableCell>
            }
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
};

export default LeetCodeQuestionsTable;
