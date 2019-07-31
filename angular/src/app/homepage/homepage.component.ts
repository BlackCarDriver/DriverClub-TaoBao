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
  //获得在主页中显示的一页商品列表的数据
  GetGoods(){
    this.server.GetHomePageGoods(this.lookingtype, this.lookingtag, this.lookingpage).subscribe(
      result => {
          this.goodsarray = result;
      })
  }

  //按照特定类型和标签获取商品列表
  GetSpecalGoods(type :string, tag:string){
    this.lookingtype = type;
    this.lookingtag = tag;
    this.GetGoods();
  }

  //搜索框搜索
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

  //获得商品的各个类型中包含的标签列表
  GetType(){
    this.server. GetHomePageType().subscribe(
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
