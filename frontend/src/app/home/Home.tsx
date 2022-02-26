import { Box, Flex } from "@chakra-ui/react";
import React from "react";
import { ChatMenu } from "../../components";
import { useAppSelector } from "../../redux/hooks";
import { Inbox } from "./Inbox";
import { MesssageArea } from "./MessageArea";

export const Home = () => {
    const { user } = useAppSelector((s) => s.auth);

    if (!user) {
        return <></>;
    }

    return (
        <Flex height="100vh" width="100vw">
            <Box
                flex={2}
                minW="400px"
                maxW="500px"
                borderRight="2px solid var(--chakra-colors-gray-300)"
            >
                <ChatMenu profile={user.profile} />
                <Inbox />
            </Box>
            <Box flex={4} bg="gray.50">
                <MesssageArea />
            </Box>
        </Flex>
    );
};
