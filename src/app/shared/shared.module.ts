import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { SubscriptionService } from './services/subscription.service';
import { ContactService } from './services/contact.service';

@NgModule({
  declarations: [],
  imports: [CommonModule, FontAwesomeModule, HttpClientModule],
  exports: [FontAwesomeModule, HttpClientModule],
  providers: [SubscriptionService, ContactService],
})
export class SharedModule {}
