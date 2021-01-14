import { object, string, ValidationError } from 'yup';
import { Err } from './error';
import { http } from './http';

const Swal = require('sweetalert2');


interface SubscriptionValidationError {
  nameError: string;
  emailError: string;
}

interface SubscriptionDTO {
  name: string;
  email: string;
}

interface Subscription {
  id: number;
  name: string;
  email: string;
  createdAt: Date;
}

const schema = object().shape({
  name: string().required().max(50),
  email: string().required().email().max(150),
});

const subscribe = async (dto: SubscriptionDTO): Promise<Subscription | Err | SubscriptionValidationError> => {
  try {

    await schema.validate(dto, { recursive: true, abortEarly: false });

    const res = await http('/subscribe', 'POST', dto);

    const data = await res.json();

    return {
      id: data.id,
      name: data.name,
      email: data.email,
      createdAt: new Date(data.createdAt),
    };
  } catch (err) {
    if (err instanceof ValidationError) {
      const validationError: SubscriptionValidationError = { nameError: '', emailError: '' };
      err.inner.forEach((item) => {
        if (item.path === 'name') {
          validationError.nameError = item.errors[0];
        } else if (item.path === 'email') {
          validationError.emailError = item.errors[0];
        }
      });

      return validationError;
    }
    return {
      message: 'Try again please',
    };
  }
};

const subscribeFormHandler = (
  nameEl: string,
  emailEl: string,
  nameErrorEl: string,
  emailErrorEl: string
) => async (e: Event) => {
  e.preventDefault();
  const name = document.getElementById(nameEl);
  const email = document.getElementById(emailEl);
  const nameError = document.getElementById(nameErrorEl);
  const emailError = document.getElementById(emailErrorEl);

  //@ts-ignore
  const res = await subscribe({ name: name.value, email: email.value });
  // Server error
  if ('message' in res) {
    Swal.fire({
      title: 'Error',
      text: res.message,
      icon: 'error',
      timerProgressBar: true,
    });
    return;
  }

  // Validation Error
  if ('nameError' in res && 'emailError' in res) {
    if (res.nameError !== '') {
      name.classList.add('input-error');

      nameError.classList.remove('hidden');
      nameError.innerText = res.nameError;
    }

    if (res.emailError !== '') {
      email.classList.add('input-error');
      emailError.classList.remove('hidden');
      emailError.innerText = res.emailError;
    }

    setTimeout(() => {
      nameError.classList.add('hidden');
      emailError.classList.add('hidden');
      email.classList.remove('input-error');
      name.classList.remove('input-error');
    }, 4000);
    return;
  }

  //@ts-ignore
  gtag('event', 'subscribe', {
    event_category: 'subscription',
    event_label: 'New user subscribed to news letters',
  })

  Swal.fire({
    title: 'Success',
    text: 'You have successfully subscribed to newsletters',
    icon: 'success',
    timerProgressBar: true,
  });
};

export { subscribeFormHandler };
