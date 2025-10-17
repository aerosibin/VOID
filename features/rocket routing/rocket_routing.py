# ====================================================
# 🚀 Priority “Rocket” Routing - Terminal Version
# Author: Narenkumar
# ====================================================

import numpy as np
import pandas as pd
import tensorflow as tf
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, Dense
import random

# -------------------------------
# Step 1: Simulate Traffic Data
# -------------------------------
np.random.seed(42)

def generate_traffic_data(num_records=1000):
    timestamps = pd.date_range("2025-01-01", periods=num_records, freq="H")
    locations = ["Zone_A", "Zone_B", "Zone_C"]
    data = []

    for t in timestamps:
        for loc in locations:
            base_congestion = {"Zone_A": 0.3, "Zone_B": 0.6, "Zone_C": 0.8}[loc]
            time_factor = np.sin(2 * np.pi * (t.hour / 24))
            weather_factor = random.choice([0.9, 1.0, 1.1])
            congestion = min(1.0, max(0.0, base_congestion + 0.3 * time_factor * weather_factor))
            data.append([t, loc, congestion])

    df = pd.DataFrame(data, columns=["timestamp", "location", "congestion_level"])
    return df

df = generate_traffic_data()
print("✅ Generated traffic dataset:", df.shape)

# -------------------------------
# Step 2: Prepare Data for LSTM
# -------------------------------
def preprocess_for_lstm(df, location="Zone_A"):
    df_zone = df[df["location"] == location].copy()
    df_zone["hour"] = df_zone["timestamp"].dt.hour
    df_zone["hour_sin"] = np.sin(2 * np.pi * df_zone["hour"] / 24)
    df_zone["hour_cos"] = np.cos(2 * np.pi * df_zone["hour"] / 24)

    features = ["hour_sin", "hour_cos", "congestion_level"]
    sequence_length = 24

    X, y = [], []
    data = df_zone[features].values
    for i in range(len(data) - sequence_length):
        X.append(data[i:i+sequence_length])
        y.append(data[i+sequence_length][2])
    return np.array(X), np.array(y)

X_A, y_A = preprocess_for_lstm(df, "Zone_A")

# -------------------------------
# Step 3: Train LSTM Model
# -------------------------------
model = Sequential([
    LSTM(32, input_shape=(X_A.shape[1], X_A.shape[2]), return_sequences=True),
    LSTM(16),
    Dense(1)
])

model.compile(optimizer='adam', loss='mse')
model.fit(X_A, y_A, epochs=5, batch_size=32, verbose=0)
print("✅ Model training complete")

# -------------------------------
# Step 4: Predict Congestion
# -------------------------------
def predict_congestion(location):
    X, _ = preprocess_for_lstm(df, location)
    pred = model.predict(X[-1].reshape(1, X.shape[1], X.shape[2]), verbose=0)
    return float(pred[0][0])

# -------------------------------
# Step 5: Route Simulation
# -------------------------------
def simulate_route(route_zones, rocket=False):
    route_info = []
    total_time = 0.0

    for zone in route_zones:
        congestion = predict_congestion(zone)
        base_speed = {"Zone_A": 40, "Zone_B": 35, "Zone_C": 25}[zone]  # km/h
        road_length = random.uniform(1.0, 2.0)

        effective_speed = base_speed * (1 - congestion * 0.7)
        time_minutes = (road_length / effective_speed) * 60

        if rocket:
            time_minutes *= 0.8  # 20% faster in rocket mode

        route_info.append({
            "zone": zone,
            "congestion_level": round(congestion, 2),
            "segment_length_km": round(road_length, 2),
            "time_minutes": round(time_minutes, 2)
        })
        total_time += time_minutes

    return route_info, round(total_time, 2)

# -------------------------------
# Step 6: Compare Routes
# -------------------------------
route_1 = ["Zone_A", "Zone_B", "Zone_C"]
route_2 = ["Zone_B", "Zone_A", "Zone_C"]

for mode in [False, True]:
    label = "ROCKET MODE" if mode else "NORMAL MODE"
    print("\n==============================")
    print(f"🚦 {label}")
    print("==============================")

    info_1, time_1 = simulate_route(route_1, rocket=mode)
    info_2, time_2 = simulate_route(route_2, rocket=mode)

    best_route = route_1 if time_1 < time_2 else route_2
    best_time = min(time_1, time_2)
    best_info = info_1 if time_1 < time_2 else info_2

    print(f"Best Route: {best_route}")
    print(f"Total Estimated Time: {best_time} minutes\n")
    for seg in best_info:
        print(f"  - {seg['zone']} | Congestion: {seg['congestion_level']} | "
              f"Length: {seg['segment_length_km']} km | Time: {seg['time_minutes']} min")

print("\n✅ Simulation complete.")
