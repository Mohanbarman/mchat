import { AUTH } from "../constants";
import api from "../instance";
import {
    IGetMeResponse,
    ILoginPayload,
    ILoginResponse,
    IRefreshTokenResponse,
    IRegisterPayload,
    IRegisterResponse,
    IReturn,
} from "./auth.types";

export const login = async (payload: ILoginPayload): IReturn<ILoginResponse> => {
    try {
        const res = await api.post(AUTH.LOGIN, payload);
        return { success: res };
    } catch (e: any) {
        return { error: e.response };
    }
};

export const register = async (data: IRegisterPayload): IReturn<IRegisterResponse> => {
    try {
        const res = await api.post(AUTH.REGISTER, data);
        return { success: res };
    } catch (e: any) {
        return { error: e.response };
    }
};

export const getMe = async (): IReturn<IGetMeResponse> => {
    try {
        const res = await api.get(AUTH.GET_ME);
        return { success: res };
    } catch (e: any) {
        return { error: e.response };
    }
};

export const getAccessToken = async (token: string): IReturn<IRefreshTokenResponse> => {
    try {
        const res = await api.post(AUTH.REFRESH_TOKEN, { token });
        return { success: res };
    } catch (e: any) {
        return { error: e.reponse };
    }
};
