import { Component, OnInit } from '@angular/core';
import { ServerService } from '../server.service';
import { LoginData, MyStatus, RequertResult, RequestProto } from '../struct';
import { AppComponent } from '../app.component';
// import { LocalStorage } from '../localstorge';
//  Property 'collapse' does not exist on type 'JQuery<HTMLElement>'....
import * as bootstrap from 'bootstrap';
// import * as $ from 'jquery';
declare let $: any;

// regex of account name 
const namereg = /^[\u4e00-\u9fa5_a-zA-Z0-9]{2,15}$/;
// regex of password
const passwordreg = /^[a-zA-Z._0-9]{6,20}$/;

// the return state 
const scuess = 1;
const enable = 2;
const disable = -2;
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
  }

  //===================== component control  =====================
  //show sing/regist box when click singin/reginst
  showsinginbox() {
    $("#exampleModal").modal('show');
  }
  //clear the cookie 🍓
  logout() {
    if (confirm("你确定要清楚登录状态并退出此账号？")) {
      //TODO: clear the cookie
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
  //=========================== safety verification ===================== 

  // check the intput box content in login box 🍓
  // canll autotily after it have been change
  InitloginChech() {
    $("#loginname").change(function () {
      if (namereg.test($("#loginname").val()) == false) {
        this.app.showMsgBox(1,"用户名格式不正确,提示：不包含空格,符号,长度为2~15");
      }
    });
    $("#loginpassword").change(function () {
      if (passwordreg.test($("#loginpassword").val()) == false) {
        this.app.showMsgBox(1,"密码格式不正确,提示：6~20个字母或数字或._组成");
      }
    });
  }
  //check the input of login input🍓
  checkLogin() {
    let worngnum = 0;
    if (namereg.test($("#loginname").val()) == false) {
      worngnum++;
    }
    if (passwordreg.test($("#loginpassword").val()) == false) {
      worngnum++;
    }
    return (worngnum == 0);
  }
  //========================= request function =====================
  //load userid from cookie if it is not empty then  🍋🍇🍓
  //select the style of nav according to login history
  setstate() {
    if (this.server.userid != "") {
      let postdata: RequestProto = {
        api: "naving",
        targetid: this.server.userid,
      };
      this.server.GetCredentMsg(postdata).subscribe(result => {
        if (result.statuscode == 0) {
          this.usermsg = result.data;
          this.server.username = this.usermsg.name;
          this.showStat(true);  //show user message in naving bar
        } else {
          this.app.showMsgBox(-1, "获取登录数据失败,请稍后重试：" + result.msg);
        }
      }, error => { this.app.showMsgBox(-1, "GetMymsg fail:" + error) });
    }
  }

  //user login, note that the username input can be id or username 🍓
  loging() {
    this.lidata.name = $("#loginname").val();
    this.lidata.password = $("#loginpassword").val();
    this.setstate();
    if (this.checkLogin() != true) {
      this.app.showMsgBox(1,"输入的格式不正确,请检查");
      return;
    }
    let postdata: RequestProto = {
      api: "login",
      targetid: this.lidata.name, //note that it can be username or true id
      data:this.lidata.password,
    };
    this.server.Entrance(postdata).subscribe(result => {
      if (result.statuscode!=0){
        this.app.showMsgBox(-1, "登录失败："+result.msg);
        return;
      }
      this.app.showMsgBox(0, "登录成功！");
      this.usermsg = result.data;
      this.server.userid = this.usermsg.id;
      this.server.username = this.usermsg.name;
      //TODO: save cookie
    });
  }

}