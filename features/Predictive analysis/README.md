# 🗺️ Real-Time Geospatial & Fleet Analytics

This feature set is designed to process and visualize real-time data streams (`stream.py`) for critical operational insights, including demand heatmaps and rider performance monitoring.

## 🚀 Key Objectives

* **Real-Time Demand Visualization:** Generate and update heatmaps to show where demand (orders) is highest across the service area.
* **Fleet Performance Monitoring:** Track and analyze key metrics for riders to optimize dispatch and service quality.
* **Data Aggregation:** Process raw streaming data and aggregate it into structured files (CSV) for analysis, reporting, and dashboard generation.

---

## ⚙️ Implementation and Data Files

The `stream.py` script likely acts as the central processing unit, consuming data (potentially from a live source matching the CSV schemas) and updating the derived analytic files.

| File Name | Role in Feature | Data Type & Use |
| :--- | :--- | :--- |
| **`stream.py`** | **Core Processor.** Connects to a raw data stream (e.g., live orders, GPS pings), performs aggregations, and writes the results to the various CSV files below. |
| **`heatmap_grid.csv`** | **Geospatial Input.** Contains aggregated data on demand intensity per geographical grid cell. This is the primary input for the visualization. |
| **`heatmap.html`** | **Visualization Output.** The final, self-contained HTML file used to render the dynamic demand heatmap based on the data in `heatmap_grid.csv`. |
| **`orders_extended.csv`** | **Detailed Order Data.** Contains rich, transactional data for every order, likely including delivery time, pick-up/drop-off coordinates, and other logistics details. |
| **`peak_hours.csv`** | **Demand Analysis.** Aggregated data showing demand volume broken down by time slots (e.g., hour of the day, day of the week), critical for forecasting and dynamic pricing. |
| **`rider_metrics.csv`** | **Performance Output.** Key performance indicators (KPIs) calculated per rider (e.g., average delivery time, acceptance rate, completion rate). |
| **`riders.csv`** | **Fleet Master Data.** Information on all active drivers/riders, including ID, current status, and perhaps vehicle type. |
| **`sales.csv` & `skus.csv`** | **Supporting Data.** Used for enriching order data (`orders_extended.csv`) to calculate value-based metrics or to tie sales volume to demand analysis. |

---

## 🧠 Data Flow and Analytics Pipeline

The feature operates as a continuous pipeline:

1.  **Data Ingestion:** `stream.py` receives raw data (e.g., a new order, a rider's location update).
2.  **Transformation:**
    * **Orders:** New order data is enriched using `skus.csv` and written to `orders_extended.csv`.
    * **Time Analysis:** Order data is aggregated to update time slots in `peak_hours.csv`.
    * **Geospatial Aggregation:** Order locations are mapped to grid cells, updating the demand counts in `heatmap_grid.csv`.
    * **Rider Performance:** Rider activity is analyzed against `riders.csv` to calculate and update `rider_metrics.csv`.
3.  **Visualization:** The `heatmap.html` file consumes the latest `heatmap_grid.csv` data to display the real-time demand landscape for operational dispatchers.

This structure ensures that real-time operational decisions (like where to stage riders) are driven by the freshest available data.
<img width="1600" height="800" alt="delay_hist" src="https://github.com/user-attachments/assets/28e2dc45-fbf7-460b-9d9c-42c23d6a5bb3" />
<img width="1600" height="800" alt="peak_hours" src="https://github.com/user-attachments/assets/a5c366a8-32a6-4707-87b4-2681b128ad21" />
