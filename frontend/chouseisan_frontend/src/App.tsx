import "./App.css";
import "./App.css";
import Box from "@mui/material/Box";
import CssBaseline from "@mui/material/CssBaseline";
import Toolbar from "@mui/material/Toolbar";
import { Link } from "react-router-dom";
import RouteSetting from "./utils/RouteSetting";
import HistorySimpler from "./components/HistorySimpler";
import { AppBar, Button } from "@mui/material";
import NotificationsNoneRoundedIcon from "@mui/icons-material/NotificationsNoneRounded";
import logoIcon from "./images/chousei_logo.png";
export default function App() {
  const location = window.location.href.split("/").pop();
  return (
    <>
      <Box
        sx={{
          display: "flex",
        }}
      >
        <CssBaseline />
        <AppBar position="relative" sx={{ backgroundColor: "#34a21a" }}>
          <Toolbar sx={{ width: 1000, margin: "0 auto" }}>
            <Link to="/">
              <img
                src={logoIcon}
                alt="Logo"
                height={45}
                style={{ marginRight: "12px" }}
              />
            </Link>
            <Box
              sx={{
                fontSize: 14,
                fontWeight: "bolder",
                fontFamily: "sans-serif",
                width: 200,
                margintop: 12,
              }}
            >
              <p style={{ margin: 0 }}>Host an event without hassle</p>
              <p style={{ margin: 0 }}>Esay scheduling!</p>
            </Box>
            <div
              style={{
                alignItems: "flex-start",
                backgroundColor: "white",
                borderRadius: "5px",
                display: "flex",
                flexDirection: "row",
                fontSize: 14,
                overflow: "visible",
                padding: "5px 5px 5px 10px",
                position: "absolute",
                right: 20,
                top: 10,
                height: 44,
                justifyContent: "space-between ",
              }}
            >
              <p
                className="pr-p"
                style={{ color: "#666", margin: 0, width: 210 }}
              >
                Login is not required but it gives you more convenience!
              </p>
              <Button
                variant="contained"
                sx={{
                  color: "#fff",
                  textAlign: "center",
                }}
              >
                REGISTER/LOGIN
              </Button>
              <Button
                sx={{
                  borderRadius: 3,
                  color: "#34a21a",
                  textAlign: "center",
                  borderColor: "#34a21a",
                  border: "1px solid",
                  marginLeft: 1,
                }}
              >
                Notice
                <NotificationsNoneRoundedIcon />
              </Button>
            </div>
          </Toolbar>
        </AppBar>
      </Box>
      <RouteSetting />
      {location !== "history" && <HistorySimpler />}
    </>
  );
}
