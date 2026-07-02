import { IconButton } from "@mui/material";
import { type Project } from "../../Managers/ProjectManager";
import "./__dashboard.scss";
import { FolderIcon, Size } from "../../components/Icons";

export const PinnedPanel = ({ projectList }: { projectList: Project[] }) => {
  return (
    <div className="pinned-projects">
      <span>Pinned Projects</span>
      <div className="project-list">
        {projectList.map((x, i) => (
          <ProjectSummaryTile project={x} key={i} />
        ))}
      </div>
    </div>
  );
};

const ProjectSummaryTile = ({ project }: { project: Project }) => {
  // TO-DO: Navigate to project page.
  return (
    <div className="project-tile">
      <IconButton
        style={{
          backgroundColor: "transparent",
        }}
      >
        <FolderIcon size={Size.xlarge} />
      </IconButton>
      <div className="project-overview">
        <span
          style={{
            fontSize: "var(--font-size-medium)",
          }}
        >
          {project.name}
        </span>
        <span
          style={{
            fontSize: "var(--font-size-small)",
            color: "var(--secondary-text)",
          }}
        >
          {project.status}
        </span>
      </div>
    </div>
  );
};
