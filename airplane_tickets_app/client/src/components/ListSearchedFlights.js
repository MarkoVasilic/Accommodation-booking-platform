import { IconButton, Button, Typography } from "@mui/material";
import * as React from "react";
import Box from "@mui/material/Box";
import { DataGrid } from "@mui/x-data-grid";
import Paper from "@mui/material/Paper";
import { useEffect, useState, Controller } from "react";
import TextField from "@mui/material/TextField";
import Stack from "@mui/material/Stack";
import CachedIcon from "@mui/icons-material/Cached";
import { blue } from "@mui/material/colors";
import FormControl from "@mui/material/FormControl";
import axiosApi from "../api/axios";
import InputLabel from "@mui/material/InputLabel";
import MenuItem from "@mui/material/MenuItem";
import Select from "@mui/material/Select";
import ReadMoreIcon from "@mui/icons-material/ReadMore";
import { useNavigate } from "react-router-dom";
import InputTextField from "./InputTextField";
import Grid from "@mui/material/Grid";
import { useForm } from "react-hook-form";


const columns = [
    {
        field: "_id",
        headerName: "ID",
        width: 80,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "name",
        headerName: "Name",
        type: "string",
        width: 180,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "taking_off_date",
        headerName: "Taking off date",
        type: "date",
        width: 150,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "start_location",
        headerName: "Start location",
        type: "string",
        width: 200,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "end_location",
        headerName: "End location",
        type: "string",
        width: 200,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "price",
        headerName: "Price number",
        type: "number",
        width: 150,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "number_of_tickets",
        headerName: "Number of tickets",
        type: "number",
        min: "1",
        max: "150",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "total_price",
        headerName: "Total Price",
        type: "number",
        width: 150,
        sortable: false,
        filterable: false,
        editable: false,
    },
];

function rowAction(navigate, buttonName, buttonUrl) {
    return {
        field: "action",
        headerName: buttonName,
        align: "center",
        headerAlign: "center",
        sortable: false,
        renderCell: (params) => {
            const onClick = (e) => {
                e.stopPropagation(); // don't select this row after clicking

                const api = params.api;
                const thisRow = {};

                api.getAllColumns()
                    .filter((c) => c.field !== "__check__" && !!c)
                    .forEach(
                        (c) =>
                        (thisRow[c.field] = params.getValue(
                            params.id,
                            c.field
                        ))
                    );

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

function ListSearchedFlights() {
    const { handleSubmit, control } = useForm();
    const [date, setDate] = useState(new Date().toISOString());
    const [start, setStart] = useState("");
    const [end, setEnd] = useState("");
    const [tickets, setTickets] = useState(1);  //obrisati
    const [flights, setFlights ] = useState([]);
    const navigate = useNavigate();
    useEffect(() => {
        getData();
    }, [date, start, end, tickets, flights]);

    let getData = async () => { //izmeniti da dodaje sve?
        axiosApi
            .get(`/flights/all?${`taking_off_date=${date}&`}${`start_location=${start}&`}${`end_location=${end}&`}${`number_of_tickets=${tickets}`}`)
            .then((response) => {
                setFlights(response.data);
            });
    };

    const onSubmit = async (data) => {
        try {
            //data.transfusion_center = user.userprofile.tranfusion_center
            //data.staff = [data.staff]
            //console.log(data)
            let searchDate = new Date(Date.parse(data.taking_off_date))
            let res = await axiosApi
            .get(`/flights/all?${`taking_off_date=${searchDate.toISOString()}&`}${`start_location=${data.start_location}&`}${`end_location=${data.end_location}&`}${`number_of_tickets=${data.number_of_tickets}`}`)
            .then((response) => {
                setFlights(response.data);
            });
        }
        catch (err) {
            console.log(err)
            //setError(true);
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
            <form onSubmit={handleSubmit(onSubmit)}>
                <Grid
                    container
                    rowSpacing={2}
                    marginTop={2}
                    sx={{ padding: "5px 20px 10px 50px", textAlign: "left" }}
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
                                rules={{ required: "This field is required" }}
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Enter end location:</Typography>
                            <InputTextField
                                name="end_location"
                                control={control}
                                type="text"
                                rules={{ required: "This field is required" }}
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Choose number of tickets:</Typography>
                            <InputTextField
                                name="number_of_tickets"
                                control={control}
                                type="number"
                                rules={{ required: "This field is required" }}
                            />
                        </Grid>
                    </Grid>
                        <Button
                            type="submit"
                            variant="contained"
                            sx={{
                                background: "#5B63F5",
                                marginTop: "-50px",
                                marginRight: "50px",
                                marginLeft: "1000px",
                                marginBottom: "5px",
                                width: "160px",
                                height: "40px",
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
                <Box sx={{ height: 700, width: "100%", marginTop: "20px" }}>
                    <DataGrid
                        rows={flights}
                        disableColumnFilter
                        columns={[...columns, rowAction(navigate)]}
                        autoHeight
                        density="comfortable"
                        disableSelectionOnClick
                        rowHeight={50}
                        pageSize={5}
                        headerHeight={35}
                    />
                </Box>
            </Paper>
        </div>
    );
}

export default ListSearchedFlights;
