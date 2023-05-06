import Navbar from "../pages/Navbar";
import ListSearchedAvailability from "../components/ListSearchedAvailability";

export default function SearchAvailability(){
    return(
        <div>
            <Navbar/>
            <ListSearchedAvailability buttonUrl={"/accomodation-details/"}></ListSearchedAvailability>
        </div>
    );
}