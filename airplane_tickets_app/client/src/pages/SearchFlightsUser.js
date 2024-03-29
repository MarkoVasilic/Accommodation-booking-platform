import Navbar from "../pages/Navbar";
import ListSearchedFlightsUser from "../components/ListSearchedFlightsUser";
import AllowedUsers from "../components/AllowedUsers";

export default function SearchFlightsUser(){
    var allowedUsers = ["REGULAR"]
    return(
        <div>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <Navbar/>
            <ListSearchedFlightsUser buttonUrl={"/flight-details/"}></ListSearchedFlightsUser>
        </div>
    );
}