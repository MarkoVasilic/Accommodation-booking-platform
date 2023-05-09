import Navbar from "../pages/Navbar";
import ListSearchedAvailabilityGuest from "../components/ListSearchedAvailabilityGuest";
import AllowedUsers from "../components/AllowedUsers";

export default function SearchAvailabilityGuest(){
    //var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <ListSearchedAvailabilityGuest buttonUrl={"/accomodation-details/"}></ListSearchedAvailabilityGuest>
        </div>
    );
}