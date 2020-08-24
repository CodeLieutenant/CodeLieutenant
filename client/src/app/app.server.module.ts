import { NgModule } from '@angular/core';
import { ServerModule } from '@angular/platform-server';

import { CoreModule } from './core/core.module';

import { AppModule } from './app.module';
import { AppComponent } from './app.component';

@NgModule({
  imports: [AppModule, ServerModule, CoreModule],
  bootstrap: [AppComponent],
})
export class AppServerModule {}
