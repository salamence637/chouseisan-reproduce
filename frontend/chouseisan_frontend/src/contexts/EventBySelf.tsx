import React, { createContext, useState, useEffect } from "react";
export const SelfEventContext = createContext(
  {} as {
    selfEventList: string[];
    setSelfEventList: React.Dispatch<React.SetStateAction<string[]>>;
  }
);
const SelfEventListProvider = ({ children }: { children: any }) => {
  const initialState = JSON.parse(
    localStorage.getItem("myState") || "[]"
  ) as string[];

  const [selfEventList, setSelfEventList] = useState<string[]>(initialState);
  useEffect(() => {
    localStorage.setItem("myState", JSON.stringify(selfEventList));
  }, [selfEventList]);
  return (
    <SelfEventContext.Provider value={{ selfEventList, setSelfEventList }}>
      {children}
    </SelfEventContext.Provider>
  );
};
export default SelfEventListProvider;
