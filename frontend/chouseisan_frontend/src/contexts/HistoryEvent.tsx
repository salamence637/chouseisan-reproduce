import React, { createContext, useState, useEffect } from "react";
import { historyEvent } from "../types/Event";
export const HistoryEventContext = createContext(
  {} as {
    historyEvent: historyEvent[];
    setHistoryEvent: React.Dispatch<React.SetStateAction<historyEvent[]>>;
  }
);
const HistoryEventProvider = ({ children }: { children: any }) => {
  let str = localStorage.getItem("historyState");
  if (!str) str = "[]";
  const initialState = JSON.parse(str);
  const [historyEvent, setHistoryEvent] =
    useState<historyEvent[]>(initialState);
  useEffect(() => {
    localStorage.setItem("historyState", JSON.stringify(historyEvent));
  }, [historyEvent]);
  return (
    <HistoryEventContext.Provider value={{ historyEvent, setHistoryEvent }}>
      {children}
    </HistoryEventContext.Provider>
  );
};
export default HistoryEventProvider;
