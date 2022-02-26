import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { useDispatch } from "react-redux";
import { login, useLogin } from "../../http";
import { actions } from "../../redux/auth/authSlice";
import { handleFieldErrors } from "../../helpers/validation";
import * as yup from "yup";
import { useNavigate } from "react-router-dom";
import { useToast } from "@chakra-ui/react";
import React from "react";

const schema = yup.object().shape({
    email: yup.string().required().email("please enter a valid email"),
    password: yup.string().required(),
});

interface IForm {
    email: string;
    password: string;
}

export const useAppLogin = () => {
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
        setError,
    } = useForm<IForm>({ resolver: yupResolver(schema) });
    const { execute, errors: apiErrors, isLoading, data } = useLogin();

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const toast = useToast();

    // on error
    React.useEffect(() => {
        if (!apiErrors) return;
        handleFieldErrors(apiErrors, setError);
    }, [isLoading, apiErrors]);

    // on success
    React.useEffect(() => {
        if (!data) return;
        const { access_token, refresh_token, ...user } = data;

        dispatch(actions.authenticate());
        dispatch(actions.setAccessToken(access_token));
        dispatch(actions.setRefreshToken(refresh_token))
        dispatch(actions.setUser(user));

        toast({
            title: "Logged in successfully",
            description: "Congratulations You've successfully logged in.",
            status: "success",
            duration: 4000,
            isClosable: true,
        });

        navigate("/");
    }, [data]);

    const onSubmit = async (form: IForm) => execute(form);

    return {
        registerField: register,
        onSubmit: handleSubmit(onSubmit),
        errors,
        loading: isSubmitting,
    };
};
