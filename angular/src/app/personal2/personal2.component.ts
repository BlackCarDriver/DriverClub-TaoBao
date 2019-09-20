import { Component, OnInit } from '@angular/core';
import { UserMessage, RequestProto } from '../struct';
import { ServerService } from '../server.service';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-personal2',
  templateUrl: './personal2.component.html',
  styleUrls: ['./personal2.component.css']
})
export class Personal2Component implements OnInit {
  data = new UserMessage();
  targetid = "";          //user which is showing in the page🍈
  btn_concern_sho = "关注";
  btn_like_sho = "点赞";
  is_concern = true;
  is_like= false;
  constructor(
    private server: ServerService,
    private app:AppComponent,
  ) { }

  ngOnInit() {
    this.targetid = this.server.LastSection();
    if (this.targetid==""){
      this.app.showMsgBox(-1, "无法获取用户id，请稍后再试");
      window.history.back();
      return;
    }
    this.getOtherMsg(this.targetid);
    this.getStatement();
  }

  //get some other message need to show in the page 🍍🔥🍈
  getOtherMsg(uid: string) {
    let postdata : RequestProto = {
      api:"othermsg",
      targetid:this.targetid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode==0){
        this.data = result.data;
        console.log(this.data)
      }else{
        this.app.showMsgBox(-1, "获取页面数据失败，请刷新试试" , result.msg );
      }
    }, error=>{console.log("GetMymsg() fail" + error)});
  }

  // add a like to a user profile  🍍🔥🍈🍑
  updateLike() {
    if(this.server.IsNotLogin()){
      return;
    }
    let postdata : RequestProto = {
      api:"likeuser",
      userid:this.server.userid,
      targetid:this.targetid, 
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){
        this.app.showMsgBox(0, "点赞成功");
        this.is_like =true;
      }else{
        this.app.showMsgBox(-1, "点赞失败，请稍后再试");
      }
    },error=>{
      this.app.showMsgBox(-1, "请求错误，请稍后再试" , error);
    });
  }

  //add or remove a user from concern list 🍍🔥🍈🍑
  addConcern() {
    if(this.server.IsNotLogin()){
      return;
    }
    let postdata : RequestProto = {
      userid:this.server.userid,
      targetid:this.targetid,
    };
    if(this.is_concern==false){  //cancel concern
      postdata.api = "addconcern";
      this.server.SmallUpdate(postdata).subscribe(result => {
        if(result.statuscode==0){
          this.app.showMsgBox(0, "关注成功");
          this.is_concern = true;
      }else{
        this.app.showMsgBox(-1, "关注失败，请稍后再试");
      }
      },err=>{
        this.app.showMsgBox(-1, "请求失败，请稍后再试", err);
      });
    }else{  //add into concern list
      postdata.api = "uncollectuser";
      this.server.DeleteMyData(postdata).subscribe(result => {
        if(result.statuscode==0){
          this.app.showMsgBox(0, "取消关注成功");
          this.is_concern = false;
        }else{
          this.app.showMsgBox(-1, "取消关注失败，请稍后再试",result.msg);
        }
      },err=>{
        this.app.showMsgBox(-1, "请求错误", err);
      });
    }
   
  }

  //send a private message to owner 🍍🔥🍈
  sendMessage() {
    if(this.server.IsNotLogin()){
      return;
    }
    let message = $("#messagesender").val().toString();
    if (message=="" || message.length > 200){
      this.app.showMsgBox(1, "消息太长或为空");
      return;
    }
    let postdata : RequestProto = {
      api:"sendmessage",
      userid:this.server.userid,
      targetid:this.targetid,  
      data:{message:message},
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){
        this.app.showMsgBox(0, "发送成功");
      }else{
        this.app.showMsgBox(-1, "发送失败，请稍后再试");
      }
    }, error=>{
        this.app.showMsgBox(-1, "请求失败，请稍后再试",error);
    });
  }
  
  //get concern and like statement  🍉🍈
  getStatement(){
    let postdata : RequestProto = {
      api:"getuserstatement",
      targetid:this.targetid,
      userid:this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result=>{
      if(result.statuscode!=0){
        this.app.showMsgBox(-1, "获取用户状态失败，请稍后再试", result.msg);
      }else{
        let state = {concern:false, like:false};
        state = result.data;
        this.is_concern = state.concern;
        this.is_like = state.like;
        if(this.is_concern) this.btn_concern_sho = "已关注";
        if(this.is_like) this.btn_like_sho="已点赞";
      }
    },error=>{
      console.log("getStatement() fail: "+error);
    });
  }

}
