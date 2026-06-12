export enum Size {
  "small" = 16,
  "medium" = 24,
  "large" = 32,
  "xlarge" = 48,
}

export const IconBase = ({
  path,
  size = Size.medium,
  color = "var(--primary-color)",
  className,
  viewBox = "0 -960 960 960",
}: {
  path: string;
  size?: Size;
  color?: string;
  className?: string;
  viewBox?: string;
}) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    height={size}
    width={size}
    viewBox={viewBox}
    fill={color}
    className={className}
  >
    <path d={path} />
  </svg>
);

// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:add:FILL@0;wght@400;GRAD@0;opsz@24&icon.query=plus&icon.size=24&icon.color=%23000000
export const AddIcon = ({ ...props }) => (
  <IconBase {...props} path={"M440-440H200v-80h240v-240h80v240h240v80H520v240h-80v-240Z"} />
);

// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:code_xml:FILL@0;wght@400;GRAD@0;opsz@24&icon.query=code&icon.size=24&icon.color=%231f1f1f
export const CodeIcon = ({ ...props }) => (
  <IconBase
    {...props}
    path={
      "M240-280 40-480l200-200 56 56-143 144 143 144-56 56Zm178 132-76-24 200-640 76 24-200 640Zm302-132-56-56 143-144-143-144 56-56 200 200-200 200Z"
    }
  />
);

// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:team_dashboard:FILL@0;wght@400;GRAD@0;opsz@24&icon.query=dashboard&icon.size=24&icon.color=%231f1f1f
export const DashboardIcon = ({ ...props }) => (
  <IconBase
    {...props}
    path={
      "M200-120q-33 0-56.5-23.5T120-200v-560q0-33 23.5-56.5T200-840h560q33 0 56.5 23.5T840-760v560q0 33-23.5 56.5T760-120H200Zm200-80v-240H200v240h200Zm80 0h280v-240H480v240ZM200-520h560v-240H200v240Z"
    }
  />
);

// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:folder_open:FILL@0;wght@400;GRAD@0;opsz@24&icon.query=folder&icon.size=24&icon.color=%231f1f1f
export const FolderIcon = ({ ...props }) => (
  <IconBase
    {...props}
    path={
      "M160-160q-33 0-56.5-23.5T80-240v-480q0-33 23.5-56.5T160-800h240l80 80h320q33 0 56.5 23.5T880-640H160v400l96-320h684L837-217q-8 26-29.5 41.5T760-160H160Z"
    }
  />
);

// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:search:FILL@0;wght@400;GRAD@0;opsz@24&icon.query=search&icon.size=24&icon.color=%231f1f1f
export const SearchIcon = ({ ...props }) => (
  <IconBase
    {...props}
    path={
      "M784-120 532-372q-30 24-69 38t-83 14q-109 0-184.5-75.5T120-580q0-109 75.5-184.5T380-840q109 0 184.5 75.5T640-580q0 44-14 83t-38 69l252 252-56 56ZM380-400q75 0 127.5-52.5T560-580q0-75-52.5-127.5T380-760q-75 0-127.5 52.5T200-580q0 75 52.5 127.5T380-400Z"
    }
  />
);

// https://fonts.google.com/icons?icon.query=code&selected=Material+Symbols+Outlined:terminal:FILL@0;wght@400;GRAD@0;opsz@24&icon.size=24&icon.color=%231f1f1f
export const TerminalIcon = ({ ...props }) => (
  <IconBase
    {...props}
    path={
      "M160-160q-33 0-56.5-23.5T80-240v-480q0-33 23.5-56.5T160-800h640q33 0 56.5 23.5T880-720v480q0 33-23.5 56.5T800-160H160Zm0-80h640v-400H160v400Zm140-40-56-56 103-104-104-104 57-56 160 160-160 160Zm180 0v-80h240v80H480Z"
    }
  />
);

// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:sentiment_excited:FILL@0;wght@400;GRAD@0;opsz@24&icon.query=user&icon.size=24&icon.color=%231f1f1f
export const UserIcon = ({ ...props }) => (
  <IconBase
    {...props}
    path={
      "M320-480v80q0 66 47 113t113 47q66 0 113-47t47-113v-80H320Zm160 180q-42 0-71-29t-29-71v-20h200v20q0 42-29 71t-71 29ZM272.5-652.5Q243-625 231-577l58 14q6-26 20-41.5t31-15.5q17 0 31 15.5t20 41.5l58-14q-12-48-41.5-75.5T340-680q-38 0-67.5 27.5Zm280 0Q523-625 511-577l58 14q6-26 20-41.5t31-15.5q17 0 31 15.5t20 41.5l58-14q-12-48-41.5-75.5T620-680q-38 0-67.5 27.5ZM324-111.5Q251-143 197-197t-85.5-127Q80-397 80-480t31.5-156Q143-709 197-763t127-85.5Q397-880 480-880t156 31.5Q709-817 763-763t85.5 127Q880-563 880-480t-31.5 156Q817-251 763-197t-127 85.5Q563-80 480-80t-156-31.5ZM480-480Zm227 227q93-93 93-227t-93-227q-93-93-227-93t-227 93q-93 93-93 227t93 227q93 93 227 93t227-93Z"
    }
  />
);

// https://fonts.google.com/icons?selected=Material+Symbols+Outlined:more_vert:FILL@0;wght@400;GRAD@0;opsz@24&icon.query=ver&icon.size=24&icon.color=%23000000
export const VertKebab = ({ ...props }) => (
  <IconBase
    {...props}
    path={
      "M480-160q-33 0-56.5-23.5T400-240q0-33 23.5-56.5T480-320q33 0 56.5 23.5T560-240q0 33-23.5 56.5T480-160Zm0-240q-33 0-56.5-23.5T400-480q0-33 23.5-56.5T480-560q33 0 56.5 23.5T560-480q0 33-23.5 56.5T480-400Zm0-240q-33 0-56.5-23.5T400-720q0-33 23.5-56.5T480-800q33 0 56.5 23.5T560-720q0 33-23.5 56.5T480-640Z"
    }
  />
);
