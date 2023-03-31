import Navbar from "../pages/Navbar";
import ListSearchedFlightsUser from "../components/ListSearchedFlightsUser";

export default function SearchFlightsUser(){
    return(
        <div>
            <Navbar/>
            <ListSearchedFlightsUser></ListSearchedFlightsUser>
        </div>
    );
}