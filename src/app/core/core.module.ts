import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { LoaderComponent } from './components/loader/loader.component';
import { SidebarComponent } from './components/sidebar/sidebar.component';
import { NavbarComponent } from './components/navbar/navbar.component';
import { SubscribeComponent } from './components/subscribe/subscribe.component';
import { FooterComponent } from './components/footer/footer.component';
import { RouterModule } from '@angular/router';

@NgModule({
  declarations: [
    LoaderComponent,
    SidebarComponent,
    NavbarComponent,
    SubscribeComponent,
    FooterComponent,
  ],
  imports: [CommonModule, RouterModule],
  exports: [LoaderComponent, SidebarComponent, FooterComponent],
})
export class CoreModule {}
