import React, { useState, useEffect } from "react";
import { Link, useParams } from "react-router-dom";
import { TextField, Button } from "@mui/material";
import "./CreateComplete.css";
import axios from "../utils/axios";
import Nonexist from "./Nonexist";

export default function CreateComplete() {
  const params = useParams();
  const [isExisted, setIsExisted] = useState(true);
  const textUrl =
    "http://localhost:3000/scheduler/view_event/" + params.eventId;
  const [url, setUrl] = useState<string | undefined>(textUrl);
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
  });

  return (
    <>
      {isExisted ? (
        <div className="container">
          <h2 className="form-header">
            Your event page is ready to be shared!
          </h2>
          <h4 className="description">
            Your event page is created! You can start inviting people by sharing
            the URL below! Using the URL, your peers can submit when they are
            available to meet.
          </h4>
          <TextField
            size="small"
            fullWidth
            // defaultValue={"https://chouseisan.com"}
            value={url}
            onChange={(e) => {
              setUrl(e.target.value);
            }}
          />
          <Button
            size="large"
            variant="contained"
            component={Link}
            to={textUrl}
            sx={{
              width: 300,
              height: 50,
              marginTop: 10,
              left: "50%",
              position: "absolute",
              transform: "translate(-50%, -50%)",
              borderRadius: 3,
            }}
          >
            Go to your event page
          </Button>
        </div>
      ) : (
        <Nonexist />
      )}
    </>
  );
}
