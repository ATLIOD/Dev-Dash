export type Project = {
  name: string;
  status: Status;
};

export enum Status {
  Planning = "Planning",
  Active = "Active",
  Maintaining = "Maintaining",
  Complete = "Complete",
}
