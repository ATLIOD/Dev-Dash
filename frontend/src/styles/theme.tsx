import { createTheme } from "@mui/material/styles";

export const theme = createTheme({
  components: {
    MuiMenu: {
      defaultProps: {
        anchorOrigin: {
          vertical: "bottom",
          horizontal: "left",
        },
        transformOrigin: {
          vertical: "top",
          horizontal: "left",
        },
      },
      styleOverrides: {
        paper: {
          borderRadius: "var(--spacing-medium)",
          border: "1px solid color-mix(in srgb, var(--border) 60%, var(--primary-text))",
          boxShadow: "none",
          marginLeft: "16px",
        },

        list: {
          padding: 0,
          background: "var(--surface-color)",
        },
      },
    },

    MuiMenuItem: {
      styleOverrides: {
        root: {
          color: "var(--primary-text)",
          fontSize: "var(--font-size-small)",
          fontFamily: "Inter",
          minHeight: 24,
          backgroundColor: "var(--surface-color)",
          "&:not(:last-of-type)": {
            borderBottom: "1px solid color-mix(in srgb, var(--border) 60%, var(--primary-text))",
          },
        },
      },
    },
  },
});
