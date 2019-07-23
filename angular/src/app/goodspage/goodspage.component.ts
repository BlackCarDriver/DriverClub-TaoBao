import { Component, OnInit } from '@angular/core';
import {GoodsDetail ,UpdateResult} from'../struct';
import{ServerService} from'../server.service';
import { asElementData } from '@angular/core/src/view';

@Component({
  selector: 'app-goodspage',
  templateUrl: './goodspage.component.html',
  styleUrls: ['./goodspage.component.css']
})
export class GoodspageComponent implements OnInit {
  //一个类不可以只声明，然后直接用，否则出现undefine error
  goodsdt = new GoodsDetail;
  userid = "00001"
  goodid = "";
  constructor(private server : ServerService) { }
  ngOnInit() {
    let rawStr = window.location.pathname;
    this.goodid = rawStr.substring(13,23);
    this.getItPage(this.goodid);
  } 

  getItPage(id:string){
    this.server.GetGoodsDeta(id, "goodsmessage").subscribe(result=>{
      this.goodsdt = result;
      $("#text-targer").html(this.goodsdt.detail);
    });
  }
  
  //点赞商品
  likeGoods(){
    let tre = new UpdateResult;
    this.server.SmallUpdate("likegoods",this.userid, this.goodid, "", 1 ).subscribe(resutl=>{
        tre = resutl;
        if (tre.status>=0) {
          alert("点赞成功!");
        }else{
          alert(tre.describe);
        }
    });
  }
  //发送私信
  sendMessage(){
    let message = $("#messagesender").val().toString();
    let tre = new UpdateResult;
    this.server.SmallUpdate("sendmessage", this.userid, this.goodsdt.userid, message,0).subscribe(result=>{
        tre = result;
        if(tre.status>=0) {
          alert("发送成功！");
        }else{
          alert(tre.describe);
        }
    });
  }
  //收藏商品
  collect(){
    let tre = new UpdateResult;
    this.server.SmallUpdate("addcollect", this.userid, this.goodid, "", 0).subscribe(result=>{
      tre = result;
      if(tre.status >=0){
        alert("收藏成功!");
      }else{
        alert(tre.describe);
      }
    });
  }

}