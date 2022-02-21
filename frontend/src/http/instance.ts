import axios from "axios";

const instance = axios.create({
    baseURL: import.meta.env.VITE_BASE_URL as string,
});

instance.interceptors.request.use((config) => {
    const token = localStorage.getItem("accessToken");
    if (token) {
        config.headers = {
            ...config.headers,
            Authorization: `Bearer ${token}`,
        };
    }
    return config
});

export default instance;
