import React from "react";
import { WS_BASE_URL } from "../config";
import { actions } from "../redux/conversations/conversationSlice";
import { actions as messageAction } from "../redux/messages/messagesSlice";
import { useAppDispatch, useAppSelector } from "../redux/hooks";
import { WsClient } from "./client";

interface IWsContext {
    client: WsClient;
    connect: () => void;
    isConnected: boolean;
}

export const WsContext = React.createContext<Partial<IWsContext>>({});

export const WsProvider = (props: any) => {
    const [isConnected, setIsConnected] = React.useState(false);
    const [wsClient, setWsClient] = React.useState<WsClient>();
    const { isAuthenticated, accessToken } = useAppSelector((s) => s.auth);
    const dispatch = useAppDispatch();

    const connect = () => {
        let socket = new WebSocket(WS_BASE_URL);

        socket.addEventListener("open", () => {
            const client = new WsClient(socket);
            if (!accessToken) return;

            setWsClient(client);
            setIsConnected(true);

            client.login(accessToken);
        });

        socket.addEventListener("error", () => {
            setIsConnected(false);
        });

        socket.addEventListener("close", () => {
            setIsConnected(false);
        });

        socket.addEventListener("message", (event) => {
            const data = JSON.parse(event.data);

            switch (data.event) {
                case "conversation/update":
                    dispatch(actions.update(data.payload));
                    break;
                case "conversation/new_message":
                    dispatch(messageAction.new(data.payload));
                    break;
                case "conversation/sent":
                    dispatch(messageAction.new(data.payload))
                    break;
            }
        });
    };

    React.useEffect(() => {
        if (isAuthenticated && accessToken) {
            connect();
        }
    }, []);

    return <WsContext.Provider value={{ client: wsClient, connect, isConnected }} {...props} />;
};
