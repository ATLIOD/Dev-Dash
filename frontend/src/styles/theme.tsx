import { alpha, createTheme } from "@mui/material/styles";

export const theme = createTheme({
  colorSchemes: {
    light: {
      palette: {
        primary: {
          main: "#ad7eff",
        },
        secondary: {
          main: "#6a13ff",
        },
        warning: {
          main: "#c3a900",
        },
        success: {
          main: "#01b54c",
        },
        info: {
          main: "#01a5c9",
        },
        text: {
          primary: "#6b6375",
          secondary: "#9f9ca2",
        },
        background: {
          default: "#fff",
          paper: "#fafafa",
        },
        divider: "#e5e4e7",
      },
    },
    dark: {
      palette: {
        primary: {
          main: "#ad7eff",
        },
        secondary: {
          main: "#6a13ff",
        },
        warning: {
          main: "#f6d500",
        },
        success: {
          main: "#00f666",
        },
        info: {
          main: "#00c9f6",
        },
        text: {
          primary: "#d9d9d9",
          secondary: "#595959ab",
        },
        background: {
          default: "#0f1416",
          paper: "#15191D",
        },
        divider: "#2b3139ab",
      },
    },
  },

  shape: {
    borderRadius: "8px",
  },
  typography: {
    fontFamily: `"Inter", sans-serif`,
  },

  components: {
    //#region Menu
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

    //#region Chip
    MuiChip: {
      styleOverrides: {
        root: {
          height: "100%",
        },
        label: {
          padding: "0 4px",
          fontSize: "var(--font-size-small)",
        },
      },
      variants: [
        {
          props: { variant: "outlined", color: "default" },
          style: ({ theme }) => ({
            color: theme.palette.text.primary,
            backgroundColor: alpha(theme.palette.text.primary, 0.12),
          }),
        },
        {
          props: { variant: "outlined", color: "primary" },
          style: ({ theme }) => ({
            backgroundColor: alpha(theme.palette.primary.main, 0.12),
          }),
        },
        {
          props: { variant: "outlined", color: "secondary" },
          style: ({ theme }) => ({
            backgroundColor: alpha(theme.palette.secondary.main, 0.12),
          }),
        },
        {
          props: { variant: "outlined", color: "success" },
          style: ({ theme }) => ({
            backgroundColor: alpha(theme.palette.success.main, 0.12),
          }),
        },
        {
          props: { variant: "outlined", color: "warning" },
          style: ({ theme }) => ({
            backgroundColor: alpha(theme.palette.warning.main, 0.12),
          }),
        },
        {
          props: { variant: "outlined", color: "info" },
          style: ({ theme }) => ({
            backgroundColor: alpha(theme.palette.info.main, 0.12),
          }),
        },
        {
          props: { variant: "outlined", color: "error" },
          style: ({ theme }) => ({
            backgroundColor: alpha(theme.palette.error.main, 0.12),
          }),
        },
      ],
    },

    //#region Buttons
    MuiButton: {
      styleOverrides: {
        root: {
          borderRadius: "6px",
          boxShadow: "none",
          "&:hover": {
            boxShadow: "none",
          },
          "&:active": {
            boxShadow: "none",
          },
          "&.Mui-focusVisible": {
            boxShadow: "none",
          },
          textTransform: "none",
          padding: "4px 8px",
        },
      },
      variants: [
        {
          props: { variant: "contained", color: "primary" },
          style: ({ theme }) => ({
            color: "#fff",
            backgroundColor: theme.palette.secondary.main,
          }),
        },
      ],
    },

    //#region Inputs
    MuiOutlinedInput: {
      styleOverrides: {
        adornedStart: {
          paddingLeft: "4px",
        },
      },
    },

    MuiInputBase: {
      styleOverrides: {
        root: {
          fontSize: "14px",
          backgroundColor: "var(--bg)",

          input: {
            padding: "8px",
          },
        },
        multiline: {
          padding: "8px",
        },
      },
    },

    MuiSelect: {
      styleOverrides: {
        select: {
          padding: "8px",
        },
      },
    },

    //#region Dialog
    MuiDialog: {
      styleOverrides: {
        paper: {
          backgroundColor: "var(--surface-color-darken)",
          backgroundImage: "none",
          padding: "8px",
          border: "1px solid color-mix(in srgb, var(--border) 60%, var(--primary-text))",
          color: "var(--primary-text)",
          fontFamily: "Inter",
        },
      },
    },
    MuiDialogTitle: {
      styleOverrides: {
        root: {
          fontSize: "var(--font-size-large)",
          padding: "8px",
        },
      },
    },

    MuiDialogContent: {
      styleOverrides: {
        root: {
          fontSize: "var(--font-size-medium)",
          color: "var(--secondary-text)",
          padding: "8px",
        },
      },
    },

    MuiDialogActions: {
      styleOverrides: {
        root: {
          fontSize: "var(--font-size-medium)",
        },
      },
    },
  },
});
