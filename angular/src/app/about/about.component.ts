import { Component, OnInit } from '@angular/core';
// import * as $ from "jquery";
declare let $ : any;

@Component({
  selector: 'app-about',
  templateUrl: './about.component.html',
  styleUrls: ['./about.component.css']
})
export class AboutComponent implements OnInit {

  constructor() { }

  ngOnInit() {
    $(function () {
      $('.container').scrollspy({ target: '#navbar' });
     });
  }


  Uptotop(id){
    // var t = $(id).offset();
    // var top = $("#right").offest();
    // var y = top - t.top;
    // var $targer = $("#right");
    // $targer.scrollTop(y);
    // return false;
  }
}
