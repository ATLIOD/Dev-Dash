import { IconButton } from "../components/Buttons/Buttons";
import { CodeIcon, DashboardIcon, SearchIcon, Size, TerminalIcon } from "../components/Icons";
import { IconTextField } from "../components/Inputs/Inputs";
import "./_layout.scss";
import { UserProfile } from "./UserProfile";

export const NavBar = () => {
  return (
    <header>
      <div className="header-wrapper">
        <div className="nav-buttons">
          <IconButton
            baseClass="text"
            icon={<TerminalIcon size={Size.xlarge} />}
            style={{ fontSize: "var(--font-size-large)" }}
          >
            Dev Dash
          </IconButton>
          <IconButton baseClass="text" icon={<DashboardIcon />}>
            Dashboard
          </IconButton>
          <IconButton baseClass="text" icon={<CodeIcon />}>
            Snippets
          </IconButton>
        </div>

        <div className="nav-buttons">
          {/* TODO: Move search bar to a separate component. */}
          <IconTextField icon={<SearchIcon size={Size.small} />} placeholder="Search..." />
          <UserProfile />
        </div>
      </div>
    </header>
  );
};
