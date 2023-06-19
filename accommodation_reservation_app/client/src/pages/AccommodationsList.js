import Navbar from "../pages/Navbar";
import ListAccommodations from "../components/ListAccommodations";
import AllowedUsers from "../components/AllowedUsers";

export default function AccommodationsList(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListAccommodations></ListAccommodations>
        </div>
    );
}