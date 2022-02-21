import React from "react";
import { WsClient } from "./client";
import { WsContext } from "./provider";

export const useWsClient = () => {
    const { client, isConnected } = React.useContext(WsContext);

    return {
        ws: client as WsClient,
        isConnected,
    };
};
