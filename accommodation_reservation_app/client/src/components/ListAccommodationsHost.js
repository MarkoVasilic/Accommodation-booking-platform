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

//proslediti nekako id accomodationa
const RenderCreateButton = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                onClick={() => {
                    return navigate('/availability/create', {state: params.row.Id});
                }}
            >
                Create availability
            </Button>
        </strong>
    )
};

const RenderAvailabilitiesButton = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                align="left"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                onClick={() => {
                    return navigate('/availabilities', {state: params.row.Id});
                }}
            >
                All availabilities
            </Button>
        </strong>
    )
};



const columns = [
    {
        field: "Name",
        headerName: "Name",
        type: "string",
        width: 350,
        sortable: false,
        filterable: false,
        editable: false
    },
    {
        field: "Location",
        headerName: "Location",
        type: "string",
        width: 350,
        sortable: false,
        filterable: false,
        editable: false
    },
    {
        field: "Create",
        align: "left",
        headerName: "Create Availability",
        width: 350,
        renderCell: RenderCreateButton,
        disableClickEventBubbling: true   
    },
    {
        field: "Get",
        align: "left",
        headerName: "All Availabilities",
        width: 350,
        renderCell: RenderAvailabilitiesButton,
        disableClickEventBubbling: true   
    }
];


function rowAction(navigate, buttonName, buttonUrl) {
    return {
        field: "Details",
        headerName: buttonName,
        align: "left",
        headerAlign: "left",
        sortable: false,
        renderCell: (params) => {
            const onClick = (e) => {
                e.stopPropagation(); // don't select this row after clicking

                const api = params.api;
                const thisRow = params.row;

                


                return navigate("/accommodation-details", { state: thisRow });
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

function AvailabilityList(props) {
    const { handleSubmit, control } = useForm();
    const [accommodations, setAccomodations ] = useState([]);
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
        console.log('Id',res.data.user.Id);

        axiosApi
            .get(`/accommodation/all/`+res.data.user.Id)
            .then((response) => {
                console.log(response.data)
                setAccomodations(response.data);
            }).catch(er => {
                console.log(er.response);
                setAccomodations([]);
            });
        }catch (err) {
                console.log(err)
                setAccomodations([]);
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
                    Accommodations
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
                                        setAccomodations([])
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
                        rows={accommodations}
                        getRowId={(row) => row.Id}
                        disableColumnFilter
                        columns={[...columns, rowAction(navigate, props.buttonName, props.buttonUrl)]}
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
