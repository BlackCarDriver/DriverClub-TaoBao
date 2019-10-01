import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MessageRoutingModule } from './message-routing.module';

import { GoodspageComponent } from '../components/goodspage/goodspage.component';
import {Personal2Component} from '../components/personal2/personal2.component';


@NgModule({
  declarations: [
    GoodspageComponent,
    Personal2Component,
  ],
  imports: [
    CommonModule,
    MessageRoutingModule
  ]
})
export class MessageModule { }
