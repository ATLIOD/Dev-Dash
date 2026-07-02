import { useState } from "react";
import { Status, type Project } from "../Managers/ProjectManager.ts";
import "./_pages.scss";
import { PinnedPanel } from "./Dashboard/PinnedPanel.tsx";
import { ProjectPanel } from "./Dashboard/ProjectPanel.tsx";

export const Home = () => {
  // TODO: Move all this state logic into the ProjectManager.
  const [projectList, setProjectList] = useState<Project[]>([
    {
      uuid: crypto.randomUUID(),
      name: "Project 1",
      description: "Project description",
      status: Status.Active,
      stack: ["React", "Go"],
      repo: "www.repo.com",
      deploy: "www.deploy.com",
    },
    { uuid: crypto.randomUUID(), name: "Project 2", status: Status.Planning },
    { uuid: crypto.randomUUID(), name: "Project 3", status: Status.Maintaining },
    { uuid: crypto.randomUUID(), name: "Project 4", status: Status.Complete },
    { uuid: crypto.randomUUID(), name: "Project 5", status: Status.Active },
  ]);

  function HandleOnCreate(value: Project) {
    setProjectList([...projectList, value]);
  }

  function HandleOnDelete(value: Project) {
    setProjectList(projectList.filter((x) => x.uuid !== value.uuid));
  }

  function HandleOnEdit(value: Project) {
    let updated = projectList.filter((x) => x.uuid !== value.uuid);
    setProjectList([...updated, value]);
  }

  return (
    <div className="project-list-wrapper">
      <PinnedPanel projectList={projectList} />
      <ProjectPanel
        projectList={projectList}
        onCreate={HandleOnCreate}
        onDelete={HandleOnDelete}
        onEdit={HandleOnEdit}
      />
    </div>
  );
};
