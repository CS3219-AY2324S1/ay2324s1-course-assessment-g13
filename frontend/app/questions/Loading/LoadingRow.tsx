import { TableCell, TableRow } from "@nextui-org/table";
import { Skeleton } from "@nextui-org/skeleton"

export default function LoadingRow() {
    return (
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
    )
}