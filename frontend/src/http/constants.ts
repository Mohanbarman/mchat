export const AUTH = {
    LOGIN: "/auth/login",
    REGISTER: "/auth/register",
    GET_ME: "/auth/me",
};

export const CONVERSATIONS = {
    GET_ALL: "/conversations",
};

export const MESSAGES = {
    GET_ALL: (id: string) => `/messages/${id}`,
};
