import React from "react";

type TextFieldProps = React.InputHTMLAttributes<HTMLInputElement> & {
  className?: string;
};

export const TextField = ({ className, ...props }: TextFieldProps) => {
  return <input className={className} {...props} />;
};

type IconTextFieldProps = TextFieldProps & {
  icon?: React.ReactNode;
};

export const IconTextField = ({ icon, className = "", ...props }: IconTextFieldProps) => {
  return (
    // TODO: implement wrapper.
    <div className="icon-textfield-wrapper">
      {icon && <span className="textfield-icon">{icon}</span>}
      <TextField className={className} {...props} />
    </div>
  );
};
