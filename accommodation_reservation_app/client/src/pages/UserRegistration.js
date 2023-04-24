import Typography from '@mui/material/Typography';
import { Stack } from '@mui/system';
import RegistrationForm from "../components/RegistrationForm";
import Navbar from "../pages/Navbar";
import React from 'react';

export default function UserRegistration() {
    return (
        <div>
            <Navbar />
            <Stack height={"100vh"} justifyContent={"center"} alignItems={"center"}>
                <Typography component="h1" variant="h4" color={"#5B63F5"} textAlign={'center'}>
                    Welcome
                </Typography>
                <RegistrationForm/>
            </Stack>
        </div>
    );
}