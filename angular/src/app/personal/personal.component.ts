import { Component, OnInit } from '@angular/core';
import { UserMessage , GoodsShort, MyMessage, Rank,User  } from '../struct';
import { ServerService } from '../server.service';
declare let $: any;

@Component({
  selector: 'app-personal',
  templateUrl: './personal.component.html',
  styleUrls: ['./personal.component.css']
})

export class PersonalComponent implements OnInit {
  username = "19070010";
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
    // this.username = this.server.Getusername();
    this.getmymsg();
    this.getmymgoods();
    this.getmycollect();
    this.getmymessage();
    this.getrank();
    this.getcare();
  }

  //获取详细信息
  getmymsg(){
    this.server.GetMyMsg(this.username,"mymsg").subscribe(result=>{
      this.msg = result;
    });
  }
   //获取我的商品信息
  getmymgoods(){
    this.server.GetMyMsg(this.username,"mygoods").subscribe(result=>{
      this.mygoodslist = result;
    });
  }
  //获取我的收藏数据
  getmycollect(){
    this.server.GetMyMsg(this.username,"mycollect").subscribe(result=>{
      this.mycollectlist = result;
    });
  }
  //获取我的消息数据
  getmymessage(){
    this.server.GetMyMsg(this.username,"message").subscribe(result=>{
      this.mymessagelist = result;
  });
  }
  //获取用户等级排行数据
  getrank(){
    this.server.GetMyMsg(this.username,"rank").subscribe(result=>{
      this.hero = result;
  });
  }
   //获取用户等级排行数据
   getcare(){
    this.server.GetMyMsg(this.username,"mycare").subscribe(result=>{
      this.icare = result[0];
      this.carei = result[1];
  });
  }
}
