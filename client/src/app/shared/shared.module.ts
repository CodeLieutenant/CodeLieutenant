import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { UrlInterceptor } from './interceptors/url.interceptor';
import { SubscriptionService } from './services/subscription.service';
import { ContactService } from './services/contact.service';

import { HTTP_INTERCEPTORS } from '@angular/common/http';

@NgModule({
  declarations: [],
  imports: [CommonModule, FontAwesomeModule, HttpClientModule],
  exports: [FontAwesomeModule, HttpClientModule],
  providers: [
    SubscriptionService,
    ContactService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: UrlInterceptor,
      multi: true,
    },
  ],
})
export class SharedModule {}
