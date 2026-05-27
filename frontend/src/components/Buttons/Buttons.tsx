import type { ButtonHTMLAttributes, ReactNode } from "react";
import "./buttons.css";

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  className?: string;
  baseClass?: "text" | "primary" | "outlined";
};

export const Button = ({ baseClass = "text", className, children, ...props }: ButtonProps) => {
  return (
    <button className={`${baseClass} ${className ?? ""}`} {...props}>
      {children}
    </button>
  );
};

type IconButtonProps = ButtonProps & {
  icon: ReactNode;
};

export const IconButton = ({ icon, baseClass, className, children, ...props }: IconButtonProps) => {
  return (
    <Button baseClass={baseClass} className={`icon-button ${className ?? ""}`} {...props}>
      {icon}
      {children}
    </Button>
  );
};
