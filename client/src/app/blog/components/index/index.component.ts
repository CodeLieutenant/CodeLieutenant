import { Component, OnInit, Inject } from '@angular/core';
import { LoaderService } from 'src/app/core/services/loader.service';
import { Meta } from '@angular/platform-browser';

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.scss'],
})
export class IndexComponent implements OnInit {
  constructor(
    private loaderService: LoaderService,
    private metaService: Meta
  ) {}

  public ngOnInit(): void {
    this.metaService.addTags(
      [
        { name: 'og:title', content: 'Blog - Dusan Malusev' },
        {
          name: 'keywords',
          content:
            'blog,development,database,nosql,sql,web development,software engineer,php,go,crytography',
        },
      ],
      true
    );
    this.loaderService.hideLoaded();
  }
}
