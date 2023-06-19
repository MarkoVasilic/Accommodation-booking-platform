import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography } from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import axiosApi from "../api/axios";
import Alert from "@mui/material/Alert";
import InputTextField from "./InputTextField";
import { useForm } from "react-hook-form";
import { Checkbox, FormControlLabel } from "@mui/material";


const url = "/accommodation/grade";

function RateHostForm(props) {
    const {state} = useLocation()
    const { handleSubmit, control, setValue } = useForm();
    const [successAlert, setSuccessAlert] = React.useState("hidden");
    const [errorAlert, setErrorAlert] = React.useState("hidden");
    const [alert, setAlert] = React.useState("");
    let navigate = useNavigate();
    

    const onSubmit = async (data) => {
        
        try {
            const res = await axiosApi.get('/user/logged');
            console.log("ID", res.data.user.Id);
            data.Grade = parseInt(data.Grade);
            data.AccommodationID = state.Id
            data.GuestID = res.data.user.Id
            data.DateOfGrade = "2020-11-30T14:20:28.000+07:00"

            const resp = await axiosApi.post('/accommodation/grade', data)
                    .then((response)=>{

                            /////// Notifikacije 
                            axiosApi
                            .get('/user/notificationsOn/'+state.HostId)
                            .then((response3) => {
                                console.log("upao u dobavljanje notOn za hosta "+state.HostId, response3.data)

                                response3.data.forEach(nottificationON =>{
                                    console.log(nottificationON.Type,nottificationON.on)
                                    if(nottificationON.Type == "GRADED_ACC" && nottificationON.on){
                                        console.log("pravi not")
                                        let userId = state.HostId
                                        let type = "GRADED_ACC"
                                        let message = "Guest rated accommodation."
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

                            /////// Notifikacije 

                    }).catch(er => {
                        console.log(er) 
                    });
            setSuccessAlert("visible");
            setErrorAlert("hidden");
            setAlert("success");
            //dodati id ulogovanog usera da se prikazu njegove ocene za hostove
            navigate('/user/accommodations-grades');
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
                Rate Accommodation
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
                            Accommodation was rated successfully!
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

export default RateHostForm;
