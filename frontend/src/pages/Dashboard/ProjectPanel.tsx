import { useState } from "react";
import { IconButton } from "../../components/Buttons/Buttons";
import { Chip, ChipType } from "../../components/Chips";
import { AddIcon, VertKebab } from "../../components/Icons";
import { Status, type Project } from "../../Managers/Project";
import { Menu, MenuItem } from "../../components/Menu";

export const ProjectPanel = ({ projectList }: { projectList: Project[] }) => {
  return (
    <div className="all-projects">
      <div className="panel-heading">
        <span>All Projects</span>
        <IconButton
          baseClass="primary"
          icon={<AddIcon color="#fff" />}
          style={{ fontSize: "var(--font-size-medium)", paddingRight: "var(--font-size-small)" }}
        >
          New Project
        </IconButton>
      </div>

      <div className="project-list-panel">
        {projectList.map((x) => (
          <ProjectTile project={x} />
        ))}
      </div>
    </div>
  );
};

const ProjectTile = ({ project }: { project: Project }) => {
  const status = project.status;
  const chipType: ChipType =
    status === Status.Active
      ? ChipType.Active
      : status === Status.Planning
        ? ChipType.Planning
        : status === Status.Maintaining
          ? ChipType.Maintaining
          : status === Status.Complete
            ? ChipType.Primary
            : ChipType.Secondary;

  const [anchorEl, setAnchorEl] = useState<HTMLElement | null>(undefined);

  function HandleOnMenuClick(e: React.MouseEvent<HTMLButtonElement>) {
    return setAnchorEl(e.currentTarget);
  }

  function HandleOnDelete() {}
  function HandleOnEdit() {}
  function HandleOnPin() {}

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
          <Chip baseClass={chipType}>{project.status}</Chip>
          <IconButton
            baseClass="text"
            style={{ backgroundColor: "transparent" }}
            icon={<VertKebab color="var(--primary-text)" />}
            onClick={HandleOnMenuClick}
          />
        </div>
      </div>

      <Menu anchorEl={anchorEl} open={Boolean(anchorEl)} onClose={() => setAnchorEl(null)}>
        <MenuItem onClick={HandleOnPin}>Pin Project</MenuItem>
        <MenuItem onClick={HandleOnEdit}>Edit Project</MenuItem>
        <MenuItem onClick={HandleOnDelete}>Delete Project</MenuItem>
      </Menu>

      {/* Update with actual project details. */}
      <div className="project-details">
        This is a project description.
        <div className="project-chips">
          <Chip>React</Chip>
          <Chip>Go</Chip>
        </div>
      </div>
    </div>
  );
};
