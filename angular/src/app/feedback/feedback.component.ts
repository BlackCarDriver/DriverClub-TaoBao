import { Component, OnInit } from '@angular/core';
import { AppComponent } from '../app.component';
import { ServerService } from '../server.service';
import { RequestProto } from '../struct';

@Component({
  selector: 'app-feedback',
  templateUrl: './feedback.component.html',
  styleUrls: ['./feedback.component.css']
})
export class FeedbackComponent implements OnInit {

  selectFileName = "未选择任意图片...";

  constructor(
    private server: ServerService,
    private app: AppComponent,
  ) { }

  ngOnInit() {
    this.server.setupHight();
    this.initSelectImg();
  }

  //change the showing name after select a images
  initSelectImg() {
    let input = this.server.getEle("inputfile");
    input.addEventListener('change', function () {
      let fileName = $("#inputfile").val().toString();
      //check the type of upload images
      if (/(!?)^.*\.(jpg)|(png)|(jpeg)$/.test(fileName) == false) {
        this.app.showMsgBox(1, "请上传png 或 jpg 格式的图片哦");
        $("#inputfile").val("");
        return false;
      }
      //check the size of upload iamge
      if ((<HTMLInputElement>document.getElementById('inputfile')).files[0].size > 200 << 10) {
        this.app.showMsgBox(1, "网络压力大,请上传200kb以下的图片哦");
        $("#inputfile").val("");
        return false;
      }
      this.selectFileName = fileName;
      return true;
    }.bind(this));
  }
  //post feedback data to server to add a record
  postFeedbackForm() {
    let fb_type = $("#fbtype").val().toString();
    let fb_location = $("#fblocation").val().toString();
    let email = $("#fbemail").val().toString();
    let fbdescribe = $("#fbdescribe").val().toString();
    let image = (<HTMLInputElement>document.getElementById('inputfile')).files[0];
    let userid = this.server.userid;
    if (fb_type == "" || fb_location == "" || fbdescribe == "") {
      this.app.showMsgBox(1, "类型,位置,问题描述不能为空哦！");
      return;
    }
    let from = new FormData();
    from.append('api', "feedback");
    from.append('fb_type', fb_type);
    from.append('fb_location', fb_location);
    from.append('email', email);
    from.append('userid', userid);
    from.append('describes', fbdescribe);
    from.append('images', image);
    this.server.postFormApi(from).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, result.msg);
        return;
      }
      this.app.showMsgBox(0, "Thank You! Feedback Success!");
      $("#fblocation").val("");
      $("#fbemail").val("");
      $("#fbdescribe").val("");
      $("#inputfile").val("");
      setTimeout(() => { window.history.back();}, 3000);
    }, err => { this.app.cFail(err); })
  }
}

