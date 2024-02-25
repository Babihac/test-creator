import React, { useState, useEffect } from "react";

export const Header: React.FC = () => <h1>React component Headerrrrr</h1>;

export const ToggleTheme: React.FC = () => {
  const [theme, setTheme] = useState("");

  const toggleTheme = () => {
    const newTheme = theme === "light" ? "dark" : "light";
    setTheme(newTheme);
    localStorage.setItem("theme", newTheme);
    document.documentElement.dataset.theme = newTheme;
  };

  useEffect(() => {
    const theme = localStorage.getItem("theme") || "light";
    setTheme(theme);

    document.documentElement.dataset.theme = theme;
  }, []);

  return (
    <button
      onClick={toggleTheme}
      aria-label="Toggle theme"
      className={`size-[24px] bg-transparent bg-no-repeat bg-center btn-ghost btn btn-sm mr-6 ${
        theme === "light" ? "bg-sun" : "bg-moon"
      }`}
    ></button>
  );
};