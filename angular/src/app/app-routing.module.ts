import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { HomepageComponent } from './components/homepage/homepage.component';
import { UploadgoodsComponent } from './components/uploadgoods/uploadgoods.component';

//import { from } from 'rxjs';

const routes: Routes = [ 
  { path: 'homepage', component: HomepageComponent },
  { path: 'upload' , component :UploadgoodsComponent},
  
  { path: 'myself',  loadChildren: () => import('./login/login.module').then(mod => mod.LoginModule) },
  { path: 'public',  loadChildren: () => import('./public/public.module').then(mod => mod.PublicModule) },
  { path: 'message',  loadChildren: () => import('./message/message.module').then(mod => mod.MessageModule) },

  { path: '', redirectTo: 'homepage', pathMatch: 'full' },
  // { path: '**', component: HomepageComponent  },  //TODO: page not found component
];

 
@NgModule({
   imports: [ RouterModule.forRoot(routes) ],
   exports: [ RouterModule ]
})

export class AppRoutingModule { }


