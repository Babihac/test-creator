import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = ["toggleTheme"];

  static values = {
    theme: { type: String, default: localStorage.getItem("theme") || "light" },
    classes: { type: Array, default: ["bg-moon", "bg-sun"] },
  };

  declare readonly toggleThemeTarget: HTMLButtonElement;
  declare readonly hasToggleThemeTarget: boolean;
  declare readonly toggleThemeTargets: HTMLButtonElement[];

  declare themeValue: string;
  declare classesValue: string[];
  declare readonly hasThemeValue: boolean;

  initialize(): void {
    console.log("Initialize toggle theme");
    document.documentElement.setAttribute("data-theme", this.themeValue);
    const iconClass = this.themeValue === "light" ? "bg-sun" : "bg-moon";

    this.toggleThemeTargets.forEach((target) => {
      target.classList.add(iconClass);
    });
  }

  toggle(event: Event): void {
    console.log("Toggle theme");
    event.preventDefault();
    const newTheme = this.themeValue === "light" ? "dark" : "light";
    this.themeValue = newTheme;
    document.documentElement.setAttribute("data-theme", newTheme);
    localStorage.setItem("theme", newTheme);

    this.toggleThemeTargets.forEach((target) => {
      this.classesValue.forEach((className) => {
        target.classList.toggle(className);
      });
    });
  }
}
