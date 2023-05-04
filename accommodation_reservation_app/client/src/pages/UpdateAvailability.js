import UpdateAvailabilityForm from "../components/UpdateAvailabilityForm";
//import AllowedUsers from "../components/AllowedUsers";
import Navbar from "../pages/Navbar";

export default function UpdateAvailability(){
    //var listOfAllowedUsers = ["Admin", "TranfusionCenterStaff"]
    return(
        <div>
        <Navbar></Navbar>
        <UpdateAvailabilityForm></UpdateAvailabilityForm>
        </div>
    );
}