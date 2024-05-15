import React, { useContext, useRef, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import {
  Button,
  Link,
  TextField,
  ToggleButton,
  ToggleButtonGroup,
} from "@mui/material";
import {
  DataGrid,
  GridColDef,
  GridColumnHeaderParams,
  GridRenderCellParams,
} from "@mui/x-data-grid";
import noIcon from "../images/no.png";
import CheckIcon from "@mui/icons-material/Check";
import QuestionMarkIcon from "@mui/icons-material/QuestionMark";
import ClearIcon from "@mui/icons-material/Clear";
// import axios from "../utils/axios";
import "./DateProposalGrid.css";
import axios from "../utils/axios";

import { event, addAttendence, nameId } from "../types/Event";
import { Controller, SubmitHandler, useForm } from "react-hook-form";
import dayjs from "dayjs";
import { HistoryEventContext } from "../contexts/HistoryEvent";
import { isValidEmail } from "../utils/Validation";

interface rowData {
  id: number;
  Schedule: string;
  yes: number | string;
  unknown: number | string;
  no: number | string;
  annotation: number;
}

export default function DateProposalGrid(props: any) {
  const japanTime = dayjs();
  const [rows, setRows] = useState<rowData[]>([]);
  const [columns, setColumns] = useState<GridColDef[]>([]);
  const [showAddForm, setShowAddForm] = useState(false);
  const [idList, setIdList] = useState<number[]>([]);
  const [nameList, setNameList] = useState<nameId>({});
  const [schedule, setSchedule] = useState([0]);
  const navigate = useNavigate();
  const params = useParams();
  const [expiration, setExpiration] = useState("");
  const { historyEvent, setHistoryEvent } = useContext(HistoryEventContext);
  let formMethods = useForm<addAttendence>({
    mode: "onChange",
    defaultValues: { result: [] },
  });
  const {
    register,
    handleSubmit,
    control,
    reset,
    formState: { errors },
  } = formMethods;
  const targetRef = useRef<HTMLDivElement | null>(
    document.querySelector(".add-header")
  );

  React.useEffect(() => {
    axios
      .get(`/attendance/${props.uuid}`)
      .then((response) => {
        console.log(response.data);
        setExpiration(response.data.due_edit);
        setRows(generateRows(response.data));
        setColumns(generateColumns(response.data));
        setSchedule(Array(response.data.scheduleList.length).fill(undefined));
      })
      .catch((reason) => {
        console.log(reason);
        console.log("ERROR connecting backend service");
      });
  }, []);
  const getAvailability = (result: number[], idList: number[]) => {
    let availability = Object.fromEntries(
      idList.map((key, index) => [key, result[index]])
    );
    return availability;
  };
  const onSubmit: SubmitHandler<addAttendence> = (data) => {
    if (Object.keys(nameList).includes(data.name)) {
      axios
        .put(`/attendance/${props.uuid}`, {
          name: data.name,
          availability: getAvailability(data.result, idList),
          comment: data.comment,
          user_id: nameList[data.name],
        })
        .then(function (response) {
          navigate(`/view_event/${params.eventId}`);
          window.location.reload();
        })
        .catch(function (response) {
          console.log("ERROR connecting backend service");
        });
    } else {
      axios
        .post(`/attendance/${props.uuid}`, {
          name: data.name,
          availability: getAvailability(data.result, idList),
          comment: data.comment,
          email: data.email,
        })
        .then(function (response) {
          navigate(`/view_event/${params.eventId}`);
          window.location.reload();
        })
        .catch(function (response) {
          console.log("ERROR connecting backend service");
        });
    }
  };
  const generateRows = (eventObject: event) => {
    let rows: rowData[] = [];
    eventObject.scheduleList.forEach((schedule, index) => {
      setIdList((idList) => {
        if (idList.includes(schedule.id)) return [...idList];
        else return [...idList, schedule.id];
      });

      let [yesNum, unknownNum, noNum] = [0, 0, 0];
      eventObject.participants.forEach((obj) => {
        if (obj.result[index] === 1) {
          noNum += 1;
        } else if (obj.result[index] === 2) {
          unknownNum += 1;
        } else if (obj.result[index] === 3) {
          yesNum += 1;
        }
      });
      rows.push({
        id: index + 1,
        Schedule: schedule.name,
        yes: yesNum,
        unknown: unknownNum,
        no: noNum,
        annotation: schedule.annotation,
      });
    });
    rows.push({
      id: eventObject.scheduleList.length + 1,
      Schedule: "Comment",
      yes: "",
      unknown: "",
      no: "",
      annotation: 0,
    });
    return rows;
  };
  const generateColumns = (eventObject: event) => {
    let columns: GridColDef[] = [
      {
        field: "Schedule",
        width: 100,
        headerAlign: "center",
        headerClassName: "dataForm-header",
        cellClassName: "dataForm-header",
        sortable: false,
      },
      {
        field: "yes",
        width: 100,
        headerName: "✔",
        headerAlign: "center",
        headerClassName: "dataForm-header",
        cellClassName: "dataForm-cell",
        sortable: false,
        renderHeader: (params: GridColumnHeaderParams) => <span>✔</span>,
      },
      {
        field: "unknown",
        headerName: "?",
        width: 100,
        headerAlign: "center",
        cellClassName: "dataForm-cell",
        headerClassName: "dataForm-header",
        sortable: false,
      },
      {
        field: "no",
        headerName: "X",
        width: 100,
        headerAlign: "center",
        cellClassName: "dataForm-cell",
        headerClassName: "dataForm-header",
        sortable: false,
      },
    ];
    const getObejctPosition = (n: number) => {
      if (n === 3) {
        return "0";
      } else if (n === 2) {
        return "52%";
      } else if (n === 1) {
        return "102%";
      }
    };

    eventObject.participants.forEach((obj) => {
      setNameList((nameList) => {
        if (Object.keys(nameList).includes(obj.name)) return { ...nameList };
        else return { ...nameList, [obj.name]: obj.user_id };
      });

      columns.push({
        field: obj.name,
        width: 100,
        headerAlign: "center",
        headerClassName: "dataForm-header",
        headerName: "",
        sortable: false,
        renderHeader: (params: GridColumnHeaderParams) => (
          <div style={{ display: "flex", justifyContent: "center" }}>
            <Link
              color={"#A52A2A"}
              onClick={() => {
                reset({
                  name: obj.name,
                  result: obj.result,
                  comment: obj.comment,
                  email:obj.email
                });
                setShowAddForm((showAddForm) => true);

                setTimeout(() => {
                  if (targetRef.current) {
                    targetRef.current.scrollIntoView({ behavior: "smooth" });
                  }
                }, 100);
              }}
              component="button"
            >
              {params.field}
            </Link>
          </div>
        ),
        renderCell: (params: GridRenderCellParams) => {
          if (params.id === obj.result.length + 1)
            return <span style={{ margin: "0 auto" }}>{obj.comment}</span>;
          else
            return (
              <img
                alt="noIcon"
                src={noIcon}
                style={{
                  height: 25,
                  width: 27,
                  objectPosition: getObejctPosition(
                    obj.result[(params.id as number) - 1]
                  ),
                  objectFit: "cover",
                  margin: "0 auto",
                }}
              />
            );
        },
      });
    });
    return columns;
  };

  return (
    <>
      {japanTime.isAfter(expiration) && (
        <p style={{ color: "red" }}>This event is over now.</p>
      )}
      <div
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
        }}
      >
        <DataGrid
          rows={rows}
          columns={columns}
          getRowId={(row) => row.id}
          getRowClassName={(params) => {
            if (params.row.annotation === 1) return "yellow-row";
            else if (params.row.annotation === 2) return "lightgreen-row";
            else if (params.row.annotation === 3) return "darkgreen-row";
            else return "";
          }}
          rowHeight={60}
          showCellVerticalBorder
          showColumnVerticalBorder
          disableColumnFilter
          disableColumnSelector
          disableColumnMenu
          sx={{
            maxWidth: 100 * columns.length,
            "& .dataForm-header": { backgroundColor: "#f1f1f1" },
            "& .dataForm-cell": {
              justifyContent: "center",
            },
            "& .MuiDataGrid-root": {
              whiteSpace: "normal",
              wordWrap: "break-word",
            },
            "& .MuiDataGrid-row:not(.MuiDataGrid-row--dynamicHeight)>.MuiDataGrid-cell":
              {
                whiteSpace: "normal",
                wordWrap: "break-word",
              },
            "& .MuiDataGrid-row:hover & .MuiDataGrid-row:active'": {
              backgroundColor: "initial", // 使用 'initial' 或其他颜色取消悬浮时的背景色
            },
          }}
        />
      </div>

      <div
        className="add-container"
        style={{
          display: "flex",
          flexDirection: "column",
          margin: "0 auto",
          width: "500px",
        }}
      >
        <p className="yellow-font">Yellow rows mean maximum number of '✔'</p>
        <p className="lightgreen-font">
          lightgreen rows mean minimum number of 'X'
        </p>
        <p className="darkgreen-font">darkgreen rows mean optimal option</p>
        {!showAddForm && (
          <Button
            className="btn-add"
            variant="contained"
            disabled={japanTime.isAfter(expiration)}
            sx={{
              width: 170,
              height: 170,
              borderRadius: "50%",
              margin: "0 auto",
              marginTop: "20px",
              fontSize: "22px",
              fontWeight: "600",
              textTransform: "none",
            }}
            onClick={() => {
              setShowAddForm((showAddForm) => true);
            }}
          >
            Add Attendance
          </Button>
        )}
        {showAddForm && (
          <>
            <div className="add-header" ref={targetRef}>
              Add attendance
            </div>
            <p className="event-detail">Name</p>
            <TextField
              size="small"
              fullWidth
              {...register("name", {
                required: "this field is required",
              })}
              error={"name" in errors}
              helperText={errors.name?.message}
            />
            <p className="event-detail">Email</p>
            <TextField
              size="small"
              fullWidth
              {...register("email", {
                required: "this field is required",
                validate: (value) =>
                  isValidEmail(value) || "Invalid email address.",
              })}
              error={"email" in errors}
              helperText={errors.email?.message}
            />
            <p className="event-detail">Schedule</p>
            {errors?.result && (
              <span className="error-message">
                Please select all the schedule!
              </span>
            )}
            {rows.slice(0, -1).map((obj, index) => (
              <div
                className={
                  errors?.result?.[index]
                    ? "schedule-row-error "
                    : "schedule-row"
                }
              >
                <span style={{ marginLeft: "20px" }}>{obj.Schedule}</span>
                <Controller
                  // defaultValue={schedu/e[index]}
                  control={control}
                  {...register(`result.${index}`, {
                    validate: (value) => {
                      if (!value) return "please select";
                      else return true;
                    },
                  })}
                  render={({
                    field: { onChange, value = schedule[index] },
                  }) => (
                    <ToggleButtonGroup
                      value={value}
                      exclusive
                      onChange={(e, newValue) => {
                        const updatedSchedule = [...schedule];
                        updatedSchedule[index] = newValue;
                        setSchedule(updatedSchedule);
                        onChange(newValue);
                      }}
                      aria-label="circular buttons"
                      sx={{ position: "absolute", right: 0 }}
                    >
                      <ToggleButton value={3} className="toggle-button">
                        <CheckIcon />
                      </ToggleButton>
                      <ToggleButton value={2} className="toggle-button">
                        <QuestionMarkIcon />
                      </ToggleButton>
                      <ToggleButton
                        value={1}
                        className="toggle-button"
                        sx={{
                          marginRight: "20px",
                        }}
                      >
                        <ClearIcon />
                      </ToggleButton>
                    </ToggleButtonGroup>
                  )}
                />
              </div>
            ))}
            <p className="event-detail">Comment</p>
            <TextField
              size="small"
              fullWidth
              {...register("comment", { required: "this field is required" })}
              error={"comment" in errors}
              helperText={errors.comment?.message}
            />
            <Button
              variant="contained"
              sx={{
                width: 170,
                height: 170,
                borderRadius: "50%",
                margin: "0 auto",
                marginTop: "20px",
                fontSize: "22px",
                fontWeight: "600",
                textTransform: "none",
              }}
              onClick={handleSubmit(onSubmit)}
            >
              Submit
            </Button>
          </>
        )}
      </div>
    </>
  );
}
