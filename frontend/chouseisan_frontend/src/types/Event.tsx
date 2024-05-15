export type proposal = {
  name: string;
  comment: string;
  result: number[];
  user_id: number
  email: string
};
export type event = {
  scheduleList: schedule[];
  participants: proposal[];
};
export type nameId={
  [name:string]: number
}
export type timeslots={[eventId:number]:string}
export type schedule={
  name:string;
  id: number;
  annotation:number
}
export type addAttendence = {
  name: string;
  email:string
  comment: string;
  result: number[];
};
export type historyEvent = {
  title: string
  scheduleList: string[]
  uuid: string
}