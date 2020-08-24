import { Component, OnInit, Inject } from '@angular/core';
import { Title, Meta } from '@angular/platform-browser';

import { ENVIRONMENT, Env } from 'src/environments/environment';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  public title: string = 'Dušan Malušev - Software Developer';

  public constructor(
    private titleService: Title,
    private metaService: Meta,
    @Inject(ENVIRONMENT) private environment: Env
  ) {}

  public ngOnInit(): void {
    this.titleService.setTitle(this.title);
    this.metaService.addTags([
      {
        name: 'keywords',
        content:
          'developer,web developer,software engineer,student,php,go,brossquad,BrosSquad,crytography',
      },
      {
        name: 'description',
        content:
          'Open Source enthusiast currenlty working for Nano Interactive as backend developer with Phalcon Framework. PHP and Go developer',
      },
      { name: 'og:site_name', content: 'Dusan Malusev - Software Developer' },
      { name: 'og:url', content: '' },
      { name: 'og:title', content: 'Dusan Malusev - Software Developer' },
      {
        name: 'og:description',
        content:
          'Open Source enthusiast currenlty working for Nano Interactive as backend developer with Phalcon Framework. PHP and Go developer',
      },
      {
        name: 'og:image',
        content: `${this.environment.baseUrl}/malusev.png`,
      },
    ]);
  }
}
