import { Component, OnInit } from '@angular/core';
import {GoodsDetail } from'../struct';
import{ServerService} from'../server.service';

@Component({
  selector: 'app-goodspage',
  templateUrl: './goodspage.component.html',
  styleUrls: ['./goodspage.component.css']
})
export class GoodspageComponent implements OnInit {
  //一个类不可以只声明，然后直接用，否则出现undefine error
  goodsdt = new GoodsDetail;
  
  constructor(private server : ServerService) { }
  ngOnInit() {
    let rawStr = window.location.pathname;
    let pid = rawStr.substring(13,23);
    this.getItPage(pid);
  } 

  getItPage(id:string){
    this.server.GetGoodsDeta(id, "goodsmessage").subscribe(result=>{
      this.goodsdt = result;
      $("#text-targer").html(this.goodsdt.detail);
  });
  }

}