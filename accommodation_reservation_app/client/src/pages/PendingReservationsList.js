import Navbar from "../pages/Navbar";
import ListPendingReservations from "../components/ListPendingReservations";
import AllowedUsers from "../components/AllowedUsers";

export default function PendingReservationsList(){
    var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListPendingReservations></ListPendingReservations>
        </div>
    );
}