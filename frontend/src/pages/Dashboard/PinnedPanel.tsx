import { IconButton } from "../../components/Buttons/Buttons";
import { FolderIcon, Size } from "../../components/Icons";
import { type Project } from "../../Managers/Project";
import "./__dashboard.scss";

export const PinnedPanel = ({ projectList }: { projectList: Project[] }) => {
  return (
    <div className="pinned-projects">
      <span>Pinned Projects</span>
      <div className="project-list">
        {projectList.map((x) => (
          <ProjectSummaryTile project={x} />
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
        icon={<FolderIcon size={Size.xlarge} />}
        style={{
          backgroundColor: "var(--surface-color)",
        }}
      />
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
