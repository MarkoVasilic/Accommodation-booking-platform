import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography } from "@mui/material";
import { useNavigate } from "react-router-dom";
import axiosApi from "../api/axios";
import Alert from "@mui/material/Alert";
import InputTextField from "./InputTextField";
import { useForm } from "react-hook-form";
import { Radio, RadioGroup, FormControlLabel } from '@mui/material';

const url = "/availability";

function CreateAvailabilityForm() {
    const { handleSubmit, control } = useForm();
    const [successAlert, setSuccessAlert] = React.useState("hidden");
    const [errorAlert, setErrorAlert] = React.useState("hidden");
    const [alert, setAlert] = React.useState("");
    let navigate = useNavigate();
    

    const onSubmit = async (data) => {
        //console.log("Podaci", data);
        let startDate = new Date(Date.parse(data.start_date));
        data.start_date = startDate.toISOString();
        let endDate = new Date(Date.parse(data.end_date));
        data.end_date = endDate.toISOString();
        data.price = parseFloat(data.price);
        try {
            const resp = await axiosApi.post(url, data);
            setSuccessAlert("visible");
            setErrorAlert("hidden");
            setAlert("success");
            //navigate("/flights-admin/all");
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
                Create Availability
            </Typography>
            <form onSubmit={handleSubmit(onSubmit)}>
            <Grid
                container
                rowSpacing={2}
                sx={{ padding: "55px 550px 0px 550px" }}
            >
                <Grid item xs={12}>
                        <InputTextField
                            name="start_date"
                            control={control}
                            type="datetime-local"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="end_date"
                            control={control}
                            type="datetime-local"
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
                <Grid container spacing={2} alignItems="center" justify="center">
                <Grid item xs={12}>
                    <Typography >
                        Is price per guest?
                    </Typography>
                </Grid>
                <Grid item container xs={12} justify="center" alignItems="center" direction="row">
                    <RadioGroup name="is_price_per_guest" defaultValue="true" control={control}>
                    <FormControlLabel value="true" control={<Radio />} label="Yes" />
                    <FormControlLabel value="false" control={<Radio />} label="No" />
                    </RadioGroup>
                </Grid>
                </Grid>
                    <Grid item xs={12}>
                    {alert === "success" ? (
                        <Alert
                            sx={{ visibility: successAlert }}
                            severity="success"
                        >
                            Availability was created successfully!
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

export default CreateAvailabilityForm;
