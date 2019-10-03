import { Component, OnInit } from '@angular/core';
import { ServerService } from '../../server.service';
import { LoginData, MyStatus, RequestProto } from '../../struct';
import { AppComponent } from '../../app.component';
// import { LocalStorage } from '../localstorge';
//  Property 'collapse' does not exist on type 'JQuery<HTMLElement>'....
import * as bootstrap from 'bootstrap';
// import * as $ from 'jquery';
declare let $: any;

@Component({
  selector: 'app-navig',
  templateUrl: './navig.component.html',
  styleUrls: ['./navig.component.css']
})

export class NavigComponent implements OnInit {
  lidata = new LoginData();
  usermsg = new MyStatus();

  constructor(
    // private localdata: LocalStorage,
    private server: ServerService,
    private app: AppComponent,
  ) { }

  ngOnInit() {
    this.InitloginChech();
    this.showStat(false);
    this.setstate();
    this.initNav();
  }

  //===================== component control  =====================
  //show sing/regist box when click singin/reginst
  showsinginbox() {
    $("#exampleModal").modal('show');
  }
  //clear the cookie 🍓
  logout() {
    if (confirm("你确定要清楚登录状态并退出此账号？")) {
      this.server.clearAllCookie();
      window.location.reload();
    }
  }
  //decide the style of nav-user box
  showStat(islogin: boolean) {
    if (islogin) {
      $('#singin').attr("style", "display:none;");
      $('#userbox').attr("style", "display:normal;");
    } else {
      $('#singin').attr("style", "display:normal;");
      $('#userbox').attr("style", "display:none;");
    }
  }
  //hide login box
  hidelib() {
    this.server.getEle("libox-hide").click();
  }
  initNav() {
    $(".navbar-inverse").mouseleave(function () {
      $('#navhide').collapse('hide');
    });
  }
  //=========================== safety verification ===================== 
  // check the intput box content in login box 🍓🍖
  // canll autotily after it have been change
  InitloginChech() {
    $("#loginname").change(this.checkname.bind(this));
    $("#loginpassword").change(this.checkpassword.bind(this));
  }
  //check input username🍖
  checkname() {
    let res = this.server.checkUerName($("#loginname").val());
    if (res != "") this.app.showMsgBox(1, res);
  }
  //check input password🍖
  checkpassword() {
    let res = this.server.checkPassword($("#loginpassword").val());
    if (res != "") this.app.showMsgBox(1, res);
  }
  //check the content of login input🍓🍖🍚
  checkLogin() {
    let err = this.server.checkUerName($("#loginname").val());
    if (err != "") {
      this.app.showMsgBox(1, err);
      return false;
    }
    err = this.server.checkPassword($("#loginpassword").val());
    if (err != "") {
      this.app.showMsgBox(1, err);
      return false;
    }
    return true;
  }

  //========================= request function =====================
  //load userid from cookie if it is not empty then  🍋🍇🍓🍄🍔
  //select the style of nav according to login history
  setstate() {
    let userid = this.server.getCookie("ui");
    if (userid == "") return;
    this.server.userid = userid;
    this.server.token = this.server.getCookie("tk");
    let postdata: RequestProto = {
      api: "naving",
      targetid: userid,
      cachetime: 120,
    };
    postdata.cachekey = "nav_" + postdata.targetid;
    this.server.GetCredentMsg(postdata).subscribe(result => {
      if (result.statuscode == 0) {
        this.usermsg = result.data;
        this.server.username = this.usermsg.name;
        this.showStat(true);
      } else if (result.statuscode == -1000) {  //token was reject
        this.app.showMsgBox(-1, result.msg);
        this.server.clearAllCookie();
        setTimeout(() => {
          window.location.reload();
        }, 2000);
      } else {
        this.app.showMsgBox(-1, "获取登录数据失败,请稍后重试：" + result.msg);
      }
    }, error => { this.app.showMsgBox(-1, "GetMymsg fail:" + error) });
  }

  //user login, note that the username input can be id or username and email🍓🍚
  login() {
    if (this.checkLogin() != true) {
      return;
    }
    this.lidata.name = $("#loginname").val();
    this.lidata.password = $("#loginpassword").val();
    this.setstate();
    let postdata: RequestProto = {
      api: "login",
      targetid: this.lidata.name, //note that it can be username or true id
      data: this.lidata.password,
    };
    this.server.Entrance(postdata).subscribe(result => {
      if (result.statuscode != 0) {
        this.app.showMsgBox(-1, "登录失败：" + result.msg);
        return;
      }
      this.app.showMsgBox(0, "登录成功！");
      setTimeout( function(){document.location.reload();}, 2000);
      this.hidelib();
      this.usermsg = result.data;
      this.server.userid = this.usermsg.id;
      this.server.username = this.usermsg.name;
      this.server.setCookie("tk", result.msg);
      this.server.setCookie("un", this.usermsg.name);
      this.server.setCookie("up", this.lidata.password);
      this.server.setCookie("ui", this.server.userid);
      this.showStat(true);
    });
  }

}