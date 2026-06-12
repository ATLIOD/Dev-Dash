import { IconButton } from "../../components/Buttons/Buttons";
import { Chip } from "../../components/Chips";
import { AddIcon, VertKebab } from "../../components/Icons";
import type { Project } from "../../Managers/Project";

export const ProjectPanel = ({ projectList }: { projectList: Project[] }) => {
  return (
    <div className="all-projects">
      <div className="panel-heading">
        <span>All Projects</span>
        <IconButton
          baseClass="primary"
          icon={<AddIcon color="var(--bg)" />}
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
          <Chip>{project.status}</Chip>
          <IconButton
            baseClass="text"
            style={{ backgroundColor: "var(--surface-color)" }}
            icon={<VertKebab color="var(--primary-text)" />}
          />
        </div>
      </div>

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
