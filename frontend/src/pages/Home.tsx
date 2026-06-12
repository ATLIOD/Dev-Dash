import { Status, type Project } from "../Managers/Project.ts";
import "./_pages.scss";
import { PinnedPanel } from "./Dashboard/PinnedPanel.tsx";
import { ProjectPanel } from "./Dashboard/ProjectPanel.tsx";

export const Home = () => {
  const projectList: Project[] = [
    { name: "Project 1", status: Status.Active },
    { name: "Project 2", status: Status.Planning },
    { name: "Project 3", status: Status.Maintaining },
    { name: "Project 4", status: Status.Complete },
    { name: "Project 5", status: Status.Active },
  ];

  return (
    <div style={{ padding: "var(--spacing-large)", gap: "var(--spacing-large);" }}>
      <PinnedPanel projectList={projectList} />
      <ProjectPanel projectList={projectList} />
    </div>
  );
};
