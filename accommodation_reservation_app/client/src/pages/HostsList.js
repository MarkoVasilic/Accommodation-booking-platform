import Navbar from "../pages/Navbar";
import ListHosts from "../components/ListHosts";
import AllowedUsers from "../components/AllowedUsers";

export default function AvailabilityList(){
    //var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <ListHosts></ListHosts>
        </div>
    );
}