import { IconButton, Button, Typography,Input } from "@mui/material";
import * as React from "react";
import Box from "@mui/material/Box";
import { DataGrid } from "@mui/x-data-grid";
import Paper from "@mui/material/Paper";
import { useEffect, useState, Controller,useRef } from "react";
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

function rowAction(navigate, buttonName, buttonUrl) {
    return {
        field: "Details",
        headerName: buttonName,
        align: "center",
        headerAlign: "left",
        sortable: false,
        renderCell: (params) => {
            const onClick = (e) => {
                e.stopPropagation(); // don't select this row after clicking

                const api = params.api;
                const thisRow = params.row;
               
                return navigate(buttonUrl, { state: thisRow });
            };
            return (
                <Button
                    variant="contained"
                    color="secondary"
                    size="small"
                    onClick={onClick}
                >
                    {" "}
                    <ReadMoreIcon />{" "}
                </Button>
            );
        },
    };
}

function ListSearchedFlightsUser(props) {
    const { handleSubmit, control } = useForm();
    const [flights, setFlights ] = useState([]);
    const [ error, setError ] = React.useState(false);
    const [er, setEr] = React.useState("");
    const [ proba, setProba] = useState(0);

    const navigate = useNavigate();
    useEffect(() => {
        getData();
        
    }, [],{});
    const date = new Date().toISOString();

    let getData = async () => {
    try{
    axiosApi
        .get(`/flights/all/?${`taking_off_date=${date}&`}${`start_location=&`}${`end_location=&`}${`number_of_tickets=1`}`)
        .then((response) => {
            setFlights(response.data);
        }).catch(er => {
            console.log(er.response);
            setFlights([]);
        });


    }catch (err) {
            console.log(err)
            
        }
    };

    const onSubmit = async (data) => {
        try {
            let searchDate = new Date(Date.parse(data.taking_off_date))
            setProba(data.number_of_tickets)
            let res = await axiosApi
            .get(`/flights/all/?${`taking_off_date=${searchDate.toISOString()}&`}${`start_location=${data.start_location}&`}${`end_location=${data.end_location}&`}${`number_of_tickets=${data.number_of_tickets}`}`)
            .then((response) => {
                setFlights(response.data);
                console.log(response.data);
                console.log("LETOVI POSLE: ", flights);
            }).catch(er => {
                console.log(er.response);
                setFlights([]);
                setError(true)
                setEr(er.response.data.error)
            });
        }
        catch (err) {
            console.log(err)
            setFlights([]);
            setError(true);
            setEr(er.response.data.error)
        }
    };

    const columns = [
        {
            field: "Name",
            headerName: "Name",
            type: "string",
            width: 210,
            sortable: false,
            filterable: false,
            editable: false,
        },
        {
            field: "Taking_Off_Date",
            headerName: "Taking off date",
            type: "datetime-local",
            width: 210,
            sortable: false,
            filterable: false,
            editable: false,
            valueFormatter: params => moment(params?.value).format("DD/MM/YYYY hh:mm A"),
        },
        {
            field: "Start_Location",
            headerName: "Start location",
            type: "string",
            width: 210,
            sortable: false,
            filterable: false,
            editable: false,
        },
        {
            field: "End_Location",
            headerName: "End location",
            type: "string",
            width: 210,
            sortable: false,
            filterable: false,
            editable: false,
        },
        {
            field: "Price",
            headerName: "Price per passenger",
            type: "number",
            width: 210,
            headerAlign: "left",
            align: "left",
            sortable: false,
            filterable: false,
            editable: false,
        },
        {
            field: "Total_Price",
            headerName: "Total Price",
            type: "number",
            width: 210,
            headerAlign: "left",
            align: "left",
            sortable: false,
            filterable: false,
            editable: false,
        },
        {
            field: "delete",
            headerName: "Book a flight",
            width: 150,
            renderCell: (params) => {
                return (
                    <Button
                        variant="contained"
                        color="error"
                        size="small"
                        style={{ marginLeft: 16 }}
                        onClick={() => {
                            axiosApi.get(`/users/logged/`).then((response) => {
                                axiosApi.put(`/tickets/buy/${response.data.user_id}`,{ flight: params.row.ID,
                                    number_of_tickets: parseInt(proba)} )
                                    .then(response => {
                                        console.log(response)
                                        navigate("/tickets-user");
                                    })
                                    .catch(error => {
                                        console.error('Problem with the booking request:', error);
                                    });
                            }); 
                            
                        }}
                    >
                        Book a flight
                    </Button>
            )
        },
            disableClickEventBubbling: true
        }
    ];

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
                    Search flights
                </Typography>
            </Stack>
            <form onSubmit={handleSubmit(onSubmit)}>
                <Grid
                    container
                    rowSpacing={2}
                    marginTop={2}
                    sx={{ padding: "0px 0px 10px 180px", textAlign: "left" }}
                >
                    <Grid container spacing={5}>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Choose taking off date:</Typography>
                            <InputTextField
                                name="taking_off_date"
                                control={control}
                                type="date"
                                rules={{ required: "This field is required" }}
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Enter start location:</Typography>
                            <InputTextField
                                name="start_location"
                                control={control}
                                type="text"
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Enter end location:</Typography>
                            <InputTextField
                                name="end_location"
                                control={control}
                                type="text"
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Choose number of tickets:</Typography>
                            <InputTextField
                                name="number_of_tickets"
                                control={control}
                                type="number"
                                min="0"
                                rules={{
                                    required: "This field is required",
                                    min: {
                                      value: 0,
                                      message: "The value cannot be less that 1"
                                    }
                                  }}

                            />
                        </Grid>
                    </Grid>
                        <Button
                            type="submit"
                            variant="contained"
                            sx={{
                                background: "#5B63F5",
                                marginTop: "30px",
                                marginRight: "50px",
                                marginLeft: "1000px",
                                marginBottom: "5px",
                                width: "160px",
                                height: "40px",
                                position: "absolute",
                                "&.MuiButtonBase-root": {
                                    "&:hover": {
                                        backgroundColor: blue[600],
                                    },
                                },
                            }}
                        >
                            Search
                        </Button>
                </Grid>
            </form>
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
                                        setFlights([])
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
                        rows={flights}
                        getRowId={(row) => row.ID}
                        disableColumnFilter
                        columns={[...columns,rowAction(navigate, props.buttonName, props.buttonUrl)]}
                        autoHeight
                        density="comfortable"
                        disableSelectionOnClick
                        rowHeight={50}
                        pageSize={5}
                        headerHeight={45}
                        headerAlign= "left"
                        align="left"
                    />
                </Box>
            </Paper>
        </div>
    );
}

export default ListSearchedFlightsUser;
