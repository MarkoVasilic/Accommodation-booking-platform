import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography } from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import axiosApi from "../api/axios";
import Alert from "@mui/material/Alert";
import InputTextField from "./InputTextField";
import { useForm } from "react-hook-form";
import { Checkbox, FormControlLabel } from "@mui/material";



function UpdateAccommodationGradeForm(props) {
    const {state} = useLocation()
    const { handleSubmit, control, setValue } = useForm();
    const [successAlert, setSuccessAlert] = React.useState("hidden");
    const [errorAlert, setErrorAlert] = React.useState("hidden");
    const [alert, setAlert] = React.useState("");
    let navigate = useNavigate();
    //

    const onSubmit = async (data) => {
       
        data.Grade = parseInt(data.Grade);
        data.id = state.ID
        console.log(data)
        try {
            //URL, id i podaci
            const resp = await axiosApi.put(`/accommodation/guest/grades/${state.ID}`, data)
            .then((response)=>{

                axiosApi
                        .get('/accommodation/all/64580a2e9f857372a34602c2')
                        .then((response2) => {
                            console.log("dobavio sve acc")
                                response2.data.forEach(accommodation =>{
                                    console.log(accommodation.Name, data.AccommodationName)
                                    if(accommodation.Name == state.AccommodationName){
                                        
                                        axiosApi
                                        .get('/user/notificationsOn/'+accommodation.HostId)
                                        .then((response3) => {
                                            console.log("upao u dobavljanje notOn za hosta", response3.data)

                                            response3.data.forEach(nottificationON =>{
                                                console.log(nottificationON.Type,nottificationON.on)
                                                if(nottificationON.Type == "GRADED_ACC" && nottificationON.on){
                                                    console.log("pravi not")
                                                    let userId = accommodation.HostId
                                                    let type = "GRADED_ACC"
                                                    let message = "Guest update grade for "+state.AccommodationName+"."
                                                    const d={
                                                        userId,
                                                        type,
                                                        message
                                                    }
                                                    axiosApi
                                                    .post(`/user/notification`,d)
                                                    .then((response) => {
                                                        
                                                    }).catch(er => {
                                                        console.log(er.response);
                                                    });
                                                }
                                            })



                                            }).catch(er => {
                                                console.log('greska u notificationOn') 
                                            });
                                    }
                                });

                            }).catch(er => {
                                console.log('greska u notificationOn') 
                            });
            });
            setSuccessAlert("visible");
            setErrorAlert("hidden");
            setAlert("success");
            navigate(-1);
        } catch (error) {
            setErrorAlert("visible");
            setSuccessAlert("hidden");
            setAlert("error");

        }
    };

    return (
        <div>
            <Typography
                component="h1"
                variant="h4"
                color={"#5B63F5"}
                marginTop={2}
                sx={{ textAlign: "center" }}
            >
                Update Accommodation Grade
            </Typography>
            <form onSubmit={handleSubmit(onSubmit)}>
            <Grid
                container
                rowSpacing={2}
                sx={{ padding: "55px 550px 0px 550px" }}
            >
                <Grid item xs={12}>
                        <InputTextField
                            name="Grade"
                            label="Grade"
                            control={control}
                            type="number"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid container spacing={2} alignItems="center" justify="center">
                <Grid item container xs={12} justify="center" alignItems="center" direction="row">
                </Grid>
                </Grid>
                    <Grid item xs={12}>
                    {alert === "success" ? (
                        <Alert
                            sx={{ visibility: successAlert }}
                            severity="success"
                        >
                            Grade was updated successfully!
                        </Alert>
                    ) : (
                        <Alert sx={{ visibility: errorAlert }} severity="error">
                            The grade must be between 1 and 5!
                        </Alert>
                    )}
                </Grid>
                
                <Grid item xs={12}>
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        onClick={onSubmit}
                        sx={{
                            mt: 3,
                            mb: 2,
                            background: "#6fbf73",
                            height: "30",
                            "&.MuiButtonBase-root": {
                                "&:hover": {
                                    backgroundColor: "#5B63F5"
                                }
                                }
                        }}
                    >
                        Submit
                    </Button>
                </Grid>
            </Grid>
            </form>
        </div>
    );
}

export default UpdateAccommodationGradeForm;
