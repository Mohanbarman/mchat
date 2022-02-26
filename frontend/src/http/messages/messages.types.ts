import { IMessage } from "../../types";
import { IPaginatedPayload } from "../../types/response.type";

export interface IGetMessagesPayload extends IPaginatedPayload {
    conversationId: string;
}

export interface IGetMessagesResponse extends IMessage {}
