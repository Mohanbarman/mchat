import { MoreVertical, Plus, LogOut, Settings } from "react-feather";
import {
    Avatar,
    Box,
    Flex,
    IconButton,
    Menu,
    MenuButton,
    MenuItem,
    MenuList,
} from "@chakra-ui/react";
import React from "react";
import { IProps } from "./ChatMenu.types";

export const ChatMenu: React.FC<IProps> = (props) => {
    return (
        <Box
            padding="13px 12px"
            bg="gray.100"
            width="100%"
            borderBottom="1px solid var(--chakra-colors-gray-200)"
        >
            <Flex
                direction="row"
                justifyContent="space-between"
                alignItems="center"
            >
                <Box cursor='pointer'>
                    <Avatar src={props.profile} />
                </Box>
                <Menu>
                    <MenuButton
                        as={IconButton}
                        aria-label="Options"
                        icon={<MoreVertical />}
                    />
                    <MenuList>
                        <MenuItem icon={<Plus size="18px" />}>
                            Add user
                        </MenuItem>
                        <MenuItem icon={<Settings size="18px" />}>
                            Settings
                        </MenuItem>
                        <MenuItem icon={<LogOut size="18px" />}>
                            Logout
                        </MenuItem>
                    </MenuList>
                </Menu>
            </Flex>
        </Box>
    );
};
