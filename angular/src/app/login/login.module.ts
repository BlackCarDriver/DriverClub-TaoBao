import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { LoginRoutingModule } from './login-routing.module';

import { PersonalComponent } from '../components/personal/personal.component';
// import { UploadgoodsComponent } from '../components/uploadgoods/uploadgoods.component';
import { ChgmymsgComponent } from '../components/chgmymsg/chgmymsg.component';

@NgModule({
  declarations: [
    PersonalComponent,
    ChgmymsgComponent,
    // UploadgoodsComponent,
  ],
  imports: [
    CommonModule,
    LoginRoutingModule
  ]
})
export class LoginModule { }
