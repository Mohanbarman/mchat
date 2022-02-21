import { useToast } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import { actions } from "../../redux/auth/authSlice";
import { useAppDispatch, useAppSelector } from "../../redux/hooks";
import { login, register, getMe } from "./auth.apis";

export const useLogin = () => {};
export const useRegister = () => {};

export const useGetMe = () => {
    const authState = useAppSelector((s) => s.authReducer);
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

        const [data, err] = await getMe();

        if (err) {
            if (err.status === 401) {
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
        dispatch(actions.setUser(data.data.data));
    };

    return { execute };
};
