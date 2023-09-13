export enum Complexity {
  EASY = 'Easy',
  MEDIUM = 'Medium',
  HARD = 'Hard',
}

export enum Category {
  ALGORITHMS = 'Algorithms',
  DATA_STRUCTURES = 'Data Structures',
  BRAIN_TEASER = 'BrainTeaser',
  STRINGS = 'Strings',
  BIT_MANIPULATION = 'Bit Manipulation',
  DYNAMIC_PROGRAMMING = 'Dynamic Programming',
}

export enum ComplexityColor {
  DEFAULT = 'default',
  PRIMARY = 'primary',
  SECONDARY = 'secondary',
  SUCCESS = 'success',
  WARNING = 'warning',
  DANGER = 'danger',
}

export const ComplexityToColor: Record<Complexity, ComplexityColor> = {
  [Complexity.EASY]: ComplexityColor.SUCCESS,
  [Complexity.MEDIUM]: ComplexityColor.WARNING,
  [Complexity.HARD]: ComplexityColor.DANGER,
};

export interface Question {
  id: string;
  title: string;
  categories: Category[] | string;
  complexity: Complexity;
  description: string;
}
