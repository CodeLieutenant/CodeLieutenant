import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { switchMap, catchError } from 'rxjs/operators';
import { ReCaptchaV3Service } from 'ng-recaptcha';
import { Observable, of } from 'rxjs';

import { SubscriptionModel } from '../models/subscription.model';
import { throwNetworkError } from '../errors/network.error';

export interface Subscription {
  email: string;
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
      catchError((httpError: HttpErrorResponse) =>
        throwNetworkError(httpError.status, httpError.error)
      )
    );
  }
}
