import { CloseOutlined } from "@mui/icons-material";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  IconButton,
} from "@mui/material";
import type { ReactNode } from "react";

const DialogHeader = ({ children, onClose }: { children: ReactNode; onClose: () => void }) => {
  return (
    <DialogTitle
      sx={{
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
      }}
    >
      {children}
      <IconButton>
        <CloseOutlined onClick={onClose} />
      </IconButton>
    </DialogTitle>
  );
};

export const DialogBase = ({
  className,
  confirmText = "Confirm",
  isValid,
  onClose,
  onConfirm,
  open,
  title,
  children,
  ...props
}: {
  className?: string;
  confirmText?: string;
  isValid: boolean;
  onClose: () => void;
  onConfirm: () => void;
  open: boolean;
  title: string;
  children: ReactNode;
}) => {
  return (
    <Dialog
      open={open}
      onClose={onClose}
      className={"dialog" + (className ? ` ${className}` : "")}
      {...props}
      maxWidth={false}
    >
      <DialogHeader onClose={onClose}>{title}</DialogHeader>
      <DialogContent>{children}</DialogContent>

      <DialogActions>
        <Button onClick={onClose} color="secondary">
          Cancel
        </Button>
        <Button onClick={onConfirm} variant="contained" disabled={!isValid}>
          {confirmText}
        </Button>
      </DialogActions>
    </Dialog>
  );
};
