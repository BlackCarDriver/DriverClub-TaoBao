import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { GoodspageComponent } from '../components/goodspage/goodspage.component';
import {Personal2Component} from '../components/personal2/personal2.component';


const routes: Routes = [
  { path: 'goodsdetail/:gid' , component:GoodspageComponent },
  { path: 'personals/:uid' , component :Personal2Component},
  { path: '', redirectTo: 'goodsdetail', pathMatch: 'full' },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class MessageRoutingModule { }
