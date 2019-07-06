import { Component, OnInit } from '@angular/core';
import {ServerService} from '../server.service';
import {PersonalBase} from '../struct';

declare var $: any;


@Component({
  selector: 'app-chgmymsg',
  templateUrl: './chgmymsg.component.html',
  styleUrls: ['./chgmymsg.component.css']
})
export class ChgmymsgComponent implements OnInit {
   headimgurl = "http://localhost:8090/source/images?tag=headimg&&name=BlackCarDriver.png";
   userid = "1234567";
   data = new PersonalBase();
  //绑定到表单的数据的默认值
   username = "未设置";
   usersex = "BOY";
   sign = "Welcome to BlackCarDriver.cn";
   grade = "2019";
   colleage = "未设置";
   email = "保密";
   qq = "保密";
   phone = "保密";
  //上传到服务器和请求获取的数据
   maindata = new PersonalBase();
  
  constructor(private server : ServerService) {}

  ngOnInit() {

    $(document).ready(function(){

      //解决下拉菜单按钮不能下拉
      $(".dropdown-toggle").on('click',function(){
          $('.dropdown-toggle').dropdown();
      });

      //选择头像后检查类型,上传头像,获取url连接
      $("#uploadheadimg").change(function(evt){
        if($(this).val() == ''){ 
          return; 
        } 
       //判断文件类型，并获取文件名到页面
       var filename = $(this).val().replace(/.*(\/|\\)/, "");
       var pos = filename.lastIndexOf(".");
       var filetype = filename.substring(pos,filename.length)  //此处文件后缀名也可用数组方式获得str.split(".") 
       if (filetype.toLowerCase()!=".jpg" && filetype.toLowerCase()!=".png"){
          alert("请上传 png 或 jpg 格式的图片");
          return;
       }
       //判断文件大小
       var files = evt.currentTarget.files;
       var filesize = files[0].size;
       if(filesize> 50 * 1024){
         alert("请上传50kb 以下的图片");
         return;
       }
        //检查无误，可以上传,通过按钮点击时间间接激发
        $("#upload").trigger("click");
      });
      //当表单被改变是显示取消按钮
      $(".baseinput").change(function(){
        $("#cancel1").removeClass("hidden");
      });
      $(".contactinput").change(function(){
        $("#cancel2").removeClass("hidden");
      });
    })//document.read() over


    //获取页面信息
    this.getmymsg();
  }

  //获取用户的基本信息
  getmymsg(){
    this.server.Getmymsg(this.userid).subscribe(
      result=>{
        this.data = result;
        this.username = this.data.username;
        this.userid = this.data.userid;
        this.usersex = this.data.usersex;
        this.sign = this.data.sign;
        this.grade = this.data.grade;
        this.colleage = this.data.colleage;
        this.email = this.data.email;
        this.qq = this.data.qq;
        this.phone = this.data.phone;
        if(this.usersex=="GIRL"){
          $("#girlbtn").removeClass("isnot");
          $("#boybtn").addClass("isnot");
          this.usersex = "GIRL";
        }else{
          $("#boybtn").removeClass("isnot");
          $("#girlbtn").addClass("isnot");
          this.usersex = "BOY";
        }
      });
  }
  //设置年级按钮事件
  setgrade(grade:number){
    $("#cancel1").removeClass("hidden");
   this.grade= grade.toString();
  }
  //选择性别按钮事件
  setboy(state :number){
    $("#cancel1").removeClass("hidden");
    if(state == 1){
      $("#boybtn").removeClass("isnot");
      $("#girlbtn").addClass("isnot");
      this.usersex = "BOY";
    }else{
      $("#girlbtn").removeClass("isnot");
      $("#boybtn").addClass("isnot");
      this.usersex = "GIRL";
    }
  }
  //点击修改头像后激活input
  selectImg(){
    $("#cancel1").removeClass("hidden");
    $("#uploadheadimg").trigger("click");
  }
  //上传选中的头像文件并更新imgurl的值
  upload(){
    var imgfiles = $("#uploadheadimg").prop('files');
    console.log(imgfiles[0]);
    this.server.UploadHeadImg(this.username,imgfiles[0]).subscribe(result=>{
      this.headimgurl = result;
      alert(result);
    });
  }
  //修改或设置基本信息并上传到服务器
  ChangeBaseMsg(){
    this.data.username = $("#myname").val();
    this.data.colleage = $("mycolleage").val();
    this.data.sign = $("#mysign").val();
    this.data.usersex =  this.usersex;
    this.data.grade = this.grade;
    this.server.UploadMyBaseMsg(this.data).subscribe(result=>{
      console.log(result);
    })
  }
  //修改或设置基本信息并上传到服务器
  ChangeContact(){
      this.data.email = $("#myemail").val();
      this.data.qq = $("#myqq").val();
      this.data.phone = $("#myphone").val();
      this.server.UploadContactMsg(this.data).subscribe(result=>{
        console.log(result);
      })
  }
  //还原输入框内容
  ClearBaseMsg(){
    $("#cancel1").addClass("hidden");
    $("#myname").val("");
    $("#mysign").val("");
    $("#mycolleage").val("");
    this.usersex = this.data.usersex;
    this.grade = this.data.grade;
    if(this.usersex=="GIRL"){
      this.setboy(0);
    }else this.setboy(1);
  }
  //还原联系方式输入框
  ClearContactMsg(){
    $("#myemail").val("");
    $("#myqq").val("");
    $("#myphone").val("");
    $("#cancel2").addClass("hidden");
  }
}
