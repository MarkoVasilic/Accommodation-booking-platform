import UpdateAvailabilityForm from "../components/UpdateAvailabilityForm";
import AllowedUsers from "../components/AllowedUsers";
import Navbar from "../pages/Navbar";

export default function UpdateAvailability(){
    var allowedUsers = ["HOST"]
    return(
        <div>
        <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
        <Navbar></Navbar>
        <UpdateAvailabilityForm></UpdateAvailabilityForm>
        </div>
    );
}