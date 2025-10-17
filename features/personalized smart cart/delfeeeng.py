# Delivery Fee Engine - Single File Implementation

def calculate_base_fee():
    base_fee = 20.0  # Fixed starting delivery fee
    return base_fee

def calculate_surge_fee(num_orders, peak_orders, base_fee, max_surge_multiplier=2.0):
    surge_ratio = num_orders / peak_orders if peak_orders > 0 else 0
    surge_multiplier = 1 + (max_surge_multiplier - 1) * min(surge_ratio, 1)
    surge_fee = base_fee * surge_multiplier - base_fee
    return surge_fee

def calculate_distance_fee(distance_km, rate_per_km=5):
    return distance_km * rate_per_km

def calculate_weather_fee(base_fee, weather_condition):
    weather_multiplier = 1.0
    if weather_condition == 'rain':
        weather_multiplier = 1.15
    elif weather_condition == 'storm':
        weather_multiplier = 1.3
    return base_fee * (weather_multiplier - 1)

def calculate_rocket_fee(base_fee, is_rocket_order, rocket_premium_percent=30):
    if is_rocket_order:
        return base_fee * (rocket_premium_percent / 100.0)
    return 0.0

def calculate_total_fee(base_fee, surge_fee, distance_fee, weather_fee, rocket_fee, min_fee=20, max_fee=150):
    fee = base_fee + surge_fee + distance_fee + weather_fee + rocket_fee
    if fee < min_fee:
        fee = min_fee
    elif fee > max_fee:
        fee = max_fee
    return round(fee, 2)

def generate_fee_breakdown(base_fee, surge_fee, distance_fee, weather_fee, rocket_fee, total_fee):
    return {
        'base_fee': base_fee,
        'surge_fee': surge_fee,
        'distance_fee': distance_fee,
        'weather_fee': weather_fee,
        'rocket_fee': rocket_fee,
        'total_fee': total_fee
    }


# Example test / simulation
def test_fee_engine():
    base = calculate_base_fee()
    surge = calculate_surge_fee(num_orders=50, peak_orders=100, base_fee=base)
    distance = calculate_distance_fee(distance_km=7)
    weather = calculate_weather_fee(base_fee=base, weather_condition='rain')
    rocket = calculate_rocket_fee(base_fee=base, is_rocket_order=True)
    
    total = calculate_total_fee(base, surge, distance, weather, rocket)
    breakdown = generate_fee_breakdown(base, surge, distance, weather, rocket, total)
    print(breakdown)

if __name__ == "__main__":
    test_fee_engine()
