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
  userid = "00001"    //当前浏览者的id
  goodid = "";
  collectbtnshow = " 收藏 ";
  likebtnshow = " 点赞 "
  commentdata: comment[] = [];

  constructor(
    private server: ServerService,
  ) { }

  ngOnInit() {
    let rawStr = window.location.pathname;
    this.goodid = rawStr.substring(13, 23);
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

  //get goods statement 🍌🔥
  getStatement() {
    let postdata : RequestProto = {
      api:"usergoodsstate",
      targetid:this.goodid,
      userid:this.userid,
    };
    this.server.GetGoodsDeta(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        alert("获取数据失败: "+result.msg);
      } else {
        console.log(result.data)
        this.state = result.data;
        if (this.state.collect) { $("#collect-btn").removeClass('btn-info'); this.collectbtnshow=" 已收藏 ";}
        if (this.state.like) { $("#like-btn").removeClass('btn-info'); this.likebtnshow=" 已点赞 "}
      }
    }, err => {
      alert("getStatement unresponse:" + err);
    })
  }

  //########################## SmallUpdate() #############################################

  //user like specified goods  🍍🔥
  likeGoods() { 
    let postdata : RequestProto = {
      api:"likegoods",
      userid:this.userid,
      targetid:this.goodid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        alert("点赞成功!");
      } else {
        alert("点赞失败："+result.statuscode+result.msg);
      }
    },error=>{
        alert("error happen in likeGoods: "+error)
    });
  }

  //user add a goods to favorite 🍍🔥
  collect() {
    let postdata : RequestProto = {
      api:"addcollect",
      userid:this.userid,
      targetid:this.goodid,
    };
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){alert("收藏成功！");}
      else{alert("收藏失败："+result.msg);}
    },error=>{
      alert("error happen in  collect():"+error);
    });
  }

  // user send a message to owner 🍍🔥
  sendMessage() {
    let message = $("#messagesender").val().toString();
    let postdata : RequestProto = {
      api:"sendmessage",
      userid:this.userid,
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

  //user comment on a goods 🍍🔥
  sendComment() {
    let comment = $("#comment-area").val().toString();
    if (comment == "") {
      alert("内容不能为空");
      return;
    }
    let postdata : RequestProto = {
      api:"addcomment",
      userid:this.userid,
      targetid:this.goodid,
      data:{comment:comment},
    };
    //todo:检查评论内容
    this.server.SmallUpdate(postdata).subscribe(result => {
      if (result.statuscode==0){
        alert("评论成功！");
      }else{
        alert("评论失败："+result.msg);
      }
    }, error=>{
      alert("error happen in sendComment():"+error);
    });
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
  comment: string;
}
