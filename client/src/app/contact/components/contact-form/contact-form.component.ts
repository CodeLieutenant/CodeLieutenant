import { Component, OnDestroy, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, FormControl } from '@angular/forms';
import { Validators } from 'angular-reactive-validation';

import { ContactService } from '../../../shared/services/contact.service';
import { Subscription } from 'rxjs';
import { ContactModel } from 'src/app/shared/models/contact.model';
import { SwalComponent } from '@sweetalert2/ngx-sweetalert2';
import { NetworkError } from 'src/app/shared/errors/network.error';

@Component({
  selector: 'app-contact-form',
  templateUrl: './contact-form.component.html',
  styleUrls: ['./contact-form.component.scss'],
})
export class ContactFormComponent implements OnDestroy {
  public contactForm: FormGroup;
  public sending: boolean = false;

  @ViewChild('dialog', { static: true })
  private contactDialog: SwalComponent;
  private contactSubscription: Subscription | null = null;

  public constructor(
    builder: FormBuilder,
    private contactService: ContactService
  ) {
    this.contactForm = builder.group({
      name: [
        null,
        [
          Validators.required('Name is required'),
          Validators.maxLength(50, 'Maximum length for name field exceeded'),
          Validators.pattern(
            /^[a-zA-Z]+$/,
            'Name can contain only alpha characters'
          ),
        ],
      ],
      email: [
        null,
        [
          Validators.required('Name is required'),
          Validators.email('Field is not valid email address'),
          Validators.maxLength(150, 'Maximum length for email field exceeded'),
        ],
      ],
      subject: [
        null,
        [
          Validators.required('Subject is required'),
          Validators.maxLength(
            150,
            'Maximum length for subject field exceeded'
          ),
        ],
      ],
      message: [
        null,
        [
          Validators.required('Message is required'),
          Validators.maxLength(
            500,
            'Maximum length for message field exceeded'
          ),
        ],
      ],
    });
  }

  public send() {
    if (!this.contactForm.valid || this.sending) {
      return;
    }
    this.sending = true;
    this.clearSubscription();
    this.contactSubscription = this.contactService
      .contact({
        name: this.name.value,
        email: this.email.value,
        subject: this.subject.value,
        message: this.message.value,
      })
      .subscribe(
        async (contactModel: ContactModel) => {
          await this.contactDialog.update({
            title: `${contactModel.name}, your message has been sent.`,
            text: `I will answer as soon as i can. Thank you`,
            icon: 'success',
          });
          await this.contactDialog.fire();
          this.contactForm.reset();
          this.sending = false;
        },
        async (error: NetworkError) => {
          if (error.type === 'validation') {
            this.contactForm.setErrors(error.validatioErrors);
          } else {
            await this.contactDialog.update({
              title: 'An error has occurred',
              text: error.message,
              icon: 'error',
            });
            await this.contactDialog.fire();
          }
          this.sending = false;
        }
      );
  }

  private clearSubscription() {
    if (this.contactSubscription !== null) {
      this.contactSubscription.unsubscribe();
    }
  }

  public ngOnDestroy(): void {
    this.clearSubscription();
  }

  public get name(): FormControl {
    return this.contactForm.get('name') as FormControl;
  }

  public get email(): FormControl {
    return this.contactForm.get('email') as FormControl;
  }

  public get subject(): FormControl {
    return this.contactForm.get('subject') as FormControl;
  }

  public get message(): FormControl {
    return this.contactForm.get('message') as FormControl;
  }
}
