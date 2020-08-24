import { Component } from '@angular/core';

export interface NavItem {
  name: string;
  path: string;
}

export type Navbar = Array<NavItem>;

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.scss'],
})
export class NavbarComponent {
  public routes: Navbar = [
    {
      name: 'Home',
      path: '/',
    },
    {
      name: 'Blog',
      path: '/blog',
    },
    {
      name: 'Contact',
      path: '/contact',
    },
    {
      name: 'Login',
      path: '/auth/login',
    },
    {
      name: 'Register',
      path: '/auth/register',
    },
  ];
}
