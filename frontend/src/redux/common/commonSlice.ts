import { createSlice } from "@reduxjs/toolkit";

export interface ICommonState {
    isAddUserModalOpen: boolean;
}

const initialState: ICommonState = {
    isAddUserModalOpen: false,
};

export const commonSlice = createSlice({
    name: "common",
    initialState,
    reducers: {
        openAddUserModal: (state) => {
            state.isAddUserModalOpen = true;
        },
        closeAddUserModal: (state) => {
            state.isAddUserModalOpen = false;
        },
    },
});

export const actions = commonSlice.actions;

export default commonSlice.reducer;
