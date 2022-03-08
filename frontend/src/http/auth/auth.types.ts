import { IUser } from "../../types";
import { IApiResponse } from "../../types/response.type";

// request types
export interface ILoginPayload {
    email: string;
    password: string;
}

export interface IRegisterPayload {
    name: string;
    email: string;
    password: string;
    status: string;
}

export interface IRefreshToken {
    accessToken: string;
}

// response types
export interface ILoginResponse extends IUser {
    access_token: string;
    refresh_token: string;
}

export interface IRegisterResponse extends IUser {}
export interface IGetMeResponse extends IUser {}

export interface IRefreshTokenResponse {
    token: string
}

export type IReturn<T> = Promise<IApiResponse<T>>;
