import Navbar from "../pages/Navbar";
import ListAcceptedReservations from "../components/ListAcceptedReservations";
import AllowedUsers from "../components/AllowedUsers";

export default function AcceptedReservationsList(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListAcceptedReservations></ListAcceptedReservations>
        </div>
    );
}