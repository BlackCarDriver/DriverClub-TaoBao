import { Component, OnInit } from '@angular/core';
import { ServerService } from '../../server.service';
import { AppComponent } from '../../app.component';
import {  HomePageGoods,GoodsType,GoodSubType } from '../../struct';

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
  //mainly data, goodslist
  goodsarray = HomePageGoods[100];  
   //different tag in specified goods type
   secondhandTags = GoodSubType[100];
   packageTags = GoodSubType[100];
   studyTags = GoodSubType[100];
   mediaTags = GoodSubType[100];
   mesageTags = GoodSubType[100];
   virtualtype = GoodSubType[100];
   otherTags = GoodSubType[100];
   allTags = GoodsType[600];
   showingTag = GoodsType[600];

   //the tag, type and page user is looking 
   lookingtype = "all"; 
   lookingtag="all";
   lookingpage=1;
   offsetpage=0;
   //total page can be shown in present type and tag 🍇
   totalpage = 0; 
   pageboxarray = new Array;

  constructor(
    private server : ServerService,
    private app:AppComponent,
  ) { }
 
  ngOnInit() {
    // $(".goods-area").mouseenter(function(){ $('.gg').collapse('hide');})
    this.GetGoods();
    this.GetType();
  }

  //get a page of goods list data 🍋🔥🍇🌽
  //note taht request protocal is write in server.service.ts
  GetGoods(){
    this.server.GetHomePageGoods(this.lookingtype, this.lookingtag, this.lookingpage).subscribe(
      result=>{
        if(result.statuscode==0){
          if (result.rows==0){
            this.app.showMsgBox(1,"没有找到数据!")
          }
          let temp:HomePageGoods[] = result.data; //let homepage show those out-of-focu images
          temp.forEach(row => {
            row.headimg = this.server.changeImgUrl(row.headimg);
          });
          this.goodsarray = temp;
          this.totalpage = Math.ceil(result.sum / this.server.homepage_goods_perpage);
          this.pageboxarray = new Array;
          for (let i=1; i<=this.totalpage && i<=5 ;i++){
            this.pageboxarray.push(i);
          }
        }else{
          this.app.showMsgBox(-1,"获取数据失败："+result.msg)
        }
      },
      error=>{console.log("GetHomePageGoods() fail: "+ error);}
    )
  }

  //get specified type or tag of goods
  GetSpecalGoods(type :string, tag:string){
    this.lookingtype = type;
    this.lookingtag = tag;
    this.lookingpage = 1;
    this.GetGoods();
  }

  //search goods by input the keyword
  SearchGoods(){
    let input :string =  $('#searchgoods').val();
    if (input==""){
      return;
    }
    if (input.length > 20) {
      this.app.showMsgBox(1,"名字太长！");
      return;
    }
    this.lookingtype = "like";
    this.lookingtag = input;
    this.GetGoods();
  }

  //show the type and tag information into page
  GetType(){
    let typearray = GoodsType[10];
    this.server.GetHomePageType().subscribe(
      result => {
          typearray = result;
          this.secondhandTags = typearray[0].list; //二手商品
          this.packageTags  = typearray[1].list;   //软件程序
          this.studyTags = typearray[2].list;      //学习资料
          this.mediaTags = typearray[3].list;      //音乐视频
          this.mesageTags  = typearray[4].list;    //文章推送
          this.otherTags  = typearray[5].list;     //其他资源
          this.allTags = this.secondhandTags.concat(this.packageTags,this.studyTags, this.mediaTags, this.mesageTags, this.otherTags);
          this.showingTag = this.allTags;
      })
  }
  
  //display goods list of speficed type
  selectType(typeid:string){
    switch(typeid){
      case "secondtype":
          if (this.lookingtype=="二手商品"){
            this.lookingtype = "all";
            this.showingTag = this.allTags;
          }else{
            this.lookingtype = "二手商品";
            this.showingTag = this.secondhandTags;
          }
          break;
      case "packagetype":
          if (this.lookingtype=="软件程序"){
            this.lookingtype = "all";
            this.showingTag = this.allTags;
          }else{
          this.lookingtype = "软件程序";
          this.showingTag = this.packageTags;
          }
          break;
      case "studytype":
          if (this.lookingtype=="学习资料"){
            this.lookingtype = "all";
            this.showingTag = this.allTags;
          }else{
          this.lookingtype = "学习资料";
          this.showingTag = this.studyTags;
          }
          break;
      case "mediatype":
          if (this.lookingtype=="音乐视频"){
            this.lookingtype = "all";
            this.showingTag = this.allTags;
          }else{
          this.lookingtype = "音乐视频";
          this.showingTag = this.mediaTags;
          }
          break;
      case "messagetype":
          if (this.lookingtype=="消息推送"){
            this.lookingtype = "all";
            this.showingTag = this.allTags;
          }else{
          this.lookingtype = "消息推送";
          this.showingTag = this.mediaTags;
          }
          break;
      case "othertype":
          if (this.lookingtype=="其他资源"){
            this.lookingtype = "all";
            this.showingTag = this.allTags;
          }else{
          this.lookingtype = "其他资源";
          this.showingTag = this.otherTags;
          }
          break;
    }
    this.lookingtag = "all";
    this.lookingpage  = 1;
    this.offsetpage = 0;
    this.GetGoods();
  }

  //get goods list by specified tag
  selectTag(tagName:string) {
    this.lookingtag = tagName;
    this.lookingpage  = 1;
    this.offsetpage = 0;
    this.GetGoods();
  }

  //================= change page function =====================
  //display previous page
  prepage(){ 
    if(this.lookingpage==0) return;
    this.lookingpage--;
    this.GetGoods();
    this.adjustPage()
  }
  //display next page
  nextpage(){
    if(this.lookingpage+1>this.totalpage) return;
    this.lookingpage++;
    this.GetGoods();
    this.adjustPage()
  }
  //display specified page
  gotopage(topage?:number){
    if(topage==undefined){
      let pageVal = $("#whichpage").val();
      if(pageVal==undefined) return;
      if(pageVal<=0 || pageVal>this.totalpage) return;
      this.lookingpage = pageVal;
      this.GetGoods();
      this.adjustPage()
    }else{
      if(topage<=0 || topage>this.totalpage) return;
      this.lookingpage = topage;
      this.GetGoods();
      this.adjustPage()
    }
  }

  //adject the pagebox display
  adjustPage(){
    this.server.totop();
    if(this.totalpage<=5) return;
    if(this.lookingpage<=3){
      this.offsetpage = 0;
    }
    if(this.lookingpage >3 && this.totalpage - this.lookingpage>=2 ){
      this.offsetpage = this.lookingpage - 3;
    }

  }
showsinginbox(){
    $("#exampleModal").modal('show');
}
}
