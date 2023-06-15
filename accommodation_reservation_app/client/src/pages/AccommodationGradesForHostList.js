import Navbar from "../pages/Navbar";
import ListAccommodationGradesForHost from "../components/ListAccommodationGradesForHost";
import AllowedUsers from "../components/AllowedUsers";

export default function AccommodationGradesHost(){
    //var allowedUsers = ["HOST"]
    return(
        <div>
            <Navbar/>
            <ListAccommodationGradesForHost></ListAccommodationGradesForHost>
        </div>
    );
}