import UpdateAvailabilityForm from "../components/UpdateAvailabilityForm";
import AllowedUsers from "../components/AllowedUsers";
import Navbar from "../pages/Navbar";

export default function UpdateAvailability(){
    var allowedUsers = ["HOST"]
    return(
        <div>
        <Navbar></Navbar>
        <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
        <UpdateAvailabilityForm></UpdateAvailabilityForm>
        </div>
    );
}