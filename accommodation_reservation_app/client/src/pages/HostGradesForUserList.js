import Navbar from "../pages/Navbar";
import ListHostGradesForUser from "../components/ListHostGradesForUser";
import AllowedUsers from "../components/AllowedUsers";

export default function HostGradesUser(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListHostGradesForUser></ListHostGradesForUser>
        </div>
    );
}