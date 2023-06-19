import Navbar from "../pages/Navbar";
import ListUserGradesForAccommodation from "../components/ListUserGradesForAccommodations";
import AllowedUsers from "../components/AllowedUsers";

export default function UserGradesForAccommodations(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListUserGradesForAccommodation></ListUserGradesForAccommodation>
        </div>
    );
}