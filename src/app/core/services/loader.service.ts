import { Injectable } from '@angular/core';
import { BehaviorSubject, Subscription, PartialObserver } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class LoaderService {
  private displayLoader$: BehaviorSubject<boolean> = new BehaviorSubject(true);

  public displayLoader() {
    this.displayLoader$.next(true);
  }

  public hideLoaded() {
    this.displayLoader$.next(false);
  }

  public subscribe(observer: PartialObserver<boolean>): Subscription {
    return this.displayLoader$.subscribe(observer);
  }
}
