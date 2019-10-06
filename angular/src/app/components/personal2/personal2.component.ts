import { Component, OnInit } from '@angular/core';
import { UserMessage, RequestProto } from '../../struct';
import { ServerService } from '../../server.service';
import { AppComponent } from '../../app.component';

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
  is_like = false;
  myid = "";
  constructor(
    private server: ServerService,
    private app: AppComponent,
  ) { }

  ngOnInit() {
    this.targetid = this.server.LastSection();
    this.myid = this.server.userid;
    if (this.targetid == "") {
      this.app.showMsgBox(-1, "无法获取用户id，请稍后再试");
      return;
    }
    this.getOtherMsg(this.targetid);
  }

  //get some other message need to show in the page 🍍🔥🍈🌽🍙
  getOtherMsg(uid: string) {
    let postdata: RequestProto = {
      api: "othermsg",
      targetid: uid,
      cachetime: 600,
      cachekey: "otmsg_" + uid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "获取页面数据失败，请刷新试试", result.msg);
        return;
      }
      //compress images size
      let temp:UserMessage = result.data;
      temp.headimg = this.server.changeImgUrl(temp.headimg);
      this.data = temp;
    }, err => { this.app.cFail(err); return; });
    this.getStatement();
  }
  //get concern and like statement  🍉🍈🌽
  getStatement() {
    if (this.server.userid==""){
      return;
    }
    let postdata: RequestProto = {
      api: "getuserstatement",
      targetid: this.targetid,
      userid: this.server.userid,
      cachetime: 300,
      cachekey: "gusm_" + this.targetid + "_" + this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "获取用户状态失败，请稍后再试", result.msg);
        return;
      }
      let state = { concern: false, like: false };
      state = result.data;
      this.is_concern = state.concern;
      this.is_like = state.like;
      if (this.is_concern) this.btn_concern_sho = "已关注";
      if (this.is_like) this.btn_like_sho = "已点赞";
    }, err => {
      this.app.cFail(err);
    });
  }
  // add a like to a user profile  🍍🔥🍈🍑
  updateLike() {
    if (this.server.IsNotLogin()) {
      return;
    }
    let postdata: RequestProto = {
      api: "likeuser",
      userid: this.server.userid,
      targetid: this.targetid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "点赞成功");
        this.is_like = true;
      } else {
        this.app.showMsgBox(-1, "点赞失败，请稍后再试");
      }
    }, err => {
        this.app.cFail(err);
    });
  }
  //add or remove a user from concern list 🍍🔥🍈🍑
  addConcern() {
    if (this.server.IsNotLogin()) {
      return;
    }
    let postdata: RequestProto = {
      userid: this.server.userid,
      targetid: this.targetid,
    };
    if (this.is_concern == false) {  //cancel concern
      postdata.api = "addconcern";
      this.server.SmallUpdate(postdata).subscribe(result => {
        if (result.statuscode == 0) {
          this.app.showMsgBox(0, "关注成功");
          this.is_concern = true;
        } else {
          this.app.showMsgBox(-1, "关注失败，请稍后再试");
        }
      }, err => {
        this.app.cFail(err);
      });
    } else {  //add into concern list
      postdata.api = "uncollectuser";
      this.server.DeleteMyData(postdata).subscribe(result => {
        if (result.statuscode == 0) {
          this.app.showMsgBox(0, "取消关注成功");
          this.is_concern = false;
        } else {
          this.app.showMsgBox(-1, "取消关注失败，请稍后再试", result.msg);
        }
      }, err => {
        this.app.cFail(err);
      });
    }

  }
  //send a private message to owner 🍍🔥🍈🍙
  sendMessage() {
    if (this.server.IsNotLogin()) {
      return;
    }
    if(this.server.userid==this.targetid){
      this.app.showMsgBox(1,"不能发消息给自己哦 :)");
      return;
    }
    let message = $("#messagesender").val().toString();
    let err = this.server.checkMessage(message);
    if (err!="") {
      this.app.showMsgBox(1, err);
      return;
    }
    let postdata: RequestProto = {
      api: "sendmessage",
      userid: this.server.userid,
      targetid: this.targetid,
      data: { message: message },
      cachekey:"uspvmsg_"+this.server.userid,
      cachetime:60,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "发送成功");
      } else {
        this.app.showMsgBox(-1, "发送失败，请稍后再试："+result.msg);
      }
    }, error => {
     this.app.cFail(error);
    });
  }
  showimg(url:string){
    this.app.ShowImg(url);
  }
}
