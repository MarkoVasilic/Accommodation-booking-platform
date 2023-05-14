import AccommodationForm from "../components/AccommodationForm";
import Navbar from "../pages/Navbar";
import AllowedUsers from "../components/AllowedUsers";

export default function CreateAvailability(){
    var allowedUsers = ["HOST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <AccommodationForm />
        </div>
    );
}