import { Component, OnInit } from '@angular/core';
import { ServerService } from '../server.service';
import {RequestProto} from '../struct';
import { AppComponent } from '../app.component';

@Component({
  selector: 'app-goodspage',
  templateUrl: './goodspage.component.html',
  styleUrls: ['./goodspage.component.css']
})

export class GoodspageComponent implements OnInit {
  goodsdt = new GoodsDetail();
  state = { collect: false, like: false };
  goodid = "";
  collectbtnshow = " 收藏 ";
  likebtnshow = " 点赞 "
  commentdata: comment[] = [];

  constructor(
    private server: ServerService,
    private app:AppComponent,
  ) { }

  ngOnInit() {
    this.goodid = this.server.LastSection();
    if(this.goodid==""){
      this.app.showMsgBox(-1, "无法获取商品ip,请刷新试试" );
      window.history.back();
    }
    this.getItPage(this.goodid);
    this.getComment(this.goodid);
    this.getStatement();
  }

    //######################## GetGoodsDeta() #######################################

  //get mainly message of it goods 🍌🔥🌽
  getItPage(id: string) {
    let postdata : RequestProto = {
      api:"goodsmessage",
      targetid:id,
      cachetime:60,
      cachekey:"gitp_"+id,
    };
    this.server.GetGoodsDeta(postdata).subscribe(result => {
      if(result.statuscode!=0){ 
          this.app.showMsgBox(-1, "获取页面数据失败，请刷新试试" , result.msg );
      }else{
        this.goodsdt = result.data;
        $("#text-targer").html(this.goodsdt.detail);
      }
    },err=>{
        this.app.showMsgBox(-1, "请求失败，请稍后再试" , err);
    });
  }

  // get comment data of it goods 🍌🔥🌽
  getComment(gid: string, latest?: boolean) {
    let postdata : RequestProto = {
      api:"goodscomment",
      targetid:gid,
      cachetime:60,
      cachekey:"gscm_"+gid,
    };
    if(latest==true){
      postdata.cachetime = 0;
    }
    this.server.GetGoodsDeta(postdata).subscribe(result=>{
      if(result.statuscode!=0){
        this.app.showMsgBox(-1, "获取评论数据失败，请稍后再试" , result.msg);
      }else{
        this.commentdata = result.data;
      }
    }, err=>{
      this.app.showMsgBox(-1, "请求失败，请稍后再试" , err);
    });
  }

  //get goods statement 🍌🔥🍈🌽
  getStatement() {
    let postdata : RequestProto = {
      api:"usergoodsstate",
      targetid:this.goodid,
      userid:this.server.userid,
      cachetime:60,
      cachekey:"usgs_"+this.goodid+"_"+this.server.userid,
    };
    this.server.GetGoodsDeta(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "获取商品状态失败，请稍后再试" , result.msg);
      } else {
        this.state = result.data;
        if (this.state.collect) { $("#collect-btn").removeClass('btn-info'); this.collectbtnshow=" 已收藏 ";}
        if (this.state.like) { $("#like-btn").removeClass('btn-info'); this.likebtnshow=" 已点赞 "}
      }
    }, err => {
      this.app.showMsgBox(-1, "获取数据失败，请稍后再试", err);
    })
  }

  //########################## SmallUpdate() #############################################

  //user like specified goods  🍍🔥🍈
  likeGoods() { 
    if(this.server.IsNotLogin()){
      return;
    }
    let postdata : RequestProto = {
      userid:this.server.userid,
      api:"likegoods",
      targetid:this.goodid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "点赞成功");
        $("#like-btn").removeClass('btn-info');
      } else {
        this.app.showMsgBox(-1, "点赞失败，请稍后再试" , result.msg);
      }
    },error=>{
      this.app.showMsgBox(-1, "获取数据失败，请稍后再试" , error);
    });
  }

  //user add a goods to favorite 🍍🔥🍈
  collect() {
    if(this.server.IsNotLogin()){
      return;
    }
    let postdata : RequestProto = {
      api:"addcollect",
      userid:this.server.userid,
      targetid:this.goodid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){
        this.app.showMsgBox(0, "收藏成功");
        $("#collect-btn").removeClass('btn-info');
        }else{
        this.app.showMsgBox(-1, "收藏失败,请稍后再试");
      }
    },error=>{
      this.app.showMsgBox(-1, "请求收藏失败，请稍后再试" , error);
    });
  }

  // user send a message to owner 🍍🔥🍈
  sendMessage() {
    if(this.server.IsNotLogin()){
      return;
    }
    let message = $("#messagesender").val().toString();
    if (message.length==0 || message.length>200){
      this.app.showMsgBox(1, "消息太长或为空");
      $("#sendcancel").click();
      return;
    }
    let postdata : RequestProto = {
      api:"sendmessage",
      userid:this.server.userid,
      targetid:this.goodsdt.userid,
      data:{message:message},
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
        if (result.statuscode==0){
          this.app.showMsgBox(0, "发送成功");
          $("#messagesender").val("");
          $("#sendcancel").click();
        }
        else{
          this.app.showMsgBox(-1, "发送失败" , result.msg);
          $("#sendcancel").click();
        }
    },error=>{
      this.app.showMsgBox(-1, "发送信息，请稍后再试" , error);
      $("#sendcancel").click();
    });
  }

  //user comment on a goods 🍍🔥🍈
  sendComment() {
    if(this.server.IsNotLogin()){
      return;
    }
    let comment = $("#comment-area").val().toString();
    if (comment == "") {
      this.app.showMsgBox(1, "发送内容不能为空");
      return;
    }
    let postdata : RequestProto = {
      api:"addcomment",
      userid:this.server.userid,
      targetid:this.goodid,
      data:{comment:comment},
    };
    //todo:检查评论内容
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){
        this.app.showMsgBox(0, "评论成功");
        this.getComment(this.goodid, true);
        $("#comment-area").val("");
      }else{
        this.app.showMsgBox(-1, "评论失败，请扫后再试" , result.msg);
      }
    }, error=>{
      this.app.showMsgBox(-1, "连接错误，请稍后再试" , error);
    });
  }

  //return a random color 🍏
  randomColor(name:string){
    let array = ["#f9a0a0", "#ea7b7b", "7bd54d","#57d2b3","#2594c8","#b325c8","#c1578d", "#d52c43", "#d4e814"];
    let random = name.length * 47;
    return array[random%array.length];
  }
}

//detail data response from server
class GoodsDetail {
  headimg: string;
  userid: string;
  username: string;
  time: any;
  title: string;
  price: number;
  id: string;
  name: string;
  visit: number;
  like: number;
  collect: number;     //precial
  talk: number;        //precial
  detail: string;
  type: string;
  tag: string;
}

type comment = {
  time: string;
  username: string;
  userid:string;
  comment: string;
}
