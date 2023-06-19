import Navbar from "../pages/Navbar";
import ListAccommodationGradesForUser from "../components/ListAccommodationGradesForUser";
import AllowedUsers from "../components/AllowedUsers";

export default function AccommodationGradesUser(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListAccommodationGradesForUser></ListAccommodationGradesForUser>
        </div>
    );
}