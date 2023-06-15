import Navbar from "./Navbar";
import UpdateAccommodationGradeForm from "../components/UpdateAccommodationGradeForm"
import AllowedUsers from "../components/AllowedUsers";

export default function UpdateAccommodationGrade(){
    //const listOfAllowedUsers = ["GUEST"];
    
    return(
        <div>
            <Navbar />
            <UpdateAccommodationGradeForm></UpdateAccommodationGradeForm>
        </div>

    );
}