import { useToast } from "@chakra-ui/react";
import React from "react";
import { getConversations } from "./conversation.apis";
import { IGetConversationsResponse } from "./conversation.types";

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
                    description: error
                        ? error.data.message
                        : "Something went wrong",
                    status: "error",
                    duration: 4000,
                });
                return
            }

            setData(success.data.data);
            setLoading(false);
        })();
    }, []);

    return { data, isLoading };
};
