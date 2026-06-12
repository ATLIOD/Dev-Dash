import "./_components.scss";

export const Chip = ({ children, ...props }) => {
  return (
    <span {...props} className="chip">
      {children}
    </span>
  );
};
