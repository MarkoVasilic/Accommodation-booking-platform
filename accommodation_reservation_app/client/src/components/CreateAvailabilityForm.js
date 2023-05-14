import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography } from "@mui/material";
import { useLocation, useNavigate } from "react-router-dom";
import axiosApi from "../api/axios";
import Alert from "@mui/material/Alert";
import InputTextField from "./InputTextField";
import { useForm } from "react-hook-form";
import { Checkbox, FormControlLabel } from "@mui/material";


const url = "/availability";

function CreateAvailabilityForm(props) {
    const {state} = useLocation()
    const { handleSubmit, control, setValue } = useForm();
    const [successAlert, setSuccessAlert] = React.useState("hidden");
    const [errorAlert, setErrorAlert] = React.useState("hidden");
    const [alert, setAlert] = React.useState("");
    let navigate = useNavigate();
    

    const onSubmit = async (data) => {
        //console.log("Podaci", data);
        //id ccomodation ubaciti u data jos
        console.log(state)
        //setValue("IsPricePerGuest", data.IsPricePerGuest === "true");
        let startDate = new Date(Date.parse(data.StartDate));
        data.StartDate = startDate
        let endDate = new Date(Date.parse(data.EndDate));
        data.EndDate = endDate
        console.log(data.Price)
        data.Price = parseFloat(data.Price);
        console.log(data.Price)
        console.log(data.IsPricePerGuest)
        data.AccommodationId = state
        try {
            const resp = await axiosApi.post(url, data);
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
                            name="StartDate"
                            control={control}
                            type="date"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="EndDate"
                            control={control}
                            type="date"
                            rules={{ required: "This field is required" }}
                        />
                </Grid>
                <Grid item xs={12}>
                        <InputTextField
                            name="Price"
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
                <FormControlLabel
                control={
                  <Checkbox
                    name="IsPricePerGuest"
                    defaultChecked={false}
                    onChange={(event) => {
                      setValue("IsPricePerGuest", event.target.checked);
                    }}
                  />
                }
                label="Is price per guest?"
                />

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
                            Overlap with existing availability, please choose another dates!
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
