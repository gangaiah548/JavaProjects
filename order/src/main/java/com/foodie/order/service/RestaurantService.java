package com.foodie.order.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.foodie.order.Entity.Restaurant;
import com.foodie.order.repository.RestaurantRepository;

import java.util.List;
import java.util.Optional;

@Service
public class RestaurantService {

    @Autowired
    private RestaurantRepository restaurantRepository;

    // Create a new restaurant
    public Restaurant createRestaurant(Restaurant restaurant) {
        return restaurantRepository.save(restaurant);
    }

    // Get all restaurants
    public List<Restaurant> getAllRestaurants() {
        return restaurantRepository.findAll();
    }

    // Get a restaurant by ID
    public Optional<Restaurant> getRestaurantById(Long id) {
        return restaurantRepository.findById(id);
                //.orElseThrow(() -> new ResourceNotFoundException("Restaurant not found"));
    }

    // Update a restaurant
    public Restaurant updateRestaurant(Long id, Restaurant restaurant) throws Exception {
        if (!restaurantRepository.existsById(id)) {
            throw new Exception("Restaurant not found");
        }
        restaurant.setId(id); // Ensure the ID is set for the update
        return restaurantRepository.save(restaurant);
    }

    // Delete a restaurant
    public void deleteRestaurant(Long id) throws Exception {
        if (!restaurantRepository.existsById(id)) {
            throw new Exception("Restaurant not found");
        }
        restaurantRepository.deleteById(id);
    }
}

