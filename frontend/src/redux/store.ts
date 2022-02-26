import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import authReducer from "./auth/authSlice";
import conversationReducer from "./conversations/conversationSlice";
import messagesReducer from "./messages/messagesSlice";

export const store = configureStore({
    reducer: {
        auth: authReducer,
        conversations: conversationReducer,
        messages: messagesReducer,
    },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<
    ReturnType,
    RootState,
    unknown,
    Action<string>
>;
