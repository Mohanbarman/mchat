import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { IUser } from "../../types";

export interface IAuthState {
    isLoading: boolean;
    isAuthenticated: boolean;
    user?: IUser;
    accessToken?: string;
}

const getInitialState = (): IAuthState => {
    const state: IAuthState = {
        isLoading: false,
        isAuthenticated: false,
    };

    const userString = localStorage.getItem("user");
    const accessToken = localStorage.getItem("accessToken");

    if (!userString || !accessToken) {
        return state;
    }

    const user = JSON.parse(userString);
    state.isAuthenticated = true;
    state.user = user;
    state.accessToken = accessToken;

    return state;
};

const initialState: IAuthState = getInitialState();

export const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        setUser: (state, action: PayloadAction<IUser>) => {
            state.user = action.payload;
            localStorage.setItem("user", JSON.stringify(action.payload));
        },
        authenticate: (state) => {
            state.isAuthenticated = true;
        },
        unauthenticate: (state) => {
            state.isAuthenticated = false;
            state.user = undefined;
            state.accessToken = undefined;
            localStorage.clear();
        },
        setLoading: (state, action: PayloadAction<boolean>) => {
            state.isLoading = action.payload;
        },
        updateUser: (state, action: PayloadAction<IUser>) => {
            state.user = action.payload;
            localStorage.setItem("user", JSON.stringify(action.payload));
        },
        setAccessToken: (state, action: PayloadAction<string>) => {
            state.accessToken = action.payload;
            localStorage.setItem("accessToken", state.accessToken);
        },
    },
});

export const actions = authSlice.actions;

export default authSlice.reducer;
