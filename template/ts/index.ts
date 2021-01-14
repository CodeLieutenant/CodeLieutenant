import '../css/style.css';

const SNOW_EFFECT = 'snow-effect';

const SUBSCRIBE_FORM = 'subscribe-form';
const SUBSCRIBE_FORM_NAME = 'subscribe-form-name';
const SUBSCRIBE_FORM_EMAIL = 'subscribe-form-email';
const SUBSCRIBE_FORM_NAME_ERROR = 'subscribe-form-name-error';
const SUBSCRIBE_FORM_EMAIL_ERROR = 'subscribe-form-email-error';

const CONTACT_FORM = 'contact-form';
const CONTACT_FORM_NAME = 'contact-form-name';
const CONTACT_FORM_EMAIL = 'contact-form-email';
const CONTACT_FORM_SUBJECT = 'contact-form-subject';
const CONTACT_FORM_MESSAGE = 'contact-form-message';
const CONTACT_FORM_NAME_ERROR = 'contact-form-name-error';
const CONTACT_FORM_EMAIL_ERROR = 'contact-form-email-error';
const CONTACT_FORM_SUBJECT_ERROR = 'contact-form-subject-error';
const CONTACT_FORM_MESSAGE_ERROR = 'contact-form-message-error';


document.addEventListener('DOMContentLoaded', async () => {
  if (document.getElementById(SNOW_EFFECT)) {
    const { particlesSnowEffect } = await import('./snow');
    await particlesSnowEffect(SNOW_EFFECT);
  }

  const subscribeForm = document.getElementById(SUBSCRIBE_FORM);

  if (subscribeForm) {
    const { subscribeFormHandler } = await import('./subscribe');

    subscribeForm.addEventListener('submit', subscribeFormHandler(
      SUBSCRIBE_FORM_NAME,
      SUBSCRIBE_FORM_EMAIL,
      SUBSCRIBE_FORM_NAME_ERROR,
      SUBSCRIBE_FORM_EMAIL_ERROR
    ));
  }

  const contactForm = document.getElementById(CONTACT_FORM);

  if (contactForm) {
    const { contactFormHandler } = await import('./contact');

    contactForm.addEventListener('submit', contactFormHandler(
      CONTACT_FORM_NAME,
      CONTACT_FORM_NAME_ERROR,
      CONTACT_FORM_EMAIL,
      CONTACT_FORM_EMAIL_ERROR,
      CONTACT_FORM_SUBJECT,
      CONTACT_FORM_SUBJECT_ERROR,
      CONTACT_FORM_MESSAGE,
      CONTACT_FORM_MESSAGE_ERROR,
    ))
  }
});
