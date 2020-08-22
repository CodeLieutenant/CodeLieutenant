import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { IndexComponent } from './components/index/index.component';
import { HeaderComponent } from './components/header/header.component';
import { PaginationComponent } from './components/pagination/pagination.component';
import { SearchComponent } from './components/search/search.component';
import { CategoriesComponent } from './components/categories/categories.component';
import { PopularComponent } from './components/popular/popular.component';
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
  declarations: [IndexComponent, HeaderComponent, PaginationComponent, SearchComponent, CategoriesComponent, PopularComponent, EntryComponent, BlogComponent],
  imports: [CommonModule, RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class BlogModule {}
