import Navbar from "../pages/Navbar";
import ListAcceptedReservations from "../components/ListAcceptedReservations";
import AllowedUsers from "../components/AllowedUsers";

export default function AcceptedReservationsList(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <Navbar/>
            <ListAcceptedReservations></ListAcceptedReservations>
        </div>
    );
}