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

  //the data need to map on html file
   goodsarray = HomePageGoods[100];  
   typearray = GoodsType[10];
   studytype = GoodSubType[100];
   sporttype = GoodSubType[100];
   daliytype = GoodSubType[100];
   electritype = GoodSubType[100];
   diytype = GoodSubType[100];
   virtualtype = GoodSubType[100];
   othertype = GoodSubType[100];

  constructor(
    private server : ServerService
  ) { }

  ngOnInit() {
    $(".goods-area").mouseenter(function(){ $('.collapse').collapse('hide');})
    this.GetGoods();
    this.GetType();
    this.set_mainbody_height();
  }


  GetGoods(){
    this.server.GetHomePageGoods("haha",11).subscribe(
      result => {
          this.goodsarray = result;
      })
  }

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
  $(".main-body").css("min-height",hight-240+"px")
}

    
  collapse(id:string){
    $('.collapse').collapse('hide');
    $(id).collapse('show');
  }

  showsinginbox(){
    $("#exampleModal").modal('show');
  }
  

}
