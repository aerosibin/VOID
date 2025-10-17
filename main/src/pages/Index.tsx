import { useState } from "react";
import { ShoppingCart, Zap, TrendingUp, Clock, MapPin } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Input } from "@/components/ui/input";
import Navbar from "@/components/Navbar";
import ProductGrid from "@/components/ProductGrid";
import CartSidebar from "@/components/CartSidebar";
import { Product } from "@/types";

const Index = () => {
  const [cartOpen, setCartOpen] = useState(false);
  const [cart, setCart] = useState<{ product: Product; quantity: number }[]>([]);

  const addToCart = (product: Product) => {
    setCart(prev => {
      const existing = prev.find(item => item.product.id === product.id);
      if (existing) {
        return prev.map(item =>
          item.product.id === product.id
            ? { ...item, quantity: item.quantity + 1 }
            : item
        );
      }
      return [...prev, { product, quantity: 1 }];
    });
  };

  const updateQuantity = (productId: string, quantity: number) => {
    if (quantity === 0) {
      setCart(prev => prev.filter(item => item.product.id !== productId));
    } else {
      setCart(prev =>
        prev.map(item =>
          item.product.id === productId ? { ...item, quantity } : item
        )
      );
    }
  };

  const cartItemsCount = cart.reduce((sum, item) => sum + item.quantity, 0);

  return (
    <div className="min-h-screen bg-gradient-subtle">
      <Navbar />
      
      {/* Hero Section */}
      <section className="bg-gradient-hero text-primary-foreground py-20 px-4">
        <div className="container mx-auto max-w-6xl">
          <div className="flex flex-col md:flex-row items-center gap-12">
            <div className="flex-1 space-y-6 animate-fade-in">
              <Badge className="bg-accent text-accent-foreground border-0 text-sm px-4 py-1.5">
                <Zap className="w-4 h-4 mr-1.5" />
                AI-Powered Delivery
              </Badge>
              <h1 className="text-5xl md:text-6xl font-bold leading-tight">
                Fresh Groceries, <br/>Delivered Smart
              </h1>
              <p className="text-lg md:text-xl opacity-90 max-w-xl">
                Experience the future of grocery delivery with AI-optimized routing, 
                dynamic pricing, and personalized recommendations.
              </p>
              <div className="flex flex-wrap gap-6 pt-4">
                <div className="flex items-center gap-2">
                  <Clock className="w-5 h-5" />
                  <span className="text-sm font-medium">15-min Rocket Delivery</span>
                </div>
                <div className="flex items-center gap-2">
                  <TrendingUp className="w-5 h-5" />
                  <span className="text-sm font-medium">Smart Recommendations</span>
                </div>
                <div className="flex items-center gap-2">
                  <MapPin className="w-5 h-5" />
                  <span className="text-sm font-medium">Live Tracking</span>
                </div>
              </div>
            </div>
            <div className="flex-1 relative">
              <div className="relative w-full h-[400px] bg-primary-foreground/10 rounded-2xl backdrop-blur-sm border border-primary-foreground/20 flex items-center justify-center">
                <ShoppingCart className="w-32 h-32 text-primary-foreground/40" />
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Search Bar */}
      <section className="container mx-auto max-w-6xl px-4 -mt-8">
        <div className="bg-card shadow-hover rounded-lg p-4">
          <Input 
            placeholder="Search for groceries..." 
            className="text-lg h-12 border-0 focus-visible:ring-1"
          />
        </div>
      </section>

      {/* Products Section */}
      <section className="container mx-auto max-w-6xl px-4 py-12">
        <div className="flex items-center justify-between mb-8">
          <h2 className="text-3xl font-bold">Fresh Picks for You</h2>
          <Button 
            onClick={() => setCartOpen(true)} 
            className="bg-primary hover:bg-primary-hover"
          >
            <ShoppingCart className="w-5 h-5 mr-2" />
            Cart {cartItemsCount > 0 && `(${cartItemsCount})`}
          </Button>
        </div>
        <ProductGrid onAddToCart={addToCart} />
      </section>

      <CartSidebar 
        open={cartOpen} 
        onClose={() => setCartOpen(false)}
        cart={cart}
        onUpdateQuantity={updateQuantity}
      />
    </div>
  );
};

export default Index;