import { useToast } from "@chakra-ui/react";
import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { useNavigate } from "react-router-dom";
import { register } from "../../http";
import { handleFieldErrors } from "../../helpers/validation";
import * as yup from "yup";

const schema = yup.object().shape({
    name: yup.string().required(),
    email: yup.string().required().email("please enter a valid email"),
    password: yup
        .string()
        .required()
        .min(8, "password must be greater than 8 characters")
        .max(50, "password must be smaller than 50 characters"),
    confirmPassword: yup
        .string()
        .oneOf([yup.ref("password")], "Password didn't matched"),
});

interface IForm {
    email: string;
    password: string;
    confirmPassword: string;
    name: string;
}

export const useRegister = () => {
    const {
        register: registerField,
        handleSubmit,
        formState: { errors, isSubmitting },
        setError,
    } = useForm<IForm>({ resolver: yupResolver(schema) });

    const toast = useToast();
    const navigate = useNavigate();

    const onSubmit = async (data: IForm) => {
        const { error } = await register({
            email: data.email,
            name: data.name,
            password: data.password,
            status: "Hey there I'm using Mchat",
        });
        if (error && error.data.errors) {
            handleFieldErrors(error.data.errors, setError);
            return;
        }
        toast({
            title: "Account created.",
            description: "We've created your account for you.",
            status: "success",
            duration: 9000,
            isClosable: true,
        });
        navigate("/");
    };

    return {
        registerField,
        onSubmit: handleSubmit(onSubmit),
        errors,
        loading: isSubmitting,
    };
};
