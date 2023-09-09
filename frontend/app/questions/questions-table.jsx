"use client"

import { Table, TableHeader, TableColumn, TableBody, TableRow, TableCell } from "@nextui-org/table";
import { Pagination } from "@nextui-org/pagination";
import { useState, useMemo, useCallback } from "react";
import { rows, columns } from './data'
import { styleCell } from "./style-cell";

export default function QuestionsTable() {
  const [page , setPage] = useState(1);
  const rowsPerPage = 10;

  const noOfPages = Math.ceil(rows.length / rowsPerPage);
  const items = useMemo(() => {
    const start = (page - 1) * rowsPerPage;
    const end = start + rowsPerPage;

    return rows.slice(start, end);
  }, [page, rows]);

  const renderCell = useCallback(styleCell, []);

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
            {(columnKey) => <TableCell>{renderCell(item, columnKey)}</TableCell>}
          </TableRow>
        )}
      </TableBody>
    </Table>
  );
}
