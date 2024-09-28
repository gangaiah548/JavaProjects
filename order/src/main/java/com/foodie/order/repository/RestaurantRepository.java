package com.foodie.order.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.foodie.order.Entity.Restaurant;

public interface RestaurantRepository extends JpaRepository<Restaurant,Long>{

}
