import {
    FormControl,
    FormLabel,
    Button,
    Input,
    FormErrorMessage,
    InputGroup,
    InputRightElement,
} from "@chakra-ui/react";
import { Link } from "react-router-dom";
import { Container } from "./Container";
import { useAppLogin } from "./useAppLogin";
import React from "react";

export const Login = () => {
    const { errors, loading, onSubmit, registerField } = useAppLogin();

    const [show, setShow] = React.useState(false);
    const handleClick = () => setShow(!show);

    return (
        <Container heading="Login">
            <form onSubmit={onSubmit}>
                <FormControl isInvalid={!!errors.email}>
                    <FormLabel htmlFor="email">Email</FormLabel>
                    <Input
                        {...registerField("email")}
                        colorScheme="teal"
                        id="email"
                        placeholder="john@example.com"
                    />
                    {errors.email && (
                        <FormErrorMessage>
                            {errors.email.message}
                        </FormErrorMessage>
                    )}
                    <FormControl isInvalid={!!errors.password}>
                        <FormLabel mt="20px" htmlFor="password">
                            Password
                        </FormLabel>
                        <InputGroup size="md">
                            <Input
                                {...registerField("password")}
                                pr="4.5rem"
                                type={show ? "text" : "password"}
                                placeholder="*********"
                            />
                            <InputRightElement width="4.5rem">
                                <Button
                                    h="1.75rem"
                                    size="sm"
                                    onClick={handleClick}
                                >
                                    {show ? "Hide" : "Show"}
                                </Button>
                            </InputRightElement>
                        </InputGroup>
                        {errors.password && (
                            <FormErrorMessage>
                                {errors.password.message}
                            </FormErrorMessage>
                        )}
                    </FormControl>
                    <Button
                        w="100%"
                        isLoading={loading}
                        type="submit"
                        colorScheme="teal"
                        mt="30px"
                    >
                        Submit
                    </Button>
                    <Link to="/register">
                        <Button w="100%" mt="30px">
                            Create an account
                        </Button>
                    </Link>
                </FormControl>
            </form>
        </Container>
    );
};
