import { Center, ChakraProvider } from "@chakra-ui/react";
import { ComponentMeta, ComponentStory } from "@storybook/react";
import { ChatMenu } from "./ChatMenu";

export default {
    title: "Components/ChatMenu",
    component: ChatMenu,
} as ComponentMeta<typeof ChatMenu>;

const Template: ComponentStory<typeof ChatMenu> = (args) => (
    <ChakraProvider>
        <Center>
            <ChatMenu {...args} />
        </Center>
    </ChakraProvider>
);

export const Normal = Template.bind({});
Normal.args = {
    profile: "https://api.uifaces.co/our-content/donated/bUkmHPKs.jpg",
};
