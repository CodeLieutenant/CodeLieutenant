import { NgModule, PLATFORM_ID } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CommonModule, isPlatformBrowser } from '@angular/common';
import { AgmCoreModule } from '@agm/core';

import { ContactComponent } from './components/contact/contact.component';
import { CoreModule } from '../core/core.module';
import { ContactFormComponent } from './components/contact-form/contact-form.component';
import { MapComponent } from './components/map/map.component';
import { InfoComponent } from './components/info/info.component';
import { HeaderComponent } from './components/header/header.component';

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
    CoreModule,
    RouterModule.forChild(routes),
    AgmCoreModule,
  ],
  exports: [RouterModule],
})
export class ContactModule {}
