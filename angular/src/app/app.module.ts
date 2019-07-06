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
import { AboutComponent } from './about/about.component';
import { CharthpComponent } from './charthp/charthp.component';
import { Personal2Component } from './personal2/personal2.component';

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
    AboutComponent,
    CharthpComponent,
    Personal2Component
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
