export enum Complexity {
  EASY = 'Easy',
  MEDIUM = 'Medium',
  HARD = 'Hard',
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
  categories: string[];
  complexity: Complexity;
  description: string;
}
