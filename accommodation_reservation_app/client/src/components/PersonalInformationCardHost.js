import * as React from 'react';
import Typography from '@mui/material/Typography';
import { useEffect, useState } from "react";
import axiosApi from "../api/axios";
import Grid from '@mui/material/Grid';

function PersonalInformationCardHostComponent() {
    const [profile, setProfile] = useState({});
    const [prominent, setProminent] = useState(false);

    let getData = async () => {
        axiosApi
            .get('/user/logged')
            .then((response) => {
                console.log(response.data.user)
                setProfile(response.data.user);
        axiosApi
                .get('/user/prominent/'+response.data.user.Id)
                .then((response1) => {
                    //console.log("RESP AFTER PROMINENT",response1)
                    setProminent(response1.data);
                    //console.log('PROMINENT', prominent)
                });})
    };


    useEffect(() => {
        getData();
    }, []);

    return (
        <Grid container spacing={2} marginTop="-10px" marginBottom="0px" alignContent={"center"}>
            <Grid item xs={3}>
                <Typography variant="h5" align='left' marginLeft={"100px"} marginRight={"5px"}>First Name: </Typography>
            </Grid>
            <Grid item xs={9}>
                <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{profile.FirstName}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Last Name: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{profile.LastName}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Username: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{profile.Username}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Email: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{profile.Email}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Address: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{profile.Address}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Profile Type: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{profile.Role}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Prominent: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>
                 {prominent ? 'Yes' : 'No'}
            </Typography>
            </Grid>

        </Grid>
    );
}

export default PersonalInformationCardHostComponent;
