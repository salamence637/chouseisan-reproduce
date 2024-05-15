import { Route, Routes } from "react-router-dom";
import InputForm from "../components/InputForm";
import CreateComplete from "../components/CreateComplete";
import ViewEvent from "../components/ViewEvent";
import History from "../components/History";

import EditEvent from "../components/EditEvent";

export default function RouteSetting() {
  return (
    <>
      <Routes>
        <Route path="/" element={<InputForm />} />
        <Route path="/create_complete">
          <Route path=":eventId" element={<CreateComplete />} />
        </Route>
        <Route path="/view_event">
          <Route path=":eventId" element={<ViewEvent />} />
        </Route>
        <Route path="/history" element={<History />} />
        <Route path="/edit">
          <Route path=":eventId" element={<EditEvent />} />
        </Route>
      </Routes>
    </>
  );
}
