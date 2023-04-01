import Navbar from "../pages/Navbar";
import ListSearchedFlights from "../components/ListSearchedFlights";

export default function SearchFlights(){
    return(
        <div>
            <Navbar/>
            <ListSearchedFlights buttonUrl={"/flight-details/"}></ListSearchedFlights>
        </div>
    );
}