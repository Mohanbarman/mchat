import axios from "axios";

const api = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL as string,
});

export const login = async (email: string, password: string) => {
    try {
        const res = await api.post("/auth/login", { email, password });
        return [res, null];
    } catch (e: any) {
        return [null, e.response];
    }
};

interface IRegisterPayload {
    name: string;
    email: string;
    password: string;
    status: string;
}

export const register = async (data: IRegisterPayload) => {
    try {
        const res = await api.post("/auth/register", data);
        return [res, null];
    } catch (e: any) {
        return [null, e.response];
    }
};

export const getMe = async (token: string) => {
    try {
        const res = await api.get("/auth/me", {
            headers: { Authorization: `Bearer ${token}` },
        });
        return [res, null];
    } catch (e: any) {
        return [null, e.response];
    }
};
