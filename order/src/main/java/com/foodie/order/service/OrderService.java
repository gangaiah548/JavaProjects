package com.foodie.order.service;



import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.foodie.order.Entity.Dish;
import com.foodie.order.Entity.MenuItem;
import com.foodie.order.Entity.Order;
import com.foodie.order.Exception.MenuItemNotFoundException;
import com.foodie.order.Exception.OrderNotFoundException;
import com.foodie.order.repository.MenuItemRepository;
import com.foodie.order.repository.OrderRepository;

import java.util.List;
import java.util.Optional;

@Service
public class OrderService {

	 @Autowired
	    private OrderRepository orderRepository;

	    @Autowired
	    private MenuItemRepository menuItemRepository;

	    public List<Order> getAllOrders() {
	        return orderRepository.findAll();
	    }

	    public Order getOrderById(Long id) {
	        return orderRepository.findById(id).orElseThrow(() -> new OrderNotFoundException(id));
	        }

	    public Order createOrder(Order order, List<Long> menuItemIds, List<Integer> quantities) {
	        // For each dish in the order, fetch the MenuItem from RestaurantDB
	        for (int i = 0; i < menuItemIds.size(); i++) {
	            Long menuItemId = menuItemIds.get(i);
	            int quantity = quantities.get(i);

	            MenuItem menuItem = menuItemRepository.findById(menuItemId)
	                .orElseThrow(() -> new MenuItemNotFoundException("MenuItem not found with id: " + menuItemId));

	            Dish dish = new Dish();
	            dish.setMenuItem(menuItem);
	            dish.setQuantity(quantity);
	            dish.setOrder(order);

	            order.getDishes().add(dish);
	        }

	        // Save the order and associated dishes
	        return orderRepository.save(order);
	    }

	    public void deleteOrder(Long id) {
	    	 Order order = orderRepository.findById(id).orElseThrow(() -> new OrderNotFoundException(id));
	         orderRepository.delete(order);
	    }
}

