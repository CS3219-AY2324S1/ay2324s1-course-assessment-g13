"use client"
import { useRouter } from "next/navigation";
import useAuth from "../hooks/useAuth";
import { useEffect, useMemo, useState } from "react";
import { notifyError, notifySuccess } from "../components/toast/notifications";
import { GET, POST } from "../libs/axios/axios";
import { AuthResponse } from "../(auth)/login/page";
import { AxiosResponse } from "axios";
import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from "@nextui-org/table";
import { Pagination } from "@nextui-org/pagination";
import { Chip } from "@nextui-org/chip";
import { Button } from "@nextui-org/button";

interface GetAuthUsersResponse {
    users : AuthResponse[]
}

interface GetUserResponse {
    user : Omit<User, "oauthId" | "oauthProvider" | "role">
}

interface ChangeRoleRequest {
    oauth_id: number;
    oauth_provider: string
}

interface MessageResponse {
    message: string
}

type User = {
    oauthId: number;
    oauthProvider: string;
    role: string;
    username: string;
}

const UserTableColumn = [
    {key:"index", label: "INDEX"},
    {key:"username", label: "USERNAME"},
    {key:"oauth_id", label: "OAUTH ID"},
    {key:"oauth_provider", label: "OAUTH PROVIDER"},
    {key:"role", label: "ROLE"},
    {key:"action", label: "ACTION"},
]

export default function ManageUser() {
    const { role } = useAuth();
    const [page, setPage] = useState(1);
    const rowsPerPage = 10;
    const [users, setUsers] = useState<User[]>([]);
    const router = useRouter();
    let noOfPages: number;

    const getUsers = async () => {
        try {
            setUsers([]);
            const authUsersResponse: AxiosResponse<GetAuthUsersResponse> = await GET("/auth/users");
            const { users: authUsers } = authUsersResponse.data
            authUsers.map(async authUser => {
                const authUserId = authUser.ID;
                try {
                    const userResponse: AxiosResponse<GetUserResponse> = await GET(`/users/${authUserId}`);
                    const { user } = userResponse.data
                    const _user : User = {
                        oauthId: authUser.oauth_id,
                        oauthProvider: authUser.oauth_provider,
                        role: authUser.role,
                        username: user.username
                    }
                    setUsers(prev => [...prev, _user])
                } catch (error) {
                    const message = error.message.data.message;
                    notifyError(message);
                }
            })
        } catch (error) {
            const message = error.message.data.message;
            notifyError(message);
        }
    }

    useEffect(() => {
        if (role !== "super admin") {
            router.back();
            notifyError("Not Super Admin");
        }
    }, [role])

    useEffect(() => {
        getUsers();
    }, [])

    useEffect(() => {
        noOfPages = Math.ceil(users.length/rowsPerPage) | 0
    }, [users])

    const items = useMemo(() => {
      const start = (page - 1) * rowsPerPage;
      const end = start + rowsPerPage;
      const paginatedQuestions = users.slice(start, end);
      const paginatedQuestionsArr = [...paginatedQuestions]
      return paginatedQuestionsArr.filter(user => user.role !== "super admin").map((user: User, i: number) => {
        return {
          ...(user as User),
          listId: i + 1 + start,
        };
      });
    }, [page, users]);
    
    const handleChangeRole = async (user: User, isDowngrade: boolean) => {
        const url = `/auth/user/${isDowngrade ? "downgrade" : "upgrade"}`
        const requestBody : ChangeRoleRequest = {
            oauth_id: user.oauthId,
            oauth_provider: user.oauthProvider
        }
        try {
            const response : AxiosResponse<MessageResponse> = await POST(url, requestBody);
            const { message } = response.data;
            notifySuccess(message);
            getUsers();
        } catch (error) {
            const message = error.message.data.message;
            notifyError(message);
        }
    }

return (
    <div className="questions mx-auto max-w-7xl px-6 h-4/5 my-10">
      <div className="questions-header flex justify-between items-center mb-5">
        <span className="text-3xl">Active Users</span>
      </div>
      <div className="table w-full">
        <Table
        aria-label="Users Table"
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
        <TableHeader columns={UserTableColumn}>
            {column => (
              <TableColumn key={column.key} align="center">
                {column.label}
              </TableColumn>
            )}
        </TableHeader>
        <TableBody items={items} emptyContent={"No Active Users"}>
            {(item) => (
            <TableRow key={item.listId}>
                <TableCell>{item.listId}</TableCell>
                <TableCell>{item.username}</TableCell>
                <TableCell>{item.oauthId}</TableCell>
                <TableCell>{item.oauthProvider.toUpperCase()}</TableCell>
                <TableCell>
                  <Chip color={item.role === "admin" ? "secondary" : "primary"}>  
                    {item.role.toUpperCase()}
                  </Chip>
                </TableCell>
                <TableCell>
                  <Button 
                    className="text-white text-small w-4/5"
                    size="sm"
                    color={item.role === "admin" ? "danger" : "success"}
                    onClick={() => handleChangeRole(item, item.role === "admin")}
                >
                    {item.role === "admin" ? "DOWNGRADE" : "UPGRADE"}
                  </Button>
                </TableCell>
            </TableRow>
            )}
        </TableBody>
        </Table>
      </div>
    </div>
  );
}
