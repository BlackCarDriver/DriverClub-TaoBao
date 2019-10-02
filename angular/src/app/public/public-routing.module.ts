import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import {RegisterComponent} from '../components/register/register.component';
import {FeedbackComponent} from '../components/feedback/feedback.component';
import {AboutComponent} from '../components/about/about.component';

const routes: Routes = [
  { path: 'signup', component: RegisterComponent},
  { path: 'resetpassword', component: RegisterComponent},
  { path: 'feedback' , component :FeedbackComponent},
  { path: 'about' , component :AboutComponent},
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PublicRoutingModule { }
