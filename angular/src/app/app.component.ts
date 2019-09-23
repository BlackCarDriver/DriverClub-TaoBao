import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  showMsg = ""; //the text show in model box
  showStatus = 0; //the status of model box
  ngOnInit() {
    
  }
  public showMsgBox(status:number, msg:string, err?:string){
    this.showMsg = msg;
    this.showStatus = status;
    console.log(msg+":"+err);
    $('#showbtn').click();
  }

  //display a wang dialogy to show that connect fail🍄
  public cFail(reason:string){
      this.showMsgBox(1,"请求失败，请稍后再试 :(" , reason);
  }
}
