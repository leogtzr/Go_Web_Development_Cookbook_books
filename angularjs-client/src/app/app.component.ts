import { Component, OnInit } from '@angular/core';
import { EmployeeService } from './services/employee.service';
import { Employee } from './models/employee';

@Component
({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  title = 'app';
  employee = new Employee('', '', '');
  employees;
  constructor(private employeeService: EmployeeService) { }
  ngOnInit(): void {
    this.getEmployees();
  }
  getEmployees(): void {
    this.employeeService.getEmployees()
    .subscribe(employees => this.employees = employees);
  }
  addEmployee(): void {
    this.employeeService.addEmployee(this.employee)
    .subscribe
    (
      employee => {
        this.getEmployees();
        this.reset();
      }
    );
  }
  private reset() {
    this.employee.id = null;
    this.employee.firstName = null;
    this.employee.lastName = null;
  }
}
