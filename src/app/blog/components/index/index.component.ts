import { Component, OnInit } from '@angular/core';
import { LoaderService } from 'src/app/core/services/loader.service';

@Component({
  selector: 'app-index',
  templateUrl: './index.component.html',
  styleUrls: ['./index.component.scss'],
})
export class IndexComponent implements OnInit {
  constructor(private loaderService: LoaderService) {}

  public ngOnInit(): void {
    this.loaderService.hideLoaded();
  }
}
