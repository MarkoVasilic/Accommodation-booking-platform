import Navbar from "../pages/Navbar";
import ListUserGradesForHost from "../components/ListUserGradesForHosts";
import AllowedUsers from "../components/AllowedUsers";

export default function UserGradesForHosts(){
   // var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <ListUserGradesForHost></ListUserGradesForHost>
        </div>
    );
}