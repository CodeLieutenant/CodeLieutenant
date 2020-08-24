export interface Env {
  baseUrl: string;
  production: boolean;
  recaptchaKey: string;
  googleMapsKey: string;
  api: string;
  apiBase: string;
}

export const ENVIRONMENT = 'environment';

export const environment: Env = {
  baseUrl: 'http://localhost:4200',
  production: false,
  recaptchaKey: '6Le3FcIZAAAAAEuqx3rtGyjkmLjfJz_QzBlfyfoT',
  googleMapsKey: 'AIzaSyCZvhYQF8uUody5Y1eTeB6zs_eIvE0tm7c',
  api: 'http://localhost:3000/api',
  apiBase: 'http://localhost:3000',
};

// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
