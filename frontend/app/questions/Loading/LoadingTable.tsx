"use client";
import { Table, TableBody, TableHeader, TableColumn, TableCell, TableRow } from "@nextui-org/table";
import { Skeleton } from "@nextui-org/skeleton";
import { useMemo } from "react";

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
    return (
        <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
            <div className="questions-header flex justify-between items-center mb-5">
            <span className="text-3xl">Question Bank</span>
                <Skeleton className="w-32 rounded-xl">
                    <div className="h-10 bg-default-300"></div>
                </Skeleton>
            </div>
            <div className="table w-full">
                <Table>
                <TableHeader columns={columns}>
                    {column => (
                    <TableColumn key={column.key} align="center">
                        {column.label}
                    </TableColumn>
                    )}
                </TableHeader>
                    <TableBody>
                        <TableRow>
                            <TableCell>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-28 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-28 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-20 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-16 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                        </TableRow>
                        <TableRow>
                            <TableCell>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-48 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-20 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-20 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                        </TableRow>
                        <TableRow>
                            <TableCell>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-40 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-16 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-24 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-24 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                        </TableRow>
                        <TableRow>
                            <TableCell>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-32 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-28 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-24 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-20 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                        </TableRow>
                        <TableRow>
                            <TableCell>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-36 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-20 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-24 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell>
                                <Skeleton className="w-28 rounded-xl">
                                    <div className="h-7 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                            <TableCell className="flex flex-row gap-3">
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                                <Skeleton className="w-5 rounded-lg">
                                    <div className="h-5 bg-default-300"></div>
                                </Skeleton>
                            </TableCell>
                        </TableRow>
                    </TableBody>
                </Table>
            </div>
        </div>
    )
}
