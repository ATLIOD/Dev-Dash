import { Chip } from "@mui/material";

export enum ChipType {
  Planning = "warning",
  Active = "success",
  Maintaining = "info",
  Primary = "primary",
  Secondary = "secondary",
}

export const ChipBase = ({ ...props }) => {
  return <Chip variant="outlined" {...props} />;
};
