import Typography from '@mui/material/Typography';
import Image from 'mui-image'
import { Stack } from '@mui/system';
import Navbar from "../pages/Navbar";
import React from 'react';

export default function Dashboard() {
    return (
        <div>
            <Navbar />
            <Stack height={"80vh"} justifyContent={"center"} alignItems={"center"}>
                <Typography variant="h4" component="h2" textAlign={'center'}>
                At Last, This Is What You've Been Searching For.
                </Typography>
                <Stack alignItems={"center"}>
                    <Image src="accommodation.png" width={200} height={150} />
                </Stack>
            </Stack>
        </div>
    );
}