import { useState } from "react";
import { Status, type Project } from "../../Managers/ProjectManager";
import { Button, Menu, MenuItem } from "@mui/material";
import "./__dashboard.scss";
import { AddOutlined } from "@mui/icons-material";
import { MenuButton } from "../../components/Buttons/Buttons";
import { ChipBase, ChipType } from "../../components/Chips";
import { ProjectDetailsDialog } from "./ProjectDetailsDialog";

export const ProjectPanel = ({
  projectList,
  onCreate,
  onDelete,
  onEdit,
}: {
  projectList: Project[];
  onCreate: (p: Project) => void;
  onDelete: (p: Project) => void;
  onEdit: (p: Project) => void;
}) => {
  const [open, setOpen] = useState<boolean>(false);

  function HandleOnClose() {
    setOpen(!open);
  }

  function HandleOnCreate(p: Project) {
    onCreate(p);
    HandleOnClose();
  }

  return (
    <div className="all-projects">
      <div className="panel-heading">
        <span>All Projects</span>
        <Button
          onClick={() => setOpen(!open)}
          color="primary"
          startIcon={<AddOutlined />}
          variant="contained"
        >
          New Project
        </Button>
      </div>

      <div className="project-list-panel">
        {projectList.map((x, i) => (
          <ProjectTile onDelete={onDelete} onEdit={onEdit} project={x} key={i} />
        ))}
      </div>

      <ProjectDetailsDialog
        open={open}
        onClose={HandleOnClose}
        onConfirm={HandleOnCreate}
        key={crypto.randomUUID()}
      />
    </div>
  );
};

const ProjectTile = ({
  project,
  onDelete,
  onEdit,
}: {
  project: Project;
  onDelete: (p: Project) => void;
  onEdit: (p: Project) => void;
}) => {
  const chipType: ChipType = GetChipTypeFromStatus(project.status);

  const [anchorEl, setAnchorEl] = useState<HTMLElement | undefined>(undefined);
  const [open, setOpen] = useState<boolean>(false);

  function EditProject(project: Project) {
    onEdit(project);
    ToggleDialog();
  }

  function HandleOnMenuClick(e: React.MouseEvent<HTMLButtonElement>) {
    return setAnchorEl(e.currentTarget);
  }

  function HandleOnDelete() {
    onDelete(project);
    setAnchorEl(undefined);
  }
  function HandleOnPin() {
    setAnchorEl(undefined);
  }

  function ToggleDialog() {
    setAnchorEl(undefined);
    setOpen(!open);
  }

  // TO-DO: Navigate to project page.
  return (
    <div className="project-overview">
      <div className="project-header">
        <span
          style={{
            fontSize: "var(--font-size-medium)",
          }}
        >
          {project.name}
        </span>
        <div className="project-actions">
          <ChipBase label={project.status} color={chipType} />
          <MenuButton onClick={HandleOnMenuClick} />
        </div>
      </div>

      <Menu anchorEl={anchorEl} open={Boolean(anchorEl)} onClose={() => setAnchorEl(undefined)}>
        <MenuItem onClick={HandleOnPin}>Pin Project</MenuItem>
        <MenuItem onClick={ToggleDialog}>Edit Project</MenuItem>
        <MenuItem onClick={HandleOnDelete}>Delete Project</MenuItem>
      </Menu>

      {/* Update with actual project details. */}
      <div className="project-details">
        {project.description}
        <div className="project-chips">
          {project?.stack?.map((x) => (
            <ChipBase label={x} />
          ))}
        </div>
      </div>

      <ProjectDetailsDialog
        key={project.uuid}
        open={open}
        onClose={ToggleDialog}
        project={project}
        onConfirm={EditProject}
      />
    </div>
  );
};

function GetChipTypeFromStatus(status: Status) {
  switch (status) {
    case Status.Active:
      return ChipType.Active;
    case Status.Planning:
      return ChipType.Planning;
    case Status.Complete:
      return ChipType.Primary;
    case Status.Maintaining:
      return ChipType.Maintaining;
  }
}
