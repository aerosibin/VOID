# ============================================================
# Unified AI Grocery Platform - Performance Insights Dashboard
# ============================================================

import streamlit as st
import pandas as pd
import plotly.express as px
import base64
st.set_page_config(page_title="AI Grocery - Performance Insights", layout="wide")
# -------------------------------
# Load Data
# -------------------------------
@st.cache_data
def load_data():
    orders = pd.read_csv("orders_extended.csv")
    rider_metrics = pd.read_csv("rider_metrics.csv")
    peak_hours = pd.read_csv("peak_hours.csv")
    heatmap_grid = pd.read_csv("heatmap_grid.csv")
    return orders, rider_metrics, peak_hours, heatmap_grid

orders, rider_metrics, peak_hours, heatmap_grid = load_data()

# -------------------------------
# Page Config
# -------------------------------

st.title("🧠 Unified AI Grocery Platform - Performance Insights")

# Sidebar navigation
tabs = ["Overview", "Demand Heatmap", "Rider Performance", "Delivery Analysis"]
selected_tab = st.sidebar.radio("📊 Select View", tabs)

# -------------------------------
# TAB 1: Overview
# -------------------------------
if selected_tab == "Overview":
    st.header("📋 Operational Overview")

    total_orders = len(orders)
    avg_time = round(orders["actual_delivery_time_min"].mean(), 2)
    on_time_rate = round((orders["delay"] <= 0).mean() * 100, 2)
    peak_hour = int(peak_hours.loc[peak_hours["orders_count"].idxmax(), "hour_of_day"])

    col1, col2, col3, col4 = st.columns(4)
    col1.metric("Total Orders", total_orders)
    col2.metric("Avg Delivery Time (min)", avg_time)
    col3.metric("On-Time Rate (%)", on_time_rate)
    col4.metric("Peak Hour", f"{peak_hour}:00")

    st.subheader("Orders by Hour of Day")
    fig = px.bar(
        peak_hours,
        x="hour_of_day",
        y="orders_count",
        title="Order Volume by Hour of Day",
        color="orders_count",
        color_continuous_scale="Tealgrn",
    )
    st.plotly_chart(fig, use_container_width=True)

# -------------------------------
# TAB 2: Demand Heatmap
# -------------------------------
elif selected_tab == "Demand Heatmap":
    st.header("🗺️ Demand Heatmap")
    st.markdown("Visualizing areas with highest order density.")

    # Load folium map as iframe
    with open("heatmap.html", "r", encoding="utf-8") as f:
        html_data = f.read()
        b64 = base64.b64encode(html_data.encode()).decode()
        html_iframe = f'<iframe src="data:text/html;base64,{b64}" width="100%" height="600"></iframe>'
        st.components.v1.html(html_iframe, height=600)

    st.markdown("---")
    st.subheader("Top 10 High-Demand Grid Cells")
    st.dataframe(heatmap_grid.sort_values("order_count", ascending=False).head(10))

# -------------------------------
# TAB 3: Rider Performance
# -------------------------------
elif selected_tab == "Rider Performance":
    st.header("🚴 Rider Efficiency Dashboard")

    st.subheader("Rider Performance Table")
    st.dataframe(rider_metrics)

    st.markdown("---")
    st.subheader("Average Delivery Time by Rider")
    fig = px.bar(
        rider_metrics,
        x="rider_id",
        y="avg_actual_time",
        color="avg_delay",
        title="Rider Delivery Time vs Delay",
        color_continuous_scale="OrRd",
    )
    st.plotly_chart(fig, use_container_width=True)

# -------------------------------
# TAB 4: Delivery Analysis
# -------------------------------
elif selected_tab == "Delivery Analysis":
    st.header("📦 Delivery Time Analysis")

    st.subheader("Delivery Delay Distribution")
    fig = px.histogram(
        orders,
        x="delay",
        nbins=20,
        title="Distribution of Delivery Delays (mins)",
        color_discrete_sequence=["#2E86AB"],
    )
    st.plotly_chart(fig, use_container_width=True)

    st.markdown("---")
    st.subheader("Delay vs Hour of Day")
    avg_delay_by_hour = orders.groupby("hour_of_day")["delay"].mean().reset_index()
    fig2 = px.line(
        avg_delay_by_hour,
        x="hour_of_day",
        y="delay",
        markers=True,
        title="Average Delay by Hour of Day",
    )
    st.plotly_chart(fig2, use_container_width=True)

# -------------------------------
# Footer
# -------------------------------
st.markdown("---")
st.caption("© 2025 Unified AI Grocery Prototype | Powered by Streamlit + Pandas + Plotly")
