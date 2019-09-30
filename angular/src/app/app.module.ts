import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';

import { AppComponent } from './app.component';
import { HomepageComponent } from './homepage/homepage.component';
import { PersonalComponent } from './personal/personal.component';
import { NavigComponent } from './navig/navig.component';
import { FooterComponent } from './footer/footer.component';
import { AppRoutingModule } from './app-routing.module';
import { ChgmymsgComponent } from './chgmymsg/chgmymsg.component';
import { UploadgoodsComponent } from './uploadgoods/uploadgoods.component';
import { GoodspageComponent } from './goodspage/goodspage.component';
import { HttpClientModule }    from '@angular/common/http';
import { Personal2Component } from './personal2/personal2.component';
import { RegisterComponent } from './register/register.component';
import { FeedbackComponent } from './feedback/feedback.component';

@NgModule({
  declarations: [
    AppComponent,
    HomepageComponent,
    PersonalComponent,
    NavigComponent,
    FooterComponent,
    ChgmymsgComponent,
    UploadgoodsComponent,
    GoodspageComponent,
    Personal2Component,
    RegisterComponent,
    FeedbackComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
