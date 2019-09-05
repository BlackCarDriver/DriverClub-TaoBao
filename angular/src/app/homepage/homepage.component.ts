import { Component, OnInit } from '@angular/core';
import { ServerService } from '../server.service';
import {  HomePageGoods,GoodsType,GoodSubType } from '../struct';

// Property 'collapse' does not exist on type 'JQuery<HTMLElement>'....
// import * as bootstrap from 'bootstrap';
declare let $ : any;
// import * as $ from 'jquery';
// declare var $ :any; 


@Component({
  selector: 'app-homepage',
  templateUrl: './homepage.component.html',
  styleUrls: ['./homepage.component.css']
})

export class HomepageComponent implements OnInit {
  //商品主页封面列表
  goodsarray = HomePageGoods[100];  
  //商品类型
   typearray = GoodsType[10];
   //商品类型对应的标签
   studytype = GoodSubType[100];
   sporttype = GoodSubType[100];
   daliytype = GoodSubType[100];
   electritype = GoodSubType[100];
   diytype = GoodSubType[100];
   virtualtype = GoodSubType[100];
   othertype = GoodSubType[100];
   //当前浏览的商品类型和标签
   lookingtype = "all"; 
   lookingtag="all";
   lookingpage=1;
  constructor(
    private server : ServerService
  ) { }

  ngOnInit() {
    $(".goods-area").mouseenter(function(){ $('.collapse').collapse('hide');})
    this.GetGoods();
    this.GetType();
    this.set_mainbody_height();
  }

  //get a page of goods list data 🍋🔥
  GetGoods(){
    this.server.GetHomePageGoods(this.lookingtype, this.lookingtag, this.lookingpage).subscribe(
      result=>{
        if(result.statuscode==0){
          this.goodsarray = result.data;
        }else{
          alert("获取数据失败："+result.msg);
        }
      },
      error=>{console.log("GetHomePageGoods() fail: "+ error);}
    )
  }

  //get specified type or tag of goods
  GetSpecalGoods(type :string, tag:string){
    this.lookingtype = type;
    this.lookingtag = tag;
    this.GetGoods();
  }

  //search goods by input the keyword
  SearchGoods(){
    let input :string =  $('#searchgoods').val();
    if (input==""){
      return;
    }
    if (input.length > 20) {
      alert("名字太长！");
      return;
    }
    this.lookingtype = "like";
    this.lookingtag = input;
    this.GetGoods();
  }

  //show the type and tag information into page
  GetType(){
    this.server.GetHomePageType().subscribe(
      result => {
          this.typearray = result;
          this.studytype = this.typearray[0].list;
          this.sporttype  = this.typearray[1].list;
          this.daliytype = this.typearray[2].list;
          this.electritype = this.typearray[3].list;
          this.diytype  = this.typearray[4].list;
          this.virtualtype  = this.typearray[5].list;
          this.othertype  = this.typearray[6].list;  
      })
  }

set_mainbody_height(){
  var hight=  $(window).height();
  $(".main-body").css("min-height",hight-240+"px");
}
 
collapse(id:string){
    $('.collapse').collapse('hide');
    $(id).collapse('show');
}

showsinginbox(){
    $("#exampleModal").modal('show');
}
  

}
