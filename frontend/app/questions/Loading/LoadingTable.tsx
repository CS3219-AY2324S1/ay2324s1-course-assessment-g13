"use client";
import { Table, TableBody, TableHeader, TableColumn, TableCell, TableRow } from "@nextui-org/table";
import { Skeleton } from "@nextui-org/skeleton";
import { useMemo } from "react";

interface LoadingRowDimentsion {
    w: number;
    h: number;
}

export default function LoadingTable() {

    const columns = useMemo(() => {
        return [
          { key: 'id', label: 'ID' },
          { key: 'title', label: 'TITLE' },
          { key: 'category', label: 'CATEGORY' },
          { key: 'complexity', label: 'COMPLEXITY' },
          { key: 'actions', label: 'ACTIONS' },
        ];
      }, []);

    const LoadingRowDimension : LoadingRowDimentsion[][][]= [
        [[{w:5,h:5}], [{w:25,h:5}], [{w:28,h:7}, {w:20,h:7}], [{w:16,h:7}], [{w:5,h:5}, {w:5,h:5}]],
        [[{w:5,h:5}], [{w:48,h:5}], [{w:20,h:7}], [{w:20,h:7}], [{w:5,h:5}, {w:5,h:5}]],
        [[{w:5,h:5}], [{w:40,h:5}], [{w:16,h:7}, {w:24,h:7}], [{w:24,h:7}], [{w:5,h:5}, {w:5,h:5}]],
        [[{w:5,h:5}], [{w:32,h:5}], [{w:28,h:7}, {w:24,h:7}], [{w:20,h:7}], [{w:5,h:5}, {w:5,h:5}]],
        [[{w:5,h:5}], [{w:36,h:5}], [{w:20,h:7}, {w:24,h:7}], [{w:28,h:7}], [{w:5,h:5}, {w:5,h:5}]]
    ]

    return (
        <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
            <div className="questions-header flex justify-between items-center mb-5">
            <span className="text-3xl">Question Bank</span>
                <Skeleton className="w-32 rounded-xl">
                    <div className="h-10 bg-default-300"></div>
                </Skeleton>
            </div>
            <div className="table w-full">
                <Table aria-label="Loading Questions Table">
                <TableHeader columns={columns}>
                    {column => (
                    <TableColumn key={column.key} align="center">
                        {column.label}
                    </TableColumn>
                    )}
                </TableHeader>
                    <TableBody>
                        {LoadingRowDimension.map((row, rowIndex) => (
                            <TableRow key={`row ${rowIndex}`}>
                                {row.map((col, colIndex) => (
                                    <TableCell key={`col ${colIndex}`} className={(colIndex === 2 || colIndex === 4) && "flex gap-3"}>
                                        {col.map((cell, cellIndex) => (
                                            <Skeleton key={cellIndex} className={`w-${cell.w} rounded-lg`}>
                                                <div className={`h-${cell.h} bg-default-300`}></div>
                                            </Skeleton>
                                        ))}
                                    </TableCell>
                                ))
                                }
                            </TableRow>
                        ))
                        }
                    </TableBody>
                </Table>
            </div>
        </div>
    )
}
