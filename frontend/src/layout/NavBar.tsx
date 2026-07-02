import { InputAdornment, TextField } from "@mui/material";
import {
  CodeOutlined,
  SearchOutlined,
  SpaceDashboardOutlined,
  TerminalOutlined,
} from "@mui/icons-material";
import "./_layout.scss";
import { UserProfile } from "./UserProfile";
import { NavButton } from "../components/Buttons/Buttons";
import { Size } from "../components/Icons";

export const NavBar = () => {
  return (
    <header>
      <div className="header-wrapper">
        <div className="nav-buttons">
          <NavButton startIcon={<TerminalOutlined />} size={Size.xlarge}>
            <span style={{ fontSize: "var(--font-size-large)" }}>Dev Dash</span>
          </NavButton>
          <NavButton startIcon={<SpaceDashboardOutlined />}>Dashboard</NavButton>
          <NavButton startIcon={<CodeOutlined />}>Snippets</NavButton>
        </div>

        <div className="nav-buttons">
          {/* TODO: Move search bar to a separate component. */}
          <TextField
            slotProps={{
              input: {
                startAdornment: (
                  <InputAdornment position="start">
                    <SearchOutlined />
                  </InputAdornment>
                ),
              },
            }}
            placeholder="Search..."
          />
          <UserProfile />
        </div>
      </div>
    </header>
  );
};
