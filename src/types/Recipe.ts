import type { StrapiMediaCollection } from "./Strapi";

export interface Recipe {
  Title: string;
  meal_type: MealType;
  cuisine: Cuisine;
  duration: Duration;
  requirements: Requirement[];
  Ingredients: string;
  Images: StrapiMediaCollection;
  Method: string;
  Foreword: string;
}

interface MealType {
  data: { attributes: { Type: string } };
}

interface Cuisine {
  data: { attributes: { Cuisine: string } };
}

interface Duration {
  data: { attributes: { Time: string } };
}

interface Requirement {
  data: { attributes: { Requirement: string } };
}
