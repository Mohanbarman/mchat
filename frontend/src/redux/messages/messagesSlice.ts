import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { IMessage } from "../../types";

export interface IMessageState {
    data: Record<string, IMessage[]>;
}

const initialState: IMessageState = { data: {} };

export const messageSlice = createSlice({
    name: "messages",
    initialState,
    reducers: {
        add: (state, action: PayloadAction<IMessage[]>) => {
            for (const message of action.payload) {
                if (state.data[message.conversation]) {
                    state.data[message.conversation].unshift(message);
                } else {
                    state.data[message.conversation] = [message];
                }
            }
        },
        new: (state, action: PayloadAction<IMessage>) => {
            state.data[action.payload.conversation].push(action.payload);
        },
    },
});

export const actions = messageSlice.actions;

export default messageSlice.reducer;
