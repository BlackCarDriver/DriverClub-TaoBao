import { Component, OnInit } from '@angular/core';
import { UserMessage, GoodsShort, MyMessage, Rank, User, RequestProto } from '../struct';
import { ServerService } from '../server.service';
declare let $: any;

@Component({
  selector: 'app-personal',
  templateUrl: './personal.component.html',
  styleUrls: ['./personal.component.css']
})

export class PersonalComponent implements OnInit {
  userid = "19070010";
  // userid = "19070009";
  key = "itisuserkey..";

  msg = new UserMessage(); //基本信息
  mygoodslist = GoodsShort[100];      //我的商品
  mycollectlist = GoodsShort[100];    //我收藏的商品
  mymessagelist = MyMessage[100]; //我的消息
  hero = Rank[20];             //等级排行榜
  icare = User[100];   //我关注的和关注我的
  carei = User[100];  //关注我的用户
  show_no_goods = false;
  show_no_message = false;
  show_no_collect = false;
  constructor(private server: ServerService) { }

  ngOnInit() {
    // this.userid = this.server.Getusername();
    this.getmymsg();
    this.getmymgoods();
    this.getmycollect();
    this.getmymessage();
    this.getrank();
    this.getcare();
  }

  //get detail information 🍍
  getmymsg() {
    let postdata: RequestProto = {
      api: "mymsg",
      targetid: this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) { this.msg = result.data; }
      else { alert("get mymsg fail: " + result.msg); }
    }, error => {
      console.log("GetMyMsg() fail: " + error)
    });
  }

  //get the list of user i care and which acre me🍍
  getcare() {
    let postdata: RequestProto = {
      api: "mycare",
      targetid: this.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.icare = result.data[0];
        this.carei = result.data[1];
      } else {
        alert("GetMyMsg fail:" + result.msg);
      }
    }, error => { console.log(error) });
  }

  //get my goods information 🍍 🍉
  getmymgoods() {
    let postdata: RequestProto = {
      api: "mygoods",
      targetid: this.userid,
      offset:0,
      limit:25,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.mygoodslist = result.data;
        if (result.rows==0) this.show_no_goods=true;
      } else {
        alert("get goods msg fail:" + result.msg);
      }
    }, error => { console.log("GetMyMsg" + error) });
  }

  //get my collect goods information 🍍 🍉 
  getmycollect() {
    let postdata: RequestProto = {
      api: "mycollect",
      targetid: this.userid,
      offset:0,
      limit:25,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.mycollectlist = result.data;
        if (result.rows==0) this.show_no_collect=true;
      } else {
        alert("get my collect message fail:" + result.msg);
      }
    }, error => { console.log("GetMyMsg fail: " + error) });
  }

  // get my mail message  🍍 🍉
  getmymessage() {
    let postdata: RequestProto = {
      api: "message",
      targetid: this.userid,
      offset:0,
      limit:25,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.mymessagelist = result.data;
        if(result.rows==0) this.show_no_message=true;
      } else {
        alert("get my messges fail:" + result.msg);
      }
    }, error => { console.log("GetMyMsg() fail:" + error); });
  }

  //get users rank message  🍍 
  getrank() {
    let postdata: RequestProto = {
      api: "rank",
      targetid: this.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.hero = result.data;
      } else {
        alert("get userrank fail:" + result.msg);
      }
    }, error => { console.log("GetMyMsg() fail: " + error) });
  }

}
