export type NetErrorType =
  | 'recaptcha'
  | 'not-found'
  | 'validation'
  | 'server'
  | 'unauthorized'
  | 'forbidden'
  | 'too-many-requests'
  | null;

export class NetworkError {
  public type: NetErrorType;
  public message = '';
  public validatioErrors: { [key: string]: string | string[] } | null;
}

export const throwNetworkError = (status: number, error: any | null) => {
  const err = new NetworkError();

  switch (status) {
    case 400:
      err.type = 'recaptcha';
      err.message = 'Invalid ReCAPTCHA. Plaase try again.';
      break;
    case 401:
      err.type = 'unauthorized';
      err.message = 'You need to login before you do this action.';
      break;
    case 403:
      err.type = 'forbidden';
      err.message = 'You dont have permission for this actions';
      break;
    case 404:
      err.type = 'not-found';
      err.message = 'Server is not responding. Plaase try again later.';
      break;
    case 422:
      err.type = 'validation';
      err.validatioErrors = JSON.parse(error).errors;
      break;
    case 429:
      err.type = 'too-many-requests';
      err.message = JSON.parse(error).message;
      break;
    default:
      err.message = 'An error has occurred';
      err.type = 'server';
      break;
  }

  throw err;
};
