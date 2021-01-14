import { ValidationError, object, string } from "yup";
import { Err } from "./error";
import { http } from "./http";

const Swal = require('sweetalert2');


interface ContactDto {
    name: string;
    email: string;
    subject: string;
    message: string;
}

interface Contact {
    ID: number;
    name: string;
    email: string;
    subject: string;
    text: string;
    createdAt: Date;
}

interface ContactValidationError {
    nameError: string;
    emailError: string;
    subjectError: string;
    messageError: string;
}

const schema = object().shape({
    name: string().required().max(50),
    email: string().required().email().max(150),
    subject: string().required().max(150),
    message: string().required().max(2000),
});


const contact = async (dto: ContactDto): Promise<Contact | Err | ContactValidationError> => {
    try {

        await schema.validate(dto, { recursive: true, abortEarly: false });

        const res = await http('/contact', 'POST', dto);

        const data = await res.json();

        return {
            ID: data.id,
            name: data.name,
            email: data.email,
            subject: data.subject,
            text: data.message,
            createdAt: new Date(data.createdAt),
        }
    } catch (err) {
        if (err instanceof ValidationError) {
            const validationError: ContactValidationError = {
                nameError: '',
                emailError: '',
                messageError: '',
                subjectError: '',
            };

            err.inner.forEach((item) => {
                if (item.path === 'name') {
                    validationError.nameError = item.errors[0];
                } else if (item.path === 'email') {
                    validationError.emailError = item.errors[0];
                } else if (item.path === 'subject') {
                    validationError.subjectError = item.errors[0];
                } else if (item.path === 'message') {
                    validationError.messageError = item.errors[0];
                }
            });

            return validationError;
        }
        return {
            message: 'Try again please',
        };
    }

}

const contactFormHandler = (
    nameEl: string,
    nameErrorEl: string,
    emailEl: string,
    emailErrorEl: string,
    subjectEl: string,
    subjectErrorEl: string,
    messageEl: string,
    messageErrorEl: string
) => async (e: Event) => {
    e.preventDefault();

    const name = document.getElementById(nameEl);
    const email = document.getElementById(emailEl);
    const subject = document.getElementById(subjectEl);
    const message = document.getElementById(messageEl);
    const nameError = document.getElementById(nameErrorEl);
    const emailError = document.getElementById(emailErrorEl);
    const subjectError = document.getElementById(subjectErrorEl);
    const messageError = document.getElementById(messageErrorEl);


    //@ts-ignore
    const res = await contact({ name: name.value, email: email.value, message: subject.value, subject: message.value });

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
    if ('nameError' in res && 'emailError' in res && 'subjectError' in res && 'messageError' in res) {
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

        if (res.subjectError !== '') {
            subject.classList.add('input-error');
            subjectError.classList.remove('hidden');
            subjectError.innerText = res.subjectError;
        }

        if (res.messageError !== '') {
            message.classList.add('input-error');
            messageError.classList.remove('hidden');
            messageError.innerText = res.messageError;
        }


        setTimeout(() => {
            nameError.classList.add('hidden');
            emailError.classList.add('hidden');
            subjectError.classList.add('hidden');
            messageError.classList.add('hidden');
            email.classList.remove('input-error');
            name.classList.remove('input-error');
            subject.classList.remove('input-error');
            message.classList.remove('input-error');
        }, 4000);
        return;
    }

    //@ts-ignore
    gtag('event', 'contact', {
        event_category: 'reach',
        event_label: 'New contact email',
    });

    Swal.fire({
        title: 'Success',
        text: 'You have successfully sent message',
        icon: 'success',
        timerProgressBar: true,
    });
}


export { contact, contactFormHandler };
