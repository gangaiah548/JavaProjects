package com.example.mysql;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.data.jpa.repository.config.EnableJpaAuditing;

@SpringBootApplication
@EnableJpaAuditing
public class JavaMysqlApplication {

	public static void main(String[] args) {
		SpringApplication.run(JavaMysqlApplication.class, args);
	}

}
