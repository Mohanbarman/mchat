import { Center, ChakraProvider } from "@chakra-ui/react";
import { ComponentMeta, ComponentStory } from "@storybook/react";
import Message from "./Message";

export default {
    title: "Components/Message",
    component: Message,
} as ComponentMeta<typeof Message>;

const Template: ComponentStory<typeof Message> = (args) => (
    <ChakraProvider>
        <Center>
            <Message {...args} />
        </Center>
    </ChakraProvider>
);

export const Text = Template.bind({});
Text.args = {
    text: "Hello, How are you doing ?",
    type: "text",
    time: new Date(),
    state: "seen",
};

export const LongText = Template.bind({});
LongText.args = {
    text: "Storybook helps you build UI components in isolation from your app's business logic, data, and context. That makes it easy to develop hard-to-reach states. Save these UI states as stories to revisit during development, testing, or QA. ",
    type: "text",
    time: new Date(),
    state: "seen",
};
