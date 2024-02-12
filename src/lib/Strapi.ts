interface Props {
  endpoint: string;
  query?: Record<string, string> | string;
  wrappedByKey?: string;
  wrappedByList?: boolean;
}

export default async function fetchApi<T>({
  endpoint,
  query,
  wrappedByKey,
}: Props): Promise<T> {
  if (endpoint.startsWith("/")) {
    endpoint = endpoint.slice(1);
  }

  const url = new URL(`${import.meta.env.STRAPI_URL}/api/${endpoint}`);

  if (query) {
    Object.entries(query).forEach(([key, value]) => {
      url.searchParams.append(key, value);
    });

    url.search = decodeURIComponent(url.search);
  }

  const res = await fetch(url.toString(), {
    headers: { Authorization: `Bearer ${import.meta.env.STRAPI_KEY}` },
  });

  const json = await res.json();

  if (json["error"]) {
    throw new Error(
      `Strapi Error: ${json["error"].status} - ${json["error"].message}`
    );
  }

  let data = json["data"];

  if (wrappedByKey) {
    data = data[wrappedByKey];
  }

  if (Array.isArray(data)) {
    data = data.map((d) => d.attributes);
  }

  return data as T;
}

export function getStrapiMediaUrl(url: string | undefined | null) {
  if (!url) {
    return "/hb-mobile.jpeg";
  }

  return `${import.meta.env.STRAPI_URL}${url}`;
}
