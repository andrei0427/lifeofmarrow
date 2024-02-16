export interface StrapiMedia {
  data: { attributes: StrapiMediaType };
}

export interface StrapiMediaCollection {
  data: { attributes: StrapiMediaType }[];
}

export interface StrapiMediaType {
  name: string;
  width: number;
  height: number;
  url: string;
}

export interface CTA {
  Title: string;
  Link: string;
}

export interface ContentSection {
  Title: string;
  Image: StrapiMedia;
  Text: string;
  CTALink?: CTA;
}
