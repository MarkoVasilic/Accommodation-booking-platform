import RateAccommodationForm from "../components/RateAccommodationForm";
import Navbar from "../pages/Navbar";
import AllowedUsers from "../components/AllowedUsers";

export default function RateAccommodation(){
    //var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <RateAccommodationForm />
        </div>
    );
}