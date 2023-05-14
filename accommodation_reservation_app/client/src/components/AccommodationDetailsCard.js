import React from 'react';
import Typography from '@mui/material/Typography';
import { useLocation } from 'react-router-dom';
import Grid from '@mui/material/Grid';
import moment from "moment";

function check(value){
    if (value == 'true')
        return 'YES'
    else
    return 'NO'
}

export default function FlightDetailsCard(props) {
    const { state } = useLocation(); 
    return (
        <Grid container spacing={2} marginTop="0px" marginBottom="10px" alignContent={"center"}>
            <Grid item xs={3}>
                <Typography variant="h5" align='left' marginLeft={"100px"} marginRight={"5px"}>Image: </Typography>
            </Grid>
            <Grid item xs={9}>
              <img src={state.Images[0]} width="150" height="150"/>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Wifi: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{check(state.Wifi)}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Kitchen: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{check(state.Kitchen)}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>AC: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{check(state.AC)}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Parking lot: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{check(state.AC)}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Max number of guests: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{state.MaxGuests}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Min number of guests: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{state.MinGuests}</Typography>
            </Grid>
        </Grid>
    );
}