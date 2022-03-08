import { Box } from "@chakra-ui/react";
import { useGetConversations } from "../../http";
import { Chat } from "../../components";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { actions } from "../../redux/conversations/conversationSlice";
import React from "react";
import { useWsClient } from "../../ws";

interface IProps {}

export const Inbox: React.FC<IProps> = (props) => {
    const { data } = useGetConversations();
    const { data: conversations, active } = useAppSelector((s) => s.conversations);
    const dispatch = useAppDispatch();
    const { ws } = useWsClient();

    React.useEffect(() => {
        if (data) dispatch(actions.set(data));
    }, [data]);

    React.useEffect(() => {
        if (active && ws) {
            ws.readConversation(active.id);
        }
    }, [active]);

    return (
        <Box>
            {Object.values(conversations).map((i) => (
                <Chat
                    key={i.id}
                    id={i.id}
                    avatar={i.user.profile}
                    name={i.user.name}
                    messageTime={new Date(i.last_message_time).getTime()}
                    isUnread={i.is_unread}
                    message={i.last_message}
                    unreadCount={i.unread_count}
                    onClick={(id) => dispatch(actions.setActive(id))}
                />
            ))}
        </Box>
    );
};
