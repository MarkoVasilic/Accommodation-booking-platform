import "./App.css";
import React from 'react';
import Dashboard from "./pages/Dashboard";
import UserRegistration from "./pages/UserRegistration";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { setAuthToken } from "./helpers/sethAuthToken";
import Login from "./pages/Login";

function App() {
    const token = localStorage.getItem("token");
    if (token) {
        setAuthToken(token);
    }
    
    return (
        <div className="App">
            <Router>
                <div className="content">
                    <Routes>
                        <Route path="/" element={<Dashboard />} />
                        <Route path="/signup" element={<UserRegistration />} />
                        <Route path="/login" element={<Login />} />
                    </Routes>
                </div>
            </Router>
        </div>
    );
}

export default App;
