import { Component, OnInit } from '@angular/core';
import { LoaderService } from 'src/app/core/services/loader.service';

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.scss'],
})
export class IndexComponent implements OnInit {
  public social: Array<{ url: string; icon: string }> = [
    {
      url: 'https://github.com/malusev998',
      icon: 'fab fa-github',
    },
    {
      url: 'https://github.com/malusev998',
      icon: 'fab fa-facebook',
    },
    {
      url: 'https://github.com/malusev998',
      icon: 'fab fa-instagram',
    },
    {
      url: 'https://github.com/malusev998',
      icon: 'fab fa-stack-overflow',
    },
  ];

  public constructor(private loaderService: LoaderService) {}

  public ngOnInit(): void {
    this.loaderService.hideLoaded();
  }
}
