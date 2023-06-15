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
import AccommodationsHost from './pages/AccommodationsHost'
import ReservationRequestsHost from './pages/ReservationRequestsHost';
import AccommodationDetailsCard from './components/AccommodationDetailsCard';
import HostsList from './components/ListHosts';
import RateHost from './pages/RateHost';
import HostGradesUser from './pages/HostGradesForUserList';
import UserGradesForHosts from './pages/UserGradesForHostList';
import UpdateHostGrade from './pages/UpdateHostGrade';
import HostGradesHost from './pages/HostGradesForHostList';
import AccommodationsList from './pages/AccommodationsList';
import RateAccommodation from './pages/RateAccommodation';
import AccommodationGradesUser from './pages/AccommodationGradesForUserList';
import UserGradesForAccommodations from './pages/UserGradesForAccommodationList';
import UpdateAccommodationGrade from './pages/UpdateAccommodationGrade';
import AccommodationGradesHost from './pages/AccommodationGradesForHostList';
import UserProfileHost from './pages/UserProfileHost';

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
                        <Route path="/availability/update" element={<UpdateAvailability />} />
                        <Route path="/availabilities" element={<AvailabilityList />} />
                        <Route path="/accepted-reservations" element={<AcceptedReservationsList />} />
                        <Route path="/pending-reservations" element={<PendingReservationsList />} />
                        <Route path="/accommodations/all" element={<SearchAvailability />} />
                        <Route path="/accommodations/all/guest" element={<SearchAvailabilityGuest />} />
                        <Route path="/accomodations/create" element={<CreateAccommodation />} />
                        <Route path="/accommodations/host" element={<AccommodationsHost />} />
                        <Route path="/reservation-requests" element={<ReservationRequestsHost />} />
                        <Route path="/accommodation-details" element={<AccommodationDetailsCard />} />
                        <Route path="/user/hosts-list" element={<HostsList/>} />
                        <Route path="/host/rate" element={<RateHost/>} />
                        <Route path="/host/all-grades" element={<HostGradesUser/>} />
                        <Route path="/user/hosts-grades" element={<UserGradesForHosts/>} /> 
                        <Route path="/host-grade/update" element={<UpdateHostGrade/>} /> 
                        <Route path="/host/my-grades" element={<HostGradesHost/>} /> 
                        <Route path="/user/accommodations-list" element={<AccommodationsList/>} />
                        <Route path="/accommodation/rate" element={<RateAccommodation/>} />
                        <Route path="/accommodation/all-grades" element={<AccommodationGradesUser/>} />
                        <Route path="/user/accommodations-grades" element={<UserGradesForAccommodations/>} /> 
                        <Route path="/accommodation-grade/update" element={<UpdateAccommodationGrade/>} /> 
                        <Route path="/host/accommodation-grades" element={<AccommodationGradesHost/>} />
                        <Route path="/user-profile-host" element={<UserProfileHost/>} />
                    </Routes>
                </div>
            </Router>
        </div>
    );
}

export default App;



