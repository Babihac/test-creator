import { Application } from "@hotwired/stimulus";
import ToggleThemeController from "./toggleThemeController";
import CreateTestController from "./createTest";

declare global {
  interface Window {
    Stimulus: Application;
  }
}

window.Stimulus = Application.start();

window.Stimulus.register("toggle-theme", ToggleThemeController);
window.Stimulus.register("create-test", CreateTestController);
