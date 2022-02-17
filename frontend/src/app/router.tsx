import { Login, Register } from "./auth";
import { Home } from "./home";
import { Routes, Route } from "react-router-dom";

export const Router = () => {
    return (
        <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/" element={<Home />}></Route>
        </Routes>
    );
};
