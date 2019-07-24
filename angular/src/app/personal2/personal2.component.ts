import { Component, OnInit } from '@angular/core';
import { UserMessage, UpdateResult } from '../struct';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-personal2',
  templateUrl: './personal2.component.html',
  styleUrls: ['./personal2.component.css']
})
export class Personal2Component implements OnInit {

  data = new UserMessage();
  userid = "";  //页面显示信息的用户id
  lookcerid = "00001";  //浏览者id
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

  //点赞用户
  updateLike(){
    let tre = new UpdateResult;
    this.server.SmallUpdate("likeuser","",this.userid,"",0 ).subscribe(result=>{
        tre = result;
        if (tre.status>=0) {
          alert("点赞成功！");
        }else{
          alert(tre.describe);
        }
    });
  }

  //关注用户
  addConcern(){
    let tre = new UpdateResult;
    //需要先获取浏览者的id，否则提示其登录
    this.server.SmallUpdate("addconcern",this.lookcerid, this.userid, "",0).subscribe(result=>{
        tre = result;
        if (tre.status>=0){
          alert("关注成功！");
        }else{
          alert(tre.describe);
        }
    });
  }
  //发送私信
  sendMessage(){
    //需要先登录
    let message = $("#messagesender").val().toString();
    let tre = new UpdateResult;
    this.server.SmallUpdate("sendmessage", this.lookcerid, this.userid, message,0).subscribe(result=>{
        tre = result;
        if(tre.status>=0) {
          alert("发送成功！");
        }else{
          alert(tre.describe);
        }
    });
  }
}
