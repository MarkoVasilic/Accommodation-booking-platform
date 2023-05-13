import React from "react";
import Grid from "@mui/material/Grid";
import { Button, Typography, IconButton } from "@mui/material";
import { green } from "@mui/material/colors";
import { useEffect, useState} from 'react';
import { useLocation, useNavigate, useParams } from 'react-router-dom';
import axiosApi from "../api/axios";
import {useForm} from "react-hook-form";
import InputTextField from "./InputTextField";
import Box from "@mui/material/Box";
import Collapse from "@mui/material/Collapse";
import Alert from "@mui/material/Alert";
import CloseIcon from "@mui/icons-material/Close";
import { Checkbox, FormControlLabel } from "@mui/material";


function UpdateAvailabilityForm() {


    let navigate = useNavigate();
    const params = useParams();
    const {control, handleSubmit, setValue, reset} = useForm();
    const [alert, setAlert] = React.useState(false);
    const [failed, setFailed] = React.useState(false);
    const [err, setErr] = React.useState("");
    const {state} = useLocation();

      function formatTimestamp(timestamp) {
        const dateObj = new Date(timestamp * 1000);
        const year = dateObj.getFullYear();
        const month = dateObj.getMonth() + 1;
        const date = dateObj.getDate();
        const hours = dateObj.getHours();
        const minutes = dateObj.getMinutes();
        const seconds = dateObj.getSeconds();
      
        return `${year}-${padNumber(month)}-${padNumber(date)} ${padNumber(hours)}:${padNumber(minutes)}:${padNumber(seconds)} +0000 UTC`;
      }
      
     // console.log(formatDate("Mon Jul 24 2023 23:25:43 GMT+0200"));
      function padNumber(num) {
        return num.toString().padStart(2, '0');
      }
      
      function formatSecondsToDate(seconds) {
        const date = new Date(seconds * 1000 - 7200*1000);
        //console.log('Sec', seconds)
        //console.log('date', date)
        return date;
      }
    const handleUpdate = async (data) => {
        try {
            //data.StartDate = (formatTimestamp(state.StartDate.seconds)).toString
            //data.EndDate = (formatTimestamp(state.EndDate.seconds)).toString
            data.StartDate = formatSecondsToDate(state.StartDate.seconds)
            data.EndDate = formatSecondsToDate(state.EndDate.seconds)
            data.Price = parseFloat(data.Price)
            console.log('A',data)
            await axiosApi.put(`/availability/${state.Id}` ,data).then(res => {
                console.log(res)
                setAlert(true)
                navigate(-1)
            }).catch(err => {
                console.log('Err');
                setFailed(true)
                setErr('Cannot update availability when there are accepted reservations!')
            });

        }
        catch (error) {
            console.log(error)
            setFailed(true)
        }
    };

    /*const getAvailability = async (e) => {
        try {
            const res = await axiosApi.get(`/availability/get/${params.availability}/`); //dodati u global handleru valjda 
            return res.data;
        } catch (error) {
            console.log(error.response);
        }
    };

    useEffect(() => {
        getAvailability().then(reset);
    },[]);*/



    return (
        <div>
            <Typography variant="h4" color={"#5B63F5"} marginTop={2}  sx={{ textAlign: "center" }}>
                Update Availability
            </Typography>
            <Grid
                container
                rowSpacing={2}
                sx={{ padding: "180px 550px 0px 550px" }}
            >

            <Grid item xs={12}>
                <InputTextField
                    name="Price"
                    label="Price"
                    control={control}
                    type="number"
                    variant="filled"
                    autoFocus
                    fullWidth
                />
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
                                        navigate('/list-centers-update');
                                    }}
                                >
                                    <CloseIcon fontSize="inherit" />
                                </IconButton>
                            }
                            sx={{ mb: 2 }}
                        >
                            Successfuly updated availability!
                        </Alert>
                    </Collapse>
                </Box>
                <Box sx={{ width: "100%" }}>
                    <Collapse in={failed}>
                        <Alert
                            severity="error"
                            action={
                                <IconButton
                                    aria-label="close"
                                    color="inherit"
                                    size="small"
                                    onClick={() => {
                                        setFailed(false);
                                    }}
                                >
                                    <CloseIcon fontSize="inherit" />
                                </IconButton>
                            }
                            sx={{ mb: 2 }}
                        >
                            {err}
                        </Alert>
                    </Collapse>
                </Box>
        </div>
    );
}

export default UpdateAvailabilityForm;
