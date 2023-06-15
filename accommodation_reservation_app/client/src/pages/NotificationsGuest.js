import Navbar from "../pages/Navbar";
import ListNotificationsGuest from "../components/ListNotificationsGuest";

export default function NotificationsGuest(){
    return(
        <div>
            <Navbar/>
            <ListNotificationsGuest></ListNotificationsGuest>
        </div>
    );
}