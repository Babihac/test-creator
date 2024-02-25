import { createRoot, hydrateRoot } from "react-dom/client";
import { Header, Body, Hello } from "./components";
import { ToggleTheme } from "toggleTheme";

const toggleRoot = document.getElementById("react-toggle-theme");
const mobileToggleRoot = document.getElementById("react-toggle-theme-mobile");

if (toggleRoot) {
  const toggleReactRoot = createRoot(toggleRoot);
  toggleReactRoot.render(<ToggleTheme />);
}

if (mobileToggleRoot) {
  const mobileToggleReactRoot = createRoot(mobileToggleRoot);
  mobileToggleReactRoot.render(<ToggleTheme />);
}

export function renderToggleTheme(id: string) {
  const root = document.getElementById(id);
  if (root) {
    const reactRoot = createRoot(root);
    reactRoot.render(<ToggleTheme />);
  }
}
