import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { switchMap, catchError } from 'rxjs/operators';
import { ReCaptchaV3Service } from 'ng-recaptcha';
import { Observable, of } from 'rxjs';

import { SubscriptionModel } from '../models/subscription.model';

export interface Subscription {
  email: string;
}

export type SubError =
  | 'recaptcha'
  | 'not-found'
  | 'validation'
  | 'server'
  | null;

export class SubscriptionError {
  public type: SubError;
  public message: string = '';
  public validatioErrors: { [key: string]: string } | null;
}

@Injectable()
export class SubscriptionService {
  public constructor(
    private recaptchaService: ReCaptchaV3Service,
    private httpClient: HttpClient
  ) {}

  public subscribe(sub: Subscription): Observable<SubscriptionModel> {
    return this.recaptchaService.execute('subscirbe').pipe(
      switchMap((recaptcha: string) =>
        this.httpClient.post<SubscriptionModel>('/subscribe', {
          ...sub,
          recaptcha,
        })
      ),
      catchError((httpError: HttpErrorResponse) => {
        console.log('I am HERE');
        let err = new SubscriptionError();

        console.log('Here');
        switch (httpError.status) {
          case 400:
            err.type = 'recaptcha';
            err.message = 'Invalid ReCAPTCHA. Plaase try again.';
            break;
          case 404:
            console.log('Here');
            err.type = 'not-found';
            err.message = 'Server is not responding. Plaase try again later.';
            break;
          case 422:
            err.type = 'validation';
            err.validatioErrors = JSON.parse(httpError.error).errors;
            break;
          default:
            err.message = 'An error has occurred';
            err.type = 'server';
            break;
        }

        console.log(err.message);
        throw err;
      })
    );
  }
}
