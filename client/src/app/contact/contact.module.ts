import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule } from '@angular/common';
import { AgmCoreModule } from '@agm/core';

import { ContactComponent } from './components/contact/contact.component';
import { ContactFormComponent } from './components/contact-form/contact-form.component';
import { MapComponent } from './components/map/map.component';
import { InfoComponent } from './components/info/info.component';
import { HeaderComponent } from './components/header/header.component';
import { ReactiveFormsModule } from '@angular/forms';
import { ReactiveValidationModule } from 'angular-reactive-validation';
import { SharedModule } from '../shared/shared.module';
import { SweetAlert2Module } from '@sweetalert2/ngx-sweetalert2';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: ContactComponent,
  },
];

@NgModule({
  declarations: [
    ContactComponent,
    ContactFormComponent,
    MapComponent,
    InfoComponent,
    HeaderComponent,
  ],
  imports: [
    CommonModule,
    SharedModule,
    ReactiveFormsModule,
    SweetAlert2Module,
    ReactiveValidationModule,
    RouterModule.forChild(routes),
    AgmCoreModule,
  ],
  exports: [RouterModule],
})
export class ContactModule {}
