import { Send } from "react-feather";
import { Box, Flex, FormControl, IconButton, Input, Textarea } from "@chakra-ui/react";
import React from "react";
import { IProps } from "./MessageInput.types";

export const MessageInput: React.FC<IProps> = (props) => {
    const [value, setValue] = React.useState("");
    const newLines = value.split("\n");
    const rows = newLines.length - 1 > 4 ? 5 : newLines.length;

    const onSubmit = () => {
        props.onSubmit(value);
        setValue("");
    };

    return (
        <Box bg="gray.200" width="100%" padding="10px">
            <Flex gap="10px" alignItems="end">
                <FormControl bg="white" borderRadius="10px">
                    <Textarea
                        rows={rows}
                        textOverflow="none"
                        variant="outline"
                        padding="10px"
                        resize="none"
                        value={value}
                        onChange={(e) => setValue(e.target.value)}
                        onKeyDown={(e) => {
                            if ((e.keyCode === 10 || e.keyCode === 13) && e.ctrlKey) {
                                onSubmit();
                            }
                        }}
                    />
                </FormControl>
                <IconButton
                    color="gray.600"
                    variant="ghost"
                    aria-label="send message"
                    outline="none"
                    onClick={onSubmit}
                >
                    <Send />
                </IconButton>
            </Flex>
        </Box>
    );
};
