import { AxiosResponse } from "axios";

export interface IResponse<T> {
    code: number;
    data: T;
    success: boolean;
    page?: IPage;
}

export interface IErrorResponse {
    code: number;
    errors?: Record<string, string>;
    success: boolean;
    message?: string;
}

export interface IPage {
    next: string | undefined;
}

export interface IApiResponse<T> {
    success?: AxiosResponse<IResponse<T>>;
    error?: AxiosResponse<IErrorResponse>;
}

export interface IPaginatedPayload {
    limit: number;
    cursor?: string;
}

export type IApiReturn<T> = Promise<IApiResponse<T>>;
