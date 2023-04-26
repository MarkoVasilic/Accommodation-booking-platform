import PersonalInformationCard from "../components/PersonalInformationCard";
import { green } from '@mui/material/colors';
import { Typography, Paper, Button } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import { Stack } from '@mui/system';
import axiosApi from "../api/axios";
import Navbar from "./Navbar";
import AllowedUsers from "../components/AllowedUsers";
import React from "react";
import Box from "@mui/material/Box";
import Alert from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import CloseIcon from "@mui/icons-material/Close";

export default function UserProfile() {
    const listOfAllowedUsers = ["HOST", "GUEST"];
    let navigate = useNavigate();
    const [failedAlert, setFailedAlert] = React.useState(false);
    const routeChange = () =>{ 
        let path = `/user-profile/update/`; 
        navigate(path);
      }

    const changePasswordRoute = () =>{ 
        let path = '/user-profile/password/'; 
        navigate(path);
      }
    
    const deleteProfile = async () => {
        try {
            const res = await axiosApi.get('/user/logged');
            await axiosApi.delete(`/user/${res.data.user.Id}`);
            localStorage.removeItem("token");
            delete axiosApi.defaults.headers.common[
                "Authorization"
            ];
            navigate("/login/");
        }
        catch (err) {
            console.log(err)
            setFailedAlert(true)
        }
    };

    return (
        <div>
            <AllowedUsers userRole = {listOfAllowedUsers}/>
            <Navbar />
            <Stack marginTop={"10px"} justifyContent={"center"}>
            <Typography align="center"  marginBottom={"20px"}  component="h1" variant="h4" color={green[800]}>
                Profile
            </Typography>
            <Button
                        variant="contained"
                        onClick={routeChange}
                        style={{
                            width: 200,
                            marginLeft: 1300,
                            marginBottom: 10,
                            marginTop: -45
                        }}
                        sx={{
                            background: "#6fbf73",
                            height: "100",
                            "&.MuiButtonBase-root": {
                                "&:hover": {
                                    backgroundColor: green[600],
                                },
                            },
                        }}
                    >
                        Update Profile
                    </Button>
                    <Button
                        variant="contained"
                        onClick={changePasswordRoute}
                        style={{
                            width: 200,
                            marginLeft: 1300,
                            marginBottom: -10
                        }}
                        sx={{
                            background: "#6fbf73",
                            height: "100",
                            "&.MuiButtonBase-root": {
                                "&:hover": {
                                    backgroundColor: green[600],
                                },
                            },
                        }}
                    >
                        Change Password
                    </Button>
                    <Button
                        variant="contained"
                        onClick={deleteProfile}
                        style={{
                            width: 200,
                            marginLeft: 1300,
                            marginTop: 20,
                            marginBottom: -10
                        }}
                        sx={{
                            background: "#6fbf73",
                            height: "100",
                            "&.MuiButtonBase-root": {
                                "&:hover": {
                                    backgroundColor: green[600],
                                },
                            },
                        }}
                    >
                        Delete Profile
                    </Button>
            <Typography align="left" marginTop={"-10px"} marginLeft={"110px"} marginBottom={"5px"} component="h4" variant="h4" color={green[800]}>
                Personal information
            </Typography>
            <Paper elevation={10} sx={{ p: { sm: 2, xs: 2 } }}>
            <PersonalInformationCard></PersonalInformationCard>
            </Paper>
        </Stack>
        <Box sx={{ width: "100%" }}>
                <Collapse in={failedAlert}>
                    <Alert
                        severity="error"
                        action={
                            <IconButton
                                aria-label="close"
                                color="inherit"
                                size="small"
                                onClick={() => {
                                    setFailedAlert(false);
                                }}
                            >
                                <CloseIcon fontSize="inherit" />
                            </IconButton>
                        }
                        sx={{ mb: 2 }}
                    >
                    Profile deleition failed!
                    </Alert>
                </Collapse>
            </Box>
        </div>
    );
}