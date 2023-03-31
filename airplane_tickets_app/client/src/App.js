import "./App.css";
import React from 'react';
import Dashboard from "./pages/Dashboard";
import UserRegistration from "./pages/UserRegistration";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { setAuthToken } from "./helpers/sethAuthToken";
import Login from "./pages/Login";
import FlightDetails from "./pages/FlightDetails";
import SearchhFlights from "./pages/SearchFlights";
import SearchhFlightsUser from "./pages/SearchFlightsUser";

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
                        <Route path="/flight-details" element={<FlightDetails />} />
                        <Route path="/flights/all" element={<SearchhFlights />} />
                        <Route path="/flights/all-user" element={<SearchhFlightsUser />} />
                    </Routes>
                </div>
            </Router>
        </div>
    );
}

export default App;
