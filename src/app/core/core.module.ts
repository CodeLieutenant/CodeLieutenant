import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ReactiveFormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';

import { ReactiveValidationModule } from 'angular-reactive-validation';
import { SharedModule } from '../shared/shared.module';

import { LoaderComponent } from './components/loader/loader.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { SubscribeComponent } from './components/subscribe/subscribe.component';
import { FooterComponent } from './components/footer/footer.component';

@NgModule({
  declarations: [
    LoaderComponent,
    SidebarComponent,
    NavbarComponent,
    SubscribeComponent,
    FooterComponent,
  ],
  imports: [
    CommonModule,
    SharedModule,
    RouterModule,
    ReactiveFormsModule,
    ReactiveValidationModule,
  ],
  exports: [LoaderComponent, SidebarComponent, FooterComponent],
})
export class CoreModule {}
