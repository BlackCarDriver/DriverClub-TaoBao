import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { UploadgoodsComponent } from './uploadgoods.component';

describe('UploadgoodsComponent', () => {
  let component: UploadgoodsComponent;
  let fixture: ComponentFixture<UploadgoodsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ UploadgoodsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(UploadgoodsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
