export enum Complexity {
  EASY = 'Easy',
  MEDIUM = 'Medium',
  HARD = 'Hard',
}

export enum Category {
  ALGORITHMS = 'Algorithms',
  DATA_STRUCTURE = 'Data Structure',
  BRAIN_TEASER = 'BrainTeaser',
  STRINGS = 'Strings',
  BIT_MANIPULATION = 'Bit Manipulation',
  DYNAMIC_PROGRAMMING = 'Dynamic Programming',
}

export const ComplexityToColor: Record<Complexity, string> = {
  [Complexity.EASY]: 'success',
  [Complexity.MEDIUM]: 'warning',
  [Complexity.HARD]: 'danger',
};

export interface Question {
  id: number;
  title: string;
  categories: Category[];
  complexity: Complexity;
  description: string;
}
