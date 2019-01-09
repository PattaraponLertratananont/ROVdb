import { Component, OnInit } from '@angular/core';
import { GetHeroService } from '../get-hero.service'
import { Hero } from '../Hero'
@Component({
  selector: 'app-hero-list',
  templateUrl: './hero-list.component.html',
  styleUrls: ['./hero-list.component.css']
})

export class HeroListComponent implements OnInit {


  class: any[] = []

  constructor(private serviceGet: GetHeroService) {

  }
  ngOnInit() {
    this.get()
  }
  get() {
    this.serviceGet.getClassMage()
      .subscribe(data => {this.class = data,console.log(data)})
  }
}
interface CLASS {
  class: string;
}