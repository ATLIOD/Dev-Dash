import { IconButton } from "../components/Buttons";
import { CodeIcon, DashboardIcon, SearchIcon, TerminalIcon } from "../components/Icons";
import { IconTextField } from "../components/Inputs";
import "./layout.css";

export const NavBar = () => {
  return (
    <header
      style={{
        display: "flex",
        alignItems: "center",
        height: "var(--navbar-height)",
        borderBottom: "1px solid var(--border)",
      }}
    >
      <div
        style={{
          margin: "var(--spacing-small) var(--spacing-medium)",
          gap: "var(--spacing-small)",
        }}
      >
        {/* TODO: onclick routes to given page. */}
        <IconButton icon={<TerminalIcon size={50} />}>
          <span style={{ fontSize: "25px" }}>Dev Dash</span>
        </IconButton>
        <IconButton icon={<DashboardIcon size={25} />}>Dashboard</IconButton>
        <IconButton icon={<CodeIcon size={25} />}>Snippets</IconButton>
      </div>
      {/* TODO: Move search bar to a separate component. */}
      <IconTextField
        icon={<SearchIcon size={18} />}
        placeholder="Search..."
        // style={{ display: "flex", left: "auto" }}
      />
      {/* TODO: Add user profile button. */}
    </header>
  );
};
