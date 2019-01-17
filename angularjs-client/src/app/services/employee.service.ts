import { Injectable } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import { Observable } from 'rxjs/Rx';
import { Employee } from '../models/employee';

@Injectable()
export class EmployeeService {
  constructor(private http: Http) { }
  getEmployees(): Observable<Employee[]> {
    return this.http.get('http://localhost:8080/employees').map((res: Response) => res.json())
    .catch((error: any) => Observable.throw(error.json().
    error || 'Server error'));
  }
  addEmployee(employee: Employee): Observable<Employee> {
    const headers = new Headers({ 'Content-Type': 'application/json' });
    const options = new RequestOptions({ headers: headers });
    return this.http.post('http://localhost:8080/employee/add', employee, options).map(this.extractData)
    .catch(this.handleErrorObservable);
  }
  private extractData(res: Response) {
    const body = res.json();
    return body || {};
  }
  private handleErrorObservable(error: Response | any) {
    console.error(error.message || error);
    return Observable.throw(error.message || error);
  }
}
