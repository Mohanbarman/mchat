import { useToast } from '@chakra-ui/react';
import React from 'react'
import { getConversations } from './conversation.apis';

export const useGetConversations = () => {
    const [isLoading, setLoading] = React.useState(false);
    const [cursor, setCursor] = React.useState("");
    const [data, setData] = React.useState([]);
    const toast = useToast();

    React.useEffect(() => {
        (async () => {
            setLoading(true);
            const [res, err] = await getConversations(15, cursor);
            if (err) {
                toast({
                    title: "Failed to load conversations",
                    description: err.data.message,
                    status: "error",
                    duration: 4000,
                });
            }
            setData(res.data.data);
            setLoading(false);
        })();
    }, []);

    return { data, isLoading };
};