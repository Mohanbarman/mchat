export const AUTH = {
    LOGIN: "/auth/login",
    REGISTER: "/auth/register",
    GET_ME: "/auth/me",
    REFRESH_TOKEN: "/auth/refresh-token"
};

export const CONVERSATIONS = {
    GET_ALL: "/conversations",
    CREATE: "/conversations"
};

export const MESSAGES = {
    GET_ALL: (id: string) => `/messages/${id}`,
};
