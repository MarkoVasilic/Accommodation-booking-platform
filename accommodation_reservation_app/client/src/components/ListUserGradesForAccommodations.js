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


function formatSecondsToDate(seconds) {
    const date = new Date(seconds * 1000 - 7200*1000);
    //console.log('Sec', seconds)
    //console.log('date', date)
    return date;
  }

const RenderUpdateAccommodationGrade = (params) => {
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
                    navigate('/accommodation-grade/update', {state: params.row});
                }}
            >
                Update
            </Button>
        </strong>
    )
};

const RenderDeleteAccommodationGrade = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                onClick={() => {
                    //proveriti url
                    console.log("params", params.row)
                    axiosApi
                    //proslediti koji treba
                    //ne reservationId nego accommodation vrv
                    .delete(`/accommodation/grade/`+params.row.ID)
                    .then((response) => {
                        refreshPage();
                    }).catch(er => {
                        console.log(er.response);
                    });
                }}
            >
                Delete
            </Button>
        </strong>
    )
};


const columns = [
    {
        field: "AccommodationName",
        headerName: "Accommodation Name",
        type: "string",
        width: 250,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "Grade",
        headerName: "Grade",
        type: "number",
        width: 250,
        sortable: false,
        filterable: false,
        editable: false,  
    },
    {
        //proveriti za tip i to ostalo
        field: "DateOfGrade",
        headerName: "Date of grade",
        type: "date",
        width: 250,
        sortable: false,
        filterable: false,
        editable: false,
        format: "DD/MM/YYYY",
        valueGetter: params => formatSecondsToDate(params.row.DateOfGrade.seconds)
    },
    {
        field: "update",
        headerName: "Update grade",
        width: 250,
        renderCell: RenderUpdateAccommodationGrade,
        disableClickEventBubbling: true   
    },
    {
        field: "delete",
        headerName: "Delete grade",
        width: 250,
        renderCell: RenderDeleteAccommodationGrade,
        disableClickEventBubbling: true   
    }
];


function AccommodationsGradesUser(props) {
    const { handleSubmit, control } = useForm();
    const [grades, setGrades ] = useState([]);
    const [ error, setError ] = React.useState(false);
    const [er, setEr] = React.useState("");
    const navigate = useNavigate();
    useEffect(() => {
        getData();
      //  onSubmit();
    }, [setGrades]);
    const date = new Date().toISOString();

        let getData = async () => {
        try{
            //console.log()
            const res = await axiosApi.get('/user/logged');
            //console.log("ID", res.data.user.Id);
        axiosApi
            //proslediti koji treba (proveriti jel ovaj)
            //znam da {id} ne treba ovako ali cisto url
            //dodati metodu za dobavljanje ocena smestaja po guestId
            .get(`/accommodation/user/grades/${res.data.user.Id}`)
            .then((response) => {
                if(response.data === null) {
                    setGrades([]);
                }
                else {
                    setGrades(response.data);
                console.log('Data', response.data)
                console.log('RES', grades)
                }
                
            }).catch(er => {
                console.log(er.response);
                setGrades([]);
            });
        }catch (err) {
                console.log(err)
                setGrades([]);
            }
            console.log('RESS',grades)

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
                    My Grades For Accommodations
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
                                        setGrades([])
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
                        rows={grades}
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

export default AccommodationsGradesUser;
