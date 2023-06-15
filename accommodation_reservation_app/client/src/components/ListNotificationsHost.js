import PersonalInformationCard from "../components/PersonalInformationCard";
import { green } from '@mui/material/colors';
import { Typography, Paper, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { Stack } from '@mui/system';
import axiosApi from "../api/axios";
import AllowedUsers from "../components/AllowedUsers";
import Box from "@mui/material/Box";
import Alert from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import CloseIcon from "@mui/icons-material/Close";
import Grid from '@mui/material/Grid';
import React, {  useEffect, useState } from "react";

export default function ListNotificationsHost() {
    const listOfAllowedUsers = ["HOST"];
    const [profile, setProfile] = useState({});
    const [create_res, setCreateRes] = useState();
    const [cancel_res, setCancelRes] = useState();
    const [graded_usr, setGradedUsr] = useState();
    const [graded_acc, setGradedAcc] = useState();
    const [prominent, setProminent] = useState();

    let updateCreateRes = async (event) => {
        try {
            setCreateRes(event)
            console.log(event)
            
          } catch (err) {
            
          }
    };

    let updateCancelRes = async (event) => {
        try {
            setCancelRes(event)
            console.log(event)
            
          } catch (err) {
            
          }
    };

    let updateGradedUsr = async (event) => {
        try {
            setGradedUsr(event)
            console.log(event)
            
          } catch (err) {
            
          }
    };

    let updateGradedAcc = async (event) => {
        try {
            setGradedAcc(event)
            console.log(event)
            
          } catch (err) {
            
          }
    };

    let updateProminent = async (event) => {
        try {
            setProminent(event)
            console.log(event)
            
          } catch (err) {
            
          }
    };

    let getData = async () => {
        try {
            axiosApi
            .get('/user/logged')
            .then((response) => {
                setProfile(response.data.user);

            }).catch(er => {
                  
                  
            });
          } catch (err) {
            
          }
    };

    useEffect(() => {
        getData();
    }, []);

    return (
        <div>
            <AllowedUsers userRole = {listOfAllowedUsers}/>
            <Stack marginTop={"10px"} justifyContent={"center"}>
            <Typography align="center"  marginBottom={"20px"}  component="h1" variant="h4" color={green[800]}>
                Notifications
            </Typography>

            <Typography align="left" marginTop={"-10px"} marginLeft={"110px"} marginBottom={"5px"} component="h4" variant="h4" color={green[800]}>
                Options
            </Typography>
            <Paper elevation={10} sx={{ p: { sm: 2, xs: 2 } }}>
            <Grid container spacing={2} marginTop="-10px" marginBottom="0px" alignContent={"center"}>

            <Grid item xs={3}>
                <Typography variant="h5" align='left' marginLeft={"100px"} marginRight={"5px"}>Create reservation: </Typography>
            </Grid>
            <Grid item xs={9}>
                <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>
                <input id="create_res" type="checkbox" checked={create_res} onChange={(event) => updateCreateRes(event.target.checked)}/>
                </Typography>
            </Grid>

            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Cancel reservation: </Typography>
            </Grid>
            <Grid item xs={9}>
                <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>
                <input id="cancel_res" type="checkbox" checked={cancel_res} onChange={(event) => updateCancelRes(event.target.checked)}/>
                </Typography>
            </Grid>

            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Graded host: </Typography>
            </Grid>
            <Grid item xs={9}>
                <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>
                <input id="graded_usr" type="checkbox" checked={graded_usr} onChange={(event) => updateGradedUsr(event.target.checked)}/>
                </Typography>
            </Grid>

            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Graded reservation: </Typography>
            </Grid>
            <Grid item xs={9}>
                <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>
                <input id="graded_acc" type="checkbox" checked={graded_acc} onChange={(event) => updateGradedAcc(event.target.checked)}/>
                </Typography>
            </Grid>

            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Prominent: </Typography>
            </Grid>
            <Grid item xs={9}>
                <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>
                <input id="prominent" type="checkbox" checked={prominent} onChange={(event) => updateProminent(event.target.checked)}/>
                </Typography>
            </Grid>

            
        </Grid>
            </Paper>
        </Stack>
        
        </div>
    );
}