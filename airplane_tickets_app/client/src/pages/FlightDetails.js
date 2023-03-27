import FlightDetailsCard from "../components/FlightDetailsCard";
import { blue } from '@mui/material/colors';
import { Typography, Paper } from '@mui/material';
import { Stack } from '@mui/system';
import Navbar from "../pages/Navbar";

export default function FlightDetails(props) {
    return (
        <div>
            <Navbar />
            <Stack marginTop={"10px"} justifyContent={"center"}>
            <Typography align="center"  marginBottom={"30px"}  component="h1" variant="h4" color="#5B63F5">
                Details
            </Typography>
            <Paper elevation={15} sx={{ p: { sm: 2, xs: 2 } }}>
            <FlightDetailsCard props></FlightDetailsCard>
            </Paper>
        </Stack>
        </div>
    );
}