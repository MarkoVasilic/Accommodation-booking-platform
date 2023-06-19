import { IconButton, Button, Typography } from "@mui/material";
import * as React from "react";
import Box from "@mui/material/Box";
import { DataGrid } from "@mui/x-data-grid";
import Paper from "@mui/material/Paper";
import { useEffect, useState, Controller } from "react";
import Stack from "@mui/material/Stack";
import { blue } from "@mui/material/colors";
import axiosApi from "../api/axios";
import ReadMoreIcon from "@mui/icons-material/ReadMore";
import { useNavigate } from "react-router-dom";
import InputTextField from "./InputTextField";
import Grid from "@mui/material/Grid";
import { useForm } from "react-hook-form";
import moment from "moment";
import CloseIcon from "@mui/icons-material/Close";
import Alert from "@mui/material/Alert";
import Collapse from "@mui/material/Collapse";


function refreshPage(){
    window.location.reload();
}

const RenderCancelReservation = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                onClick={() => {
                    let HostId = ""
                    let prominent = false
                    axiosApi
                        .get('/accommodation/all/64580a2e9f857372a34602c2')
                        .then((response2) => {
                            console.log("dobavio sve acc")
                                response2.data.forEach(accommodation =>{
                                    if(accommodation.Name == params.row.Name){
                                       HostId = accommodation.HostId
                                       console.log("Host id", HostId)

                                       axiosApi
                                       .get('/user/prominent/'+HostId)
                                       .then((response1) => {
                                           //console.log("RESP AFTER PROMINENT",response1)
                                           prominent = response1.data;
                                           console.log('PROMINENT', prominent)
                                           //////////////////////////Boze pomozi
                                           axiosApi
                                           .put(`/accommodation/reservation/cancel/`+params.row.ReservationId)
                                           .then((response) => {
                                               console.log("AAA")
                                               console.log(response.data)
                       
                                               axiosApi
                                               .get('/user/prominent/'+HostId)
                                               .then((response1) => {
                                                   //console.log("RESP AFTER PROMINENT",response1)
                                                            if(prominent != response1.data){
                                                            axiosApi
                                                                .get('/user/notificationsOn/'+HostId)
                                                                .then((response3) => {
                                                                    console.log("upao u dobavljanje notOn za hosta", response3.data)
                                
                                                                    response3.data.forEach(nottificationON =>{
                                                                        console.log(nottificationON.Type,nottificationON.on)
                                                                        if(nottificationON.Type == "PROMINENT" && nottificationON.on){
                                                                            console.log("pravi not")
                                                                            let userId = HostId
                                                                            let type = "PROMINENT"
                                                                            let message = "Your prominent host status has been changed."
                                                                            const d={
                                                                                userId,
                                                                                type,
                                                                                message
                                                                            }
                                                                            axiosApi
                                                                            .post(`/user/notification`,d)
                                                                            .then((response) => {
                                                                                
                                                                            }).catch(er => {
                                                                                console.log(er.response);
                                                                            });
                                                                        }
                                                                    })
                                                                    }).catch(er => {
                                                                        console.log('greska u notificationOn') 
                                                                    });
                                                            }
                                                            //console.log('PROMINENT', prominent)
                                                        });
                       
                                              
                                                    axiosApi
                                                    .get('/user/notificationsOn/'+HostId)
                                                    .then((response3) => {
                                                        console.log("upao u dobavljanje notOn za hosta", response3.data)
                            
                                                        response3.data.forEach(nottificationON =>{
                                                            console.log(nottificationON.Type,nottificationON.on)
                                                            if(nottificationON.Type == "CANCEL_ACC" && nottificationON.on){
                                                                console.log("pravi not")
                                                                let userId = HostId
                                                                let type = "CANCEL_ACC"
                                                                let message = "Reservation in "+params.row.Name+" has been canceled."
                                                                const d={
                                                                    userId,
                                                                    type,
                                                                    message
                                                                }
                                                                axiosApi
                                                                .post(`/user/notification`,d)
                                                                .then((response) => {
                                                                    
                                                                }).catch(er => {
                                                                    console.log(er.response);
                                                                });
                                                            }
                                                        })
                            
                                                        }).catch(er => {
                                                            console.log('greska u notificationOn') 
                                                        });
                                                           
                       
                       
                                               
                       
                                               refreshPage();
                                           }).catch(er => {
                                               console.log(er.response);
                                           });

                                           ////////////////////////////////
                                       });

                                    }
                                })
                            })

                            

                    
                }}
            >
                Cancel
            </Button>
        </strong>
    )
};

