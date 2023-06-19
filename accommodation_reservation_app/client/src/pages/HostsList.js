import Navbar from "../pages/Navbar";
import ListHosts from "../components/ListHosts";
import AllowedUsers from "../components/AllowedUsers";

export default function HostsList(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListHosts></ListHosts>
        </div>
    );
}