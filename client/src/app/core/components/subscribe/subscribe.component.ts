import { Component, OnDestroy, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { Validators } from 'angular-reactive-validation';
import { SwalComponent } from '@sweetalert2/ngx-sweetalert2';
import { SubscriptionService } from 'src/app/shared/services/subscription.service';
import { Subscription } from 'rxjs';
import { SubscriptionModel } from 'src/app/shared/models/subscription.model';
import { NetworkError } from 'src/app/shared/errors/network.error';

@Component({
  selector: 'app-subscribe',
  templateUrl: './subscribe.component.html',
  styleUrls: ['./subscribe.component.scss'],
})
export class SubscribeComponent implements OnDestroy {
  @ViewChild('dialog', { static: true })
  private dialog: SwalComponent;
  private subscription: Subscription | null = null;
  private dialogSubscription: Subscription | null = null;

  public sending = false;
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
    this.clearSubscriptions();
    this.sending = true;
    this.subscription = this.subService
      .subscribe({ email: this.email.value })
      .subscribe(
        async (value: SubscriptionModel) => {
          await this.dialog.update({
            title: 'You have subscribed successfully',
            text: `Welcome ${value.email}. You have been added to my subscription list. No spam I promise.`,
            icon: 'success',
          });
          await this.dialog.fire();
          this.subscribeForm.reset();

          this.sending = false;
        },
        async (error: NetworkError) => {
          if (error.type === 'validation') {
            this.subscribeForm.setErrors(error.validatioErrors);
          } else {
            await this.dialog.update({
              title: 'An error has occurred',
              text: error.message,
              icon: 'error',
            });
            await this.dialog.fire();
          }
          this.sending = false;
        }
      );
  }

  public get email(): FormControl {
    return this.subscribeForm.get('email') as FormControl;
  }

  public ngOnDestroy(): void {
    this.clearSubscriptions();
  }

  private clearSubscriptions(): void {
    if (this.subscription !== null) {
      this.subscription.unsubscribe();
    }

    if (this.dialogSubscription !== null) {
      this.dialogSubscription.unsubscribe();
    }
  }
}
