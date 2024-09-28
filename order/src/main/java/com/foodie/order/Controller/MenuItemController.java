package com.foodie.order.Controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.foodie.order.Entity.MenuItem;
import com.foodie.order.service.MenuItemService;

@RestController
@RequestMapping("/menu-items")
public class MenuItemController {

    @Autowired
    private MenuItemService menuItemService;

    @PostMapping
    public ResponseEntity<?> createMenuItem(@RequestBody MenuItem menuItemDTO) {
        MenuItem createdMenuItem;
		try {
			createdMenuItem = menuItemService.createMenuItem(menuItemDTO);
			return ResponseEntity.status(201).body(createdMenuItem);
		} catch (Exception e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
			return ResponseEntity.status(500).body(e.getMessage());
		}
        
    }
}
