import { Component, PLATFORM_ID, Inject } from '@angular/core';
import { isPlatformBrowser } from '@angular/common';

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.scss'],
})
export class MapComponent {
  public lat = 43;
  public lng = 21;

  public constructor(@Inject(PLATFORM_ID) private platformId: string) {}

  public get isBrowser(): boolean {
    return isPlatformBrowser(this.platformId);
  }
}
