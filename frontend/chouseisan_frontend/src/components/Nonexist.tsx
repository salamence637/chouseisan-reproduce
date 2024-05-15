import { Button } from "@mui/material";
import { useState } from "react";

export default function Nonexist() {
  const [number, setNumber] = useState(0);
  const [str, setStr] = useState("");
  return (
    <div className="container">
      <h2 className="form-header">We are sorry we couldn't locate that page</h2>
      <h4 className="description">
        The event does not exist or has expired/been deleted.
      </h4>
      <h1>yang{Number(number)}</h1>
      <Button
        variant="contained"
        onClick={() => {
          setNumber((number) => number + 1);
          setStr((str) => str + "1");
        }}
        size="large"
      >
        {number}
      </Button>
    </div>
  );
}
