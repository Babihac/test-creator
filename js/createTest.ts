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
    "selectedQuestionsChips",
  ];

  static values = {
    duration: { type: Number, default: 60 },
    testName: { type: String, default: "" },
    teacherId: { type: String, default: "" },
    teacherName: { type: String, default: "" },
    maxScore: { type: Number, default: 100 },
    minRequiredScore: { type: Number, default: 60 },
    currentStep: { type: Number, default: 0 },
    questions: { type: Array, default: [] },
  };

  declare readonly durationValueTarget: HTMLSpanElement;

  declare readonly durationInputTarget: HTMLInputElement;
  declare readonly testNameInputTarget: HTMLInputElement;
  declare readonly teacherIdInputTarget: HTMLSelectElement;
  declare readonly maxScoreInputTarget: HTMLInputElement;
  declare readonly minRequiredScoreInputTarget: HTMLInputElement;
  declare readonly stepOneTarget: HTMLDivElement;
  declare readonly stepTwoTarget: HTMLDivElement;
  declare readonly stepTargets: HTMLDivElement[];
  declare readonly selectedQuestionsChipsTarget: HTMLDivElement;

  declare readonly hasMaxScoreInputTarget: boolean;
  declare readonly hasMinRequiredScoreInputTarget: boolean;
  declare readonly hasDurationInputTarget: boolean;
  declare readonly hasTestNameInputTarget: boolean;
  declare readonly hasTeacherIdInputTarget: boolean;

  declare durationValue: number;
  declare testNameValue: string;
  declare teacherIdValue: string;
  declare teacherNameValue: string;
  declare maxScoreValue: number;
  declare minRequiredScoreValue: number;
  declare currentStepValue: number;
  declare questionsValue: {
    name: string;
    id: string;
    points: string;
  }[];

  initialize(): void {
    this.durationValueTarget.textContent = this.durationValue.toString();
    this.teacherIdValue = this.teacherIdInputTarget.value;
    this.teacherNameValue = this.teacherIdInputTarget.selectedOptions[0].text;

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
    const target = event.target as HTMLSelectElement;
    this.teacherIdValue = this.teacherIdInputTarget.value;
    this.teacherNameValue = target.selectedOptions[0].text;
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

    console.log("teacher id", this.teacherIdValue);

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
    this.currentStepValue = 0;
  }

  stepTwoTargetConnected(): void {
    this.currentStepValue = 1;
  }

  stepThreeTargetConnected(): void {
    this.currentStepValue = 2;
  }

  stepFourTargetConnected(): void {
    this.currentStepValue = 3;
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
  summaryTestNameTargetConnected(element: Element): void {
    element.textContent = this.testNameValue;
  }

  summaryTestDurationTargetConnected(element: Element): void {
    element.textContent = this.durationValue.toString();
  }

  summaryTeacherTargetConnected(element: Element): void {
    element.textContent = this.teacherNameValue;
  }

  summaryMaxPointsTargetConnected(element: Element): void {
    element.textContent = this.maxScoreValue.toString();
  }

  summaryMinPointsTargetConnected(element: Element): void {
    element.textContent = this.minRequiredScoreValue.toString();
  }

  summaryQuestionsTargetConnected(element: Element): void {
    this.questionsValue.forEach((question) => {
      const tableRow = document.createElement("tr");

      tableRow.classList.add("border-neutral-100");

      const questionName = document.createElement("td");
      const questionPoints = document.createElement("td");

      questionName.textContent = question.name;
      questionPoints.textContent = question.points;

      tableRow.appendChild(questionName);
      tableRow.appendChild(questionPoints);
      element.appendChild(tableRow);
    });
  }

  selectedQuestionsChipsTargetConnected(): void {
    console.log("selected questions chips connected");

    this.questionsValue.forEach((question) => {
      const chip = this.createQuestionChip(question.name, question.id);
      this.selectedQuestionsChipsTarget.appendChild(chip);
    });
  }

  addQuestion(event: Event): void {
    console.log("add question");
    const target = event.target as HTMLButtonElement;
    const name = target.dataset.questionName;
    const id = target.dataset.questionId;
    const points = target.dataset.questionPoints;
    const chip = this.createQuestionChip(name ?? "", id ?? "");

    if (!name || !id || !points) return;

    this.selectedQuestionsChipsTarget.appendChild(chip);

    this.questionsValue.push({ name, id, points });

    console.log("questions", this.questionsValue);
  }

  deleteQuestion(event: Event): void {
    console.log("delete question");
    const target = event.target as HTMLButtonElement;
    const id = target.dataset.questionId;

    if (id) {
      const index = this.questionsValue.findIndex(
        (question) => question.id === id
      );
      if (index > -1) {
        this.questionsValue.splice(index, 1);
      }
    }

    target.parentElement?.remove();

    console.log("questions", this.questionsValue);
  }

  createQuestionChip(question: string, id: string): HTMLDivElement {
    const chip = document.createElement("div");
    const deleteIcon = document.createElement("i");
    const textDiv = document.createElement("div");
    const hiddenInput = document.createElement("input");
    hiddenInput.type = "hidden";
    hiddenInput.name = "questions[]";
    hiddenInput.value = id;

    textDiv.textContent = question;
    deleteIcon.dataset.action = "click->create-test#deleteQuestion";
    deleteIcon.classList.add(
      "fa-solid",
      "fa-x",
      "fa-xs",
      "cursor-pointer",
      "p-1"
    );
    deleteIcon.dataset.questionId = id;
    chip.classList.add("badge", "badge-primary", "flex", "gap-1");
    chip.appendChild(textDiv);
    chip.appendChild(deleteIcon);
    chip.appendChild(hiddenInput);
    return chip;
  }
}
