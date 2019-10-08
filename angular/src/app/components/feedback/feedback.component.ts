import { Component, OnInit } from '@angular/core';
import { AppComponent } from '../../app.component';
import { ServerService } from '../../server.service';

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
    this.initSelectImg();
  }

  //change the showing name after select a images 🍙
  initSelectImg() {
    let input = this.server.getEle("inputfile");
    input.addEventListener('change', function () {
      let img:File = $("#inputfile").prop('files')[0];
      let err = this.server.checkImgFile(img, 1000);
      if (err != "") {
        alert(err);
        return;
      }
      this.selectFileName = img.name;
      return true;
    }.bind(this));
  }

  //post feedback data to server to add a record  🍙
  postFeedbackForm() {
    let fb_type = $("#fbtype").val().toString();
    let fb_location = $("#fblocation").val().toString();
    if (fb_location.length>200) {
      this.app.showMsgBox(1,"反馈位置描述超出限制哦 _(:з」∠)_")
      return;
    }
    let email = $("#fbemail").val().toString();
    let err= this.server.checkEmail(email);
    if (email !="" &&  err!="") {
      this.app.showMsgBox(1,err);
      return ;
    }
    let fbdescribe = $("#fbdescribe").val().toString();
    if (fbdescribe.length>480){
      this.app.showMsgBox(1,"问题描述长度超出限制哦 _(:з」∠)_");
      return;
    }
    if (fb_type == "" || fb_location == "" || fbdescribe == "") {
      this.app.showMsgBox(1, "类型,位置,问题描述不能为空哦！");
      return;
    }
    let image = (<HTMLInputElement>document.getElementById('inputfile')).files[0];
    let from = new FormData();
    from.append('api', "feedback");
    from.append('fb_type', fb_type);
    from.append('fb_location', fb_location);
    from.append('email', email);
    from.append('userid', this.server.userid);
    from.append('describes', fbdescribe);
    from.append('images', image);
    this.server.postFormApi(from).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "反馈被拒绝："+result.msg);
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

