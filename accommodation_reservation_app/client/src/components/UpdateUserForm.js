import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography } from "@mui/material";
import { green } from "@mui/material/colors";
import { useEffect, useState } from 'react';
import { useNavigate} from 'react-router-dom';
import axiosApi from "../api/axios";
import { useForm } from "react-hook-form";
import InputTextField from "./InputTextField";
import Box from "@mui/material/Box";
import Alert from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import CloseIcon from "@mui/icons-material/Close";

import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
const schema = yup
  .object()
  .shape({
    FirstName: yup.string().min(2, "First name must be longer than 1").required(),
    LastName: yup.string().min(2, "Last name must be longer than 1").required(),
    Username: yup.string().min(4, "Username must be longer than 3").required(),
    Address: yup.string().min(5, "Address must be longer than 4").required(),
  }).required();


function UpdateCenterForm() {

    let navigate = useNavigate();
    const {control, handleSubmit, reset} = useForm({
        resolver: yupResolver(schema)
    });
    const [user, setUser] = useState({});
    const [alert, setAlert] = React.useState(false);
    const [failedAlert, setFailedAlert] = React.useState(false);

    const getUser = async (e) => {
        try {
            const res = await axiosApi.get('/user/logged');
            console.log("User update",res.data.user);
            setUser(res.data.user);
            return res.data.user;
        } catch (error) {
            console.log(error.response);
        }
    };

    useEffect(() => {
        getUser().then(reset);
    }, [reset]);

    const handleUpdate = async (data) => {
        try {
            
            await axiosApi.put(`/user/${user.Id}`, data);
            setAlert(true)
        }
        catch (err) {
            console.log(err)
            setFailedAlert(true)
        }
    };

    

    return (
        <div>
            <div style={{ textAlign: 'center' }}>
            <Typography variant="h4" color={green[800]} marginTop={2}>
                Update profile
            </Typography>
            </div>
            <Grid
                container
                marginTop={"-40px"}
                rowSpacing={2}
                sx={{ padding: "55px 550px 0px 550px" }}
            >
                <Grid item xs={12}>
                    <InputTextField
                        name="FirstName"
                        control={control}
                        variant="filled"
                        label="First name"
                        autoFocus
                        fullWidth
                />
                </Grid>
                <Grid item xs={12}>
                <InputTextField
                        name="LastName"
                        control={control}
                        variant="filled"
                        label="Last Name"
                        fullWidth
                    />
                </Grid>
                <Grid item xs={12}>
                <InputTextField
                        name="Username"
                        control={control}
                        variant="filled"
                        label="Username"
                        fullWidth
                    />
                </Grid>
                <Grid item xs={12}>
                <InputTextField
                        name="Address"
                        control={control}
                        variant="filled"
                        label="Address"
                        fullWidth
                    />
                </Grid>
                <Grid item xs={12}>
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        onClick={handleSubmit(handleUpdate)}
                        sx={{
                            mt: 3,
                            mb: 2,
                            background: "#6fbf73",
                            height: "30",
                            "&.MuiButtonBase-root": {
                                "&:hover": {
                                    backgroundColor: green[600],
                                },
                            },
                        }}
                    >
                        Submit
                    </Button>
                </Grid>
            </Grid>
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
                        Update successfull!
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
                        Update failed!
                    </Alert>
                </Collapse>
            </Box>
        </div>
    );
}

export default UpdateCenterForm;
