export interface Env {
  production: boolean;
  recaptchaKey: string;
  googleMapsKey: string;
  api: string;
  apiBase: string;
}

export const environment: Env = {
  production: false,
  recaptchaKey: '6Le3FcIZAAAAAEuqx3rtGyjkmLjfJz_QzBlfyfoT',
  googleMapsKey: '',
  api: 'http://localhost:3000/api',
  apiBase: 'http://localhost:3000',
};

// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
