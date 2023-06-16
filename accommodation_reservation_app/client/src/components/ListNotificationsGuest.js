import { green } from '@mui/material/colors';
import { Typography, Paper, Button } from '@mui/material';
import { Stack } from '@mui/system';
import axiosApi from "../api/axios";
import AllowedUsers from "../components/AllowedUsers";
import Box from "@mui/material/Box";
import Grid from '@mui/material/Grid';
import React, {  useEffect, useState } from "react";
import { DataGrid } from "@mui/x-data-grid";

export default function ListNotificationsGuest() {
    const listOfAllowedUsers = ["GUEST"];
    const [profile, setProfile] = useState({});
    const [not, setReservationReact] = useState(false);
    const [notificationsOn, setnotificationsOn ] = useState([]);
    const [notifications, setNotifications ] = useState([]);

    let updateNotifications = async (event) => {
        try {
            setReservationReact(event)
            console.log(event)
            let id = "ttt"
            let type = "RESERVATION_REACT"
            let on = event
            const updateNotigivationOn = {
                id,
                type,
                on,
            };

            axiosApi.put('/user/notificationon/'+profile.Id, updateNotigivationOn)
                .then((response) => {
                    if(on){
                        axiosApi
                        .get('/user/notification/'+profile.Id)
                        .then((response1) => {
                            setNotifications(response1.data);
                            console.log('Notification2',response1.data) 
                            console.log('Notification',notifications)
                        }).catch(er => {
                            console.log('greska u notification') 
                            setNotifications([])
                        });
                    }else{
                        setNotifications([])
                    }
                    
                })


          } catch (err) {
            
          }
    };

    useEffect(() => {
        getData();
    }, [setNotifications]);


    let getData = async () => {
        try {
            axiosApi
            .get('/user/logged')
            .then((response) => {
                setProfile(response.data.user);

                axiosApi
                .get('/user/notificationsOn/'+response.data.user.Id)
                .then((response2) => {

                   setnotificationsOn(response2.data)
                    if(response2.data[0].on){
                        setReservationReact(true)
                        axiosApi
                        .get('/user/notification/'+response.data.user.Id)
                        .then((response1) => {
                            setNotifications(response1.data); 
                            console.log('Notification2',response1.data) 
                            console.log('Notification',notifications)
                        }).catch(er => {
                            console.log('greska u notification') 
                            setNotifications([])
                        });
                    }else{
                        setReservationReact(false)
                        setNotifications([])  
                    }

                }).catch(er => {
                    console.log('greska u notificationOn') 
                });

            }).catch(er => {
                setNotifications([])
                console.log('greska ') 
            });
          } catch (err) {
            setNotifications([])
            console.log('greska') 
          }
    };

    
    function formatSecondsToDate(seconds) {
        return new Date(seconds * 1000 - 7200*1000);;
    }
    
    const columns = [
        {
            field: "Message",
            headerName: "Message",
            type: "string",
            width: 880,
            sortable: false,
            filterable: false,
            editable: false,
        },
        {
            field: "Type",
            headerName: "Type",
            type: "string",
            width: 380,
            sortable: false,
            filterable: false,
            editable: false,
        },
        {
            field: "DateOfNotification",
            headerName: "Date",
            type: "date",
            width: 300,
            sortable: false,
            filterable: false,
            editable: false,
            format:"DD/MM/YYYY",
            valueGetter: params => formatSecondsToDate(params.row.DateOfNotification.seconds) 
             
        }
    ];

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
                <Typography variant="h5" align='left' marginLeft={"100px"} marginRight={"5px"}>Reservations react: </Typography>
            </Grid>
            <Grid item xs={9}>
                <Typography variant="h5" color="text.secondary" align='left' marginLeft={"80px"}>
                <input id="wifi" type="checkbox" checked={not} onChange={(event) => updateNotifications(event.target.checked)}/>
                </Typography>
            </Grid>

           
        </Grid>
            </Paper>
        </Stack>


        <Paper>
            <Box sx={{ height: 700, width: "100%", marginTop: "20px", marginBottom: "20px"}}>
                <DataGrid
                    rows={notifications}
                    getRowId={(row) => row.id}
                    disableColumnFilter
                    columns={columns}
                    autoHeight
                    density="comfortable"
                    disableSelectionOnClick
                    rowHeight={50}
                    pageSize={5}
                    headerHeight={35}
                    headerAlign= "left"
                    align="left"
                />
            </Box>
        </Paper>
        
        </div>
    );
}