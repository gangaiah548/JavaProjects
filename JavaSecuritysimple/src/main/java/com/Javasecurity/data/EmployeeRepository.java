package com.Javasecurity.data;

import org.springframework.data.jpa.repository.JpaRepository;

import com.Javasecurity.model.Employee;

public interface EmployeeRepository extends JpaRepository<Employee, Long>{

}