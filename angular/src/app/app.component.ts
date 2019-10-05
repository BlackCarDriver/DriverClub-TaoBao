import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  showMsg = ""; //the text show in model box
  showStatus = 0; //the status of model box
  imghref = ""; //the images showing in model box
  ngOnInit() {
    $(".window").on("click", function () {
      $('.navbar-collapse').collapse('hide'); 
    })
  }
  public showMsgBox(status: number, msg: string, err?: string) {
    this.showMsg = msg;
    this.showStatus = status;
    console.log(msg + ":" + err);
    $('#showbtn').click();
  }

  //display a wang dialogy to show that connect fail🍄
  public cFail(reason: string) {
    this.showMsgBox(1, "请求失败，请稍后再试 :(", reason);
  }
  //display a images in model box 🍛
  public ShowImg(href: string) {
    if (href == "") return;
    this.imghref = href.replace("/_", "/");
    $("#showimg").click();
  }
  //close the images model box🍛
  closeimg() {
    $("#imgbox").modal("hide");
  }
}
