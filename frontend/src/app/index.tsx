import React from "react";
import { useGetMe } from "../api/hooks";
import { Router } from "./router";

export const App = () => {
    const { execute } = useGetMe();

    React.useEffect(() => {
        execute();
    }, []);

    return (
        <>
            <Router />
        </>
    );
};
