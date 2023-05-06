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


const RenderUpdateButton = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                onClick={() => {
                    //navigate('/availability/update/'+params.row.id);
                }}
            >
                Update
            </Button>
        </strong>
    )
};



const columns = [
    {
        field: "StartDate",
        headerName: "Start date",
        type: "datetime-local",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
        //format:"DD/MM/YYYY hh:mm A",
        valueFormatter: params => moment(params?.value).add(-2, 'h').format("DD/MM/YYYY hh:mm:ss A"),
    },
    {
        field: "EndDate",
        headerName: "End Date",
        type: "datetime-local",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
        //format:"DD/MM/YYYY hh:mm A",
        valueFormatter: params => moment(params?.value).add(-2, 'h').format("DD/MM/YYYY hh:mm:ss A"),
    },
    {
        field: "price",
        headerName: "Price",
        type: "number",
        width: 300,
        headerAlign: "left",
        align: "left",
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "IsPricePerGuest",
        headerName: "Is price per guest",
        type: "string",
        width: 300,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "update",
        headerName: "Update Availability",
        width: 300,
        renderCell: RenderUpdateButton,
        disableClickEventBubbling: true   
    }
];


function AvailabilityList(props) {
    const { handleSubmit, control } = useForm();
    const [availabilities, setAvailabilities ] = useState([]);
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
        axiosApi
            //.get(`/flights/all/?${`taking_off_date=${date}&`}${`start_location=&`}${`end_location=&`}${`number_of_tickets=1`}`)
            .then((response) => {
                setAvailabilities(response.data);
            }).catch(er => {
                console.log(er.response);
                setAvailabilities([]);
            });
        }catch (err) {
                console.log(err)
                setAvailabilities([]);
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
                    Availabilities
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
                                        setAvailabilities([])
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
                        rows={availabilities}
                        getRowId={(row) => row.ID}
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

export default AvailabilityList;
