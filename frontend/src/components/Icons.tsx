export const IconBase = ({
  path,
  size = 24,
  color = "var(--primary-color)",
  className,
  viewBox = "0 -960 960 960",
}: {
  path: string;
  size?: number | string;
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
