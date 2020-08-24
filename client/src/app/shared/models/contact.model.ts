export interface ContactModel {
  id: number;
  name: string;
  email: string;
  subject: string;
  message: string;
  createdAt: Date | string;
}
