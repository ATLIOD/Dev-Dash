import type { HtmlHTMLAttributes } from "react";
import "./_components.scss";

export enum ChipType {
  Planning = "planning",
  Active = "active",
  Maintaining = "maintaining",
  Primary = "primary",
  Secondary = "secondary",
}

type ChipProps = HtmlHTMLAttributes<HTMLSpanElement> & {
  className?: string;
  baseClass?: ChipType;
};

export const Chip = ({
  baseClass = ChipType.Secondary,
  className,
  children,
  ...props
}: ChipProps) => {
  return (
    <div className="chip-wrapper">
      <span {...props} className={`${baseClass}` + (className ? ` ${className}` : "")}>
        {children}
      </span>
    </div>
  );
};
