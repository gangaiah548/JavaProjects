package com.foodie.order.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.foodie.order.Entity.Dish;

@Repository
public interface DishRepository extends JpaRepository<Dish, Long> {
}
