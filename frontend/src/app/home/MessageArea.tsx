import { Box, Flex } from "@chakra-ui/react";
import React from "react";
import { ChatMenu, Message, MessageInput } from "../../components";
import { useGetMessages } from "../../http";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { actions } from "../../redux/messages/messagesSlice";
import { useWsClient } from "../../ws";

export const MesssageArea: React.FC = () => {
    const { user } = useAppSelector((s) => s.auth);
    const { active } = useAppSelector((s) => s.conversations);
    const { data: messages } = useAppSelector((s) => s.messages);
    const { execute, data } = useGetMessages();
    const dispatch = useAppDispatch();
    const { ws } = useWsClient();

    React.useEffect(() => {
        if (!data) return;
        dispatch(actions.add(data));
    }, [data]);

    React.useEffect(() => {
        if (!active) return;
        execute(active.id);
    }, [active]);

    React.useEffect(() => {
        const container = document.getElementById("messages-container");
        if (!container) return;
        container.scrollTop = container.scrollHeight;
    }, [messages]);

    if (!user || !active) {
        return <></>;
    }

    const conversationMessages = Object.values(messages[active.id] || {});

    const status: Record<number, "sent" | "delivered" | "seen"> = {
        0: "sent",
        1: "delivered",
        3: "seen",
    };

    const onMessageSubmit = (message: string) => {
        ws.sendMessage({
            text: message,
            userID: active.user.id,
        });
    };

    return (
        <Flex direction="column" position="relative" height="100%">
            <ChatMenu onAddUser={console.log} profile={user.profile} />
            <Flex flex="1" justifyContent="flex-end" direction="column">
                <Box height="100%" id="messages-container" position="relative" overflowY="scroll" overflowX="hidden">
                    <Box
                        position="absolute"
                        top="0"
                        left="0"
                        display="flex"
                        flexDir="column"
                        padding="10px 30px"
                        gap="15px"
                        width="100%"
                    >
                        {conversationMessages.map((i) => (
                            <Message
                                key={i.id}
                                id={i.id}
                                isMe={i.is_me}
                                state={status[i.status]}
                                text={i.text}
                                time={new Date(i.created_at)}
                                type="text"
                            />
                        ))}
                    </Box>
                </Box>
            </Flex>
            <Box>
                <MessageInput onSubmit={onMessageSubmit} />
            </Box>
        </Flex>
    );
};
