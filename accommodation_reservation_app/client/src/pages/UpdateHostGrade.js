import Navbar from "./Navbar";
import UpdateHostGradeForm from "../components/UpdateHostGradeForm"
import AllowedUsers from "../components/AllowedUsers";

export default function UpdateHostGrade(){
    const listOfAllowedUsers = ["GUEST"];
    
    return(
        <div>
            <Navbar />
            <AllowedUsers userRole = {listOfAllowedUsers}></AllowedUsers>
            <UpdateHostGradeForm></UpdateHostGradeForm>
        </div>

    );
}