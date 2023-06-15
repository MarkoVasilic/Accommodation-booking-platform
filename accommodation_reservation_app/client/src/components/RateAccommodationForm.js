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
       
        data.Grade = parseInt(data.Grade);
        console.log(data.Price)
        try {
            const resp = await axiosApi.post(url, data);
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
