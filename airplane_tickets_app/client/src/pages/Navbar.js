import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Toolbar from "@mui/material/Toolbar";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import ListItemText from "@mui/material/ListItemText";
import Box from "@mui/material/Box";
import Drawer from "@mui/material/Drawer";
import List from "@mui/material/List";
import ListItem from "@mui/material/ListItem";
import ListItemButton from "@mui/material/ListItemButton";
import { useNavigate } from "react-router-dom";
import { Button } from "@mui/material";
import { useState, useEffect, useMemo } from "react";
import axiosApi from "../api/axios";

const getData = async () =>
        axiosApi
            .get(`/users/logged/`)
            .then((response) => {
                return response.data
            })
            .catch(function (error) {
                if (error.response) {
                    // Request made and server responded
                    console.log(error.response.data);
                    console.log(error.response.status);
                    console.log(error.response.headers);
                } else if (error.request) {
                    // The request was made but no response was received
                    console.log(error.request);
                } else {
                    // Something happened in setting up the request that triggered an Error
                    console.log("Error", error.message);
                }
                return {};
            });

export default function Navbar() {
    const [state, setState] = React.useState(false);
    let navigate = useNavigate();
    const [user, setUser] = useState("");

    
    useEffect(() => {
        getData().then(setUser);
    },[]);
    const sidemenu = useMemo(() => chooseSideMenu(user.role), [user.role]);
    const button1 = useMemo(() => chooseButton1(user.role), [user.role]);
    const button2 = useMemo(() => chooseButton2(user.role), [user.role]);

    const toggleDrawer = (anchor, open) => (event) => {
        if (
            event.type === "keydown" &&
            (event.key === "Tab" || event.key === "Shift")
        ) {
            return;
        }

        setState({ ...state, [anchor]: open });
    };

    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static" sx={{ background: "#5B63F5" }}>
                <Toolbar sx={{ justifyContent: "space-between" }}>
                    {
                        <div>
                            <IconButton
                                size="large"
                                aria-label="account of current user"
                                aria-controls="big-menu-appbar"
                                aria-haspopup="true"
                                onClick={toggleDrawer("left", true)}
                                color="inherit"
                            >
                                <MenuIcon />
                            </IconButton>
                            <Drawer
                                anchor={"left"}
                                open={state["left"]}
                                onClose={toggleDrawer("left", false)}
                            >
                                <Box
                                    role="presentation"
                                    onClick={toggleDrawer("left", false)}
                                    onKeyDown={toggleDrawer("left", false)}
                                >
                                    <List>
                                        {Object.keys(sidemenu).map((k) => (
                                            <ListItem key={k} disablePadding>
                                                <ListItemButton>
                                                    <ListItemText
                                                        primary={k}
                                                        onClick={() =>
                                                            navigate(
                                                                sidemenu[k]
                                                            )
                                                        }
                                                    />
                                                </ListItemButton>
                                            </ListItem>
                                        ))}
                                    </List>
                                </Box>
                            </Drawer>
                        </div>
                    }
                <Box
                     m={1}
                     display="flex"
                     justifyContent="flex-end"
                     alignItems="flex-end"
                >
                <Button color="inherit" onClick={() => navigate("/")}>
                        Home
                    </Button>
                    <Button
                        color="inherit"
                        sx={{ marginLeft: "auto" }}
                        onClick={() => {
                            if (button1.name === "LogOut") {
                                localStorage.removeItem("token");
                                delete axiosApi.defaults.headers.common[
                                    "Authorization"
                                ];
                                navigate("/login/");
                            } else {
                                navigate(button1.url);
                            }
                        }}
                    >
                        {button1.name}
                    </Button>
                    <Button
                        color="inherit"
                        onClick={() => navigate(button2.url)}
                    >
                        {button2.name}
                    </Button>
                    </Box>
                </Toolbar>
            </AppBar>
        </Box>
    );
}

const chooseSideMenu = (role) => {
    if (!role) return { "Search flights": "/flights/all" };
    if (role === "ADMIN") {
        return {
            "Search flights": "/flights-admin/all",
            "Create flight": "/flights/create/"
        };
    } else if (role === "REGULAR") {
        return {
            "Search flights": "/flights/all-user",
            "My tickets":"/tickets-user"
        };
    } else {
        return { "Search flights": "/flights/all" };
    }
};

const chooseButton1 = (role) => {
    if (!role) return { name: "Login", url: "/login" };
    else {
        return { name: "LogOut", url: "/logout" };
    }
};

const chooseButton2 = (role) => {
    if (!role) return { name: "SignUp", url: "/signup" };
    else {
        return { name: "", url: "/" };
    }
};
