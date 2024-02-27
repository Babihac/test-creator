import { Application } from "@hotwired/stimulus";
import HelloController from "./helloController";

declare global {
  interface Window {
    Stimulus: Application;
  }
}

window.Stimulus = Application.start();

window.Stimulus.register("hello", HelloController);
