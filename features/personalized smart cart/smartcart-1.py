import pandas as pd
from collections import Counter
import itertools
from sklearn.metrics.pairwise import cosine_similarity
import numpy as np

# Load datasets
# NOTE: You MUST have 'sales.csv' and 'skus.csv' in the same directory.
try:
    sales_df = pd.read_csv('sales.csv')  # date, sku, units_sold, price
    skus_df = pd.read_csv('skus.csv')    # sku, price, base_daily_demand
except FileNotFoundError as e:
    print(f"Error: Missing CSV file. Make sure '{e.filename}' is present.")
    # Create dummy data for demonstration if files are missing
    sales_data = {
        'date': ['2023-01-01', '2023-01-01', '2023-01-02', '2023-01-02', '2023-01-03', '2023-01-03', '2023-01-03'],
        'sku': ['SKU0009', 'SKU0012', 'SKU0009', 'SKU0034', 'SKU0012', 'SKU0034', 'SKU0050'],
        'units_sold': [2, 1, 3, 1, 5, 2, 1],
        'price': [10.0, 5.0, 10.0, 20.0, 5.0, 20.0, 15.0]
    }
    sku_data = {
        'sku': ['SKU0009', 'SKU0012', 'SKU0034', 'SKU0050', 'SKU0070', 'SKU0080'],
        'price': [10.0, 5.0, 20.0, 15.0, 30.0, 50.0],
        'base_daily_demand': [50, 40, 60, 25, 10, 5]
    }
    sales_df = pd.DataFrame(sales_data)
    skus_df = pd.DataFrame(sku_data)


# --- Helper Functions (Unchanged) ---

def get_habit_reminders(user_history, current_cart, freq_threshold=2):
    """Suggests items frequently purchased by the user but not in the current cart."""
    freq_items = Counter(user_history)
    return [sku for sku, freq in freq_items.items() if freq >= freq_threshold and sku not in current_cart]

def get_trending_skus(skus_df, top_n=5):
    """Suggests the top N trending SKUs based on base daily demand."""
    trending = skus_df.sort_values('base_daily_demand', ascending=False).head(top_n)
    return trending['sku'].tolist()

def get_frequent_pairs(sales_df, min_count=2): # Reduced min_count for dummy data
    """Finds frequently purchased item pairs (used for association rules)."""
    # Use 'date' as a proxy for the 'basket' or transaction ID
    basket_dict = sales_df.groupby('date')['sku'].apply(list) 
    pair_counter = Counter()
    for basket in basket_dict:
        # Generate combinations of unique items in the basket
        for a, b in itertools.combinations(sorted(set(basket)), 2):
            pair_counter[(a, b)] += 1
    frequent = {pair: count for pair, count in pair_counter.items() if count >= min_count}
    return frequent

def get_pairing_suggestions(current_cart, frequent_pairs):
    """Suggests items that are frequent complements to items in the cart."""
    suggestions = set()
    for item in current_cart:
        for (a, b) in frequent_pairs.keys():
            if item == a and b not in current_cart:
                suggestions.add(b)
            elif item == b and a not in current_cart:
                suggestions.add(a)
    return list(suggestions)


# --- ML Replacement Functions (scikit-learn/Pandas) ---

def train_recommender(sales_df):
    """
    Calculates and returns a SKU-to-SKU similarity matrix using Cosine Similarity.
    This replaces the surprise KNNBasic model.
    """
    print("Training Item-Similarity Model...")
    # Create the Item-User Matrix (User = date, Item = sku, Value = units_sold)
    item_user_matrix = sales_df.pivot_table(
        index='sku', 
        columns='date', 
        values='units_sold'
    ).fillna(0) # Fill NaN (no purchase) with 0

    # Calculate the Item-to-Item Similarity Matrix (Cosine Similarity)
    item_similarity_matrix = cosine_similarity(item_user_matrix)
    
    # Convert back to DataFrame for easy lookup (The 'algo' object)
    item_sim_df = pd.DataFrame(
        item_similarity_matrix, 
        index=item_user_matrix.index, 
        columns=item_user_matrix.index
    )
    return item_sim_df

def get_ml_recommendations(item_sim_df, current_cart, top_n=5):
    """
    Generates recommendations based on accumulated similarity scores.
    This replaces the surprise prediction method.
    """
    # Initialize a Series to store accumulated scores for potential recommendations
    recommendation_scores = pd.Series(dtype=float).reindex(item_sim_df.index).fillna(0)
    
    # Accumulate scores from all items currently in the cart
    for item_in_cart in current_cart:
        if item_in_cart in item_sim_df.index:
            # Add similarity scores of other items to the current item
            recommendation_scores += item_sim_df[item_in_cart]

    # Remove items that are already in the cart
    recommendation_scores = recommendation_scores.drop(current_cart, errors='ignore')
    
    # Sort and return top N SKUs
    top_recs = recommendation_scores.sort_values(ascending=False).head(top_n).index.tolist()
    
    return top_recs

# --- Main Logic Function (Updated to use new ML functions) ---

def smart_cart_recommendations(user_history, current_cart, sales_df, skus_df, item_sim_df):
    """Combines different recommendation strategies into one list."""
    
    frequent_pairs = get_frequent_pairs(sales_df)
    trending_skus = get_trending_skus(skus_df)

    habit_reminders = get_habit_reminders(user_history, current_cart)
    pairings = get_pairing_suggestions(current_cart, frequent_pairs)
    trending = [sku for sku in trending_skus if sku not in current_cart]
    
    # Use the new scikit-learn based function
    ml_recs = get_ml_recommendations(item_sim_df, current_cart) 

    return {
        'habit_reminders': habit_reminders,
        'pairing_suggestions': pairings,
        'trending_nudge': trending,
        'ml_recommendations': ml_recs
    }

# ----------------- EXECUTION -----------------

# 1. Train recommendation model (run once)
# The 'algo' object is now the Item Similarity DataFrame
item_sim_df = train_recommender(sales_df)

# 2. Example user data (replace with actual user session data)
user_history = ['SKU0009', 'SKU0009', 'SKU0012', 'SKU0012', 'SKU0012', 'SKU0034']
current_cart = ['SKU0009', 'SKU0034']

# 3. Get smart cart recommendations
recommendations = smart_cart_recommendations(user_history, current_cart, sales_df, skus_df, item_sim_df)
print("\n--- Smart Cart Recommendations ---")
print(recommendations)