import { GET } from "./axios";

export const refreshToken = async () => {
    try {
        console.log("Refreshing Token...")
        await GET('/auth/refresh');
    } catch (error) {
        throw error;
    }
}
