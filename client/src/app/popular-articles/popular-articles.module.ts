import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ArticlesComponent } from './components/articles/articles.component';
import { ArticleComponent } from './components/article/article.component';
import { SharedModule } from '../shared/shared.module';

@NgModule({
  declarations: [ArticlesComponent, ArticleComponent],
  imports: [CommonModule, SharedModule],
  exports: [ArticlesComponent],
})
export class PopularArticlesModule {}
