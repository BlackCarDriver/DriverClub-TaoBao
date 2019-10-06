import { Component, OnInit } from '@angular/core';
import { ServerService } from '../../server.service';
import { AppComponent } from '../../app.component';
import { RequestProto } from '../../struct';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {
  model = "";
  wait = 120;
  sendMsg = "发送验证码";
  isComfirm = false;
  pd:signupbody ={};
  constructor(
    private server: ServerService,
    private app: AppComponent,
  ) { }

  ngOnInit() {
    this.model = this.server.LastSection();
    document.getElementById("gohome").click();
  }

  //send the input message to server and get a comfirm code🍖
  getComfirmCode() {
    if (this.wait != 120) {
      return;
    }
    this.pd.name = $("#username").val().toString();
    let err = this.server.checkUerName(this.pd.name);
    if (err != "") {
      this.app.showMsgBox(1, err);
      return;
    }
    this.pd.password = $("#password1").val().toString();
    err = this.server.checkPassword(this.pd.password);
    if (err != "") {
      this.app.showMsgBox(1, err);
      return;
    }
    let password2 = $("#password2").val().toString();
    if (this.pd.password != password2) {
      this.app.showMsgBox(1, "输入的两个密码不一致,请检查");
      return;
    }
    this.pd.email = $("#email").val().toString();
    err = this.server.checkEmail( this.pd.email);
    if (err != "") {
      this.app.showMsgBox(1, err);
      return;
    }
    let postdata: RequestProto = {
      api: "getcomfirmcode",
      targetid:this.pd.name,
      data: this.pd,
    };
    this.server.Entrance(postdata).subscribe(result=>{
      if(result.statuscode!=0){
        this.app.showMsgBox(1,"🙈 请求未成功："+result.msg);
        return;
      }
      this.setUnchange();
      this.startTimer();
      this.isComfirm = true;
      $("#sendcode").addClass("noaction");
      this.app.showMsgBox(0, "验证码已发出，30分钟内有效，请注意查收");
    }, err=>{ this.app.cFail(err); })
  }

  //send final comfirm code and sign up message to register a new account🍖
  register() {
    if(!this.isComfirm){
      this.app.showMsgBox(1,"请先填写好注册信息并获取验证码")
      return;
    }
    this.pd.code = $('#comfirmcode').val().toString();
    let err = this.server.checkCode(this.pd.code);
    if (err!=""){
      this.app.showMsgBox(1,err);
      return;
    }
    let postdata: RequestProto = {
      api: "comfirmAndRegisit",
      targetid:this.pd.name,
      data: this.pd,
    };
    this.server.Entrance(postdata).subscribe(result=>{
      if (result.statuscode!=0){
        this.app.showMsgBox(-1, result.msg)
        return;
      }
      this.app.showMsgBox(0, "注册成功,即将前往主页！");
      setTimeout(() => { this.server.gohome(); }, 3000);
    }, err=>{this.app.cFail(err);})
  }
  //================= element control function ======================
  //make the input can't changed after comfirm code sending is request🍖
  setUnchange() {
    $("#username").attr("readonly", "true")
    $("#password1").attr("readonly", "true")
    $("#password2").attr("readonly", "true")
    $("#email").attr("readonly", "true")
    $("#email").attr("readonly", "true")
  }
  //wait 120 second before let user send comirm code again🍖
  startTimer() {
    this.wait--;
    this.sendMsg = this.wait.toString() + "s后重发";
    if (this.wait <= 0) {
      $("#sendcode").removeClass("noaction");
      this.wait = 120;
      this.sendMsg = "重新发送";
      return;
    }
    setTimeout(this.startTimer.bind(this), 1000);
  }
}

type signupbody = {
  name?:string;
  password?:string;
  email?:string;
  code?:string;
}