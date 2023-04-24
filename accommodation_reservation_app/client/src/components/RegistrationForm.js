import React from "react";
import { TextField } from '@mui/material';
import Grid from "@mui/material/Grid";
import { Button, RadioGroup, FormControlLabel, Radio } from "@mui/material";
import { useForm, Controller } from "react-hook-form";
import InputTextField from "./InputTextField";
import { sumbitRegistration } from "../api/signInOn";
import Box from "@mui/material/Box";
import Alert from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import CloseIcon from "@mui/icons-material/Close";
import { useNavigate } from "react-router-dom";

const RegistrationForm = () => {
    const { handleSubmit, control, getValues, setError } = useForm();
    const [alert, setAlert] = React.useState(false);
    const [failedAlert, setFailedAlert] = React.useState(false);
    let navigate = useNavigate();

    const onSubmit = async (data) => {
        delete data.confirm_password
        try {
            
            await sumbitRegistration(data)
            setAlert(true)
        }
        catch (err) {
            console.log(err)
            const errMes = err.response.data
            setFailedAlert(true)
            for (let key in errMes) {
                setError(key, { message: errMes[key] })
            }
        }
    };

    function isValidEmail(email) {
        return /\S+@\S+\.\S+/.test(email);
    }

    return (
        <div>
            <form onSubmit={handleSubmit(onSubmit)}>
                <Grid
                    container
                    rowSpacing={2}
                    sx={{ padding: "55px 550px 0px 550px" }}
                >
                    <Grid item xs={12}>
                        <Grid item>
                            <InputTextField
                                name="FirstName"
                                control={control}
                                label="First Name"
                                rules={{ required: "First name required",
                                    minLength: {
                                        value: 2,
                                        message:
                                            "First name needs to be longer than 1",
                                    }, }}
                            />
                        </Grid>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid item>
                            <InputTextField
                                name="LastName"
                                control={control}
                                label="Last Name"
                                rules={{ required: "Last name required",
                                minLength: {
                                    value: 2,
                                    message:
                                        "Last name needs to be longer than 1",
                                }, }}
                            />
                        </Grid>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid item>
                            <InputTextField
                                name="Username"
                                control={control}
                                label="Username"
                                rules={{ required: "Username required",
                                minLength: {
                                    value: 4,
                                    message:
                                        "Username needs to be longer than 3",
                                }, }}
                            />
                        </Grid>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid item>
                            <Controller
                                name="Email"
                                control={control}
                                defaultValue=""
                                render={({
                                    field: { onChange, value },
                                    fieldState: { error },
                                }) => (
                                    <TextField
                                        label="Email"
                                        variant="filled"
                                        value={value}
                                        fullWidth
                                        onChange={onChange}
                                        error={!!error}
                                        helperText={
                                            error ? error.message : null
                                        }
                                    />
                                )}
                                rules={{
                                    required: "Email required",
                                    validate: {
                                        validateEmail: (v) =>
                                            isValidEmail(getValues("Email")) ||
                                            "Email form not correct",
                                    },
                                }}
                            />
                        </Grid>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid item>
                            <InputTextField
                                name="Password"
                                control={control}
                                label="Password"
                                type="password"
                                rules={{
                                    required: "Password required",
                                    minLength: {
                                        value: 6,
                                        message:
                                            "Password needs to be longer than 5",
                                    },
                                }}
                            />
                        </Grid>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid item>
                            <Controller
                                name="confirm_password"
                                control={control}
                                defaultValue=""
                                render={({
                                    field: { onChange, value },
                                    fieldState: { error },
                                }) => (
                                    <TextField
                                        label="Confirm Password"
                                        variant="filled"
                                        type={"password"}
                                        value={value}
                                        fullWidth
                                        onChange={onChange}
                                        error={!!error}
                                        helperText={
                                            error ? error.message : null
                                        }
                                    />
                                )}
                                rules={{
                                    required: "Password confirmation required",
                                    validate: {
                                        confirmPassword: (v) =>
                                            v === getValues("Password") ||
                                            "Passwords doesn't match",
                                    },
                                    minLength: {
                                        value: 6,
                                        message:
                                            "Password needs to be longer than 5",
                                    },
                                }}
                            />
                        </Grid>
                    </Grid>
                    <Grid item xs={12}>
                        <Grid item>
                            <InputTextField
                                name="Address"
                                control={control}
                                label="Address"
                                rules={{ required: "Address required",
                                    minLength: {
                                        value: 5,
                                        message:
                                            "Address needs to be longer than 4",
                                    }, }}
                            />
                        </Grid>
                    </Grid>
                    <Grid item xs={12}>
                        <Controller
                            name="Role"
                            control={control}
                            defaultValue="HOST"
                            render={({ field }) => (
                            <RadioGroup row aria-label="Role" {...field}>
                                <FormControlLabel
                                value="HOST"
                                control={<Radio />}
                                label="Host"
                                />
                                <FormControlLabel
                                value="GUEST"
                                control={<Radio />}
                                label="Guest"
                                />
                            </RadioGroup>
                            )}
                        />
                    </Grid>
                    <Grid item xs={12}>
                        <Button
                            type="submit"
                            variant="contained"
                            sx={{
                                background: "#6fbf73",
                                marginTop: 2,
                                "&.MuiButtonBase-root": {
                                    "&:hover": {
                                        backgroundColor: "#5B63F5",
                                    },
                                },
                            }}
                            fullWidth
                        >
                            Sign Up
                        </Button>
                    </Grid>
                </Grid>
            </form>
            <Box sx={{ width: "100%" }}>
                <Collapse in={alert}>
                    <Alert
                        severity="success"
                        action={
                            <IconButton
                                aria-label="close"
                                color="inherit"
                                size="small"
                                onClick={() => {
                                    setAlert(false);
                                    navigate("/");
                                }}
                            >
                                <CloseIcon fontSize="inherit" />
                            </IconButton>
                        }
                        sx={{ mb: 2 }}
                    >
                        Registration successfull!
                    </Alert>
                </Collapse>
            </Box>
            <Box sx={{ width: "100%" }}>
                <Collapse in={failedAlert}>
                    <Alert
                        severity="error"
                        action={
                            <IconButton
                                aria-label="close"
                                color="inherit"
                                size="small"
                                onClick={() => {
                                    setFailedAlert(false);
                                    navigate("/");
                                }}
                            >
                                <CloseIcon fontSize="inherit" />
                            </IconButton>
                        }
                        sx={{ mb: 2 }}
                    >
                        Registration failed!
                    </Alert>
                </Collapse>
            </Box>
        </div>
    );
};

export default RegistrationForm;
