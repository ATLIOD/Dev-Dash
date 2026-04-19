import { createBrowserRouter } from "react-router";
import Home from "../pages/Home";
import MainLayout from "../layout/MainLayout";

// Additional routing information:
// https://reactrouter.com/start/data/routing

export const router = createBrowserRouter([
  {
    Component: MainLayout,
    children: [{ index: true, Component: Home }],
  },
]);
