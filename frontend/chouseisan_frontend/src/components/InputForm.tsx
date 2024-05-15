import React, { useState, FormEvent, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { TextField, Button, Box } from "@mui/material";
import dayjs from "dayjs";
import "./InputForm.css";
import topIcon from "../images/top.png";
import FlagIcon from "@mui/icons-material/Flag";
import { DemoContainer } from "@mui/x-date-pickers/internals/demo";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";

import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DateCalendar, DateTimePicker } from "@mui/x-date-pickers";
import axios from "../utils/axios";
import { SelfEventContext } from "../contexts/EventBySelf";

export default function InputForm() {
  const japanTime = dayjs();
  const [dateList, setDateList] = React.useState("");
  const [title, setTitle] = useState("");
  const [detail, setDetail] = useState("");
  const { selfEventList, setSelfEventList } = useContext(SelfEventContext);
  const [expiration, setExpiration] = useState(
    dayjs(japanTime).add(7, "day").toString()
  );
  const navigate = useNavigate();
  // localStorage.clear();
  const eventSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    axios
      // "http://localhost:8080/"
      .post(
        `/event`,
        {
          title: title,
          detail: detail,
          dateTimeProposal: dateList,
          due_edit: expiration,
        },

        {
          headers: { "Content-Type": "application/json" },
          withCredentials: true,
        }
      )
      .then(function (response) {
        let uuid = response.data.event_id.replace(/-/g, "");
        setSelfEventList((selfEventList) => [...selfEventList, uuid]);

        navigate(`/create_complete/${uuid}`);
      })
      .catch(function (response) {
        console.log("ERROR connecting backend service");
      });
  };
  return (
    <>
      <Box sx={{ backgroundColor: "#6dd643", height: 350, marginBottom: 0 }}>
        <div
          style={{
            margin: "0 auto",
            width: 970,
            position: "relative",
          }}
        >
          <img src={topIcon} height="auto" width={"100%"} alt="topicon" />
          <div
            style={{
              left: 20,
              position: "absolute",
              top: 8,
              width: 640,
              color: "white",
              borderRadius: 10,
            }}
          >
            <h2>
              Chouseisan organizes every detail about your event! It all starts
              by creating an event page!
            </h2>
          </div>
        </div>
      </Box>
      <Box
        sx={{
          backgroundColor: "#eaf4e5",
          minheight: 1000,
        }}
      >
        <form className="container" onSubmit={eventSubmit}>
          <h2 className="form-header">
            <FlagIcon />
            Create your event page
          </h2>
          <div className="form">
            <div className="box1">
              <div className="event-box">
                <p className="item-title">
                  <span className="step-label">STEP1</span>Event Title
                </p>
                <p className="item-description">
                  “Team Dinner Party”, “Project Meeting”, etc...
                </p>
                <TextField
                  // size="small"
                  label="Title"
                  fullWidth
                  inputProps={{ style: { padding: 0 } }}
                  required
                  onChange={(e) => setTitle(e.target.value)}
                ></TextField>
              </div>
              <div className="event-box">
                <p className="item-title">Event Details (Optional)</p>
                <p className="item-description">
                  Let’s schedule the party! Please respond by ______
                </p>
                <TextField
                  size="small"
                  label="Detail"
                  multiline
                  fullWidth
                  rows={16}
                  inputProps={{ style: { padding: 0 } }}
                  onChange={(e) => setDetail(e.target.value)}
                ></TextField>
              </div>
            </div>
            <div className="box2">
              <div className="event-box">
                <p className="item-title">
                  <span className="step-label">STEP2</span>Date/Time Proposals
                </p>
                <p className="item-description">
                  List the dates and corresponding times propose to host an
                  event.<br></br>*Input one proposal per line.
                </p>
                <p className="item-description">
                  Example:<br></br>　Aug 7(Mon) 20:00～<br></br>　Aug 8(Tue)
                  20:00～<br></br>　Aug 9(Wed) 21:00～
                </p>

                <TextField
                  sx={{ marginBottom: "20px" }}
                  size="small"
                  multiline
                  fullWidth
                  rows={7}
                  label="Proposal"
                  inputProps={{ style: { padding: 0 } }}
                  onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setDateList(event.target.value);
                  }}
                  value={dateList}
                  placeholder="Simply input your proposals in the Month DD(DAY) TIME format. Or you can click on the specific date(s) in the calendar."
                ></TextField>
                <p className="item-title">Expiration date</p>
                <p className="item-description">
                  Select the expiration date that you want to close this event.
                  <br></br>
                  Please notice users cannot add attendance on this event from
                  that date.
                </p>
                <LocalizationProvider dateAdapter={AdapterDayjs}>
                  <DemoContainer components={["DateTimePicker"]}>
                    <DateTimePicker
                      defaultValue={dayjs(japanTime).add(7, "day")}
                      label="expiration date picker"
                      disablePast
                      onChange={(date) => {
                        const origin = date!.toString();
                        console.log("saasdasd" + origin);
                        setExpiration((expiration) => origin);
                      }}
                    />
                  </DemoContainer>
                </LocalizationProvider>
              </div>
            </div>
            <div className="box3">
              <p className="item-description">
                ↓Click on the specific date(s) you want to propose.
              </p>
              <LocalizationProvider dateAdapter={AdapterDayjs}>
                <DateCalendar
                  defaultValue={dayjs(japanTime)}
                  sx={{ overflow: "visible" }}
                  disablePast
                  onChange={(date) => {
                    //set to asia/tokyo timezone
                    const origin = date!.add(9, "hour").toString();
                    let res = `${origin.slice(8, 11)} ${origin.slice(
                      5,
                      7
                    )}(${origin.slice(0, 3)}) ${origin.slice(17, 22)}～`;
                    setDateList((dateList) => {
                      if (dateList) dateList += `\n`;
                      dateList += `${res}`;
                      return dateList;
                    });
                  }}
                />
              </LocalizationProvider>
            </div>
          </div>
          <Button
            sx={{
              width: 700,
              position: "absolute",
              transform: "translate(-50%, -50%)",
              marginTop: 5,
              left: "50%",
              height: 50,
            }}
            variant="contained"
            size="large"
            type="submit"
          >
            <FlagIcon />
            Create An Event Page!
          </Button>
        </form>
      </Box>
    </>
  );
}
