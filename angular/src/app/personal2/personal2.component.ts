import { Component, OnInit } from '@angular/core';
import { UserMessage } from '../struct';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-personal2',
  templateUrl: './personal2.component.html',
  styleUrls: ['./personal2.component.css']
})
export class Personal2Component implements OnInit {

  data = new UserMessage();

  constructor(private server : ServerService) { }

  ngOnInit() {
    this.getOtherMsg();
  }

  //获取页面数据
  getOtherMsg(){
    this.server.GetMyMsg("12345", "othermsg").subscribe(result=>{
      this.data = result;
    });
  }

}
