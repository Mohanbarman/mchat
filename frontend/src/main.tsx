import React from "react";
import ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { ChakraProvider } from "@chakra-ui/react";
import { MemoryRouter } from "react-router-dom";
import { store } from "./redux/store";
import { App } from "./app";
import { WsProvider } from "./ws";

ReactDOM.render(
    <React.StrictMode>
        <Provider store={store}>
            <ChakraProvider>
                <MemoryRouter>
                    <WsProvider>
                        <App />
                    </WsProvider>
                </MemoryRouter>
            </ChakraProvider>
        </Provider>
    </React.StrictMode>,
    document.getElementById("root")
);
