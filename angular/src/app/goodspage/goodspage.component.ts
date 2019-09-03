import { Component, OnInit } from '@angular/core';
import { UpdateResult } from '../struct';
import { HttpClient } from '@angular/common/http';
import { ServerService } from '../server.service';
import {RequestProto} from '../struct';

@Component({
  selector: 'app-goodspage',
  templateUrl: './goodspage.component.html',
  styleUrls: ['./goodspage.component.css']
})

export class GoodspageComponent implements OnInit {
  goodsdt: GoodsDetail;
  state = { collect: false, like: false };
  userid = "00001"    //当前浏览者的id
  goodid = "";
  collectbtnshow = " 收藏 ";
  likebtnshow = " 点赞 "
  commentdata: comment[] = [];

  constructor(
    private server: ServerService,
    private http: HttpClient,
  ) { }

  ngOnInit() {
    let rawStr = window.location.pathname;
    this.goodid = rawStr.substring(13, 23);
    this.getItPage(this.goodid);
    this.getComment(this.goodid);
    this.getStatement();
  }

    //######################## GetGoodsDeta() #######################################

  //get mainly message of it goods 🍌
  getItPage(id: string) {
    let postdata : RequestProto = {
      api:"goodsmessage",
      goodsid:this.goodid,
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

  // get comment data of it goods 🍌
  getComment(gid: string) {
    let postdata : RequestProto = {
      api:"goodscomment",
      goodsid:this.goodid,
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

  //get goods statement 🍌
  getStatement() {
    let postdata : RequestProto = {
      api:"usergoodsstate",
      goodsid:this.goodid,
      userid:this.userid,
    };
    this.server.GetGoodsDeta(postdata).subscribe(result => {
      if (result.statuscode < 0) {
        alert("getStatement(): "+result.msg);
      } else {
        this.state = result.data;
        if (this.state.collect) { $("#collect-btn").removeClass('btn-info'); this.collectbtnshow=" 已收藏 ";}
        if (this.state.like) { $("#like-btn").removeClass('btn-info'); this.likebtnshow=" 已点赞 "}
      }
    }, err => {
      alert("getStatement unresponse:" + err);
    })
  }
  //#######################################################################

  //点赞商品
  likeGoods() {
    let tre = new UpdateResult;
    this.server.SmallUpdate("likegoods", this.userid, this.goodid, "", 1).subscribe(resutl => {
      tre = resutl;
      if (tre.status >= 0) {
        alert("点赞成功!");
      } else {
        alert(tre.describe);
      }
    });
  }
  //发送私信
  sendMessage() {
    let message = $("#messagesender").val().toString();
    let tre = new UpdateResult;
    this.server.SmallUpdate("sendmessage", this.userid, this.goodsdt.userid, message, 0).subscribe(result => {
      tre = result;
      if (tre.status >= 0) {
        alert("发送成功！");
      } else {
        alert(tre.describe);
      }
    });
  }
  //收藏商品
  collect() {
    let tre = new UpdateResult;
    this.server.SmallUpdate("addcollect", this.userid, this.goodid, "", 0).subscribe(result => {
      tre = result;
      if (tre.status >= 0) {
        alert("收藏成功!");
      } else {
        alert(tre.describe);
      }
    });
  }
  //发表评论
  sendComment() {
    let tre = new UpdateResult;
    let comment: string;
    comment = $("#comment-area").val().toString();
    if (comment == "") {
      alert("内容不能为空");
      return;
    }
    alert(comment);
    //todo:检查评论内容
    this.server.SmallUpdate("addcomment", this.userid, this.goodid, comment, 0).subscribe(result => {
      tre = result;
      if (tre.status >= 0) {
        alert("评论成功");
      }
    });
  }



}

//detail data response from server
type GoodsDetail = {
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

// //whether it collect have liked or collected by user 
// type statement = {
//   like: boolean;
//   collect: boolean;
// }