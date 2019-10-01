import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { PersonalComponent } from '../components/personal/personal.component';
import { ChgmymsgComponent } from '../components/chgmymsg/chgmymsg.component';
// import { UploadgoodsComponent } from '../components/uploadgoods/uploadgoods.component';

const routes: Routes = [
  { path: 'myhome', component: PersonalComponent },
  { path: 'change', component: ChgmymsgComponent },
  // { path: 'upload', component: UploadgoodsComponent },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class LoginRoutingModule { }
 