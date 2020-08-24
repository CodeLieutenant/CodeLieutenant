import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';

import { PopularArticlesModule } from '../popular-articles/popular-articles.module';
import { IndexComponent } from './components/index/index.component';
import { HeaderComponent } from './components/header/header.component';
import { PaginationComponent } from './components/pagination/pagination.component';
import { SearchComponent } from './components/search/search.component';
import { CategoriesComponent } from './components/categories/categories.component';
import { EntryComponent } from './components/entry/entry.component';
import { BlogComponent } from './components/blog/blog.component';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    component: IndexComponent,
  },
];

@NgModule({
  declarations: [
    IndexComponent,
    HeaderComponent,
    PaginationComponent,
    SearchComponent,
    CategoriesComponent,
    EntryComponent,
    BlogComponent,
  ],
  imports: [CommonModule, RouterModule.forChild(routes), PopularArticlesModule],
  exports: [RouterModule],
})
export class BlogModule {}
