package com.foodie.order.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.foodie.order.Entity.Order;

@Repository
public interface OrderRepository extends JpaRepository<Order, Long> {
}
