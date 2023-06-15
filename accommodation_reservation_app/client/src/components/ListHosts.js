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




const RenderRateHost = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                onClick={() => {
                    console.log(params.row)
                    navigate('/host/rate', {state: params.row});
                }}
            >
                Rate
            </Button>
        </strong>
    )
};

const RenderGetAllHostGrades = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                onClick={() => {
                    console.log(params.row)
                    navigate('/host/all-grades', {state: params.row});
                }}
            >
                All grades
            </Button>
        </strong>
    )
};



const columns = [
    {
        field: "FirstName",
        headerName: "Host first name",
        type: "string",
        width: 250,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "LastName",
        headerName: "Host last name",
        type: "string",
        width: 250,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "Name",
        headerName: "Accommodation name",
        type: "string",
        width: 250,
        sortable: false,
        filterable: false,
        editable: false,  
    },
    {
        field: "Location",
        headerName: "Accommodation location",
        type: "string",
        width: 250,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "rate",
        headerName: "Rate host",
        width: 250,
        renderCell: RenderRateHost,
        disableClickEventBubbling: true   
    },
    {
        field: "allGrades",
        headerName: "All grades",
        width: 250,
        renderCell: RenderGetAllHostGrades,
        disableClickEventBubbling: true   
    }
];


function HostsList(props) {
    const { handleSubmit, control } = useForm();
    const [hosts, setHosts ] = useState([]);
    const [ error, setError ] = React.useState(false);
    const [er, setEr] = React.useState("");
    const navigate = useNavigate();
    useEffect(() => {
        getData();
      //  onSubmit();
    }, [setHosts]);
    const date = new Date().toISOString();

        let getData = async () => {
        try{
            //console.log()
            //const res = await axiosApi.get('/user/logged');
            //console.log("ID", res.data.user.Id);
        axiosApi
            //proslediti koji treba (proveriti jel ovaj)
            //.get(`/user/host/all`)
            .then((response) => {
                setHosts(response.data);
                console.log('Data', response.data)
                console.log('RES', hosts)
            }).catch(er => {
                console.log(er.response);
                setHosts([]);
            });
        }catch (err) {
                console.log(err)
                setHosts([]);
            }
            console.log('RESS',hosts)

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
                    Hosts
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
                                        setHosts([])
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
                        rows={hosts}
                        //getRowId={(row) => row.ReservationId}
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

export default HostsList;
