import React, { useState, useEffect, useContext } from "react";
import { Box, Button, Link } from "@mui/material";
import "./ViewEvent.css";
import DateProposalGrid from "./DateProposalGrid";
import { useParams } from "react-router-dom";
import axios from "../utils/axios";
import Nonexist from "./Nonexist";
import { SelfEventContext } from "../contexts/EventBySelf";
import { HistoryEventContext } from "../contexts/HistoryEvent";

export default function ViewEvent() {
  const { selfEventList, setSelfEventList } = useContext(SelfEventContext);
  const { historyEvent, setHistoryEvent } = useContext(HistoryEventContext);
  const [title, setTitle] = useState("");
  const [detail, setDetail] = useState("");
  const [no, setNo] = useState(0);
  const params = useParams();
  const [isExisted, setIsExisted] = useState(false);

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
        if (response.data.message === "Event Found.") {
          setIsExisted(true);
        }
      })
      .catch((error) => {
        console.log(error);
        console.log("ERROR connecting backend service");
      });
  });
  React.useEffect(() => {
    if (isExisted) {
      // setHistoryEvent((historyEvent) => {
      //   if (historyEvent.includes(input)) return historyEvent;
      //   else if (historyEvent.length >= 5)
      //     return [input, ...historyEvent.slice(0, -1)];
      //   else return [input, ...historyEvent];
      // });
      axios
        .get(`/event/timeslots/${input}`)
        .then((response) => {
          const obj = {
            ["scheduleList"]: Object.values(
              response.data.timeslots
            ) as string[],
            ["title"]: response.data.title as string,
            ["uuid"]: input,
          };
          const flag = Object.values(historyEvent).some(
            (value) => value.uuid === input
          );

          setHistoryEvent((historyEvent) => {
            if (flag) return historyEvent;
            else if (historyEvent.length >= 5)
              return [obj, ...historyEvent.slice(0, -1)];
            else return [obj, ...historyEvent];

            // if (historyEvent.includes(input)) return historyEvent;
            // else if (historyEvent.length >= 5)
            //   return [input, ...historyEvent.slice(0, -1)];
            // else return [input, ...historyEvent];
          });
        })
        .catch((reason) => {
          console.log(reason);
          console.log("ERROR connecting backend service");
        });
      axios
        .get(`/event/basic/${input}`)
        .then((response) => {
          setTitle(response.data.title);
          setNo(response.data.num_users);
          setDetail(response.data.detail);
        })
        .catch((reason) => {
          console.log(reason);
          console.log("ERROR connecting backend service");
        });
    }
  }, [isExisted]);

  return isExisted ? (
    <>
      <div className="container1">
        <Link
          href="/scheduler"
          color={"#a46702"}
          underline="hover"
          sx={{ marginBottom: "15px" }}
        >
          Host a new event
        </Link>
        <div className="event-header">
          {title}
          {selfEventList.includes(params.eventId as string) && (
            <Button
              variant="outlined"
              // href="/edit"
              href={"/scheduler/edit/" + params.eventId}
              sx={{
                position: "absolute",
                right: 0,
                fontSize: 15,
              }}
            >
              Edit this event
            </Button>
          )}
        </div>
        <p className="event-info">
          <span className="no-label">No. of respondents</span>
          {no}
          <span style={{ marginLeft: "30px", fontSize: 14 }}>
            You are the event organizer
          </span>
        </p>
        <div style={{ minHeight: "150px" }}>
          {detail !== "" && (
            <>
              <p className="event-detail">Event Details</p>
              <p>{detail}</p>
            </>
          )}
        </div>
      </div>
      <Box sx={{ background: "#f9f9f9" }}>
        <div className="container2">
          <p className="event-detail">Date Proposals</p>
          <p>Click on the name to edit your response.</p>
          <DateProposalGrid uuid={input} />
        </div>
      </Box>
    </>
  ) : (
    <Nonexist />
  );
}
