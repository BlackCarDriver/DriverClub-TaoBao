import { Component, OnInit } from '@angular/core';
import { PersonalExpend , Goods2,MyMessage, Rank,User  } from '../struct';
import { ServerService } from '../server.service';
declare let $: any;

@Component({
  selector: 'app-personal',
  templateUrl: './personal.component.html',
  styleUrls: ['./personal.component.css']
})
export class PersonalComponent implements OnInit {
  username = "123456";
  key = "itisuserkey..";
  msg = new PersonalExpend(); //基本信息
  mygoodslist = Goods2[100];      //我的商品
  mycollectlist = Goods2[100];    //我收藏的商品
  mymessagelist = MyMessage[100]; //我的消息
  hero = Rank[20];             //等级排行榜
  icare = User[100];   //我关注的和关注我的
  carei = User[100];  //关注我的用户
  constructor(private server : ServerService) { }

  ngOnInit() {
    // this.username = this.server.Getusername();
    // this.getmymsg();
    // this.getmymgoods();
    // this.getmycollect();
    // this.getmymessage();
    // this.getrank();
    // this.getcare();
    // console.log(this.msg);
  }

  //获取详细信息
  getmymsg(){
    this.server.GetMyMsg(this.username,this.key,"mymsg").subscribe(result=>{
      this.msg = result;
      console.log("result:", result);
    });
  }
   //获取我的商品信息
  getmymgoods(){
    this.server.GetMyMsg(this.username,this.key,"mygoods").subscribe(result=>{
      this.mygoodslist = result;
    });
  }
  //获取我的收藏数据
  getmycollect(){
    this.server.GetMyMsg(this.username,this.key,"mycollect").subscribe(result=>{
      this.mycollectlist = result;
    });
  }
  //获取我的消息数据
  getmymessage(){
    this.server.GetMyMsg(this.username,this.key,"message").subscribe(result=>{
      this.mymessagelist = result;
  });
  }
  //获取用户等级排行数据
  getrank(){
    this.server.GetMyMsg(this.username,this.key,"rank").subscribe(result=>{
      this.hero = result;
  });
  }
   //获取用户等级排行数据
   getcare(){
    this.server.GetMyMsg(this.username,this.key,"care").subscribe(result=>{
      this.icare = result[0];
      this.carei = result[1];
  });
  }
}
