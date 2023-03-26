import React from 'react';
import Typography from '@mui/material/Typography';
import { useLocation } from 'react-router-dom';
import Grid from '@mui/material/Grid';

export default function FlightDetailsCard(props) {
    //const { state } = useLocation(); state.name
    return (
        <Grid container spacing={2} marginTop="0px" marginBottom="10px" alignContent={"center"}>
            <Grid item xs={3}>
                <Typography variant="h5" align='left' marginLeft={"100px"} marginRight={"5px"}>Name: </Typography>
            </Grid>
            <Grid item xs={9}>
              <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{"A"}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Taking off date: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{"A"}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Start location: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{"A"}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>End location: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{"A"}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Price: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{"A"}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Number of tickets: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{"A"}</Typography>
            </Grid>
            <Grid item xs={3}>
            <Typography variant="h5" align='left' marginLeft={"100px"}>Remaining tickets: </Typography>
            </Grid>
            <Grid item xs={9}>
            <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>{"A"}</Typography>
            </Grid>
            
        </Grid>
    );
}
