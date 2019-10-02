import { Component, OnInit } from '@angular/core';
import { ServerService } from '../../server.service';
@Component({
  selector: 'app-notfound1',
  templateUrl: './notfound1.component.html',
  styleUrls: ['./notfound1.component.css']
})
export class Notfound1Component implements OnInit {

  constructor(private server:ServerService) { }

  ngOnInit() {
  }
  back(){
    window.history.back();
  }
    
}
