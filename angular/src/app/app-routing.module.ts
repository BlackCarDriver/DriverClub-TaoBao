import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { HomepageComponent } from '../app/homepage/homepage.component';
import { PersonalComponent} from '../app/personal/personal.component';
import { ChgmymsgComponent } from '../app/chgmymsg/chgmymsg.component';
import { UploadgoodsComponent } from '../app/uploadgoods/uploadgoods.component';
import { GoodspageComponent } from '../app/goodspage/goodspage.component';
import {AboutComponent} from '../app/about/about.component';
import { CharthpComponent  }from '../app/charthp/charthp.component';
import {Personal2Component} from '../app/personal2/personal2.component';

//import { from } from 'rxjs';

const routes: Routes = [ 
  { path: 'homepage', component: HomepageComponent },
  { path: 'changemsg' , component: ChgmymsgComponent},
  { path: 'uploadgoods' , component: UploadgoodsComponent},
  { path: 'personal', component: PersonalComponent},
  { path: 'personals/:uid' , component :Personal2Component},
  { path: 'goodsdetail/:gid' , component:  GoodspageComponent },
  { path: 'about' , component: AboutComponent},
  { path:'chat' ,component: CharthpComponent },
  { path: '', redirectTo: 'homepage', pathMatch: 'full' },
];


@NgModule({
   imports: [ RouterModule.forRoot(routes) ],
   exports: [ RouterModule ]
})

export class AppRoutingModule { }


