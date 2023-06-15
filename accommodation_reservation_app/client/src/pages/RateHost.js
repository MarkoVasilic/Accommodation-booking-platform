import RateHostForm from "../components/RateHostForm";
import Navbar from "../pages/Navbar";
import AllowedUsers from "../components/AllowedUsers";

export default function RateHost(){
    //var allowedUsers = ["GUEST"]
    return(
        <div>
            <Navbar/>
            <RateHostForm />
        </div>
    );
}