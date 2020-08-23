import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HttpClientModule } from '@angular/common/http';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';

import { SubscriptionService } from './services/subscription.service';

@NgModule({
  declarations: [],
  imports: [CommonModule, FontAwesomeModule, HttpClientModule],
  exports: [FontAwesomeModule, HttpClientModule],
  providers: [SubscriptionService],
})
export class SharedModule {}
