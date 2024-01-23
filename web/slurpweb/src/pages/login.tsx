import {
  Box,
  Button,
  Container,
  Grid,
  TextField,
  Typography,
  styled,
} from "@mui/material";
import React from "react";
import ReactDOM from "react-dom/client";
const StyledTextField = styled(TextField)({
  "& label": {
    color: "white",
  },
  "&:hover label": {
    fontWeight: 700,
  },
  "& label.Mui-focused": {
    color: "white",
  },
  "& .MuiInput-underline:after": {
    borderBottomColor: "white",
  },
  "& .MuiOutlinedInput-root": {
    "& fieldset": {
      borderColor: "white",
    },
    "&:hover fieldset": {
      borderColor: "white",
      borderWidth: 2,
    },
    "&.Mui-focused fieldset": {
      borderColor: "white",
    },
  },
});
export function Login() {
  return (
    <Container
      sx={{
        display: "flex",
        flexDirection: "column",
        justifyContent: "center",
        alignItems: "center",
        height: "100vh",
      }}
      maxWidth={false}
    >
      <Typography variant="h2" fontFamily={"Major Mono Display"}>
        slurp
      </Typography>
      <Typography variant="subtitle1" fontFamily={"Major Mono Display"}>
        login or create account
      </Typography>
      <StyledTextField
        variant="outlined"
        label="Username"
        margin="normal"
        sx={{ color: "white" }}
      >
        {" "}
      </StyledTextField>
      <StyledTextField
        variant="outlined"
        type="password"
        margin="normal"
        label="Password"
        sx={{ color: "white" }}
      >
        {" "}
      </StyledTextField>
      <Button variant="contained" href="#contained-buttons">
        Login
      </Button>
    </Container>
  );
}
