import type { StrapiMedia, StrapiMediaCollection } from "./Strapi";

export interface Home {
  CTA: string;
  Logo: StrapiMedia;
  ImageCarousel: StrapiMediaCollection;
}
