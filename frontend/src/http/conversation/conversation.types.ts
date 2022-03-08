import { IApiResponse, IPaginatedPayload } from "../../types/response.type";
import { IConversation } from "../../types";

export interface IGetConversationPayload extends IPaginatedPayload {}

export interface IGetConversationsResponse extends IConversation {}

export interface ICreateConversationResponse extends IConversation {}

export interface ICreateConversationPayload {
    email: string;
}

export type IReturn<T> = Promise<IApiResponse<T>>;
