import React from "react";
import { Box, Text } from "@chakra-ui/react";
import { IMessageProps } from "./Message.types";

const DoubleTick = () => {
    return (
        <svg viewBox="0 0 16 15" width="16" height="15">
            <path
                fill="currentColor"
                d="m15.01 3.316-.478-.372a.365.365 0 0 0-.51.063L8.666 9.879a.32.32 0 0 1-.484.033l-.358-.325a.319.319 0 0 0-.484.032l-.378.483a.418.418 0 0 0 .036.541l1.32 1.266c.143.14.361.125.484-.033l6.272-8.048a.366.366 0 0 0-.064-.512zm-4.1 0-.478-.372a.365.365 0 0 0-.51.063L4.566 9.879a.32.32 0 0 1-.484.033L1.891 7.769a.366.366 0 0 0-.515.006l-.423.433a.364.364 0 0 0 .006.514l3.258 3.185c.143.14.361.125.484-.033l6.272-8.048a.365.365 0 0 0-.063-.51z"
            ></path>
        </svg>
    );
};

const SingleTick = () => (
    <svg viewBox="0 0 16 15" width="16" height="15">
        <path
            fill="currentColor"
            d="m10.91 3.316-.478-.372a.365.365 0 0 0-.51.063L4.566 9.879a.32.32 0 0 1-.484.033L1.891 7.769a.366.366 0 0 0-.515.006l-.423.433a.364.364 0 0 0 .006.514l3.258 3.185c.143.14.361.125.484-.033l6.272-8.048a.365.365 0 0 0-.063-.51z"
        ></path>
    </svg>
);

const SeenTick = () => (
    <Box color="blue.400">
        <DoubleTick />
    </Box>
);

const DeliveredTick = () => (
    <Box color="gray.500">
        <DoubleTick />
    </Box>
);

const SentTick = () => (
    <Box color="gray.500">
        <SingleTick />
    </Box>
);
export const Message: React.FC<IMessageProps> = (props) => {
    const time = props.time.toLocaleTimeString("en-US", {
        hour: "2-digit",
        minute: "2-digit",
    });
    const direction = props.isMe ? "right" : "left";
    const bg = props.isMe ? "teal.100" : "teal.50";

    console.log(props)

    return (
        <Box
            bg={bg}
            borderRadius="10px"
            onClick={() => props.onClick(props.id)}
            padding="7px 10px"
            display="inline-block"
            float={direction}
            position="relative"
            boxShadow="#00000036 0px 1px 2px"
            maxW="600px"
            {...{
                [props.isMe ? "borderTopRightRadius" : "borderTopLeftRadius"]:
                    "0",
            }}
        >
            <Text fontSize="0.9rem" fontWeight="medium" whiteSpace="pre-wrap">
                    {props.text}
            </Text>
            <Box display="flex" justifyContent="flex-end">
                <Text
                    fontSize="0.75rem"
                    fontWeight="medium"
                    color="blackAlpha.600"
                >
                    {time}
                </Text>
                {props.isMe && (
                    <Box padding="2px" paddingLeft="4px">
                        {props.state === "delivered" && <DeliveredTick />}
                        {props.state === "seen" && <SeenTick />}
                        {props.state === "sent" && <SentTick />}
                    </Box>
                )}
            </Box>
            <Box
                color={bg}
                position="absolute"
                top="0"
                {...{
                    [direction]: "-8px",
                }}
                transform={props.isMe ? "unset" : "scaleX(-1)"}
            >
                <svg viewBox="0 0 8 13" width="8" height="13">
                    <path
                        opacity=".13"
                        d="M5.188 1H0v11.193l6.467-8.625C7.526 2.156 6.958 1 5.188 1z"
                    ></path>
                    <path
                        fill="currentColor"
                        d="M5.188 0H0v11.193l6.467-8.625C7.526 1.156 6.958 0 5.188 0z"
                    ></path>
                </svg>
            </Box>
        </Box>
    );
};
