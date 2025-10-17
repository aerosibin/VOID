import pandas as pd
from collections import Counter
import itertools
from surprise import Dataset, Reader, KNNBasic

# Load datasets
sales_df = pd.read_csv('sales.csv')  # date, sku, units_sold, price
skus_df = pd.read_csv('skus.csv')    # sku, price, base_daily_demand

def get_habit_reminders(user_history, current_cart, freq_threshold=2):
    freq_items = Counter(user_history)
    return [sku for sku, freq in freq_items.items() if freq >= freq_threshold and sku not in current_cart]

def get_trending_skus(skus_df, top_n=5):
    trending = skus_df.sort_values('base_daily_demand', ascending=False).head(top_n)
    return trending['sku'].tolist()

def get_frequent_pairs(sales_df, min_count=10):
    basket_dict = sales_df.groupby('date')['sku'].apply(list)
    pair_counter = Counter()
    for basket in basket_dict:
        for a, b in itertools.combinations(sorted(set(basket)), 2):
            pair_counter[(a, b)] += 1
    frequent = {pair: count for pair, count in pair_counter.items() if count >= min_count}
    return frequent

def get_pairing_suggestions(current_cart, frequent_pairs):
    suggestions = set()
    for item in current_cart:
        for (a, b) in frequent_pairs.keys():
            if item == a and b not in current_cart:
                suggestions.add(b)
            elif item == b and a not in current_cart:
                suggestions.add(a)
    return list(suggestions)

def train_recommender(sales_df):
    sales_df = sales_df.copy()
    sales_df['user_id'] = sales_df['date']  # placeholder user id; replace with real user id if available
    reader = Reader(rating_scale=(1, 10))
    data = Dataset.load_from_df(sales_df[['user_id', 'sku', 'units_sold']], reader)
    trainset = data.build_full_trainset()
    sim_options = {'name': 'cosine', 'user_based': False}  # item-based CF
    algo = KNNBasic(sim_options=sim_options)
    algo.fit(trainset)
    return algo

def get_ml_recommendations(algo, current_cart, top_n=5):
    dummy_user = 'dummy'
    all_skus = skus_df['sku'].tolist()
    unseen_skus = [sku for sku in all_skus if sku not in current_cart]
    preds = [(sku, algo.predict(dummy_user, sku).est) for sku in unseen_skus]
    preds.sort(key=lambda x: x[1], reverse=True)
    return [sku for sku, _ in preds[:top_n]]

def smart_cart_recommendations(user_history, current_cart, sales_df, skus_df, algo):
    frequent_pairs = get_frequent_pairs(sales_df)
    trending_skus = get_trending_skus(skus_df)

    habit_reminders = get_habit_reminders(user_history, current_cart)
    pairings = get_pairing_suggestions(current_cart, frequent_pairs)
    trending = [sku for sku in trending_skus if sku not in current_cart]
    ml_recs = get_ml_recommendations(algo, current_cart)

    return {
        'habit_reminders': habit_reminders,
        'pairing_suggestions': pairings,
        'trending_nudge': trending,
        'ml_recommendations': ml_recs
    }

# Train recommendation model (run once on system start or schedule)
algo = train_recommender(sales_df)

# Example user data (replace with actual user session data)
user_history = ['SKU0009', 'SKU0009', 'SKU0012', 'SKU0012', 'SKU0012', 'SKU0034']
current_cart = ['SKU0009', 'SKU0034']

# Get smart cart recommendations
recommendations = smart_cart_recommendations(user_history, current_cart, sales_df, skus_df, algo)
print(recommendations)
