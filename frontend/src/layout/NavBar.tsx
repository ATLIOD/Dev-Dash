import { IconButton } from "../components/Buttons/Buttons";
import { CodeIcon, DashboardIcon, SearchIcon, Size, TerminalIcon } from "../components/Icons";
import { IconTextField } from "../components/Inputs/Inputs";
import "./layout.css";
import { UserProfile } from "./UserProfile";

export const NavBar = () => {
  return (
    <header
      style={{
        display: "flex",
        height: "var(--navbar-height)",
        borderBottom: "1px solid var(--border)",
        alignItems: "center",
        margin: "var(--spacing-xsmall)",
      }}
    >
      <div
        style={{
          display: "flex",
          alignItems: "center",
          justifyContent: "space-between",
          width: "100%",
        }}
      >
        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "var(--spacing-medium)",
          }}
        >
          <IconButton baseClass="text" icon={<TerminalIcon size={Size.xlarge} />}>
            <span style={{ fontSize: "var(--font-size-large)" }}>Dev Dash</span>
          </IconButton>
          <IconButton baseClass="text" icon={<DashboardIcon size={Size.medium} />}>
            Dashboard
          </IconButton>
          <IconButton baseClass="text" icon={<CodeIcon size={Size.medium} />}>
            Snippets
          </IconButton>
        </div>

        <div
          style={{
            display: "flex",
            alignItems: "center",
            gap: "var(--spacing-medium)",
          }}
        >
          {/* TODO: Move search bar to a separate component. */}
          <IconTextField icon={<SearchIcon size={16} />} placeholder="Search..." />
          <UserProfile />
        </div>
      </div>
    </header>
  );
};
