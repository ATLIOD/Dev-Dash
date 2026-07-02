import { CloseOutlined, MoreVertOutlined } from "@mui/icons-material";
import { Button, IconButton } from "@mui/material";
import { Size } from "../Icons";
import type { ReactNode } from "react";

export const MenuButton = ({
  onClick,
}: {
  onClick: (e: React.MouseEvent<HTMLButtonElement>) => void;
}) => {
  return (
    <IconButton
      className="menu-button"
      onClick={onClick}
      style={{
        padding: 0,
      }}
    >
      <MoreVertOutlined />
    </IconButton>
  );
};

export const CloseButton = ({ onClick }: { onClick: () => void }) => {
  return (
    <IconButton className="menu-button" onClick={onClick}>
      <CloseOutlined />
    </IconButton>
  );
};

export const NavButton = ({
  size = Size.large,
  startIcon,
  endIcon,
  children,
  ...props
}: {
  size?: Size;
  startIcon?: ReactNode;
  endIcon?: ReactNode;
  children: ReactNode;
}) => {
  return (
    <Button
      sx={{
        color: "text.primary",
        "& .MuiButton-startIcon .MuiSvgIcon-root": {
          color: "primary.main",
          fontSize: size,
        },
        "& .MuiButton-startIcon": {
          margin: 0,
        },
      }}
      {...props}
      startIcon={startIcon}
      endIcon={endIcon}
    >
      {children}
    </Button>
  );
};
