import { Outlet } from "react-router";
import { NavBar } from "./NavBar";

export default function MainLayout() {
  return (
    <div>
      <NavBar />

      <main>
        <Outlet />
      </main>

      <footer>© Dev Dash {new Date().getFullYear()}</footer>
    </div>
  );
}
