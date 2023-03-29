import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";
import axiosApi from "../api/axios";
import Alert from "@mui/material/Alert";
import InputTextField from "./InputTextField";
import { useForm } from "react-hook-form";

const url = "/flights/create/";

function CreateFlightForm() {
    const { handleSubmit, control } = useForm();
    const [successAlert, setSuccessAlert] = React.useState("hidden");
    const [errorAlert, setErrorAlert] = React.useState("hidden");
    const [alert, setAlert] = React.useState("");
    let navigate = useNavigate();
    

    const onSubmit = async (data) => {
        //console.log("Podaci", data);
        let searchDate = new Date(Date.parse(data.taking_off_date));
        data.taking_off_date = searchDate.toISOString();
        data.price = parseFloat(data.price);
        data.number_of_tickets = parseInt(data.number_of_tickets);
        try {
            const resp = await axiosApi.post(url, data);
            setSuccessAlert("visible");
            setErrorAlert("hidden");
            setAlert("success");
            navigate("/flights-admin/all");
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
            >
                Create a Flight
            </Typography>
            <form onSubmit={handleSubmit(onSubmit)}>
            <Grid
                container
                rowSpacing={2}
                sx={{ padding: "55px 550px 0px 550px" }}
            >
                <Grid item xs={12}>
                        <InputTextField
                            name="name"
                            label="Name"
                            control={control}
                            type="text"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="taking_off_date"
                            control={control}
                            type="datetime-local"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="start_location"
                            label="Start location"
                            control={control}
                            type="text"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="end_location"
                            label="End location"
                            control={control}
                            type="text"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="price"
                            label="Price"
                            control={control}
                            type="number"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="number_of_tickets"
                            label="Number of tickets"
                            control={control}
                            type="number"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                    {alert === "success" ? (
                        <Alert
                            sx={{ visibility: successAlert }}
                            severity="success"
                        >
                            Flight was created successfully!
                        </Alert>
                    ) : (
                        <Alert sx={{ visibility: errorAlert }} severity="error">
                            Please try again!
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

export default CreateFlightForm;
