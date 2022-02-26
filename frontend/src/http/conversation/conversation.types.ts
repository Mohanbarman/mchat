import { IApiResponse, IPaginatedPayload } from "../../types/response.type";
import { IConversation } from "../../types";

export interface IGetConversationPayload extends IPaginatedPayload {}

export interface IGetConversationsResponse extends IConversation {}

export type IReturn<T> = Promise<IApiResponse<T>>;
