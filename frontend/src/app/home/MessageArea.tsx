import { Box, Flex } from "@chakra-ui/react";
import { ChatMenu, MessageInput } from "../../components";
import { useAppSelector } from "../../redux/hooks";

interface IProps {
    selectedConversation: string;
}

export const MesssageArea: React.FC<IProps> = (props) => {
    const { user } = useAppSelector((s) => s.authReducer);

    if (!user || !props.selectedConversation) {
        return <></>;
    }

    return (
        <Flex direction="column" position="relative" height="100%">
            <ChatMenu profile={user.profile} />
            <Flex flex="1" justifyContent="flex-end" direction="column">
                <Box maxH="calc(100vh - (75px + 64px))" bg="red"></Box>
            </Flex>
            <Box>
                <MessageInput onSubmit={(value) => console.log(value)} />
            </Box>
        </Flex>
    );
};
