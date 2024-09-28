package com.foodie.order.Controller;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.foodie.order.Entity.ModelOrder;
import com.foodie.order.Entity.Order;
import com.foodie.order.service.OrderService;





@RestController
@RequestMapping("/api/orders")
public class OrderController {

	 @Autowired
	    private OrderService orderService;

	    // Get all orders
	    @GetMapping
	    public ResponseEntity<List<Order>> getAllOrders() {
	        List<Order> orders = orderService.getAllOrders();
	        return ResponseEntity.ok(orders);
	    }

	    // Get order by ID
	    @GetMapping("/{orderId}")
	    public ResponseEntity<Order> getOrderById(@PathVariable Long orderId) {
	        Optional<Order> order = orderService.getOrderById(orderId);
	        return order.map(ResponseEntity::ok).orElseGet(() -> ResponseEntity.notFound().build());
	    }

	  

	    // Delete an order by ID
	    @DeleteMapping("/{orderId}")
	    public ResponseEntity<Void> deleteOrder(@PathVariable Long orderId) {
	        Optional<Order> order = orderService.getOrderById(orderId);
	        if (order.isPresent()) {
	            orderService.deleteOrder(orderId);
	            return ResponseEntity.noContent().build();
	        } else {
	            return ResponseEntity.notFound().build();
	        }
	    }
    @PostMapping
    public ResponseEntity<Order> createOrder(@RequestBody ModelOrder order) {
    	
    	Order newOrder = new Order();
        newOrder.setCustomerName(order.getCustomerName());

        Order savedOrder = orderService.createOrder(newOrder, order.getMenuItemIds(), order.getQuantities());
        return ResponseEntity.ok(savedOrder);
        
    }

 
}
