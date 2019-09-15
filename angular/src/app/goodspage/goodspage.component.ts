import { Component, OnInit } from '@angular/core';
import { ServerService } from '../server.service';
import {RequestProto} from '../struct';

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
  ) { }

  ngOnInit() {
    this.goodid = this.server.LastSection();
    if(this.goodid==""){
      alert("获取商品ID失败！");
      window.history.back();
    }
    this.getItPage(this.goodid);
    this.getComment(this.goodid);
    this.getStatement();
  }

    //######################## GetGoodsDeta() #######################################

  //get mainly message of it goods 🍌🔥
  getItPage(id: string) {
    let postdata : RequestProto = {
      api:"goodsmessage",
      targetid:this.goodid,
    };
    this.server.GetGoodsDeta(postdata).subscribe(result => {
      if(result.statuscode!=0){ 
          alert("getItPage():"+result.msg);
      }else{
        this.goodsdt = result.data;
        $("#text-targer").html(this.goodsdt.detail);
      }
    },err=>{
        alert("getItPage fail: "+err);
    });
  }

  // get comment data of it goods 🍌🔥
  getComment(gid: string) {
    let postdata : RequestProto = {
      api:"goodscomment",
      targetid:this.goodid,
    };
    this.server.GetGoodsDeta(postdata).subscribe(result=>{
      if(result.statuscode!=0){
        alert("getComment():"+result.msg);
      }else{
        this.commentdata = result.data;
      }
    }, err=>{
       alert("getComment fail: "+err);
    });
  }

  //get goods statement 🍌🔥🍈
  getStatement() {
    let postdata : RequestProto = {
      api:"usergoodsstate",
      targetid:this.goodid,
      userid:this.server.userid,
    };
    this.server.GetGoodsDeta(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        alert("获取数据失败: "+result.msg);
      } else {
        this.state = result.data;
        if (this.state.collect) { $("#collect-btn").removeClass('btn-info'); this.collectbtnshow=" 已收藏 ";}
        if (this.state.like) { $("#like-btn").removeClass('btn-info'); this.likebtnshow=" 已点赞 "}
        this.getStatement();
      }
    }, err => {
      alert("getStatement unresponse:" + err);
    })
  }

  //########################## SmallUpdate() #############################################

  //user like specified goods  🍍🔥🍈
  likeGoods() { 
    let postdata : RequestProto = {
      userid:this.server.userid,
      api:"likegoods",
      targetid:this.goodid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        alert("点赞成功!");
        this.getStatement();
      } else {
        alert("点赞失败："+result.statuscode+result.msg);
      }
    },error=>{
        alert("error happen in likeGoods: "+error)
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
      if (result.statuscode==0){alert("收藏成功！");}
      else{alert("收藏失败："+result.msg);}
    },error=>{
      alert("error happen in  collect():"+error);
    });
  }

  // user send a message to owner 🍍🔥🍈
  sendMessage() {
    if(this.server.IsNotLogin()){
      return;
    }
    let message = $("#messagesender").val().toString();
    if (message.length==0 || message.length>200){
      alert("消息太长或为空");
      return;
    }
    let postdata : RequestProto = {
      api:"sendmessage",
      userid:this.server.userid,
      targetid:this.goodsdt.userid,
      data:{message:message},
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
        if (result.statuscode==0){alert("发送成功！");}
        else{alert("发送失败："+result.statuscode+":"+result.msg);}
    },error=>{
        alert("error happen in sendMessage():"+error);
    });
  }

  //user comment on a goods 🍍🔥🍈
  sendComment() {
    if(this.server.IsNotLogin()){
      return;
    }
    let comment = $("#comment-area").val().toString();
    if (comment == "") {
      alert("内容不能为空");
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
        alert("评论成功！");
        this.getComment(this.goodid);
        $("#comment-area").val("");
      }else{
        alert("评论失败："+result.msg);
      }
    }, error=>{
      alert("error happen in sendComment():"+error);
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
