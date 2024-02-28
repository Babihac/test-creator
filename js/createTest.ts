import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = [
    "durationInput",
    "durationValue",
    "testNameInput",
    "teacherIdInput",
    "maxScoreInput",
    "minRequiredScoreInput",
    "step",
  ];

  static values = {
    duration: { type: Number, default: 60 },
    testName: { type: String, default: "" },
    teacherId: { type: String, default: "" },
    maxScore: { type: Number, default: 100 },
    minRequiredScore: { type: Number, default: 60 },
    currentStep: { type: Number, default: 0 },
  };

  declare readonly durationValueTarget: HTMLSpanElement;

  declare readonly durationInputTarget: HTMLInputElement;
  declare readonly testNameInputTarget: HTMLInputElement;
  declare readonly teacherIdInputTarget: HTMLInputElement;
  declare readonly maxScoreInputTarget: HTMLInputElement;
  declare readonly minRequiredScoreInputTarget: HTMLInputElement;
  declare readonly stepOneTarget: HTMLDivElement;
  declare readonly stepTwoTarget: HTMLDivElement;
  declare readonly stepTargets: HTMLDivElement[];

  declare readonly hasMaxScoreInputTarget: boolean;
  declare readonly hasMinRequiredScoreInputTarget: boolean;
  declare readonly hasDurationInputTarget: boolean;
  declare readonly hasTestNameInputTarget: boolean;
  declare readonly hasTeacherIdInputTarget: boolean;

  declare durationValue: number;
  declare testNameValue: string;
  declare teacherIdValue: string;
  declare maxScoreValue: number;
  declare minRequiredScoreValue: number;
  declare currentStepValue: number;

  initialize(): void {
    this.durationValueTarget.textContent = this.durationValue.toString();
    this.teacherIdValue = this.teacherIdInputTarget.value;

    this.stepTargets.forEach((step, index) => {
      if (index === this.currentStepValue) {
        step.classList.remove("hidden");
      } else {
        step.classList.add("hidden");
      }
    });
  }

  updateDuration(event: Event): void {
    const value = parseInt(this.durationInputTarget.value, 10);
    this.durationValue = value;
    this.durationValueTarget.textContent = value.toString();
  }

  updateTestName(event: Event): void {
    this.testNameValue = this.testNameInputTarget.value;
  }

  updateTeacherId(event: Event): void {
    this.teacherIdValue = this.teacherIdInputTarget.value;
  }

  updateMaxScore(event: Event): void {
    this.maxScoreValue = parseInt(this.maxScoreInputTarget.value, 10);
  }

  updateMinRequiredScore(event: Event): void {
    this.minRequiredScoreValue = parseInt(
      this.minRequiredScoreInputTarget.value,
      10
    );
  }

  durationInputTargetConnected(): void {
    this.durationInputTarget.value = this.durationValue.toString();
    this.durationValueTarget.textContent = this.durationValue.toString();

    this.durationInputTarget.addEventListener(
      "input",
      this.updateDuration.bind(this)
    );
  }

  durationInputTargetDisconnected(): void {
    this.hasDurationInputTarget &&
      this.durationInputTarget.removeEventListener(
        "input",
        this.updateDuration.bind(this)
      );
  }

  testNameInputTargetConnected(): void {
    this.testNameInputTarget.value = this.testNameValue;

    this.testNameInputTarget.addEventListener(
      "input",
      this.updateTestName.bind(this)
    );
  }

  testNameInputTargetDisconnected(): void {
    this.hasTestNameInputTarget &&
      this.testNameInputTarget.removeEventListener(
        "input",
        this.updateTestName.bind(this)
      );
  }

  teacherIdInputTargetConnected(): void {
    this.teacherIdInputTarget.value = this.teacherIdValue;

    this.teacherIdInputTarget.addEventListener(
      "input",
      this.updateTeacherId.bind(this)
    );
  }

  teacherIdInputTargetDisconnected(): void {
    this.hasTeacherIdInputTarget &&
      this.teacherIdInputTarget.removeEventListener(
        "input",
        this.updateTeacherId.bind(this)
      );
  }

  maxScoreInputTargetConnected(): void {
    console.log("max score connected");
    this.maxScoreInputTarget.value = this.maxScoreValue.toString();

    this.maxScoreInputTarget.addEventListener(
      "input",
      this.updateMaxScore.bind(this)
    );
  }

  maxScoreInputTargetDisconnected(): void {
    this.hasMaxScoreInputTarget &&
      this.maxScoreInputTarget.removeEventListener(
        "input",
        this.updateMaxScore.bind(this)
      );
  }

  minRequiredScoreInputTargetConnected(): void {
    this.minRequiredScoreInputTarget.value =
      this.minRequiredScoreValue.toString();

    this.minRequiredScoreInputTarget.addEventListener(
      "input",
      this.updateMinRequiredScore.bind(this)
    );
  }

  minRequiredScoreInputTargetDisconnected(): void {
    this.hasMinRequiredScoreInputTarget &&
      this.minRequiredScoreInputTarget.removeEventListener(
        "input",
        this.updateMinRequiredScore.bind(this)
      );
  }

  stepOneTargetConnected(): void {
    console.log("step one connected");
    this.currentStepValue = 0;
  }

  stepTwoTargetConnected(): void {
    console.log("step two connected");
    this.currentStepValue = 1;
  }

  currentStepValueChanged(): void {
    console.log("change", this.currentStepValue);
    this.stepTargets.forEach((step, index) => {
      if (index === this.currentStepValue) {
        step.classList.remove("hidden");
      } else {
        step.classList.add("hidden");
      }
    });
  }
}
