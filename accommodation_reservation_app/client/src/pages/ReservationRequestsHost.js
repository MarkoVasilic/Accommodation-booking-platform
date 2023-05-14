import Navbar from "../pages/Navbar";
import ListReservationRequestsHost from "../components/ListReservationRequestsHost";
import AllowedUsers from "../components/AllowedUsers";

export default function ReservationRequestsHost(){
    var allowedUsers = ["HOST"]
    return(
        <div>
            <Navbar/>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <ListReservationRequestsHost></ListReservationRequestsHost>
        </div>
    );
}