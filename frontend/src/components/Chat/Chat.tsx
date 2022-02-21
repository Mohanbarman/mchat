import { Box, Flex, Text, useDisclosure, Avatar } from "@chakra-ui/react";
import * as helpers from "../../helpers";
import React from "react";
import { IProps } from "./Chat.types";

export const Chat: React.FC<IProps> = (props) => {
    const messageTime = new Date(props.messageTime);

    const time = helpers.parseTime(messageTime);
    const date = helpers.parseFriendlyDate(messageTime);
    const showDate = helpers.getDayDiff(messageTime) > 1;

    const { isOpen, onClose, onOpen } = useDisclosure();

    return (
        <Box
            width="100%"
            maxW="600px"
            onMouseEnter={onOpen}
            onMouseLeave={onClose}
            bg={isOpen ? "gray.50" : "white"}
            transition="all .3s"
            cursor="pointer"
            borderBottom="1px solid #0000000d"
            onClick={() => props.onClick(props.id)}
        >
            <Flex
                direction="row"
                justifyContent="space-between"
                padding="10px 14px"
            >
                <Flex direction="row" gap="20px">
                    <Avatar
                        src={props.avatar}
                        width="60px"
                        height="60px"
                        borderRadius="50%"
                    />
                    <Flex direction="column" justifyContent="space-between">
                        <Text fontWeight="normal" fontSize="1.2rem">
                            {props.name}
                        </Text>
                        <Text
                            fontWeight={props.isUnread ? "medium" : "normal"}
                            color={props.isUnread ? "black" : "gray.600"}
                            fontSize="1.1rem"
                            noOfLines={1}
                        >
                            {props.message}
                        </Text>
                    </Flex>
                </Flex>
                <Flex
                    direction="column"
                    justifyContent="space-between"
                    alignItems="end"
                    minW="100px"
                >
                    <Text
                        color={props.isUnread ? "teal.500" : "gray.500"}
                        fontSize=".95rem"
                        fontWeight={props.isUnread ? "bold" : "medium"}
                    >
                        {showDate ? date : time}
                    </Text>
                    {props.isUnread && (
                        <Box
                            bg="teal.500"
                            height="25px"
                            width="25px"
                            display="flex"
                            justifyContent="center"
                            alignItems="center"
                            borderRadius="50%"
                        >
                            <Text
                                color="white"
                                fontWeight="bold"
                                fontSize=".7rem"
                            >
                                {props.unreadCount > 99
                                    ? "99"
                                    : props.unreadCount}
                            </Text>
                        </Box>
                    )}
                </Flex>
            </Flex>
        </Box>
    );
};
