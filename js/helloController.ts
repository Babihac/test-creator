import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = ["name"];

  declare readonly nameTarget: HTMLInputElement;
  declare readonly hasNameTarget: boolean;
  declare readonly nameTargets: HTMLInputElement[];
  declare readonly hasxNameTargets: boolean;

  initialize() {
    console.log("Controller loaded");
    console.log("hasNameTarget:", this.hasNameTarget);
    console.log("nameTargets:", this.nameTargets);
  }

  greet() {
    this.nameTarget.textContent = `Hello, ${this.nameTarget.value}!`;
  }
}
