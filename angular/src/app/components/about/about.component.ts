import { Component, OnInit } from '@angular/core';
import { RequestProto } from '../../struct';
import { ServerService } from '../../server.service';

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent implements OnInit {

  //static data
  data : Map[] = [];

  constructor(private server:ServerService) { }

  ngOnInit() {
    this.getData();
  }

  getData(){
    let postdata: RequestProto = {
      api: "staticdata",
      targetid:"staticdata",
    }
    this.server.Entrance(postdata).subscribe(result=>{
        if(result.statuscode!=0){
          alert("获取统计数据失败:"+ result.msg);
          return;
        }
        let temp:Map[] = result.data;
        //only let those who have login to see the static data
        if(this.server.userid==""){
          temp.forEach(row => {
              row.value = "(登录后可见)";
          });
        }
        this.data = temp;   
      })
  }
}

type Map = {
    key:string,
    value:any,
}