import { ChakraProvider } from "@chakra-ui/provider";
import { Box, Center } from "@chakra-ui/react";
import { ComponentMeta, ComponentStory } from "@storybook/react";
import { UserAddModal } from "./UserAddModal";

export default {
    title: "Components/UserAddModal",
    component: UserAddModal,
} as ComponentMeta<typeof UserAddModal>;

const Template: ComponentStory<typeof UserAddModal> = (props) => (
    <ChakraProvider>
        <Box height="95vh" display="flex" flexDirection="column" justifyContent="end">
            <UserAddModal {...props} />
        </Box>
    </ChakraProvider>
);

export const Normal = Template.bind({});
Normal.args = {
    isOpen: true,
    onClose: () => console.log("close clicked"),
    error: "",
    onSubmit: (value) => console.log(value),
};
