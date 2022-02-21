import { CONVERSATIONS } from "../constants";
import api from "../instance";

export const getConversations = async (limit: number, cursor: string) => {
    try {
        const res = await api.get(CONVERSATIONS.GET_ALL, {
            params: { limit, cursor },
        });
        return [res, null];
    } catch (e: any) {
        return [null, e.response];
    }
};
