import { Component, OnInit, OnDestroy } from '@angular/core';
import { LoaderService } from '../../services/loader.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-loader',
  templateUrl: './loader.component.html',
  styleUrls: ['./loader.component.scss'],
})
export class LoaderComponent implements OnInit, OnDestroy {
  private loaderSubscription: Subscription = null;
  public showLoader = false;

  public constructor(private loaderService: LoaderService) {}

  public ngOnInit(): void {
    this.loaderSubscription = this.loaderService.subscribe({
      next: (value: boolean) => (this.showLoader = value),
    });
  }

  public ngOnDestroy(): void {
    if (this.loaderSubscription !== null) {
      this.loaderSubscription.unsubscribe();
    }
  }
}
