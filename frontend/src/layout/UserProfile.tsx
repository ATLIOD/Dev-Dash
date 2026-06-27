import { IconButton } from "../components/Buttons/Buttons";
import { Size, UserIcon } from "../components/Icons";

export const UserProfile = () => {
  // TODO: Add menu for profile and settings.
  return (
    <IconButton style={{ marginRight: "var(--spacing-small)" }}>
      <UserIcon size={Size.large} />
    </IconButton>
  );
};
