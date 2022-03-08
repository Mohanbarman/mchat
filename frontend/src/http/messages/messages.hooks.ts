import { useToast } from "@chakra-ui/react";
import React from "react";
import { getMessages } from "./messages.apis";
import { IGetMessagesPayload, IGetMessagesResponse } from "./messages.types";

export const useGetMessages = () => {
    const [isLoading, setLoading] = React.useState(false);
    const [cursor, setCursor] = React.useState<string>();
    const [data, setData] = React.useState<IGetMessagesResponse[]>([]);
    const toast = useToast();

    const execute = async (conversationId: string) => {
        setLoading(true);
        const { success, error } = await getMessages({
            limit: 20,
            cursor: cursor,
            conversationId,
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
        setCursor(success.data.page?.next);
        setLoading(false);
    };

    return { data, isLoading, execute };
};
