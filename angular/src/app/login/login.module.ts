import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { LoginRoutingModule } from './login-routing.module';

import { PersonalComponent } from '../components/personal/personal.component';
import { ChgmymsgComponent } from '../components/chgmymsg/chgmymsg.component';
import { UploadgoodsComponent } from '../components/uploadgoods/uploadgoods.component';

@NgModule({
  declarations: [
    PersonalComponent,
    ChgmymsgComponent,
    UploadgoodsComponent,
  ],
  imports: [
    CommonModule,
    FormsModule,
    LoginRoutingModule
  ]
})
export class LoginModule { }
