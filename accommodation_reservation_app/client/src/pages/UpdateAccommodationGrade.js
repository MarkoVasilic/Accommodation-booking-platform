import Navbar from "./Navbar";
import UpdateAccommodationGradeForm from "../components/UpdateAccommodationGradeForm"
import AllowedUsers from "../components/AllowedUsers";

export default function UpdateAccommodationGrade(){
    var allowedUsers = ["GUEST"];
    
    return(
        <div>
            <Navbar />
            <AllowedUsers userRole = {allowedUsers}></AllowedUsers>
            <UpdateAccommodationGradeForm></UpdateAccommodationGradeForm>
        </div>

    );
}