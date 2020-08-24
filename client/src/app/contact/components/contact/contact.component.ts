import { Component, OnInit } from '@angular/core';
import { Meta, Title } from '@angular/platform-browser';
import { LoaderService } from 'src/app/core/services/loader.service';

@Component({
  selector: 'app-contact',
  templateUrl: './contact.component.html',
  styleUrls: ['./contact.component.scss'],
})
export class ContactComponent implements OnInit {
  public constructor(
    private titleService: Title,
    private loaderService: LoaderService
  ) {}

  public ngOnInit(): void {
    this.titleService.setTitle('Contact - Dušan Malušev');
    this.loaderService.hideLoaded();
  }
}
