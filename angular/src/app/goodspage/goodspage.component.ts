import { Component, OnInit } from '@angular/core';
import {HomePageGoods,GoodsDetail } from'../struct';
import{ServerService} from'../server.service';

@Component({
  selector: 'app-goodspage',
  templateUrl: './goodspage.component.html',
  styleUrls: ['./goodspage.component.css']
})
export class GoodspageComponent implements OnInit {
  //一个类不可以只声明，然后直接用，否则出现undefine error
  goodsdt = new GoodsDetail;
  goodsds = "";

  constructor(private server : ServerService) { }
  ngOnInit() {
    this.getItPage(111);
  } 

  getItPage(id:number){
    this.server.GetGoodsDeta(id, "message").subscribe(
      result=>{this.goodsdt = result;}
    )
    //获取描述商品的文件
    this.server.GetGoodsDeta(id, "detail").subscribe(
      result=>{
        this.goodsds=result;
        $("#text-targer").html(this.goodsds);
      })
  }

}
