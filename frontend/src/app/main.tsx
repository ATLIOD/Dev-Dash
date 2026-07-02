import { createRoot } from "react-dom/client";
import { router } from "./routes";
import { RouterProvider } from "react-router";
import { StrictMode } from "react";
import { ThemeProvider } from "@mui/material";
import { theme } from "../styles/theme";
import "../index.scss";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider theme={theme}>
      <RouterProvider router={router} />
    </ThemeProvider>
  </StrictMode>,
);
