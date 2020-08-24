import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { ReCaptchaV3Service } from 'ng-recaptcha';
import { Observable } from 'rxjs';
import { switchMap, catchError } from 'rxjs/operators';

import { ContactModel } from '../models/contact.model';
import { throwNetworkError } from '../errors/network.error';

export interface Contact {
  name: string;
  email: string;
  subject: string;
  message: string;
}

@Injectable()
export class ContactService {
  public constructor(
    private recaptchaService: ReCaptchaV3Service,
    private httpClient: HttpClient
  ) {}

  public contact(contact: Contact): Observable<ContactModel> {
    return this.recaptchaService.execute('contact').pipe(
      switchMap((recaptcha: string) =>
        this.httpClient.post<ContactModel>('/contact', {
          ...contact,
          recaptcha,
        })
      ),
      catchError((httpError: HttpErrorResponse) =>
        throwNetworkError(httpError.status, httpError.error)
      )
    );
  }
}
