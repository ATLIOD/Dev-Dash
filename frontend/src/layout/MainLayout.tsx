import { Outlet } from "react-router";
import { NavBar } from "./NavBar";

export default function MainLayout() {
  return (
    <>
      {/* TO-DO #16: Set up main layout with Navbar, footer, and page outlet.
    https://reactrouter.com/start/declarative/routing */}
      <NavBar />
      <Outlet />
    </>
  );
}
