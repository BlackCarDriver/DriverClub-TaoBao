import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { CharthpComponent } from './charthp.component';

describe('CharthpComponent', () => {
  let component: CharthpComponent;
  let fixture: ComponentFixture<CharthpComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ CharthpComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(CharthpComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
