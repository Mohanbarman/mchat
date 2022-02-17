import {
    FormErrorMessage,
    FormControl,
    FormLabel,
    Button,
    Input,
    InputGroup,
    InputRightElement,
} from "@chakra-ui/react";
import React from "react";
import { Link, useNavigate } from "react-router-dom";
import { Container } from "./Container";
import { useRegister } from "./useRegister";

export const Register = () => {
    const { onSubmit, errors, loading, registerField } = useRegister();

    const [show, setShow] = React.useState(false);
    const handleClick = () => setShow(!show);

    return (
        <Container heading="Creating a new account">
            <form onSubmit={onSubmit}>
                <FormControl isInvalid={!!errors.name} mb="20px">
                    <FormLabel htmlFor="name">Full Name</FormLabel>
                    <Input
                        {...registerField("name")}
                        colorScheme="teal"
                        id="name"
                        type="text"
                        placeholder="John"
                    />
                    {errors.name && (
                        <FormErrorMessage>
                            {errors.name.message}
                        </FormErrorMessage>
                    )}
                </FormControl>
                <FormControl isInvalid={!!errors.email} mb="20px">
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
                </FormControl>
                <FormControl isInvalid={!!errors.password} mb="20px">
                    <FormLabel htmlFor="password">Password</FormLabel>
                    <Input
                        colorScheme="teal"
                        {...registerField("password")}
                        id="password"
                        placeholder="********"
                        type="password"
                    />
                    {errors.password && (
                        <FormErrorMessage>
                            {errors.password.message}
                        </FormErrorMessage>
                    )}
                </FormControl>
                <FormControl isInvalid={!!errors.confirmPassword}>
                    <FormLabel htmlFor="confirmPassword">
                        Confirm Password
                    </FormLabel>
                    <InputGroup size="md">
                        <Input
                            {...registerField("confirmPassword")}
                            pr="4.5rem"
                            type={show ? "text" : "password"}
                            placeholder="*********"
                        />
                        <InputRightElement width="4.5rem">
                            <Button h="1.75rem" size="sm" onClick={handleClick}>
                                {show ? "Hide" : "Show"}
                            </Button>
                        </InputRightElement>
                    </InputGroup>
                    {errors.confirmPassword && (
                        <FormErrorMessage>
                            {errors.confirmPassword.message}
                        </FormErrorMessage>
                    )}
                </FormControl>
                <Button
                    type="submit"
                    isLoading={loading}
                    w="100%"
                    colorScheme="teal"
                    mt="30px"
                >
                    Submit
                </Button>
                <Link to="/login">
                    <Button w="100%" mt="30px">
                        Login with existing account
                    </Button>
                </Link>
            </form>
        </Container>
    );
};
