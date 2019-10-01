import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ChgmymsgComponent } from './chgmymsg.component';

describe('ChgmymsgComponent', () => {
  let component: ChgmymsgComponent;
  let fixture: ComponentFixture<ChgmymsgComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ChgmymsgComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ChgmymsgComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