function formatSecondsToDate(seconds) {
    const date = new Date(seconds * 1000 - 7200*1000);
    //console.log('Sec', seconds)
    //console.log('date', date)
    return date;
  }

const columns = [
    {
        field: "Name",
        headerName: "Name",
        type: "string",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "Location",
        headerName: "Location",
        type: "string",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "StartDate.seconds",
        headerName: "Start date",
        type: "date",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
        format: "DD/MM/YYYY",
        valueGetter: params => formatSecondsToDate(params.row.StartDate.seconds)
    },
    {
        field: "EndDate.seconds",
        headerName: "End date",
        type: "date",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
        format: "DD/MM/YYYY",
        valueGetter: params => formatSecondsToDate(params.row.EndDate.seconds)
    },
    {
        field: "cancel",
        headerName: "Cancel Reservation",
        width: 300,
        renderCell: RenderCancelReservation,
        disableClickEventBubbling: true   
    }
];


function AcceptedReservationsList(props) {
    const { handleSubmit, control } = useForm();
    const [reservations, setReservations ] = useState([]);
    const [ error, setError ] = React.useState(false);
    const [er, setEr] = React.useState("");
    const navigate = useNavigate();
    const [prominent, setProminent] = useState(false);

    useEffect(() => {
        getData();
      //  onSubmit();
    }, []);
    const date = new Date().toISOString();

        let getData = async () => {
        try{
            const res = await axiosApi.get('/user/logged');
            console.log("ID", res.data.user.Id);
        axiosApi
            .get(`/reservation/guest/accepted/`+res.data.user.Id)
            .then((response) => {
                console.log("Reservations: ", response.data)
                setReservations(response.data);
            }).catch(er => {
                console.log(er.response);
                setReservations([]);
            });

    
        }catch (err) {
                console.log(err)
                setReservations([]);
            }
           // console.log('Sec',new Date(reservations[0].StartDate.seconds*1000-7200*1000))
        };

    return (
        <div>
            <Stack direction={"row"} sx={{ justifyContent: "center" }}>
                <Typography
                    component="h1"
                    variant="h4"
                    color={"#5B63F5"}
                    marginBottom={3}
                    marginTop={1}
                >
                    Accepted Reservations
                </Typography>
            </Stack>
            <Paper>
            <Box sx={{ width: "100%" }}>
                    <Collapse in={error}>
                        <Alert
                            severity="error"
                            action={
                                <IconButton
                                    aria-label="close"
                                    color="inherit"
                                    size="small"
                                    onClick={() => {
                                        setError(false);
                                        setReservations([])
                                    }}
                                >
                                    <CloseIcon fontSize="inherit" />
                                </IconButton>
                            }
                            sx={{ mb: 2 }}
                        >
                            {er}
                        </Alert>
                    </Collapse>
                </Box>
                <Box sx={{ height: 700, width: "100%", marginTop: "20px", marginBottom: "20px"}}>
                    <DataGrid
                        rows={reservations}
                        getRowId={(row) => row.ReservationId}
                        disableColumnFilter
                        columns={[...columns]}
                        autoHeight
                        density="comfortable"
                        disableSelectionOnClick
                        rowHeight={50}
                        headerHeight={35}
                        headerAlign= "left"
                        align="left"
                    />
                </Box>
            </Paper>
        </div>
    );
}

export default AcceptedReservationsList;
