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
  userid = "";
  constructor(private server : ServerService) { }

  ngOnInit() {
    let rawStr = window.location.pathname;
    this.userid = rawStr.substring(11,21);
    this.getOtherMsg(this.userid);
  }

  //获取页面数据
  getOtherMsg(uid : string){
    this.server.GetMyMsg(this.userid, "othermsg").subscribe(result=>{
      this.data = result;
      
    });
  }

}
