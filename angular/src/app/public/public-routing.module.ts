import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {RegisterComponent} from '../components/register/register.component';
import {FeedbackComponent} from '../components/feedback/feedback.component';

const routes: Routes = [
  { path: 'signup', component: RegisterComponent},
  { path: 'resetpassword', component: RegisterComponent},
  { path: 'feedback' , component :FeedbackComponent},
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PublicRoutingModule { }
