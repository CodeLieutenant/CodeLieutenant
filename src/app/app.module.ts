import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { AgmCoreModule } from '@agm/core';
import { RECAPTCHA_V3_SITE_KEY, RecaptchaV3Module } from 'ng-recaptcha';
import { ReactiveValidationModule } from 'angular-reactive-validation';
import { SweetAlert2Module } from '@sweetalert2/ngx-sweetalert2';

import { environment, ENVIRONMENT } from '../environments/environment';
import { CoreModule } from './core/core.module';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [AppComponent],
  imports: [
    BrowserModule.withServerTransition({ appId: 'serverApp' }),
    AppRoutingModule,
    CoreModule,
    ReactiveFormsModule,
    ReactiveValidationModule,
    RecaptchaV3Module,
    SweetAlert2Module.forRoot(),
    AgmCoreModule.forRoot({
      apiKey: environment.googleMapsKey,
    }),
  ],
  providers: [
    {
      provide: RECAPTCHA_V3_SITE_KEY,
      useValue: environment.recaptchaKey,
    },
    {
      provide: ENVIRONMENT,
      useValue: environment,
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
