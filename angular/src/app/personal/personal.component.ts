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
  msg_maxrow = 8;
  msg_sumpage = 0;
  msg_nowat = 1;
  msg_offset = 0;
  msg_array = new Array;
  mg_maxrow = 8;
  mg_sumpage = 0;
  mg_nowat = 1;
  mg_offset = 0;
  mg_array = new Array;
  mc_maxrow = 8;
  mc_sumpage = 0;
  mc_nowat = 1;
  mc_offset = 0;
  mc_array = new Array;
  constructor(private server: ServerService) { }

  ngOnInit() {
    if(this.server.IsNotLogin()){
      window.history.back();
    }
    // this.userid = this.server.Getusername();
    this.getmymsg();
    this.getmymgoods();
    this.getmycollect();
    this.getmymessage();
    this.getrank();
    this.getcare();
  }

  //get detail information 🍍🍈
  getmymsg() {
    let postdata: RequestProto = {
      api: "mymsg",
      targetid: this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) { this.msg = result.data; }
      else { alert("get mymsg fail: " + result.msg); }
    }, error => {
      console.log("GetMyMsg() fail: " + error)
    });
  }

  //get the list of user i care and which acre me🍍🍈
  getcare() {
    let postdata: RequestProto = {
      api: "mycare",
      targetid: this.server.userid,
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

  //get my goods information 🍍 🍉🍈 🍇🍏
  getmymgoods() {
    let postdata: RequestProto = {
      api: "mygoods",
      targetid: this.server.userid,
      offset:(this.mg_nowat-1)*this.mg_maxrow,
      limit:this.mg_maxrow,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.mygoodslist = result.data;
        this.mg_sumpage = Math.ceil(result.sum/this.mg_maxrow);
        if (result.rows==0) this.show_no_goods=true;
        else if ( this.mg_sumpage>1){
          this.mg_array = new Array;
          for(let i=1;i<=this.mg_sumpage && i<=9;i++){
            this.mg_array.push(i);
          }
        }
      }
    }, error => { console.log("GetMyMsg" + error) });
  }

  //get my collect goods information 🍍 🍉 🍈 🍇🍏
  getmycollect() {
    let postdata: RequestProto = {
      api: "mycollect",
      targetid: this.server.userid,
      offset: (this.mc_nowat-1)*this.mc_maxrow,
      limit:this.mc_maxrow,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.mycollectlist = result.data;
        if (result.rows==0) this.show_no_collect=true;
        else if (result.sum>1){
          this.mc_array = new Array;
          this.mc_sumpage = Math.ceil(result.sum/this.mc_maxrow);
          for(let i=1;i<=this.mc_sumpage && i<=9;i++){
            this.mc_array.push(i);
          }
        }
      } else {
        alert("get my collect message fail:" + result.msg);
      }
    }, error => { console.log("GetMyMsg fail: " + error) });
  }

  // get my mail message  🍍 🍉🍈 🍇 🍏 🍑
  getmymessage() {
    let postdata: RequestProto = {
      api: "message",
      targetid: this.server.userid,
      offset:(this.msg_nowat-1)*this.msg_maxrow,
      limit:this.msg_maxrow,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.mymessagelist = result.data;
        if(result.rows==0) this.show_no_message=true;
        else if (result.sum>1){
          this.msg_array = new Array;
          this.msg_sumpage = Math.ceil(result.sum/this.msg_maxrow);
          for(let i=1;i<=this.msg_sumpage && i<=9;i++){
            this.msg_array.push(i);
          }
        }
      } else {
        alert("get my messges fail:" + result.msg);
      }
    }, error => { console.log("GetMyMsg() fail:" + error); });
  }

  //get users rank message  🍍🍈
  getrank() {
    let postdata: RequestProto = {
      api: "rank",
      targetid: this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.hero = result.data;
      } else {
        alert("get userrank fail:" + result.msg);
      }
    }, error => { console.log("GetMyMsg() fail: " + error) });
  }

  //delete my upload goods 🍑
  deleteMyGoods(gid : string){
    let postdata: RequestProto = {
      api: "deletemygoods",
      targetid: gid,
      userid:this.server.userid,
    };
    this.server.DeleteMyData(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        alert("删除成功!");
        this.getmymgoods();
      } else {
        alert("DeleteMyData() fail:" + result.msg);
      }
    }, error => { console.log(error) });
  }

  //cancel collect a goods 🍑
  cancelCollect(gid:string){
    let postdata: RequestProto = {
      api: "uncollectgoods",
      targetid: gid,
      userid:this.server.userid,
    };
    this.server.DeleteMyData(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        alert("取消收藏成功!");
        this.getmycollect();
      } else {
        alert("cancelCollect() fail:" + result.msg);
      }
    }, error => { console.log(error) });
  }

    //cancel collect a goods 🍑
    deleteMessage(mid:string){
      let postdata: RequestProto = {
        api: "deletemymessage",
        targetid: mid,
        userid:this.server.userid,
      };
      this.server.DeleteMyData(postdata).subscribe(result => {
        if (result.statuscode == 0) {
          alert("删除消息成功!");
          this.getmymessage();
        } else {
          alert("deleteMessage() fail:" + result.msg);
        }
      }, error => { console.log(error) });
    }
//#################### reference to pagebox #######################

  //functions reference to my_message area pageboxx 🍏
  setMsgPagebox(topage:number){
    if (topage<0 || topage>this.msg_sumpage) return;
    this.msg_nowat = topage;
    this.getmymessage();
  }
  MsgPrepage(){
    if(this.msg_nowat==0) return;
    this.msg_nowat--;
    this.getmymessage();
    this.adjustMsgPage();
  }
  MsgNextpage(){
    if(this.msg_nowat+1>this.msg_sumpage) return;
    this.msg_nowat++;
    this.getmymessage();
    this.adjustMsgPage();
  }
  adjustMsgPage(){
    if(this.msg_sumpage<=9) return;
    if(this.msg_nowat>5){
      this.msg_offset = this.msg_nowat - 5;
    }
  }
  //functions reference to my_goods area pagebox 🍏
  setMgPagebox(topage:number){
    if (topage<0 || topage>this.mg_sumpage) return;
    this.mg_nowat = topage;
    this.getmymgoods();
  }
  MgPrepage(){
    if(this.mg_nowat==0) return;
    this.mg_nowat--;
    this.getmymgoods();
    this.adjustMgPage();
  }
  MgNextpage(){
    if(this.mg_nowat+1>this.mg_sumpage) return;
    this.mg_nowat++;
    this.getmymgoods();
    this.adjustMgPage();
  }
  adjustMgPage(){
    if(this.mg_sumpage<=9) return;
    if(this.mg_nowat>5){
      this.mg_offset = this.mg_nowat - 5;
    }
  }
  //functions reference to my_collect area pagebox 🍏
  setMcPagebox(topage:number){
    if (topage<0 || topage>this.mc_sumpage) return;
    this.mc_nowat = topage;
    this.getmycollect();
  }
  McgPrepage(){
    if(this.mc_nowat==0) return;
    this.mc_nowat--;
    this.getmycollect();
    this.adjustMcgPage();
  }
  McgNextpage(){
    if(this.mc_nowat+1>this.mc_sumpage) return;
    this.mc_nowat++;
    this.getmycollect();
    this.adjustMcgPage();
  }
  adjustMcgPage(){
    if(this.mc_sumpage<=9) return;
    if(this.mc_nowat>5){
      this.mc_offset = this.mc_nowat - 5;
    }
  }
}
