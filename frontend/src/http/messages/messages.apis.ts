import { IApiReturn } from "../../types";
import { IGetMessagesPayload, IGetMessagesResponse } from "./messages.types";
import api from "../instance";
import { MESSAGES } from "../constants";

export const getMessages = async (payload: IGetMessagesPayload): IApiReturn<IGetMessagesResponse[]> => {
    try {
        const { conversationId, ...options } = payload;
        const res = await api.get(MESSAGES.GET_ALL(conversationId), {
            params: options,
        });
        return { success: res };
    } catch (e: any) {
        return { error: e.response };
    }
};
