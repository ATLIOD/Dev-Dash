import { useEffect, useRef } from "react";
import "./_components.scss";

export const Menu = ({
  anchorEl,
  children,
  className = "",
  onClose,
  open,
}: {
  anchorEl: HTMLElement | null;
  children: React.ReactNode;
  className?: string;
  onClose: () => void;
  open: boolean;
}) => {
  if (!open || !anchorEl) {
    return null;
  }

  const box = anchorEl.getBoundingClientRect();

  function HandleOnKeyDown(event: React.KeyboardEvent<HTMLDivElement>) {
    if (event.key === "Escape") {
      onClose();
    }
  }

  return (
    <>
      <div
        onClick={onClose}
        style={{
          position: "fixed",
          inset: 0,
        }}
      />
      {open && (
        <div
          className={"menu-wrapper" + (className ? ` ${className}` : "")}
          onKeyDown={HandleOnKeyDown}
          style={{
            top: box.bottom,
            left: box.left + 16,
          }}
        >
          {children}
        </div>
      )}
    </>
  );
};

export const MenuItem = ({
  children,
  onClick,
}: {
  children: React.ReactNode;
  onClick: () => void;
}) => {
  return (
    <button onClick={onClick} className="menu-item">
      {children}
    </button>
  );
};
