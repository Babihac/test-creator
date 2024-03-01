import { Controller } from "@hotwired/stimulus";

export default class extends Controller {
  static targets = [
    "questionNameInput",
    "questionTypeInput",
    "questionPointsInput",
    "questionTextInput",
    "questionAnswer",
    "answerList",
    "questionsCountInput",
    "answerTextInput",
    "answerCorrectInput",
  ];

  static values = {
    questionsCount: { type: Number, default: 0 },
  };

  declare readonly answerListTarget: HTMLDivElement;
  declare readonly questionAnswerTargets: HTMLDivElement[];
  declare readonly questionAnswerTarget: HTMLDivElement;
  declare readonly questionsCountInputTarget: HTMLInputElement;
  declare readonly answerTextInputTargets: HTMLInputElement[];
  declare readonly answerCorrectInputTargets: HTMLInputElement[];

  declare questionsCountValue: number;

  toggleCorrectAnswer(event: Event): void {
    const target = event.target as HTMLInputElement;
    const toggleLabel = target.parentElement?.querySelector(
      ".toggle-label"
    ) as HTMLElement;

    const realInput = target.previousElementSibling as HTMLInputElement;

    if (realInput) {
      realInput.value = target.checked ? "true" : "false";
    }

    if (target.checked) {
      toggleLabel.textContent = "Correct Answer";
    } else {
      toggleLabel.textContent = "Incorrect Answer";
    }
  }

  questionAnswerTargetConnected(): void {
    this.questionsCountValue++;
  }

  questionAnswerTargetDisconnected(): void {
    this.questionsCountValue--;
  }

  questionsCountValueChanged(): void {
    this.questionsCountInputTarget.value =
      this.questionAnswerTargets.length.toString();
  }

  removeAnswerDefinition(event: Event): void {
    const target = event.target as HTMLElement;
    const answerElement = target.closest(".answer-row") as HTMLElement;

    answerElement.remove();
  }

  // reindexAnswers(): void {
  //   this.answerTextInputTargets.forEach((input, index) => {
  //     input.name = `answer[${index}][text]`;
  //   });

  //   this.answerCorrectInputTargets.forEach((input, index) => {
  //     input.name = `answer[${index}][correct]`;
  //   });
  // }
}
