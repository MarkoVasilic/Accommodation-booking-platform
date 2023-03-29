import CreateFlightForm from "../components/CreateFlightForm";
import Navbar from "../pages/Navbar";
import AllowedUsers from "../components/AllowedUsers";

export default function CreateFlight(){
    var allowedUsers = ["ADMIN"]
    return(
        <div>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <Navbar/>
            <CreateFlightForm />
        </div>
    );
}