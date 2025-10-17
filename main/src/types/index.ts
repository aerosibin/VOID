export interface Product {
  id: string;
  name: string;
  price: number;
  category: string;
  image: string;
  unit: string;
  inStock: boolean;
}

export interface CartItem {
  product: Product;
  quantity: number;
}

export interface PricingBreakdown {
  baseFee: number;
  distanceFee: number;
  surgeFee: number;
  weatherFee: number;
  rocketFee: number;
  totalFee: number;
}

export interface Recommendation {
  type: 'habit' | 'pairing' | 'trending' | 'ml';
  products: Product[];
  label: string;
}