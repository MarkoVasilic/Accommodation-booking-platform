import Navbar from "../pages/Navbar";
import ListSearchedFlightsAdmin from "../components/ListSearchedFlightsAdmin";
import AllowedUsers from "../components/AllowedUsers";

export default function SearchFlightsAdmin(){
    var allowedUsers = ["ADMIN"]
    return(
        <div>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <Navbar/>
            <ListSearchedFlightsAdmin buttonUrl={"/flight-details/"}></ListSearchedFlightsAdmin>
        </div>
    );
}