import { Component, OnInit } from '@angular/core';
import {
  faFacebook,
  faInstagram,
  faGithub,
  faStackOverflow,
  faStackExchange,
  IconDefinition,
} from '@fortawesome/free-brands-svg-icons';

import { LoaderService } from 'src/app/core/services/loader.service';

interface SocialIcon {
  url: string;
  icon: IconDefinition;
}

type Icons = Array<SocialIcon>;

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.scss'],
})
export class IndexComponent implements OnInit {
  public social: Icons = [
    {
      url: 'https://github.com/malusev998',
      icon: faGithub,
    },
    {
      url: 'https://www.facebook.com/dmalusev/',
      icon: faFacebook,
    },
    {
      url: 'https://www.instagram.com/dusanmalusev/',
      icon: faInstagram,
    },
    {
      url: 'https://stackoverflow.com/users/8411483/dusan-malusevz',
      icon: faStackOverflow,
    },
    {
      url: 'https://stackexchange.com/users/11475985/dusan-malusev',
      icon: faStackExchange,
    },
  ];

  public constructor(private loaderService: LoaderService) {}

  public ngOnInit(): void {
    this.loaderService.hideLoaded();
  }
}
