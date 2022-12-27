package com.Javasecurity.config;

import org.springframework.security.config.annotation.authentication.builders.AuthenticationManagerBuilder;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityConfigurerAdapter;


public class SecurityConfig extends WebSecurityConfigurerAdapter {
	
@Override
public void configure(AuthenticationManagerBuilder auth) throws Exception {
	auth.inMemoryAuthentication().withUser("abc")
	.password("abc").roles("USER");
}
@Override
public void configure(HttpSecurity http) throws Exception {
	http.antMatcher("/**").authorizeRequests().anyRequest().hasRole("USER")
			.and().formLogin().loginPage("/login.jsp")
			.failureUrl("/login.jsp?error=1").loginProcessingUrl("/login")
			.permitAll().and().logout()
			.logoutSuccessUrl("/listEmployees.html");

}
}
