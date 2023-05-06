import Navbar from "../pages/Navbar";
import ListAvailability from "../components/ListAvailability";
import AllowedUsers from "../components/AllowedUsers";

export default function AvailabilityList(){
    var allowedUsers = ["HOST"]
    return(
        <div>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <Navbar/>
            <ListAvailability></ListAvailability>
        </div>
    );
}