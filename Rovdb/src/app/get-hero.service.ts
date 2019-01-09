import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Hero } from './Hero';

@Injectable({
  providedIn: 'root'
})
export class GetHeroService {

  apiUrl='http://localhost:1323'

  constructor(private http:HttpClient) { }

  getClassMage():Observable<Hero[]>{
    return this.http.get<Hero[]>(this.apiUrl+'/hero')
  }
}
