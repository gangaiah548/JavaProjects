package com.foodie.order.Controller;

import java.util.List;
import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
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

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.tags.Tag;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

@RestController
@RequestMapping("/api/orders")
@Tag(name = "Order Management", description = "APIs related to Order Management")
public class OrderController {

	private static final Logger logger = LoggerFactory.getLogger(OrderController.class);

	@Autowired
	private OrderService orderService;

	@GetMapping
    @Operation(summary = "Get all orders", description = "Retrieve a list of all orders")
	public ResponseEntity<List<Order>> getAllOrders() {
		logger.info("Fetching all orders");
		List<Order> orders = orderService.getAllOrders();
		logger.info("Total orders found: {}", orders.size());
		return ResponseEntity.ok(orders);
	}

	// Get order by ID
	@GetMapping("/{orderId}")
    @Operation(summary = "Get order by ID", description = "Retrieve an order by its ID")
	public ResponseEntity<Order> getOrderById(@PathVariable Long orderId) {
		logger.info("Fetching order with ID: {}", orderId);
		Order order = orderService.getOrderById(orderId);
		// return order.map(ResponseEntity::ok).orElseGet(() ->
		// ResponseEntity.notFound().build());
		logger.error("Order with ID {} not found", orderId);
		return new ResponseEntity<Order>(order, HttpStatus.OK);
	}

	// Delete an order by ID
	@DeleteMapping("/{orderId}")
	public ResponseEntity<Void> deleteOrder(@PathVariable Long orderId) {
		logger.info("Deleting order with ID: {}", orderId);
		orderService.deleteOrder(orderId);
		return ResponseEntity.noContent().build();

	}

	@PostMapping
	@Operation(summary = "Create an order", description = "Create a new order")
	public ResponseEntity<Order> createOrder(@RequestBody ModelOrder order) {
		logger.info("Creating order for customer: {}", order.getCustomerName());
		Order newOrder = new Order();
		newOrder.setCustomerName(order.getCustomerName());
		if (order.getMenuItemIds().size() != order.getQuantities().size()) {
			throw new IllegalArgumentException("The number of menu items and quantities must match.");
		}

		Order savedOrder = orderService.createOrder(newOrder, order.getMenuItemIds(), order.getQuantities());
		logger.info("Created order", order);
		return ResponseEntity.ok(savedOrder);

	}

}
