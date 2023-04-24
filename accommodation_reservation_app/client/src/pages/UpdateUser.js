import Navbar from "./Navbar";
import UpdateUserForm from "../components/UpdateUserForm"
import AllowedUsers from "../components/AllowedUsers";

export default function UpdateUser(){
    const listOfAllowedUsers = ["HOST", "GUEST"];
    
    return(
        <div>
            <AllowedUsers userRole = {listOfAllowedUsers}/>
            <Navbar />
            <UpdateUserForm></UpdateUserForm>
        </div>

    );
}