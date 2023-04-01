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



const columns = [
    {
        field: "Name",
        headerName: "Name",
        type: "string",
        width: 220,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "Taking_Off_Date",
        headerName: "Taking off date",
        type: "datetime-local",
        width: 220,
        sortable: false,
        filterable: false,
        editable: false,
        //format:"DD/MM/YYYY hh:mm A",
        valueFormatter: params => moment(params?.value).add(-2, 'h').format("DD/MM/YYYY hh:mm:ss A"),
    },
    {
        field: "Start_Location",
        headerName: "Start location",
        type: "string",
        width: 220,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "End_Location",
        headerName: "End location",
        type: "string",
        width: 220,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "Price",
        headerName: "Price per passenger",
        type: "number",
        width: 220,
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
        width: 220,
        headerAlign: "left",
        align: "left",
        sortable: false,
        filterable: false,
        editable: false,
    },
];


function ListTicketsUser(props) {
    const [flights, setFlights ] = useState([]);
    useEffect(() => {
        getData();
    }, []);
    const date = new Date().toISOString();

        let getData = async () => {
        try{
            //dobavi usera i karte umesto ovog zahteva
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
                setFlights([]);
            }
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
                    Search flights
                </Typography>
            </Stack>
            <Paper>
                <Box sx={{ height: 700, width: "100%", marginTop: "20px", marginBottom: "20px"}}>
                    <DataGrid
                        rows={flights}
                        getRowId={(row) => row.ID}
                        disableColumnFilter
                        columns={[...columns]}
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

export default ListTicketsUser;
