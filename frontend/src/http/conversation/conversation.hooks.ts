import { useToast } from "@chakra-ui/react";
import React from "react";
import { createConversation } from ".";
import { IApiResponse } from "../../types";
import { getConversations } from "./conversation.apis";
import {
    ICreateConversationPayload,
    ICreateConversationResponse,
    IGetConversationsResponse,
} from "./conversation.types";

export const useGetConversations = () => {
    const [isLoading, setLoading] = React.useState(false);
    const [cursor, setCursor] = React.useState("");
    const [data, setData] = React.useState<IGetConversationsResponse[]>([]);
    const toast = useToast();

    React.useEffect(() => {
        (async () => {
            setLoading(true);
            const { success, error } = await getConversations({
                limit: 10,
                cursor: cursor,
            });

            if (error || !success) {
                toast({
                    title: "Failed to load conversations",
                    description: error ? error.data.message : "Something went wrong",
                    status: "error",
                    duration: 4000,
                });
                return;
            }

            setData(success.data.data);
            setLoading(false);
        })();
    }, []);

    return { data, isLoading };
};

export const useCreateConversation = () => {
    const [error, setError] = React.useState("");

    const execute = async (payload: ICreateConversationPayload): Promise<IApiResponse<ICreateConversationResponse>> => {
        const { success, error } = await createConversation(payload);

        if (error || !success) {
            if (error?.data.message) {
                setError(error.data.message);
            } else if (error?.data.errors) {
                setError(error.data.errors["email"][0]);
            }
            return { error };
        }

        return { success };
    };

    return { execute, error };
};
