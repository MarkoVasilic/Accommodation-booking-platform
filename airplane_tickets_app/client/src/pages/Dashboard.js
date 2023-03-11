import Typography from '@mui/material/Typography';
import Image from 'mui-image'
import { Stack } from '@mui/system';
import Navbar from "../pages/Navbar";
import axiosApi from "../api/axios";
import { useState, useEffect } from "react";
import React from 'react';

export default function Dashboard() {
    return (
        <div>
            <Navbar />
            <Stack height={"80vh"} justifyContent={"center"}>
                <Typography variant="h4" component="h2">
                    Let your dreams take flight.
                </Typography>
                <Stack alignItems={"center"}>
                    <Image src="airplane.png" width={200} height={150} />
                </Stack>
            </Stack>
        </div>
    );
}