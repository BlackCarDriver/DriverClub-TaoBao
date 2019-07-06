import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { NavigComponent } from './navig.component';

describe('NavigComponent', () => {
  let component: NavigComponent;
  let fixture: ComponentFixture<NavigComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ NavigComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(NavigComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
