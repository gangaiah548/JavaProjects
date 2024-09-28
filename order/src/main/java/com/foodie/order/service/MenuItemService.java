package com.foodie.order.service;

import java.util.Optional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.context.config.ConfigDataResourceNotFoundException;
import org.springframework.stereotype.Service;

import com.foodie.order.Entity.MenuItem;
import com.foodie.order.Entity.Restaurant;
import com.foodie.order.repository.MenuItemRepository;
import com.foodie.order.repository.RestaurantRepository;

import jakarta.annotation.Resource;

@Service
public class MenuItemService {

    @Autowired
    private MenuItemRepository menuItemRepository;

    @Autowired
    private RestaurantRepository restaurantRepository;

    public MenuItem createMenuItem(MenuItem menuItemDTO) throws Exception {
        Optional<Restaurant> orestaurant = restaurantRepository.findById(menuItemDTO.getRestaurant().getId());
        MenuItem menuItem = new MenuItem();    
        if (orestaurant.isPresent()) {
        
        menuItem.setName(menuItemDTO.getName());
        menuItem.setPrice(menuItemDTO.getPrice());
        menuItem.setRestaurant(orestaurant.get());
        }else {
        	throw new Exception("not found");
        }

        return menuItemRepository.save(menuItem);
    }
}

