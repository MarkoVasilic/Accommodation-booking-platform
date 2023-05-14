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


/*const RenderMakeReservation = (params) => {
    let navigate = useNavigate();
    return (
        <strong>
            <Button
                variant="contained"
                color="primary"
                size="small"
                style={{ marginLeft: 16 }}
                
                onClick={() => {
                    const res = axiosApi.get('/user/logged');
                    const data = {AvailabilityID: params.row.AvailabilityID, GuestId: res.data.user.Id, StartDate: startDate, EndDate: endDate, GuestsNum: guestsNum}
                    axiosApi
                    .post(`/reservation/`) //dodati data
                    .then((response) => {
                        console.log("AAA")
                        console.log(response.data)
                        navigate('/pending-reservations')
                    }).catch(er => {
                        console.log(er.response);
                    });
                }}
            >
                Reserve it
            </Button>
        </strong>
    )
};*/

function formatPrice(pricePerGuest) {
    var IsPricePerGuest = ""
    if (pricePerGuest === true) {
        IsPricePerGuest = "Yes"
    } else {
        IsPricePerGuest = "No"
    }
    //console.log('Sec', seconds)
    //console.log('date', date)
    return IsPricePerGuest;
  }
const columns = [
    {
        field: "Name",
        headerName: "Name",
        type: "string",
        width: 200,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "Location",
        headerName: "Location",
        type: "string",
        width: 270,
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "SinglePrice",
        headerName: "Price per guest / accomodation",
        type: "number",
        width: 280,
        headerAlign: "left",
        align: "left",
        sortable: false,
        filterable: false,
        editable: false,
    },
    {
        field: "TotalPrice",
        headerName: "Total Price",
        type: "number",
        width: 200,
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
        width: 200,
        sortable: false,
        filterable: false,
        editable: false,
        valueGetter: params => formatPrice(params.row.IsPricePerGuest)
    },
    /*{
        field: "make_reservation",
        headerName: "Make a reservation",
        width: 260,
        renderCell: RenderMakeReservation,
        disableClickEventBubbling: true   
    }*/
];

function MakeReservation(navigate, startDate, endDate, guestsNum) {
    return {
        field: "Make reservation",
        headerName: "Reserve",
        align: "center",
        headerAlign: "left",
        sortable: false,
        renderCell: (params) => {
            const onClick = async(e) => {
                e.stopPropagation(); // don't select this row after clicking

                const api = params.api;
                const thisRow = params.row;

                    const res = await axiosApi.get('/user/logged');
                    console.log('Id',res.data.user.Id);
                
                console.log('User', res.data.user.Id);
                var guests = parseInt(guestsNum)
                const data = {AvailabilityID: params.row.AvailabilityID, GuestId: res.data.user.Id, StartDate: startDate, EndDate: endDate, NumGuests: guests}
                console.log('Data', data)
                axiosApi
                .post('/reservation', data) //dodati data
                .then((response) => {
                    console.log("AAA")
                    console.log(response.data)
                    navigate('/pending-reservations')
                }).catch(er => {
                    console.log(er.response);
                });
            

               // return navigate({ state: thisRow });
            };
            return (
                <Button
                    variant="contained"
                    color="secondary"
                    size="small"
                    onClick={onClick}
                >
                    Reserve
                </Button>
            );
        },
    };
}

function rowAction(navigate, buttonName, buttonUrl, startDate, endDate, guestsNum) {
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

function ListSearchedAvailabilityGuest(props) {
    const { handleSubmit, control } = useForm();
    const [accomodation, setAccomodation ] = useState([]);
    const [ error, setError ] = React.useState(false);
    const [er, setEr] = React.useState("");
    const navigate = useNavigate();
    const [ startDate, setStartDate ] = React.useState("");
    const [ endDate, setEndDate ] = React.useState("");
    const [ guestsNum, setGuestsNum ] = React.useState(0);

    useEffect(() => {
      //  getData();
        onSubmit();
    }, []);
    const date = new Date().toISOString();

        const onSubmit = async (data) => {
            try {
                console.log(data)
                data.StartDate = new Date(Date.parse(data.StartDate))
                data.EndDate = new Date(Date.parse(data.EndDate))   
                data.GuestsNum = parseInt(data.NumGuests)
                setStartDate(data.StartDate)
                setEndDate(data.EndDate)
                setGuestsNum(data.NumGuests)
                console.log('Data', data)

                await axiosApi
                .post('/availability/search', data)
                .then((response) => {
                    console.log("Res", response)
                    setAccomodation(response.data);
                }).catch(er => {
                    console.log(er.response);
                    setAccomodation([]);
                });
            }catch (err) {
                    console.log(err)
                    setAccomodation([]);
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
                    Search accomodations
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
                                Choose start date:</Typography>
                            <InputTextField
                                name="StartDate"
                                control={control}
                                type="date"
                                rules={{ required: "This field is required" }}
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Choose end date:</Typography>
                            <InputTextField
                                name="EndDate"
                                control={control}
                                type="date"
                                rules={{ required: "This field is required" }}
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Enter location:</Typography>
                            <InputTextField
                                name="Location"
                                control={control}
                                type="text"
                                rules={{ required: "This field is required" }}
                            />
                        </Grid>
                        <Grid item xs={12} md={2}>
                            <Typography>
                                Number of guests:</Typography>
                            <InputTextField
                                name="NumGuests"
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
                                        setAccomodation([])
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
                        rows={accomodation}
                        getRowId={(row) => row.AccommodationId}
                        disableColumnFilter
                        columns={[...columns, rowAction(navigate, props.buttonName, props.buttonUrl), MakeReservation(navigate, startDate, endDate, guestsNum)]}
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

export default ListSearchedAvailabilityGuest;
