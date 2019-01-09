import { Component, OnInit } from '@angular/core';
import { SelectItem } from 'primeng/api';

@Component({
  selector: 'app-hero-list',
  templateUrl: './hero-list.component.html',
  styleUrls: ['./hero-list.component.css']
})

export class HeroListComponent implements OnInit {

  class: Dropdown[];
  role: Dropdown[];

  constructor() {
    //SelectItem API with label-value pairs
    this.class = [
      { name: 'เมจ' },
      { name: 'เมจ/แทงค์' },
      { name: 'แอสซาซิน' },
    ];
    this.role = [
      {name:'ว่องไว/ฟาร์ม'},
      {name:'ควบคุม/ว่องไว'},
    ]
  }
  ngOnInit() {
  }

}
interface Dropdown {
  name: string;
}