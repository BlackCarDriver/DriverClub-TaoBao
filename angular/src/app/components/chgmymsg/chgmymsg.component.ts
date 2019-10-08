import { Component, OnInit } from '@angular/core';
import { ServerService } from '../../server.service';
import { PersonalSetting, RequestProto } from '../../struct';
import { AppComponent } from '../../app.component';

declare var $: any;

@Component({
  selector: 'app-chgmymsg',
  templateUrl: './chgmymsg.component.html',
  styleUrls: ['./chgmymsg.component.css']
})
export class ChgmymsgComponent implements OnInit {
  headimgurl = "";
  max_user_headimg_size = 100;
  data = new PersonalSetting();
  userid = "";
  username = "";
  usersex = "BOY";
  sign = "...";
  grade = "2019";
  major = "...";
  colleage = "...";
  dorm = "...";
  email = "...";
  qq = "...";
  phone = "...";
  maindata = new PersonalSetting();

  constructor(
    private server: ServerService,
    private app: AppComponent,
  ) { }

  ngOnInit() {
    if (this.server.IsNotLogin()) {
      this.server.gohome();
    }
    this.initImgInput();
    this.initSaveBtn();
    this.getmymsg();
  }

  //==================== safe vertify ================
  //check and upload user images and get the url link to it after select a images. 🍙
  initImgInput() {
    $("#uploadheadimg").change(function () {
      //check the file name and type
      let goodsImg: File = $("#uploadheadimg").prop('files')[0];
      let imgName = goodsImg.name;
      if (imgName == "") return;
      let err = this.server.checkImgFile(goodsImg, 500);
      if (err != "") {
        alert(err);
        return;
      }
      $("#upload").trigger("click");
    }.bind(this));
  }
  //show the save button when input have been change 🍆
  initSaveBtn() {
    $(".bip").change(function () {
      $("#savebtn1").removeClass("hidden");
    });
    $(".cip").change(function () {
      $("#savebtn2").removeClass("hidden");
    });
  }

  //===================== Request funciton ================
  //get a user's base information   🍍🍈🍏🍆🍚🍙
  getmymsg() {
    if (this.server.userid == "") {
      return;
    }
    let postdata: RequestProto = {
      api: "settingdata",
      targetid: this.server.userid,
      cachetime: 60,
      cachekey: "setmsg_" + this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(1, "获取我的信息失败,请稍后再试:" + result.msg);
        return;
      }
      this.userid = this.server.userid;
      this.data = result.data;
      this.headimgurl = this.data.headimg;
      this.username = this.data.name;
      this.usersex = this.data.sex;
      this.sign = this.data.sign;
      this.grade = this.data.grade;
      this.dorm = this.data.dorm;
      this.colleage = this.data.colleage;
      this.email = this.data.emails;
      this.qq = this.data.qq;
      this.phone = this.data.phone;
      this.major = this.data.major;
      if (this.usersex == "GIRL") {
        $("#girlbtn").removeClass("isnot");
        $("#boybtn").addClass("isnot");
        this.usersex = "GIRL";
      } else {
        $("#boybtn").removeClass("isnot");
        $("#girlbtn").addClass("isnot");
        this.usersex = "BOY";
      }
    }, err => { this.app.cFail(err); });
  }
  //update a profile image and get it url after saved by server 🍍🍈🍜
  //note that the it function is called autoly after the input checking is pass
  upload() {
    if (this.server.userid == "") return
    var imgfiles = $("#uploadheadimg").prop('files');
    this.server.UploadImg(this.username, imgfiles[0]).subscribe(result => {
      if (result.statuscode == 0) {
        this.data.headimg = result.data;
        this.headimgurl = result.data;
        //update database
        let postdata: RequestProto = {
          api: "MyHeadImage",
          userid: this.server.userid,
          data: result.data,
          cachekey: "chgheadimg_"+this.server.userid,
          cachetime:60,
        };
        this.server.UpdateMessage(postdata).subscribe(result => {
          if (result.statuscode == 0) {
            this.app.showMsgBox(0, "修改成功");
          } else {
            this.app.showMsgBox(-1, "修改失败，请稍后再试：" + result.msg);
          }
        }, err => { this.app.cFail(err); });
      } else {
        this.app.showMsgBox(-1, "上传失败,请稍后再试" + result.msg);
      }
    }, err => { this.app.cFail(err); });
  }
  //update user base message of profile  🍍🍈🍙🍜
  ChangeBaseMsg() {
    let err = "";
    if (this.server.userid == "") return;
    this.data.name = $("#myname").val();
    err = this.server.checkUerName(this.data.name);
    if (err!=""){
      this.app.showMsgBox(1,err);
      return;
    }
    this.data.sign = $("#mysign").val();
    if (this.data.sign.length>50){
      this.app.showMsgBox(1,"签名太长度超出限制哦")
      return;
    }
    this.data.sex = this.usersex;
    this.data.grade = this.grade;
    this.data.colleage = $("#mycolleage").val();
    if (this.data.colleage.length>50){
      this.app.showMsgBox(1,"学院名称长度超出限制哦")
      return;
    }
    this.data.dorm = $("#mydorm").val();
    if (this.data.dorm.length>50){
      this.app.showMsgBox(1,"宿舍位置长度超出限制哦")
      return;
    }
    this.data.major = $("#mymajor").val();
    if (this.data.major.length>50){
      this.app.showMsgBox(1,"专业名称长度超出限制哦")
      return;
    }
    let postdata: RequestProto = {
      api: "changemybasemsg",
      userid: this.server.userid,
      data: this.data,
      cachekey: "chgbsmsg"+this.server.userid,
      cachetime:120,
    };
    this.server.UpdateMessage(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "修改成功");
        $("#savebtn1").addClass("hidden");
      } else {
        this.app.showMsgBox(-1, "修改失败，请稍后重试:" + result.msg);
      }
    }, err => { this.app.cFail(err); })
  }
  //update user's connect message of profile  🍍🍈🍙🍜
  ChangeContact() {
    if (this.server.userid == "") return;
    this.data.emails = $("#myemail").val();
    let err = this.server.checkEmail(this.data.emails);
    if (err!=""){
      this.app.showMsgBox(1,err);
      return;
    }
    this.data.qq = $("#myqq").val();
    if (this.data.qq.length>20){
      this.app.showMsgBox(1,"qq 长度超出限制哦");
      return;
    }
    this.data.phone = $("#myphone").val();
    if (this.data.phone.length>20){
      this.app.showMsgBox(1,"电话号码长度超出限制哦");
      return;
    }
    let postdata: RequestProto = {
      api: "MyConnectMessage",
      userid: this.server.userid,
      data: this.data,
      cachekey: "chgctmsg"+this.server.userid,
      cachetime:120,
    };
    this.server.UpdateMessage(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "修改成功");
        $("#savebtn2").addClass("hidden");
      } else {
        this.app.showMsgBox(-1, "修改失败，请扫后再试" + result.msg);
      }
    })
  }
  //=================== called by element ==================
  setgrade(grade: number) {
    $("#cancel1").removeClass("hidden");
    this.grade = grade.toString();
  }
  setboy(state: number) {
    $("#cancel1").removeClass("hidden");
    if (state == 1) {
      $("#boybtn").removeClass("isnot");
      $("#girlbtn").addClass("isnot");
      this.usersex = "BOY";
    } else {
      $("#girlbtn").removeClass("isnot");
      $("#boybtn").addClass("isnot");
      this.usersex = "GIRL";
    }
  }
  selectImg() {
    $("#cancel1").removeClass("hidden");
    $("#uploadheadimg").trigger("click");
  }

}
