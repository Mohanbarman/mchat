import { IReturn } from "../auth/auth.types";
import { CONVERSATIONS } from "../constants";
import api from "../instance";
import {
    IGetConversationPayload,
    IGetConversationsResponse,
} from "./conversation.types";

export const getConversations = async (
    payload: IGetConversationPayload
): IReturn<IGetConversationsResponse[]> => {
    try {
        const res = await api.get(CONVERSATIONS.GET_ALL, {
            params: payload,
        });
        return { success: res };
    } catch (e: any) {
        return { error: e.response };
    }
};
