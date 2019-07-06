import { Component, OnInit } from '@angular/core';
import {HomePageGoods,GoodsState } from'../struct';
import{ServerService} from'../server.service';

@Component({
  selector: 'app-goodspage',
  templateUrl: './goodspage.component.html',
  styleUrls: ['./goodspage.component.css']
})
export class GoodspageComponent implements OnInit {
  //一个类不可以只声明，然后直接用，否则出现undefine error
  data1 = new HomePageGoods;
  data2 = new GoodsState;
  data3 = "";

  constructor(private server : ServerService) { }
  ngOnInit() {
    this.getItPage(111);
  } 

  getItPage(id:number){
    this.server.GetGoodsDetail_1(id).subscribe(
      result=>{this.data1 = result;}
    )
    this.server.GetGoodsDetail_2(id).subscribe(
      result=>{this.data2 = result;}
    )
    this.server.GetGoodsDetail_3(id).subscribe(
      result=>{
        this.data3=result;
        $("#text-targer").html(this.data3);
      }
      )
  }

}
