import {
    Modal,
    ModalOverlay,
    ModalContent,
    ModalHeader,
    ModalCloseButton,
    ModalBody,
    FormControl,
    FormLabel,
    Input,
    ModalFooter,
    Button,
    FormErrorMessage,
} from "@chakra-ui/react";
import { yupResolver } from "@hookform/resolvers/yup";
import React from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import * as yup from "yup";

interface IProps {
    isOpen: boolean;
    onClose: () => void;
    onSubmit: (data: string) => any;
    error: string;
}

interface IForm {
    email: string;
}

const schema = yup.object().shape({
    email: yup.string().required().email(),
});

export const UserAddModal: React.FC<IProps> = (props) => {
    const {
        register,
        handleSubmit,
        formState: { errors },
    } = useForm<IForm>({
        resolver: yupResolver(schema),
    });

    const onSubmit: SubmitHandler<IForm> = async ({ email }) => {
        await props.onSubmit(email);
    };

    return (
        <Modal isOpen={props.isOpen} onClose={props.onClose}>
            <ModalOverlay />
            <form onSubmit={handleSubmit(onSubmit)}>
                <ModalContent>
                    <ModalHeader>Start conversation with a user</ModalHeader>
                    <ModalCloseButton />
                    <ModalBody pb={6}>
                        <FormControl isInvalid={!!props.error || !!errors.email}>
                            <FormLabel>Email</FormLabel>
                            <Input {...register("email")} placeholder="john@example.com" />
                            {props.error && <FormErrorMessage>{props.error}</FormErrorMessage>}
                            {errors.email && <FormErrorMessage>{errors.email.message}</FormErrorMessage>}
                        </FormControl>
                    </ModalBody>

                    <ModalFooter>
                        <Button type="submit" colorScheme="blue" mr={3}>
                            Add
                        </Button>
                        <Button onClick={props.onClose}>Cancel</Button>
                    </ModalFooter>
                </ModalContent>
            </form>
        </Modal>
    );
};
