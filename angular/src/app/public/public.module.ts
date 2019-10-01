import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { PublicRoutingModule } from './public-routing.module';

import {RegisterComponent} from '../components/register/register.component';
import {FeedbackComponent} from '../components/feedback/feedback.component';


@NgModule({
  declarations: [
    FeedbackComponent,
    RegisterComponent,
  ],
  imports: [
    CommonModule,
    PublicRoutingModule
  ]
})
export class PublicModule { }
