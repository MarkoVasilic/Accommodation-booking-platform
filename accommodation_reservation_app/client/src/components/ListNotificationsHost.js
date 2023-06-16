import { green } from '@mui/material/colors';
import { Typography, Paper, Button } from '@mui/material';
import { Stack } from '@mui/system';
import axiosApi from "../api/axios";
import AllowedUsers from "../components/AllowedUsers";
import Box from "@mui/material/Box";
import Grid from '@mui/material/Grid';
import React, {  useEffect, useState } from "react";
import { DataGrid } from "@mui/x-data-grid";    

export default function ListNotificationsHost() {
    const listOfAllowedUsers = ["HOST"];
    const [profile, setProfile] = useState({});
    const [create_res, setCreateRes] = useState(false);
    const [cancel_res, setCancelRes] = useState(false);
    const [graded_usr, setGradedUsr] = useState(false);
    const [graded_acc, setGradedAcc] = useState(false);
    const [prominent, setProminent] = useState(false);
    const [notificationsOn, setnotificationsOn ] = useState([]);
    const [notifications, setNotifications ] = useState([]);

    let updateCreateRes = async (event) => {
        try {
            setCreateRes(event)
            console.log(event)
            let id = "ttt"
            let type = "CREATE_ACC"
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
                        }).catch(er => {
                            console.log('greska u notification') 
                            if(!cancel_res && !graded_acc && !graded_usr && !prominent){
                                setNotifications([])
                            }
                        });
                    }else{
                        if(!cancel_res && !graded_acc && !graded_usr && !prominent){
                            setNotifications([])
                        } 
                    }
                    
                }).catch(er => {
                    console.log('greska u update') 
                });


          } catch (err) {
            
          }
    };

    let updateCancelRes = async (event) => {
        try {
            setCancelRes(event)
            console.log(event)

            let id = "ttt"
            let type = "CANCEL_ACC"
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
                        }).catch(er => {
                            console.log('greska u notification') 
                            if(!create_res && !graded_acc && !graded_usr && !prominent){
                                setNotifications([])
                            }
                        });
                    }else{
                        if(!create_res && !graded_acc && !graded_usr && !prominent){
                            setNotifications([])
                        } 
                    }
                    
                }).catch(er => {
                    console.log('greska u update') 
                });

            
          } catch (err) {
            
          }
    };

    let updateGradedUsr = async (event) => {
        try {
            setGradedUsr(event)
            console.log(event)

            let id = "ttt"
            let type = "GRADED_USR"
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
                        }).catch(er => {
                            console.log('greska u notification') 
                            if(!create_res && !graded_acc && !cancel_res && !prominent){
                                setNotifications([])
                            }
                        });
                    }else{
                        if(!create_res && !graded_acc && !cancel_res && !prominent){
                            setNotifications([])
                        } 
                    }
                    
                }).catch(er => {
                    console.log('greska u update') 
                });
            
          } catch (err) {
            
          }
    };

    let updateGradedAcc = async (event) => {
        try {
            setGradedAcc(event)
            console.log(event)

            let id = "ttt"
            let type = "GRADED_ACC"
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
                        }).catch(er => {
                            console.log('greska u notification') 
                            if(!create_res && !graded_usr && !cancel_res && !prominent){
                                setNotifications([])
                            }
                        });
                    }else{
                        if(!create_res && !graded_usr && !cancel_res && !prominent){
                            setNotifications([])
                        } 
                    }
                    
                }).catch(er => {
                    console.log('greska u update') 
                });
            
          } catch (err) {
            
          }
    };

    let updateProminent = async (event) => {
        try {
            setProminent(event)
            console.log(event)

            let id = "ttt"
            let type = "PROMINENT"
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
                        }).catch(er => {
                            console.log('greska u notification') 
                            if(!create_res && !graded_acc && !graded_usr && !cancel_res){
                                setNotifications([])
                            }
                        });
                    }else{
                        if(!create_res && !graded_acc && !graded_usr && !cancel_res){
                            setNotifications([])
                        } 
                    }
                    
                }).catch(er => {
                    console.log('greska u update') 
                });
            
          } catch (err) {
            
          }
    };

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

                   response2.data.forEach(notificationOn => {
              
                    switch (notificationOn.Type) {
                      case "CREATE_ACC":
                        setCreateRes(notificationOn.on);
                        break;
                      case "CANCEL_ACC":
                        setCancelRes(notificationOn.on);
                        break;
                      case "GRADED_USR":
                        setGradedUsr(notificationOn.on);
                        break;
                      case "GRADED_ACC":
                        setGradedAcc(notificationOn.on);
                        break;
                      case "PROMINENT":
                        setProminent(notificationOn.on);
                        break;
                      default:
                        break;
                    }
                  });

                }).catch(er => {
                    console.log('greska u notificationOn') 
                });

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

            }).catch(er => {
                  
                  
            });
          } catch (err) {
            
          }
    };

    useEffect(() => {
        getData();
    }, [setNotifications]);

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