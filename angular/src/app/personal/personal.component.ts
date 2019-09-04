import { Component, OnInit } from '@angular/core';
import { UserMessage , GoodsShort, MyMessage, Rank,User,RequestProto, RequertResult  } from '../struct';
import { ServerService } from '../server.service';
import { post } from 'selenium-webdriver/http';
declare let $: any;

@Component({
  selector: 'app-personal',
  templateUrl: './personal.component.html',
  styleUrls: ['./personal.component.css']
})

export class PersonalComponent implements OnInit {
  userid = "19070010";
  key = "itisuserkey..";

  msg = new UserMessage(); //基本信息
  mygoodslist = GoodsShort[100];      //我的商品
  mycollectlist = GoodsShort[100];    //我收藏的商品
  mymessagelist = MyMessage[100]; //我的消息
  hero = Rank[20];             //等级排行榜
  icare = User[100];   //我关注的和关注我的
  carei = User[100];  //关注我的用户
  constructor(private server : ServerService) { }

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
  getmymsg(){
     let postdata : RequestProto = {
      api:"mymsg",
      userid:this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result=>{
      if (result.statuscode==0){this.msg = result.data;}
      else{alert("get mymsg fail: "+result.msg);}
    }, error=>{console.log("GetMyMsg() fail: " + error )
  });
  }

  //get my goods information 🍍
  getmymgoods(){
    let postdata : RequestProto = {
      api:"mygoods",
      userid:this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result=>{
      if(result.statuscode==0){
        this.mygoodslist = result;
      }else{
        alert("get goods msg fail:"+result.msg);
      }
    }, error=>{console.log("GetMyMsg"+error)});
  }

  //get my collect goods information 🍍
  getmycollect(){
    let postdata : RequestProto = {
      api:"mycollect",
      userid:this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result=>{
      if (result.statuscode==0){
        this.mycollectlist = result;
      }else{
        alert("get my collect message fail:"+result.msg);
      }
    }, error=>{ console.log("GetMyMsg fail: "+error)});
  }

  // get my mail message  🍍
  getmymessage(){
    let postdata : RequestProto = {
      api:"message",
      userid:this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result=>{
      if (result.statuscode==0){
        this.mymessagelist = result;
      }else{
        alert("get my messges fail:"+result.msg );
      }
  }, error=>{ console.log("GetMyMsg() fail:"+error);});
  }

  //get users rank message  🍍
  getrank(){
    let postdata : RequestProto = {
      api:"rank",
      userid:this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result=>{
      if (result.statuscode==0){
        this.hero = result;
      }else{
        alert("get userrank fail:"+result.msg);
      }
  }, error=>{console.log("GetMyMsg() fail: "+ error)});
  }

   //get user's cared numbers information 🍍
   getcare(){
    let postdata : RequestProto = {
      api:"mycare",
      userid:this.userid
    };
    this.server.GetMyMsg(postdata).subscribe(result=>{
      if (result.statuscode==0){
        this.icare = result[0];
        this.carei = result[1];
      }else{
        alert("GetMyMsg fail:"+result.msg);
      }
  }, error=>{ console.log(error)});
  }
}
