import React from "react";
import { useAppDispatch, useAppSelector } from "../redux/hooks";
import { Router } from "./router";
import { getAccessToken } from "../http";
import { actions } from "../redux/auth/authSlice";
import { useNavigate } from "react-router-dom";

export const App = () => {
    const { isAuthenticated } = useAppSelector((s) => s.auth);
    const dispatch = useAppDispatch();
    const navigate = useNavigate();

    React.useEffect(() => {
        if (isAuthenticated) {
            let timer: any = undefined;
            timer = setInterval(async () => {
                const refreshToken = localStorage.getItem("refreshToken");
                if (!refreshToken) return;
                const { success, error } = await getAccessToken(refreshToken);
                if (error || !success) {
                    dispatch(actions.unauthenticate);
                    clearInterval(timer);
                    return;
                }
                console.log("newaccess token", success.data.data);
                localStorage.setItem("accessToken", success.data.data.token);
            }, 1000 * 60);
        }
    }, [isAuthenticated]);

    React.useEffect(() => {
        const run = async () => {
            const refreshToken = localStorage.getItem("refreshToken");
            if (!refreshToken) return;
            const { success, error } = await getAccessToken(refreshToken);
            if (error || !success) {
                dispatch(actions.unauthenticate);
                return;
            }
            localStorage.setItem("accessToken", success.data.data.token);
        };
        run();
    }, []);

    React.useEffect(() => {
        if (isAuthenticated) {
            navigate("/");
        } else {
            navigate("/login");
        }
    }, []);

    return <Router />;
};
