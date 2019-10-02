import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { Notfound1Component } from './components/notfound1/notfound1.component';
import { HomepageComponent } from './components/homepage/homepage.component';

import { LoadSelect } from './LoadSelect';

const routes: Routes = [ 
  { path: 'homepage', component: HomepageComponent },
  { path: 'myself', loadChildren: "./login/login.module#LoginModule" },
  { path: 'public', loadChildren: "./public/public.module#PublicModule", data: { preload: true } },
  { path: 'message', loadChildren: "./message/message.module#MessageModule", data: { preload: true }},
  { path: '', redirectTo: '/homepage', pathMatch: 'full' },
  { path: '**', component: Notfound1Component},
];

 
@NgModule({
   imports: [ 
     RouterModule.forRoot(routes,{preloadingStrategy:LoadSelect})
    ],
   exports: [ RouterModule ],
   providers:[LoadSelect],
})

export class AppRoutingModule { }


