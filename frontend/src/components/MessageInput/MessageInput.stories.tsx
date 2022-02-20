import { ChakraProvider } from "@chakra-ui/provider";
import { Box, Center } from "@chakra-ui/react";
import { ComponentMeta, ComponentStory } from "@storybook/react";
import { MessageInput } from "./MessageInput";

export default {
    title: "Components/MessageInput",
    component: MessageInput,
} as ComponentMeta<typeof MessageInput>;

const Template: ComponentStory<typeof MessageInput> = (props) => (
    <ChakraProvider>
        <Box height="95vh" display='flex' flexDirection='column' justifyContent='end'>
            <MessageInput {...props} />
        </Box>
    </ChakraProvider>
);

export const Normal = Template.bind({});
Normal.args = {
    onSubmit: console.log,
};
