interface ImportMetaEnv {
  readonly DEV?: string;
  readonly STRAPI_URL: string;
  readonly STRAPI_KEY: string;
  readonly PUBLIC_RECAPTCHA_CLIENT_KEY: string;
  readonly RECAPTCHA_SERVER_SECRET: string;
  readonly CONTACT_FORM_TO: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
