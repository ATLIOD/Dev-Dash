import { Outlet } from "react-router";
import { NavBar } from "./NavBar";

export default function MainLayout() {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
      }}
    >
      <NavBar />

      <main style={{ flex: 1 }}>
        <Outlet />
      </main>

      <footer
        style={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          fontSize: "12px",
        }}
      >
        © Dev Dash {new Date().getFullYear()}
      </footer>
    </div>
  );
}
