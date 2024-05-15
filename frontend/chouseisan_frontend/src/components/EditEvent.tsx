import React, { useState, FormEvent, useEffect } from "react";
import { useNavigate, useParams } from "react-router-dom";

import {
  TextField,
  Button,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableRow,
  IconButton,
  Link,
  List,
  ListItem,
  ListItemText,
  ListItemSecondaryAction,
  ButtonGroup,
  Typography,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";

import dayjs from "dayjs";
import "./InputForm.css";
import "./EditEvent.css";
import { DemoContainer } from "@mui/x-date-pickers/internals/demo";
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs";
import { LocalizationProvider } from "@mui/x-date-pickers/LocalizationProvider";
import { DateCalendar, DateTimePicker } from "@mui/x-date-pickers";
import axios from "../utils/axios";
import { timeslots } from "../types/Event";
import Nonexist from "./Nonexist";

export default function EditEvent() {
  const [title, setTitle] = useState("");
  const [detail, setDetail] = useState("");
  const [clicked, setClicked] = useState(false);
  const japanTime = dayjs();
  const [dateList, setDateList] = useState<timeslots>({});
  const [deletedDates, setDeletedDates] = useState<number[]>([]);
  const [newList, setNewList] = useState("");
  const navigate = useNavigate();
  const [isExisted, setIsExisted] = useState(true);
  const [expiration, setExpiration] = useState(
    dayjs(japanTime).add(7, "day").toString()
  );
  const leftCellStyle = {
    fontSize: "16px",
    fontWeight: "bold",
    borderRight: "1px solid #ddd",
    width: "210px",
    padding: "18px",
  };
  const rightCellStyle = {
    padding: "10px",
  };
  const params = useParams();
  const input =
    params.eventId?.slice(0, 8) +
    "-" +
    params.eventId?.slice(8, 12) +
    "-" +
    params.eventId?.slice(12, 16) +
    "-" +
    params.eventId?.slice(16, 20) +
    "-" +
    params.eventId?.slice(20, 32);
  useEffect(() => {
    axios
      .get(`/event/exist/${input}`)
      .then((response) => {
        if (response.data.message === "Event Not Found.") setIsExisted(false);
      })
      .catch((error) => {
        console.log(error);
        console.log("ERROR connecting backend service");
      });
    axios
      .get(`/event/timeslots/${input}`)
      .then((response) => {
        setDateList(response.data.timeslots);
        setTitle(response.data.title);
        setDetail(response.data.detail);
        //title, detail are not completed
      })
      .catch((error) => {
        console.log(error);
        console.log("ERROR connecting backend service");
      });
  }, []);

  const eventEdit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (deletedDates.length > 0) {
      axios
        .put(`/event/deleteTimeslots/${input}`, { timeslot_ids: deletedDates })
        .then((response) => {})
        .catch((error) => {
          console.log(error);
          console.log("ERROR connecting backend service");
        });
    }
    axios
      .put(`/event/addTimeslots/${input}`, { dateTimeProposal: newList })
      .then((response) => {})
      .catch((error) => {
        console.log(error);
        console.log("ERROR connecting backend service");
      });
    axios
      .put(`/event/editTitleDetail/${input}`, {
        title: title,
        detail: detail,
        due_edit: expiration,
      })
      .then((response) => {})
      .catch((error) => {
        console.log(error);
        console.log("ERROR connecting backend service");
      });
    navigate(`/view_event/${params.eventId}`);
  };
  const deleteEvent = () => {
    axios
      .delete(`/event/${input}`)
      .then((response) => {})
      .catch((error) => {
        console.log(error);
        console.log("ERROR connecting backend service");
      });
  };
  return isExisted ? (
    <>
      <p className="firstLink">
        <Link
          href={"/scheduler/view_event/" + params.eventId}
          color={"#a46702"}
          underline="hover"
          sx={{ marginBottom: "15px" }}
        >
          {title}
        </Link>
        {" > "}Edit/Delete Event
      </p>
      <div className="container1">
        <div className="event-header1">Edit/Delete Event</div>

        {/* <p style={{}}></p> */}
        <TableContainer
          component={Paper}
          sx={{
            backgroundColor: "rgba(255, 255, 255, 0.6)",
            border: "1px solid #ccc",
          }}
        >
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableBody>
              <TableRow>
                <TableCell style={leftCellStyle}>Event Title</TableCell>
                <TableCell style={rightCellStyle}>
                  <TextField
                    fullWidth
                    helperText="Team Dinner Party”, “Project Meeting”, etc..."
                    value={title}
                    required
                    onChange={(e) => setTitle(e.target.value)}
                  ></TextField>
                </TableCell>
              </TableRow>

              <TableRow>
                <TableCell style={leftCellStyle}>Event Details</TableCell>
                <TableCell style={rightCellStyle}>
                  <TextField
                    fullWidth
                    multiline
                    rows={3}
                    helperText="*Let’s schedule the party! Please respond by ___"
                    value={detail}
                    onChange={(e) => setDetail(e.target.value)}
                  ></TextField>
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell style={leftCellStyle}>Expiration date</TableCell>
                <TableCell style={rightCellStyle}>
                  <LocalizationProvider dateAdapter={AdapterDayjs}>
                    <DemoContainer components={["DateTimePicker"]}>
                      <DateTimePicker
                        defaultValue={dayjs(japanTime).add(7, "day")}
                        label="expiration date picker"
                        disablePast
                        onChange={(date) => {
                          const origin = date!.toString();
                          setExpiration((expiration) => origin);
                        }}
                      />
                    </DemoContainer>
                  </LocalizationProvider>
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell style={leftCellStyle}>Date Proposals</TableCell>
                <TableCell style={rightCellStyle}>
                  <h2>Delete proposed dates</h2>
                  <List>
                    {Object.entries(dateList).map((value, index) => (
                      <ListItem
                        key={index}
                        sx={{
                          maxWidth: 500,
                          border: "1px solid #ccc",
                          borderRadius: "8px",
                          margin: "8px 0",
                          display: "flex",
                          justifyContent: "space-between",
                        }}
                      >
                        <ListItemSecondaryAction>
                          <IconButton
                            onClick={() => {
                              const newObject = Object.assign({}, dateList);
                              delete newObject[Number(value[0])];
                              setDateList((dateList) => newObject);
                              setDeletedDates((deletedDates) => [
                                ...deletedDates,
                                Number(value[0]),
                              ]);
                            }}
                          >
                            <DeleteIcon />
                          </IconButton>
                        </ListItemSecondaryAction>
                        <ListItemText primary={value[1]} />
                      </ListItem>
                    ))}
                  </List>
                  <h2>Add proposed dates</h2>
                  <p>Please enter the new proposed dates.</p>
                  <div className="form">
                    <TextField
                      fullWidth
                      helperText=""
                      multiline
                      rows={15}
                      onChange={(
                        event: React.ChangeEvent<HTMLInputElement>
                      ) => {
                        setNewList(event.target.value);
                      }}
                      value={newList}
                      placeholder="Simply input your proposals in the Month DD(DAY) TIME format. Or you can click on the specific date(s) in the calendar."
                    ></TextField>

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
                          setNewList((newList) => {
                            if (newList) newList += `\n`;
                            newList += `${res}`;
                            return newList;
                          });
                        }}
                      />
                    </LocalizationProvider>
                  </div>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>

        <form onSubmit={eventEdit}>
          <ButtonGroup
            sx={{
              position: "absolute",
              transform: "translate(-50%, -50%)",
              marginTop: 5,
              marginBottom: 10,
              left: "50%",
              height: 50,
            }}
          >
            <Button
              sx={{
                marginRight: "58px",
                borderRadius: 0,
                width: 150,
              }}
              variant="contained"
            >
              Go Back
            </Button>
            <Button
              sx={{ borderRadius: 0, width: 300 }}
              variant="contained"
              type="submit"
            >
              Save Changes
            </Button>
          </ButtonGroup>
        </form>
        <div className="event-header2">Edit/Delete Event</div>
        <TableContainer
          component={Paper}
          sx={{
            backgroundColor: "rgba(255, 255, 255, 0.6)",
            border: "1px solid #ccc",
            marginBottom: 20,
          }}
        >
          <Table sx={{ minWidth: 650 }} aria-label="simple table">
            <TableBody>
              <TableRow>
                <TableCell
                  style={{
                    ...leftCellStyle,
                    backgroundColor: "red",
                    color: "white",
                  }}
                >
                  Cancel your event
                </TableCell>
                <TableCell
                  style={{
                    ...rightCellStyle,
                    flexDirection: "column",
                    display: "flex",
                  }}
                >
                  {clicked && (
                    <p className="warning">
                      Are you sure you want to delete this event?<br></br>Once
                      you delete an event, it cannot be recovered.<br></br>If
                      you are sure, please press the "Delete Event" button.
                    </p>
                  )}
                  <Button
                    sx={{
                      width: 400,
                      marginTop: 1,
                      marginBottom: 3,
                      border: clicked ? "1px solid red" : "1px solid", // 根据点击状态设置边框颜色
                      "&:hover": {
                        border: clicked ? "1px solid red" : "1px solid", // 根据点击状态设置悬停时的边框颜色
                      },
                    }}
                    onClick={() => {
                      setClicked(true);
                      if (clicked) {
                        deleteEvent();
                      }
                    }}
                  >
                    Delete event
                  </Button>
                  <Typography variant="caption">
                    *Event pages cannot be restored once deleted.
                  </Typography>
                </TableCell>
              </TableRow>
            </TableBody>
          </Table>
        </TableContainer>
      </div>
    </>
  ) : (
    <Nonexist />
  );
}
