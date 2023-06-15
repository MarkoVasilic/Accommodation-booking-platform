import Navbar from "../pages/Navbar";
import ListAccommodationGradesForUser from "../components/ListAccommodationGradesForUser";
import AllowedUsers from "../components/AllowedUsers";

export default function AccommodationGradesUser(){
    //var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <ListAccommodationGradesForUser></ListAccommodationGradesForUser>
        </div>
    );
}