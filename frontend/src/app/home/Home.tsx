import { Box, Flex } from "@chakra-ui/react";
import React from "react";
import { ChatMenu } from "../../components";
import { commonActions } from "../../redux/common";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { AddUser } from "./AddUser";
import { Inbox } from "./Inbox";
import { MesssageArea } from "./MessageArea";

export const Home = () => {
    const { user } = useAppSelector((s) => s.auth);
    const dispatch = useAppDispatch();

    if (!user) {
        return <></>;
    }

    return (
        <Flex height="100vh" width="100vw">
            <AddUser />
            <Box flex={2} minW="400px" maxW="500px" borderRight="2px solid var(--chakra-colors-gray-300)">
                <ChatMenu
                    onAddUser={() => {
                        dispatch(commonActions.openAddUserModal());
                    }}
                    profile={user.profile}
                />
                <Inbox />
            </Box>
            <Box flex={4} bg="gray.50">
                <MesssageArea />
            </Box>
        </Flex>
    );
};
