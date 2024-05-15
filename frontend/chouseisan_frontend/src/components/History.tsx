import axios from "../utils/axios";
import { useState, useEffect, useContext } from "react";
import { useNavigate } from "react-router-dom";

import { Button, Grid, Link, ListItem, ListItemText } from "@mui/material";

import { HistoryEventContext } from "../contexts/HistoryEvent";

import "./History.css";
export default function History() {
  const { historyEvent, setHistoryEvent } = useContext(HistoryEventContext);
  const navigate = useNavigate();
  console.log(historyEvent);

  const buttonStyle = {
    width: "700px",
    height: "180px",
    border: "1px solid #ccc",
    "&:hover": {
      backgroundColor: "#f8f6e3",
      border: "1px solid #ccc",
    },
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    marginTop: "30px",
  };
  return (
    <>
      <div className="topBox">
        <p className="first">
          <Link
            href="/scheduler"
            color={"#a46702"}
            underline="hover"
            sx={{ marginBottom: "5px" }}
          >
            Top
          </Link>
          {" > "}Recently viewed events
        </p>
        <h1>Recently viewed events</h1>
      </div>
      <div style={{ backgroundColor: "#f1f1f1e6" }}>
        <div className="bottomBox">
          <Button
            className="history-item"
            disabled
            sx={{
              ...buttonStyle,
              marginRight: "20px",
              backgroundColor: "#fff",
              height: "50px",
              color: "green",
              fontWeight: "600",
              "&.Mui-disabled": { color: "green" },
            }}
            variant="outlined"
          >
            Upto 5 most recent events are displayed here.
          </Button>
          {historyEvent.slice(0, 5).map((value, index) => {
            // console.log(timeslotList);
            // console.log(title);
            return (
              <Button
                className="history-item"
                sx={{
                  ...buttonStyle,
                  marginRight: "20px",
                  backgroundColor: "#fff",
                }}
                variant="outlined"
                onClick={() => {
                  console.log("clicked");
                  navigate(`../view_event/${value["uuid"].replace(/-/g, "")}`);
                }}
              >
                <Grid container sx={{ height: "100%" }} spacing={1}>
                  <Grid
                    item
                    xs={12}
                    sx={{
                      fontWeight: "bold",
                      color: "black",
                      fontSize: "18px",
                    }}
                  >
                    {value["title"]}
                  </Grid>
                  {value["scheduleList"].slice(0, 6).map((timeslot, index) => (
                    <Grid item xs={4} key={index}>
                      <ListItem
                        sx={{
                          border: "1px solid #ccc",
                          borderRadius: "4px",
                          "& .css-10hburv-MuiTypography-root": {
                            fontSize: "12px", // 调整字体大小
                            color: "black",
                          },
                          maxWidth: "200px",
                        }}
                      >
                        <ListItemText
                          primary={`${timeslot}`}
                          primaryTypographyProps={{
                            style: {
                              padding: "1px",
                              maxWidth: `40ch`, // 设置最大宽度
                              overflow: "hidden",
                              textOverflow: "ellipsis",
                              whiteSpace: "nowrap",
                            },
                          }}
                        />
                      </ListItem>
                    </Grid>
                  ))}
                </Grid>
              </Button>
            );
          })}
        </div>
      </div>
    </>
  );
}
