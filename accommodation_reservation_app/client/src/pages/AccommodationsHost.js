import Navbar from "../pages/Navbar";
import ListAccommodationsHost from "../components/ListAccommodationsHost";
import AllowedUsers from "../components/AllowedUsers";

export default function PendingReservationsList(){
    var allowedUsers = ["HOST"]
    return(
        <div>
            <Navbar/>
            <ListAccommodationsHost></ListAccommodationsHost>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
        </div>
    );
}