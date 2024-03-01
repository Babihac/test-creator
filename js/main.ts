import { Application } from "@hotwired/stimulus";
import ToggleThemeController from "./toggleThemeController";
import CreateTestController from "./createTest";
import CreateQuestion from "./createQuestion";

declare global {
  interface Window {
    Stimulus: Application;
  }
}

window.Stimulus = Application.start();

window.Stimulus.register("toggle-theme", ToggleThemeController);
window.Stimulus.register("create-test", CreateTestController);
window.Stimulus.register("create-question", CreateQuestion);
