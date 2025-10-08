import React from "react";

type InputProps = {
  type?: string;
  name: string;
  value?: string;
  placeholder?: string;
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void;
  className?: string;
};

export const Input: React.FC<InputProps> = ({
  type = "text",
  name,
  value,
  placeholder,
  onChange,
  className = "",
}) => {
  return (
    <input
      type={type}
      name={name}
      value={value}
      placeholder={placeholder}
      onChange={onChange}
      className={`border rounded px-2 py-1 ${className}`}
    />
  );
};
