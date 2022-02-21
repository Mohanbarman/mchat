import React from "react";
import { useGetMe } from "../http";
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
