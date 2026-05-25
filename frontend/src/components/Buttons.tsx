import type { ButtonHTMLAttributes, ReactNode } from "react";
import "./buttons.css";

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  className?: string;
};

export const Button = ({ className, children, ...props }: ButtonProps) => {
  return (
    <button className={className} {...props}>
      {children}
    </button>
  );
};

type IconButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  icon: ReactNode;
  className?: string;
};

export const IconButton = ({ icon, className, children, ...props }: IconButtonProps) => {
  return (
    <Button className={`icon-button ${className}`} {...props}>
      {icon}
      {children}
    </Button>
  );
};
