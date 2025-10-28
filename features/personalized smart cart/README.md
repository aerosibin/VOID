# 🛒 Smart Cart & Delivery Fee Engine

This directory contains the core logic for the **Smart Cart** feature, which processes shopping cart data, optimizes product recommendations, and calculates the final dynamic **Delivery Fee**.

The system integrates sales history and product information to provide intelligent services for the customer checkout process.

## 🚀 Key Features

  * **Intelligent Product Recommendations (Smart Cart Logic):** Uses historical sales and SKU data to analyze cart contents and suggest complimentary or alternative items to increase basket size or improve margin.
  * **Dynamic Delivery Fee Calculation:** Implements the logic (`delfeeeng.py`) to determine the final delivery charge based on real-time factors like cart value, distance, weight, and potentially current demand or priority.
  * **Data Driven:** Relies on up-to-date sales and SKU data (`sales.csv`, `skus.csv`) for accurate recommendations and fee calculations.

-----

## ⚙️ Implementation Files

| File Name | Description |
| :--- | :--- |
| `smartcart.py` / `smartcart-1.py` | Contains the **main logic for the Smart Cart**. This includes functions for parsing cart contents, running recommendation algorithms, and preparing data for the delivery fee calculation. The multiple files suggest iterative development or A/B testing versions. |
| `delfeeeng.py` | The **Delivery Fee Engine**. This module takes inputs (e.g., cart total, calculated weight/volume, delivery location) and applies a set of business rules and algorithms to output the final delivery price. |
| `sales.csv` | **Historical Transaction Data.** Used by the Smart Cart logic to identify purchase patterns (e.g., "users who bought X also bought Y") for recommendations. |
| `skus.csv` | **Product Master Data.** Contains details like item weight, dimensions, category, and base price, all critical inputs for both recommendation and delivery fee calculations. |
| `README.md` | This file. |

-----

## 🛠️ Getting Started

To run the core calculations and simulations for this feature, you will need **Python 3.x** and likely the `pandas` library for handling the CSV data.

### Prerequisites

```bash
pip install pandas
```

### Usage (Conceptual)

The scripts are designed to be imported and called by the main application service. For local testing, you might run them directly:

1.  **Ensure Data is Available:** Verify that `sales.csv` and `skus.csv` are in the root directory and contain valid data.
2.  **Run Smart Cart Logic:**
    ```python
    import smartcart
    # Assume cart_data is a dictionary representing a user's current basket
    recommendations = smartcart.get_recommendations(cart_data)
    ```
3.  **Calculate Delivery Fee:**
    ```python
    import delfeeeng
    # Prepare inputs from the cart and user profile
    delivery_inputs = smartcart.prepare_delivery_inputs(cart_data, user_profile)
    final_fee = delfeeeng.calculate_fee(delivery_inputs)
    ```

-----

## 📊 Data Schema Highlights

### `skus.csv` (Product Data)

| Column Name | Type | Purpose |
| :--- | :--- | :--- |
| `sku_id` | String | Unique product identifier. |
| `weight_g` | Integer | Used by `delfeeeng.py` for weight-based shipping costs. |
| `category` | String | Used by `smartcart.py` for category-based recommendations. |
| `base_price` | Float | Item's unit price. |

### `sales.csv` (Historical Transactions)

| Column Name | Type | Purpose |
| :--- | :--- | :--- |
| `order_id` | Integer | Unique identifier for a sale. |
| `sku_id` | String | The product purchased. |
| `user_id` | Integer | Allows tracking co-purchased items for recommendations. |
| `timestamp` | Datetime | Used for time-based trend analysis. |

-----

Would you like to see a small Python snippet detailing the inputs and expected outputs for the `delfeeeng.py` script?
