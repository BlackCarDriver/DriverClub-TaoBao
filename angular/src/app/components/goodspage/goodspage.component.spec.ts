import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GoodspageComponent } from './goodspage.component';

describe('GoodspageComponent', () => {
  let component: GoodspageComponent;
  let fixture: ComponentFixture<GoodspageComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GoodspageComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GoodspageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
