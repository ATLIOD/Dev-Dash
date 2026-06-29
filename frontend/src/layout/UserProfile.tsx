import { IconButton } from "@mui/material";
import { Size, UserIcon } from "../components/Icons";

export const UserProfile = () => {
  // TODO: Add menu for profile and settings.
  return (
    <IconButton
      sx={{
        padding: 0,
      }}
    >
      <UserIcon size={Size.large} />
    </IconButton>
  );
};
