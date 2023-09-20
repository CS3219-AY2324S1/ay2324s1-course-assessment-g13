import { Category, Complexity } from "./question";

export interface Preference {
  categories: Category[] | string;
  complexities: Complexity[] | string;
}
