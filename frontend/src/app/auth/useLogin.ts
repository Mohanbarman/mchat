import { useForm } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { useDispatch } from "react-redux";
import { login } from "../../api/apis";
import { actions } from "../../redux/auth/authSlice";
import { handleFieldErrors } from "../../helpers/validation";
import * as yup from "yup";
import { useNavigate } from "react-router-dom";
import { useToast } from "@chakra-ui/react";

const schema = yup.object().shape({
    email: yup.string().required().email("please enter a valid email"),
    password: yup.string().required(),
});

interface IForm {
    email: string;
    password: string;
}

export const useLogin = () => {
    const {
        register,
        handleSubmit,
        formState: { errors, isSubmitting },
        setError,
    } = useForm<IForm>({ resolver: yupResolver(schema) });

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const toast = useToast();

    const onSubmit = async (data: Record<string, string>) => {
        const [res, err] = await login(data.email, data.password);
        if (err) {
            handleFieldErrors(err.data?.errors, setError);
            return;
        }
        const user = res.data.data
        dispatch(actions.authenticate());
        dispatch(actions.setAccessToken(user.access_token))
        delete user.access_token
        dispatch(actions.setUser(user))
        toast({
            title: "Logged in successfully",
            description: "Congratulations You've successfully logged in.",
            status: "success",
            duration: 4000,
            isClosable: true,
        });
        navigate("/");
    };

    return {
        registerField: register,
        onSubmit: handleSubmit(onSubmit),
        errors,
        loading: isSubmitting,
    };
};
