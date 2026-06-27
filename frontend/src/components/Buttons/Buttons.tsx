import type { ButtonHTMLAttributes } from "react";
import "./buttons.scss";
import { VertKebab } from "../Icons";

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  className?: string;
  baseClass?: "text" | "primary" | "outlined" | "secondary";
};

export const Button = ({ baseClass = "text", className, children, ...props }: ButtonProps) => {
  return (
    <button className={baseClass + (className ? ` ${className}` : "")} {...props}>
      {children}
    </button>
  );
};

export const IconButton = ({ baseClass, className, children, ...props }: ButtonProps) => {
  return (
    <Button baseClass={baseClass} className={`icon-button ${className ?? ""}`} {...props}>
      {children}
    </Button>
  );
};

export const MenuButton = ({
  onClick,
}: {
  onClick: (e: React.MouseEvent<HTMLButtonElement>) => void;
}) => {
  return (
    <IconButton baseClass="text" className="menu-button" onClick={onClick}>
      <VertKebab className="kebab" color="var(--primary-text)" />
    </IconButton>
  );
};
