import { GET } from "./axios";

export const refreshToken = async () => {
    await GET('/auth/refresh');
}
