import Navbar from "../pages/Navbar";
import ListTicketsUser from "../components/ListTicketsUser";
import AllowedUsers from "../components/AllowedUsers";

export default function ListTicketsUserPage(){
    var allowedUsers = ["REGULAR"]
    return(
        <div>
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <Navbar/>
            <ListTicketsUser></ListTicketsUser>
        </div>
    );
}