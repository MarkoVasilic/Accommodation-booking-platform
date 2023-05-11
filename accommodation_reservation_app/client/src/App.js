import React from 'react';
import Dashboard from "./pages/Dashboard";
import UserRegistration from "./pages/UserRegistration";
import Login from "./pages/Login";
import UpdateUser from "./pages/UpdateUser";
import UserProfile from "./pages/UserProfile";
import UserChangePassword from "./pages/UserChangePassword";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { setAuthToken } from "./helpers/sethAuthToken";
import CreateAvailability from './pages/CreateAvailability';
import UpdateAvailability from './pages/UpdateAvailability';
import AvailabilityList from './pages/AvailabilityList';
import AcceptedReservationsList from './pages/AcceptedReservationsList';
import PendingReservationsList from './pages/PendingReservationsList';
import SearchAvailability from './pages/SearchedAvailability';
import SearchAvailabilityGuest from './pages/SearchedAvailabilityGuest';
import CreateAccommodation from './pages/CreateAccommodation';

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
                        <Route path="/availability/create" element={<CreateAvailability />} />
                        <Route path="/availability/update/:availability" element={<UpdateAvailability />} />
                        <Route path="/availabilities" element={<AvailabilityList />} />
                        <Route path="/accepted-reservations" element={<AcceptedReservationsList />} />
                        <Route path="/pending-reservations" element={<PendingReservationsList />} />
                        <Route path="/accomodations/all" element={<SearchAvailability />} />
                        <Route path="/accomodations/all/guest" element={<SearchAvailabilityGuest />} />
                        <Route path="/accomodations/create" element={<CreateAccommodation />} />
                    </Routes>
                </div>
            </Router>
        </div>
    );
}

export default App;
