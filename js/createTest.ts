import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = ["durationInput", "durationValue"];

  static values = {
    duration: { type: Number, default: 60 },
  };

  declare readonly durationInputTarget: HTMLInputElement;
  declare readonly durationValueTarget: HTMLSpanElement;

  declare durationValue: number;
  declare readonly hasDurationValue: boolean;

  initialize(): void {
    this.durationValueTarget.textContent = this.durationValue.toString();
    this.durationInputTarget.value = this.durationValue.toString();
  }

  updateDuration(event: Event): void {
    const value = parseInt(this.durationInputTarget.value, 10);
    this.durationValue = value;
    this.durationValueTarget.textContent = value.toString();
  }
}
