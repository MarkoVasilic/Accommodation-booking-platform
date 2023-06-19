import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography } from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import axiosApi from "../api/axios";
import Alert from "@mui/material/Alert";
import InputTextField from "./InputTextField";
import { useForm } from "react-hook-form";
import { Checkbox, FormControlLabel } from "@mui/material";
import { useEffect, useState } from "react";


const url = "/user/grade";

function RateHostForm(props) {
    const {state} = useLocation()
    const { handleSubmit, control, setValue } = useForm();
    const [successAlert, setSuccessAlert] = React.useState("hidden");
    const [errorAlert, setErrorAlert] = React.useState("hidden");
    const [alert, setAlert] = React.useState("");
    let navigate = useNavigate();
    const [prominent, setProminent] = useState(false);
    const [hostId, setHostId] = React.useState("");

    
    let getData = async () => {

        axiosApi
        .get('/user/host/all')
        .then((response2) => {
            console.log("dobavio sve hostove")
                response2.data.forEach(host =>{
                    if(host.FirstName ==state.HostFirstName){
                        setHostId(host.id)
                        axiosApi
                            .get('/user/prominent/'+host.id)
                            .then((response1) => {
                                //console.log("RESP AFTER PROMINENT",response1)
                                setProminent(response1.data);
                                console.log('PROMINENT', prominent)
                            });
                    }
                });

            }).catch(er => {
                console.log('greska u dobavljanju hosta') 
            });
        
        
    };


    useEffect(() => {
        getData();
    }, []);
    

    const onSubmit = async (data) => {
       
        data.id = state.ID
        console.log('State', state)
        data.Grade = parseInt(data.Grade);
        console.log(data)
        try {
            const resp = await axiosApi.put(`/user/grade/${state.ID}`, data)
                    .then((response)=> {

                        axiosApi
                        .get('/user/notificationsOn/'+hostId)
                        .then((response3) => {
                            console.log("upao u dobavljanje notOn za hosta", response3.data)

                            response3.data.forEach(nottificationON =>{
                                console.log(nottificationON.Type,nottificationON.on)
                                if(nottificationON.Type == "GRADED_USR" && nottificationON.on){
                                    console.log("pravi not")
                                    let userId = hostId
                                    let type = "GRADED_USR"
                                    let message = "Guest updated rate for you."
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

                            axiosApi
                            .get('/user/prominent/'+hostId)
                            .then((response1) => {
                                if(response1.data != prominent){
                                    axiosApi
                                    .get('/user/notificationsOn/'+hostId)
                                    .then((response3) => {
                                        console.log("upao u dobavljanje notOn za hosta", response3.data)

                                        response3.data.forEach(nottificationON =>{
                                            console.log(nottificationON.Type,nottificationON.on)
                                            if(nottificationON.Type == "PROMINENT" && nottificationON.on){
                                                console.log("pravi not")
                                                let userId = hostId
                                                let type = "PROMINENT"
                                                let message = "Your prominent host status has been changed."
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
                            }).catch(er => {
                                console.log('greska u prominent host') 
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
                Update Host Grade
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

export default RateHostForm;
