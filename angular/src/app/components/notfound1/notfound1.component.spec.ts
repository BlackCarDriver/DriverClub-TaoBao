import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { Notfound1Component } from './notfound1.component';

describe('Notfound1Component', () => {
  let component: Notfound1Component;
  let fixture: ComponentFixture<Notfound1Component>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ Notfound1Component ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(Notfound1Component);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
