import { useState, useEffect, useContext } from "react";

import { ListItem, ListItemText, Button, Grid, Link } from "@mui/material";
import "./HistorySimpler.css";
import axios from "../utils/axios";
import { HistoryEventContext } from "../contexts/HistoryEvent";
import ArrowForwardIosIcon from "@mui/icons-material/ArrowForwardIos";
import { useNavigate } from "react-router-dom";
import { historyEvent } from "../types/Event";
export default function HistorySimpler() {
  const { historyEvent, setHistoryEvent } = useContext(HistoryEventContext);
  const arr = ["1", "2"];
  const navigate = useNavigate();

  const buttonStyle = {
    width: "465px",
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
  };
  // localStorage.clear();
  return (
    <div className="bg">
      <div className="history-simpler-container">
        <div className="header">
          Recently viewed events
          <p className="paragraph">Other users won't be able to see this</p>
        </div>
      </div>

      <div className="history-card">
        {historyEvent.slice(0, 2).map((value, index) => {
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
                navigate(`view_event/${value["uuid"].replace(/-/g, "")}`);
              }}
            >
              <Grid container sx={{ height: "100%" }} spacing={1}>
                <Grid
                  item
                  xs={12}
                  sx={{ fontWeight: "bold", color: "black", fontSize: "18px" }}
                >
                  {value["title"]}
                </Grid>
                {value["scheduleList"].slice(0, 6).map((timeslot, idx) => (
                  <Grid item xs={4} key={idx}>
                    <ListItem
                      key={2 * idx}
                      sx={{
                        border: "1px solid #ccc",
                        borderRadius: "4px",
                        "& .css-10hburv-MuiTypography-root": {
                          fontSize: "12px", // 调整字体大小
                          color: "black",
                        },
                        maxWidth: "130px",
                      }}
                    >
                      <ListItemText
                        primary={`${timeslot}`}
                        primaryTypographyProps={{
                          style: {
                            padding: "1px",
                            maxWidth: `20ch`, // 设置最大宽度
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
      <Link
        href="/scheduler/history"
        color={"#a46702"}
        underline="hover"
        sx={{ marginBottom: "15px" }}
      >
        <ArrowForwardIosIcon sx={{ fontSize: 11 }} />
        View complete history
      </Link>
    </div>
  );
}
