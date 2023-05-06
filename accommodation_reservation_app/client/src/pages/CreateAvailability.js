import CreateAvailabilityForm from "../components/CreateAvailabilityForm";
import Navbar from "../pages/Navbar";
import AllowedUsers from "../components/AllowedUsers";

export default function CreateAvailability(){
    var allowedUsers = ["HOST"]
    return(
        <div>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <Navbar/>
            <CreateAvailabilityForm />
        </div>
    );
}