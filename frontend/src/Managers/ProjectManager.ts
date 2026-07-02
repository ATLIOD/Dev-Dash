// TODO: Turn this into a project manager to provide project state to application.
export type Project = {
  name: string;
  uuid: string;
  description?: string;
  status: Status;
  stack?: string[];
  repo?: string;
  deploy?: string;
};

export enum Status {
  Planning = "Planning",
  Active = "Active",
  Maintaining = "Maintaining",
  Complete = "Complete",
}
