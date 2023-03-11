import * as React from "react";
import Avatar from "@mui/material/Avatar";
import Button from "@mui/material/Button";
import CssBaseline from "@mui/material/CssBaseline";
import TextField from "@mui/material/TextField";
import Box from "@mui/material/Box";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import Typography from "@mui/material/Typography";
import Container from "@mui/material/Container";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { sumbitLogin } from "../api/signInOn";
import { useForm, Controller } from "react-hook-form";
import { setAuthToken } from "../helpers/sethAuthToken";
import Alert from "@mui/material/Alert";
import IconButton from "@mui/material/IconButton";
import Collapse from "@mui/material/Collapse";
import CloseIcon from "@mui/icons-material/Close";
import { useNavigate } from "react-router-dom";
import { blue } from "@mui/material/colors";
import Navbar from "../pages/Navbar";

const theme = createTheme();

export default function Login() {
    const { handleSubmit, control, getValues } = useForm();
    const [alert, setAlert] = React.useState(false);
    let navigate = useNavigate();

    const onSubmit = async (data) => {
        data.email = data.username;
        delete data.username;
        try {
            const response = await sumbitLogin(data);
            const token = response.data.token;
            localStorage.setItem("token", token);
            setAuthToken(token);
            navigate("/");
            window.location.reload(false);
        } catch (err) {
            console.log(err)
            setAlert(true);
        }
    };

    function isValidEmail(email) {
        return /\S+@\S+\.\S+/.test(email);
    }

    return (
        <div>
            <Navbar />
            <ThemeProvider theme={theme}>
                <Container component="main" maxWidth="xs">
                    <CssBaseline />
                    <Box
                        sx={{
                            marginTop: 8,
                            display: "flex",
                            flexDirection: "column",
                            alignItems: "center",
                        }}
                    >
                        <Avatar sx={{ m: 1, bgcolor: "#5B63F5" }}>
                            <LockOutlinedIcon />
                        </Avatar>
                        <Typography component="h1" variant="h4" color={"#5B63F5"}>
                            Login
                        </Typography>
                        <form onSubmit={handleSubmit(onSubmit)}>
                            <Box sx={{ mt: 1 }}>
                                <Controller
                                    name="username"
                                    control={control}
                                    defaultValue=""
                                    render={({
                                        field: { onChange, value },
                                        fieldState: { error },
                                    }) => (
                                        <TextField
                                            margin="normal"
                                            fullWidth
                                            label="Email"
                                            autoComplete="email"
                                            autoFocus
                                            variant="filled"
                                            value={value}
                                            onChange={onChange}
                                            error={!!error}
                                            helperText={
                                                error ? error.message : null
                                            }
                                        />
                                    )}
                                    rules={{
                                        required: "Email required",
                                        validate: {
                                            validateEmail: (v) =>
                                                isValidEmail(
                                                    getValues("username")
                                                ) || "Email form not correct",
                                        },
                                    }}
                                />
                                <Controller
                                    name="password"
                                    control={control}
                                    defaultValue=""
                                    render={({
                                        field: { onChange, value },
                                        fieldState: { error },
                                    }) => (
                                        <TextField
                                            margin="normal"
                                            fullWidth
                                            label="Password"
                                            type="password"
                                            autoComplete="current-password"
                                            variant="filled"
                                            value={value}
                                            onChange={onChange}
                                            error={!!error}
                                            helperText={
                                                error ? error.message : null
                                            }
                                        />
                                    )}
                                    rules={{
                                        required: "Password required",
                                        minLength: {
                                            value: 6,
                                            message:
                                                "Password needs to be longer than 5",
                                        },
                                    }}
                                />
                                <Button
                                    type="submit"
                                    fullWidth
                                    variant="contained"
                                    sx={{ mt: 3, mb: 2, background: "#5B63F5",  "&.MuiButtonBase-root": {
                                        "&:hover": {
                                            backgroundColor: blue[600]
                                        }
                                        }}}
                                >
                                    Sign In
                                </Button>
                            </Box>
                        </form>
                    </Box>
                    <Collapse in={alert}>
                        <Alert
                            severity="error"
                            action={
                                <IconButton
                                    aria-label="close"
                                    color="inherit"
                                    size="small"
                                    onClick={() => {
                                        setAlert(false);
                                    }}
                                >
                                    <CloseIcon fontSize="inherit" />
                                </IconButton>
                            }
                            sx={{ mb: 2 }}
                        >
                            Login failed!
                        </Alert>
                    </Collapse>
                </Container>
            </ThemeProvider>
        </div>
    );
}
