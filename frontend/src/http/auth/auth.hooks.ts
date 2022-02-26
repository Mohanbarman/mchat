import { useToast } from "@chakra-ui/react";
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { actions } from "../../redux/auth/authSlice";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { IUser } from "../../types";
import { login, register, getMe } from "./auth.apis";
import { ILoginPayload, ILoginResponse, IRegisterPayload } from "./auth.types";

export const useLogin = () => {
    const [isLoading, setLoading] = useState(false);
    const [errors, setErrors] = useState<Record<string, string>>();
    const [data, setData] = useState<ILoginResponse>();

    const execute = async (payload: ILoginPayload) => {
        setLoading(true);
        const { error, success } = await login(payload);
        setLoading(false);
        if (error || !success) {
            setErrors(error?.data.errors);
            return;
        }
        setData(success.data.data);
    };

    return { execute, isLoading, errors, data };
};

export const useRegister = () => {
    const [isLoading, setLoading] = useState(false);
    const [errors, setErrors] = useState<Record<string, string>>();
    const [data, setData] = useState<IUser>();

    const execute = async (payload: IRegisterPayload) => {
        setLoading(true);
        const { error, success } = await register(payload);
        setLoading(false);
        if (error || !success) {
            setErrors(error?.data.errors);
            return;
        }
        setData(success.data.data);
    };

    return { execute, isLoading, errors, data };
};

export const useGetMe = () => {
    const authState = useAppSelector((s) => s.auth);
    const navigate = useNavigate();
    const dispatch = useAppDispatch();
    const toast = useToast();

    const execute = async () => {
        if (
            (!authState.isAuthenticated && !authState.isLoading) ||
            authState.accessToken === undefined
        ) {
            navigate("/login");
            dispatch(actions.unauthenticate());
            return;
        }

        const { success, error } = await getMe();

        if (error || !success) {
            if (error && error.status === 401) {
                dispatch(actions.unauthenticate());
                navigate("/login");
            }

            toast({
                description: "Please try again later",
                duration: 5000,
                status: "error",
                title: "Something went wrong",
            });
            return;
        }

        dispatch(actions.authenticate());
        dispatch(actions.setUser(success.data.data));
    };

    return { execute };
};
