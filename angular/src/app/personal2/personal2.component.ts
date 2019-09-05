import { Component, OnInit } from '@angular/core';
import { UserMessage, RequestProto } from '../struct';
import { ServerService } from '../server.service';

@Component({
  selector: 'app-personal2',
  templateUrl: './personal2.component.html',
  styleUrls: ['./personal2.component.css']
})
export class Personal2Component implements OnInit {

  data = new UserMessage();
  userid = "";          //user which is showing in the page
  lookcerid = "00001"; 
  constructor(private server: ServerService) { }

  ngOnInit() {
    let rawStr = window.location.pathname;
    this.userid = rawStr.substring(11, 21);
    this.getOtherMsg(this.userid);
  }

  //get some other message need to show in the page 🍍
  getOtherMsg(uid: string) {
    let postdata : RequestProto = {
      api:"othermsg",
      targetid:this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode==0){
        this.data = result.data;
        console.log(this.data)
      }else{
        alert("get other message fail: "+ result.msg);
      }
    }, error=>{console.log("GetMymsg() fail" + error)});
  }

  // add a like to a user profile  🍍
  updateLike() {
    let postdata : RequestProto = {
      api:"likegoods",
      userid:this.lookcerid,
      targetid:this.userid, 
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){
        alert("收藏成功！");
      }else{
        alert("收藏失败："+result.msg);
      }
    },error=>{
      alert("updateLike() fail: "+error); 
    });
  }

  //add a user into favorite 🍍
  addConcern() {
    //todo: must login before following operation
    let postdata : RequestProto = {
      api:"addconcern",
      userid:this.lookcerid,
      targetid:this.userid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if(result.statuscode==0){alert("关注成功！");}
      else{alert("关注失败："+result.msg);}
    },err=>{
      alert("addConcern() fail: "+err);
    });
  }

  //send a private message to owner 🍍
  sendMessage() {
    //todo: must login before following operation
    let message = $("#messagesender").val().toString();
    //TODO: check the message
    let postdata : RequestProto = {
      api:"sendmessage",
      userid:this.lookcerid,
      targetid:this.userid,  
      data:{message:message},
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){
        alert("发送成功！");
      }else{
        alert("发送失败："+result.msg);
      }
    }, error=>{
        alert("sendMessage() fail: "+error);
    });
  }
}
