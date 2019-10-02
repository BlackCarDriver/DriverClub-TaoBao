import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule }    from '@angular/common/http';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './app-routing.module';

import { NavigComponent } from './components/navig/navig.component';
import { FooterComponent } from './components/footer/footer.component';

import { HomepageComponent } from './components/homepage/homepage.component';
import { Notfound1Component } from './components/notfound1/notfound1.component';

@NgModule({
  declarations: [
    AppComponent,
    HomepageComponent,
    NavigComponent,
    FooterComponent,
    Notfound1Component,
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
