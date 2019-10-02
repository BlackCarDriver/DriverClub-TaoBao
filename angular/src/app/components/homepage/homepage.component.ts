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
  //goods type and tag list
   typearray = GoodsType[10];
   //different tag in specified goods type
   studytype = GoodSubType[100];
   sporttype = GoodSubType[100];
   daliytype = GoodSubType[100];
   electritype = GoodSubType[100];
   diytype = GoodSubType[100];
   virtualtype = GoodSubType[100];
   othertype = GoodSubType[100];
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
    $(".goods-area").mouseenter(function(){ $('.gg').collapse('hide');})
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
            this.app.showMsgBox(1,"没有找到数据！")
          }
          let temp:HomePageGoods[] = result.data; //let homepage show those out-of-focu images
          temp.forEach(row => {
            row.headimg = this.server.changeImgUrl(row.headimg);
          });
          this.goodsarray = temp;
          this.totalpage = Math.ceil(result.sum / this.server.homepage_goods_perpage);
          this.pageboxarray = new Array;
          for (let i=1;i<=this.totalpage && i<=5 ;i++){
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
    if(this.lookingpage>3){
      this.offsetpage = this.lookingpage - 3;
    }
  }

collapse(id:string){
    $('.collapse').collapse('hide');
    $(id).collapse('show');
}

showsinginbox(){
    $("#exampleModal").modal('show');
}
  

}
