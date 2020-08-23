import { Component, OnInit, OnDestroy } from '@angular/core';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { Validators } from 'angular-reactive-validation';
import { SubscriptionService } from 'src/app/shared/services/subscription.service';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-subscribe',
  templateUrl: './subscribe.component.html',
  styleUrls: ['./subscribe.component.scss'],
})
export class SubscribeComponent implements OnDestroy {
  private subscription: Subscription | null = null;

  public sending: boolean = false;
  public subscribeForm: FormGroup;

  constructor(private subService: SubscriptionService, builder: FormBuilder) {
    this.subscribeForm = builder.group({
      email: [
        null,
        [
          Validators.required('Email is required'),
          Validators.email('Input is not valid email'),
        ],
      ],
    });
  }

  public subscribe(): void {
    if (!this.subscribeForm.valid || this.sending) {
      return;
    }
    this.clearSubscription();
    this.sending = true;
    this.subscription = this.subService
      .subscribe({ email: this.email.value })
      .subscribe(
        (value) => {
          this.sending = false;
        },
        (error) => {
          console.log(error);
          this.sending = false;
        }
      );
  }

  public get email(): FormControl {
    return this.subscribeForm.get('email') as FormControl;
  }

  public ngOnDestroy(): void {
    this.clearSubscription();
  }

  private clearSubscription() {
    if (this.subscription !== null) {
      this.subscription.unsubscribe();
    }
  }
}
