import { Inject, Injectable } from '@angular/core';
import {
  HttpEvent,
  HttpInterceptor,
  HttpHandler,
  HttpRequest,
} from '@angular/common/http';
import { Observable } from 'rxjs';

import { Env, ENVIRONMENT } from '../../../environments/environment';

@Injectable()
export class UrlInterceptor implements HttpInterceptor {
  public constructor(@Inject(ENVIRONMENT) private env: Env) {}

  public intercept(
    req: HttpRequest<any>,
    next: HttpHandler
  ): Observable<HttpEvent<any>> {
    if (req.url.startsWith('http')) {
      return next.handle(req);
    }

    const clone = req.clone({
      url: `${this.env.api}${req.url}`,
    });

    return next.handle(clone);
  }
}
