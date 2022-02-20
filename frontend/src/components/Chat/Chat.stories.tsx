import { Center, ChakraProvider } from "@chakra-ui/react";
import { ComponentMeta, ComponentStory } from "@storybook/react";
import { Chat } from "./Chat";

export default {
    title: "Components/Chat",
    component: Chat,
    argTypes: {
        messageTime: {
            control: {
                type: "date",
            },
        },
    },
} as ComponentMeta<typeof Chat>;

const Template: ComponentStory<typeof Chat> = (args) => (
    <ChakraProvider>
        <Center>
            <Chat {...args} />
        </Center>
    </ChakraProvider>
);

export const NormalChat = Template.bind({});
NormalChat.args = {
    name: "John Doe",
    avatar: "https://api.uifaces.co/our-content/donated/bUkmHPKs.jpg",
    isUnread: false,
    message: "How are you doing ?",
    messageTime: Date.now(),
    unreadCount: 0,
};

export const UnreadChat = Template.bind({});
UnreadChat.args = {
    name: "John Doe",
    avatar: "https://api.uifaces.co/our-content/donated/bUkmHPKs.jpg",
    isUnread: true,
    message: "How are you doing ?",
    messageTime: Date.now(),
    unreadCount: 70,
};

export const LongMessage = Template.bind({});
LongMessage.args = {
    name: "John Doe",
    avatar: "https://api.uifaces.co/our-content/donated/bUkmHPKs.jpg",
    isUnread: true,
    messageTime: Date.now(),
    unreadCount: 70,
    message: "Storybook helps you build UI components in isolation from your app's business logic, data, and context. That makes it easy to develop hard-to-reach states. Save these UI states as stories to revisit during development, testing, or QA. ",
};
