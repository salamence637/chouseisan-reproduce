import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import { BrowserRouter } from "react-router-dom";
import { Box } from "@mui/material";
import SelfEventListProvider from "./contexts/EventBySelf";
import HistoryEventProvider from "./contexts/HistoryEvent";
const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  <React.StrictMode>
    <BrowserRouter basename="/scheduler">
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          fontFamily: "Lucida Grande",
        }}
      >
        <HistoryEventProvider>
          <SelfEventListProvider>
            <App />
            
          </SelfEventListProvider>
        </HistoryEventProvider>
      </Box>
    </BrowserRouter>
  </React.StrictMode>
  
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
