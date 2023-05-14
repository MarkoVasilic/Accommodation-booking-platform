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
                    axiosApi
                    .put(`/accommodation/reservation/cancel/`+params.row.ReservationId)
                    .then((response) => {
                        console.log("AAA")
                        console.log(response.data)
                        refreshPage();
                    }).catch(er => {
                        console.log(er.response);
                    });
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
