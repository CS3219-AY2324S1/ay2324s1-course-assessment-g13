'use client'

import {Table, TableHeader, TableColumn, TableBody, TableRow, TableCell, getKeyValue} from "@nextui-org/table";
import {Pagination} from "@nextui-org/pagination";
import { useState, useMemo } from "react";

const tableColumns = ['id', 'title', 'category', 'complexity', 'actions']

const rows = [
    {id: 1, title:  "Reverse a String", category:  ["Strings", "Algorithms"], complexity:  "Easy"},
    {id: 2, title:  "Linked List Cycle Detection", category:  ["Data Structures", "Algorithms"], complexity:  "Easy"},
    {id: 3, title:  "Roman to Integer", category:  ["Algorithms"], complexity:  "Easy" },
    {id: 4, title:  "Add Binary Bit Manipulation", category:  ["Algorithms"], complexity:  "Easy" },
    {id: 5, title:  "Fibonacci Number Recursion", category:  ["Algorithms"], complexity:  "Easy" },
    {id: 6, title:  "Implement Stack using Queues", category:  ["Data Structures"], complexity:  "Easy"},
    {id: 7, title:  "Combine Two Tables", category:  ["Databases"], complexity:  "Easy" },
    {id: 8, title:  "Repeated DNA Sequences", category:  ["Algorithms", "Bit Manipulation"], complexity:  "Medium"},
    {id: 9, title:  "Course Schedule", category:  ["Data Structures"], complexity:  "Algorithms"},
    {id: 10, title:  "LRU Cache Design", category:  ["Data Structures"], complexity:  "Medium"},
    {id: 11, title:  "Longest Common Subsequence Strings", category:  ["Algorithms"], complexity:  "Medium"},
    {id: 12, title:  "Rotate Image Arrays", category:  ["Algorithms"], complexity:  "Medium"},
    {id: 13, title:  "Airplane Seat Assignment Probability", category:  ["Brainteaser"], complexity:  "Medium"},
    {id: 14, title:  "Validate Binary Search Tree", category:  ["Data Structures", "Algorithms"], complexity:  "Medium"},
    {id: 15, title:  "Sliding Window Maximum Arrays", category:  ["Algorithms"], complexity:  "Hard" },
    {id: 16, title:  "N-Queen Problem", category:  ["Algorithms"], complexity:  "Hard"},
    {id: 17, title:  "Serialize and Deserialize a Binary Tree", category:  ["Data Structures", "Algorithms"], complexity:  "Hard"},
    {id: 18, title:  "Wildcard Matching Strings", category:  ["Algorithms"], complexity:  "Hard"  },
    {id: 19, title:  "Chalkboard XOR Game", category:  ["Brainteaser"], complexity:  "Hard"},
    {id: 20, title:  "Trips and Users", category:  ["Databases"], complexity:  "Hard" },
]

const columns = tableColumns.map(col =>  {return {key: col, label: col.toUpperCase()}})

export default function QuestionsTable() {
  const [page , setPage] = useState(1);
  const rowsPerPage = 12;

  const noOfPages = Math.ceil(rows.length / rowsPerPage);
  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;

    return rows.slice(start, end);
  }, [page, rows])

  return (
    <Table 
      aria-label="Example table with dynamic content"
      bottomContent={
        <div className="flex w-full justify-center">
          <Pagination
            isCompact
            showControls
            showShadow
            color="secondary"
            page={page}
            total={noOfPages}
            onChange={(page) => setPage(page)}
          />
        </div>
      }
    >
      <TableHeader columns={columns}>
        {(column) => <TableColumn key={column.key} align="center">{column.label}</TableColumn>}
      </TableHeader>
      <TableBody items={items} emptyContent={"No rows to display."}>
        {(item) => (
          <TableRow key={item.key}>
            {(columnKey) => <TableCell>{getKeyValue(item, columnKey)}</TableCell>}
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
