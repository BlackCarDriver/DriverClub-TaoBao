import { Component, OnInit } from '@angular/core';
import {ServerService} from '../server.service';
import {PersonalSetting, UpdateResult} from '../struct';

declare var $: any;


@Component({
  selector: 'app-chgmymsg',
  templateUrl: './chgmymsg.component.html',
  styleUrls: ['./chgmymsg.component.css']
})
export class ChgmymsgComponent implements OnInit {
   headimgurl = "https://tb1.bdstatic.com/tb/r/image/2018-02-11/7ec7062f14307db6f1728bc108c3189c.jpeg";
   userid = "00001";
   data = new PersonalSetting();
   updateresult = new UpdateResult();
  //绑定到表单的数据的默认值
   username = "未设置";
   usersex = "BOY";
   sign = "Welcome to BlackCarDriver.cn";
   grade = "2019";
   colleage = "未设置";
   dorm = "未设置";
   email = "保密";
   qq = "保密";
   phone = "保密";
  //上传到服务器和请求获取的数据
   maindata = new PersonalSetting();
  
  constructor(private server : ServerService) {}

  ngOnInit() {
    //初始化组件事件
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
    })
    //获取用户已有的信息
    this.getmymsg();
  }

  //获取用户的基本信息
  getmymsg(){
    this.server.GetMyMsg(this.userid, "setdata").subscribe(
      result=>{
        this.data = result;
        this.headimgurl = this.data.headimg;
        this.username = this.data.name;
        this.userid = this.data.id;
        this.usersex = this.data.sex;
        this.sign = this.data.sign;
        this.grade = this.data.grade;
        this.colleage = this.data.colleage;
        this.email = this.data.emails;
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
  
  //上传选中的头像文件并更新imgurl的值
  upload(){
    var imgfiles = $("#uploadheadimg").prop('files');
    console.log(imgfiles[0]);
    this.server.UploadImg(this.username,imgfiles[0]).subscribe(result=>{
      if( result.status>=0){
        this.headimgurl = result.imgurl;
        this.data.headimg = result.imgurl;
        this.server.UpdateMessage(this.userid, "MyHeadImage", this.data).subscribe(result=>{
            if(result.status>=0){
              alert("修改成功！");
            }else{
              alert(result.describe);
            }
        });
      }else{
        alert(result.describe);
      }
    });
  }

  //修改或设置基本信息并上传到服务器
  ChangeBaseMsg(){
    this.data.name = $("#myname").val();
    this.data.colleage = $("mycolleage").val();
    this.data.sign = $("#mysign").val();
    this.data.dorm =  $("#mydorm").val();
    this.data.sex =  this.usersex;
    this.data.grade = this.grade;
    this.server.UpdateMessage(this.userid, "MyBaseMessage", this.data).subscribe(result=>{
      this.updateresult = result; 
      if (this.updateresult.status >= 0) {
        alert("修改成功！");
      }else{
        alert(this.updateresult.describe);
      }
    })
  }
  
  //修改或设置联系方式信息并上传到服务器
  ChangeContact(){
    this.data.emails = $("#myemail").val();
    this.data.qq = $("#myqq").val();
    this.data.phone = $("#myphone").val();
    this.server.UpdateMessage(this.userid, "MyConnectMessage", this.data).subscribe(result=>{
      if(result.status >= 0) {
        alert("修改成功！");
      }else{
        alert(result.describe);
      }
    })
}
  //=================== 设置组件 ==================

  //设置年级选择按钮事件
  setgrade(grade:number){
    $("#cancel1").removeClass("hidden");
   this.grade = grade.toString();
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

  //还原输入框内容
  ClearBaseMsg(){
    $("#cancel1").addClass("hidden");
    $("#myname").val("");
    $("#mysign").val("");
    $("#mycolleage").val("");
    this.usersex = this.data.sex;
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
