import { createRoot } from "react-dom/client";
import { router } from "./routes";
import { RouterProvider } from "react-router";
import { StrictMode } from "react";
import "../index.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
);
