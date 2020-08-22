import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RECAPTCHA_V3_SITE_KEY, RecaptchaV3Module } from 'ng-recaptcha';

import { environment } from '../environments/environment';
import { CoreModule } from './core/core.module';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule.withServerTransition({ appId: 'serverApp' }),
    AppRoutingModule,
    CoreModule,
    RecaptchaV3Module,
  ],
  providers: [
    {
      provide: RECAPTCHA_V3_SITE_KEY,
      useValue: '',
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
