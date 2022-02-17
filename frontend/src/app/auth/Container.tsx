import {
    Box,
    Flex,
    useMediaQuery,
    Image,
    Heading,
    Text,
} from "@chakra-ui/react";
import React from "react";
import landingImg from "../../assets/landing.png";

interface IProps {
    children: any;
    heading: string;
}

export const Container: React.FC<IProps> = ({ children, heading }) => {
    const [isDesktop] = useMediaQuery("(min-width: 800px)");
    return (
        <Flex minH="100vh">
            {isDesktop && (
                <Flex
                    width="50vw"
                    minH="100vh"
                    bg="gray.100"
                    align="center"
                    justify="center"
                    flexDir="column"
                    padding="30px"
                >
                    <Image
                        padding="30px"
                        maxW="600px"
                        w="100%"
                        src={landingImg}
                    />
                    <Heading textAlign="center" mb="10px">
                        Welcome to our{"  "}
                        <Box display="inline" color="teal">
                            Mchat
                        </Box>
                        <br />A Messaging app
                    </Heading>
                    <Text fontSize="lg">
                        Talk to any person in your mother language
                    </Text>
                </Flex>
            )}
            <Box width={isDesktop ? "50vw" : "100vw"} bg="gray.50" minH="100vh">
                <Flex
                    flexDirection="column"
                    align="center"
                    justify="center"
                    padding="30px"
                    minH="100vh"
                >
                    <Heading mb="40px">{heading}</Heading>
                    <Flex flexDirection="column" width="100%" maxW="400px">
                        {children}
                    </Flex>
                </Flex>
            </Box>
        </Flex>
    );
};
