package com.foodie.order.Entity;
import java.util.List;

import io.micrometer.common.lang.NonNull;
import lombok.Data;

@Data
public class ModelOrder {
	@NonNull
	String customerName;
	@NonNull
	List<Long> menuItemIds;
	@NonNull
	List<Integer> quantities;

}
