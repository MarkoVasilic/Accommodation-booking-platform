import CreateAvailabilityForm from "../components/CreateAvailabilityForm";
import Navbar from "../pages/Navbar";
//import AllowedUsers from "../components/AllowedUsers";

export default function CreateAvailability(){
   // var allowedUsers = ["ADMIN"]
    return(
        <div>
            <Navbar/>
            <CreateAvailabilityForm />
        </div>
    );
}