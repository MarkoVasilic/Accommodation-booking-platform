import Navbar from "../pages/Navbar";
import ListHostGradesForHost from "../components/ListHostGradesForHost";
import AllowedUsers from "../components/AllowedUsers";

export default function HostGradesHost(){
    //var allowedUsers = ["HOST"]
    return(
        <div>
            <Navbar/>
            <ListHostGradesForHost></ListHostGradesForHost>
        </div>
    );
}