import React from "react";
import { useAppSelector } from "../redux/hooks";
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
    const { isAuthenticated, accessToken } = useAppSelector(
        (s) => s.authReducer
    );

    const connect = () => {
        let socket = new WebSocket("ws://localhost:8080/ws/");

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
    };

    const shouldConnect = () => {
        if (!isConnected && isAuthenticated && accessToken) {
            connect();
        }
    };

    React.useEffect(() => {
        shouldConnect();
    }, []);

    React.useEffect(() => {
        shouldConnect();
    }, [isConnected]);

    return (
        <WsContext.Provider
            value={{ client: wsClient, connect, isConnected }}
            {...props}
        />
    );
};
