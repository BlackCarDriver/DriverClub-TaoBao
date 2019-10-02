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
      window.history.back();
    }
    this.initImgInput();
    this.initSaveBtn();
    this.getmymsg();
  }

  //==================== safe vertify ================
  //check and upload user images and get the url link to it after select a images.
  initImgInput(){
      $("#uploadheadimg").change(function(evt) {
        if ($(this).val() == '') {
          return;
        }
        //check the file name and type
        var filename = $(this).val().replace(/.*(\/|\\)/, "");
        var pos = filename.lastIndexOf(".");
        var filetype = filename.substring(pos, filename.length)
        if (filetype != ".jpg" && filetype != ".png") {
          alert("请上传 jpg 或 png 格式的图片")
          return;
        }
        //check the image size
        var files = evt.currentTarget.files;
        var filesize = files[0].size;
        if (filesize > 100 * 1024) {
          alert( "请上传100kb 以下的图片");
          return;
        }
        $("#upload").trigger("click");
      });
  }
  //show the save button when input have been change 🍆
  initSaveBtn(){
    $(".bip").change(function () {
      $("#savebtn1").removeClass("hidden");
    });
    $(".cip").change(function () {
      $("#savebtn2").removeClass("hidden");
    });
  }

  //===================== Request funciton ================
  //get a user's base information   🍍🍈🍏🍆
  getmymsg() {
    let postdata: RequestProto = {
      api: "settingdata",
      targetid: this.server.userid,
    };
    this.server.GetMyMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
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
      } else {
        this.app.showMsgBox(1, "请求失败,请稍后再试", result.msg);
      }
    }, error => { console.log(error) });
  }
  //update a profile image and get it url after saved by server 🍍🍈
  //note that the it function is called autoly after the input checking is pass
  upload() {
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
        };
        this.server.UpdateMessage(postdata).subscribe(result => {
          if (result.statuscode == 0) {
            this.app.showMsgBox(0, "修改成功");
          } else {
            this.app.showMsgBox(-1, "修改失败，请稍后再试", result.msg);
          }
        }, error => { console.log("UpdateMessage() fail: " + error); });
      } else {
        this.app.showMsgBox(-1, "上传失败,请稍后再试", result.msg);
      }
    }, error => { console.log("UploadImg() fail: " + error) });
  }
  //update user base message of profile  🍍🍈
  ChangeBaseMsg() {
    this.data.name = $("#myname").val();
    this.data.sign = $("#mysign").val();
    this.data.sex = this.usersex;
    this.data.grade = this.grade;
    this.data.colleage = $("#mycolleage").val();
    this.data.dorm = $("#mydorm").val();
    this.data.major = $("#mymajor").val();
    let postdata: RequestProto = {
      api: "changemybasemsg",
      userid: this.server.userid,
      data: this.data,
    };
    this.server.UpdateMessage(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "修改成功");
        $("#savebtn1").addClass("hidden");
      } else {
        this.app.showMsgBox(-1, "修改失败，请稍后重试:"  + result.msg);
      }
    }, error => { console.log("UpdateMessage() fail: " + error); })
  }
  //update user's connect message of profile  🍍🍈
  ChangeContact() {
    this.data.emails = $("#myemail").val();
    this.data.qq = $("#myqq").val();
    this.data.phone = $("#myphone").val();
    let postdata: RequestProto = {
      api: "MyConnectMessage",
      userid: this.server.userid,
      data: this.data,
    };
    this.server.UpdateMessage(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.app.showMsgBox(0, "修改成功");
        $("#savebtn2").addClass("hidden");
      } else {
        this.app.showMsgBox(-1, "修改失败，请扫后再试", result.msg);
      }
    }, error => {
      this.app.showMsgBox(-1, "请求失败，请扫后再试", error);
    }
    )
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
