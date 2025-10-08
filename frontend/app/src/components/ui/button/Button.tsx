import React from "react";

type ButtonProps = {
  type?: "button" | "submit" | "reset";
  name?: string;
  onClick?: (e: React.MouseEvent<HTMLButtonElement>) => void;
  className?: string;
  children?: React.ReactNode;
};

export const Button: React.FC<ButtonProps> = ({
  type = "button",
  name,
  onClick,
  className = "",
  children,
}) => {
  return (
    <button
      type={type}
      name={name}
      onClick={onClick}
      className={`border rounded px-2 py-1 ${className}`}
    >
      {children}
    </button>
  );
};
