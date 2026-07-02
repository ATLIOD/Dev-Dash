import { DialogContent, MenuItem, Select, TextField, type SelectChangeEvent } from "@mui/material";
import { DialogBase } from "../../components/Modals/Dialogs";
import { Status, type Project } from "../../Managers/ProjectManager";
import { useState, type ChangeEvent } from "react";

export const ProjectDetailsDialog = ({
  open,
  onClose,
  onConfirm,
  project,
}: {
  open: boolean;
  onClose: () => void;
  onConfirm: (value: Project) => void;
  project?: Project;
}) => {
  const title = !project ? "New" : "Edit";

  const [name, setName] = useState<string>(project?.name ?? "");
  const [description, setDescription] = useState<string>(project?.description ?? "");
  const [stack, setStack] = useState<string[]>(project?.stack ?? []);
  const [repo, setRepo] = useState<string>(project?.repo ?? "");
  const [deploy, setDeploy] = useState<string>(project?.deploy ?? "");
  const [status, setStatus] = useState<Status>(project?.status ?? Status.Planning);

  const [nameError, setNameError] = useState<string>("");
  const isValid = name.trim().length > 0;

  function HandleDeployChange(e: ChangeEvent<HTMLInputElement>) {
    setDeploy(e.target.value);
  }
  function HandleDescriptionChange(e: ChangeEvent<HTMLInputElement>) {
    setDescription(e.target.value);
  }

  function HandleNameChange(e: ChangeEvent<HTMLInputElement>) {
    const value = e.target.value;
    const error = value.trim() ? "" : "Project name is required.";

    setNameError(error);
    setName(value);
  }

  function HandleRepoChange(e: ChangeEvent<HTMLInputElement>) {
    setRepo(e.target.value);
  }

  function HandleStatusChange(e: SelectChangeEvent<Status>) {
    setStatus(e.target.value);
  }

  function HandleStackChange() {
    setStack([]);
  }

  function HandleOnConfirm() {
    if (!isValid) {
      return;
    }

    const result: Project = {
      uuid: project?.uuid ?? crypto.randomUUID(),
      name,
      description,
      status,
      stack,
      repo,
      deploy,
    };

    onConfirm(result);
  }

  return (
    <DialogBase
      isValid={isValid}
      open={open}
      onClose={onClose}
      onConfirm={HandleOnConfirm}
      title={`${title} Project`}
      confirmText={project ? "Save" : "Create"}
    >
      <DialogContent className="project-details">
        <div className="details-column">
          Name*
          <TextField
            value={name}
            onChange={HandleNameChange}
            required
            error={!!nameError}
            helperText={nameError}
          />
        </div>

        <div className="details-column">
          Description
          <TextField multiline rows={4} value={description} onChange={HandleDescriptionChange} />
        </div>

        <div className="details-row">
          <div className="details-column">
            Status
            <Select value={status} onChange={HandleStatusChange}>
              <MenuItem value={Status.Planning}>Planning</MenuItem>
              <MenuItem value={Status.Active}>Active</MenuItem>
              <MenuItem value={Status.Maintaining}>Maintaining</MenuItem>
              <MenuItem value={Status.Complete}>Complete</MenuItem>
            </Select>
          </div>
          <div className="details-column">
            Stack
            <TextField value={stack} onChange={HandleStackChange} />
          </div>
        </div>

        <div className="details-column">
          Repository URL
          <TextField value={repo} onChange={HandleRepoChange} />
        </div>

        <div className="details-column">
          Deployment URL
          <TextField value={deploy} onChange={HandleDeployChange} />
        </div>
      </DialogContent>
    </DialogBase>
  );
};
