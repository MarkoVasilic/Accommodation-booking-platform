import React from 'react';
import Dashboard from "./pages/Dashboard";
import UserRegistration from "./pages/UserRegistration";
import Login from "./pages/Login";
import UpdateUser from "./pages/UpdateUser";
import UserProfile from "./pages/UserProfile";
import UserChangePassword from "./pages/UserChangePassword";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { setAuthToken } from "./helpers/sethAuthToken";

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
                        <Route path="/user-profile" element={<UserProfile />} />
                        <Route path="/user-profile/update" element={<UpdateUser />} />
                        <Route path="/user-profile/password/" element={<UserChangePassword />} />
                    </Routes>
                </div>
            </Router>
        </div>
    );
}

export default App;
