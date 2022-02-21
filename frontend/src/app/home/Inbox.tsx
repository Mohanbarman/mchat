import { Box } from "@chakra-ui/react";
import { useEffect } from "react";
import { useGetConversations } from "../../http";
import { Chat } from "../../components";
import { useWsClient } from "../../ws";

interface IProps {
    onClick: (value: string) => any
}

export const Inbox: React.FC<IProps> = (props) => {
    const { data, isLoading } = useGetConversations();
    const { ws, isConnected } = useWsClient();

    // useEffect(() => {
    //     if (isConnected) {
    //         ws.sendMessage({
    //             text: "Heloo world",
    //             userID: "3e0887fe-7a97-4cde-8aed-797397bd9724",
    //         });
    //     }
    // }, [isConnected]);

    return (
        <Box>
            {data.map((i: any) => (
                <Chat
                    key={i.id}
                    id={i.id}
                    avatar={i.user.profile}
                    name={i.user.name}
                    messageTime={new Date(i.created_at).getTime()}
                    isUnread={false}
                    message="Hello world"
                    unreadCount={1}
                    onClick={props.onClick}
                />
            ))}
        </Box>
    );
};
