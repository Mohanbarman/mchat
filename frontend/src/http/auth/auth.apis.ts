import { AUTH } from "../constants";
import api from "../instance";

export const login = async (email: string, password: string) => {
    try {
        const res = await api.post(AUTH.LOGIN, { email, password });
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
        const res = await api.post(AUTH.REGISTER, data);
        return [res, null];
    } catch (e: any) {
        return [null, e.response];
    }
};

export const getMe = async () => {
    try {
        const res = await api.get(AUTH.GET_ME);
        return [res, null];
    } catch (e: any) {
        return [null, e.response];
    }
};
