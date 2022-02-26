import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { IConversation } from "../../types";

interface IConversationState {
    isLoading: boolean;
    data: IConversation[];
    active?: IConversation;
}

const initialState: IConversationState = {
    isLoading: true,
    data: [],
    active: undefined,
};

export const conversationSlice = createSlice({
    name: "conversation",
    initialState,
    reducers: {
        setLoading: (state, action: PayloadAction<boolean>) => {
            state.isLoading = action.payload;
        },
        set: (state, action: PayloadAction<IConversation[]>) => {
            const data = action.payload.map((i) => {
                i.is_typing = false;
                return i;
            });
            state.data = data;
        },
        add: (state, action: PayloadAction<IConversation>) => {
            const data = { ...action.payload, is_typing: false };
            state.data.unshift(data);
        },
        update: (state, action: PayloadAction<IConversation>) => {
            const index = state.data.findIndex(
                (i) => i.id === action.payload.id
            );
            if (index === -1) return;
            const data = { ...action.payload, is_typing: false };
            state.data[index] = data;
        },
        setActive: (state, action: PayloadAction<string>) => {
            const index = state.data.findIndex((i) => i.id === action.payload);
            if (index === -1) return;
            state.active = state.data[index];
        },
        setTyping: (
            state,
            action: PayloadAction<{ id: string; value: boolean }>
        ) => {
            const index = state.data.findIndex(
                (i) => i.id === action.payload.id
            );
            if (index === -1) return;
            state.data[index].is_typing = action.payload.value;
        },
    },
});

export const actions = conversationSlice.actions;

export default conversationSlice.reducer;
