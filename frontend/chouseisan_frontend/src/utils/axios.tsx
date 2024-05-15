import axios from "axios";

export default axios.create({
  baseURL: "http://localhost:8080/",
  withCredentials: true,
  validateStatus: function (status) {
    return status >= 200 && status <= 302; // default
  },
});

export const authAxios = axios.create({
  baseURL: process.env.REACT_APP_API_URL,
  headers: { "Content-Type": "application/json" },
});
