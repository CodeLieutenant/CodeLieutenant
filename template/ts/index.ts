import '../css/style.css';

const SNOW_EFFECT = 'snow-effect';
const SUBSCRIBE_FORM = 'subscribe-form';
const SUBSCRIBE_FORM_NAME = 'subscribe-form-name';
const SUBSCRIBE_FORM_NAME_ERROR = 'subscribe-form-name-error';
const SUBSCRIBE_FORM_EMAIL = 'subscribe-form-email';
const SUBSCRIBE_FORM_EMAIL_ERROR = 'subscribe-form-email-error';

document.addEventListener('DOMContentLoaded', async () => {
  if (document.getElementById(SNOW_EFFECT)) {
    const { particlesSnowEffect } = await import('./snow');
    await particlesSnowEffect(SNOW_EFFECT);
  }

  const subscribeForm = document.getElementById(SUBSCRIBE_FORM);

  if (subscribeForm) {
    const { subscribeFormHandler } = await import('./subscribe');
    subscribeForm.addEventListener(
      'submit',
      subscribeFormHandler(
        SUBSCRIBE_FORM_NAME,
        SUBSCRIBE_FORM_EMAIL,
        SUBSCRIBE_FORM_NAME_ERROR,
        SUBSCRIBE_FORM_EMAIL_ERROR
      )
    );
  }
});
