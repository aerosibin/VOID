import { Plus, Heart } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Product } from "@/types";
import { useAuth } from "@/contexts/AuthContext";
import { supabase } from "@/integrations/supabase/client";
import { useToast } from "@/hooks/use-toast";
import { useNavigate } from "react-router-dom";
import { useState } from "react";

// Mock data
const mockProducts: Product[] = [
  { id: "1", name: "Organic Bananas", price: 3.99, category: "Fruits", image: "🍌", unit: "per bunch", inStock: true },
  { id: "2", name: "Fresh Milk", price: 4.49, category: "Dairy", image: "🥛", unit: "1L", inStock: true },
  { id: "3", name: "Whole Wheat Bread", price: 2.99, category: "Bakery", image: "🍞", unit: "loaf", inStock: true },
  { id: "4", name: "Free Range Eggs", price: 5.99, category: "Dairy", image: "🥚", unit: "dozen", inStock: true },
  { id: "5", name: "Roma Tomatoes", price: 4.29, category: "Vegetables", image: "🍅", unit: "per lb", inStock: true },
  { id: "6", name: "Avocados", price: 6.99, category: "Fruits", image: "🥑", unit: "4 pack", inStock: true },
  { id: "7", name: "Greek Yogurt", price: 3.79, category: "Dairy", image: "🥤", unit: "500g", inStock: true },
  { id: "8", name: "Fresh Spinach", price: 2.49, category: "Vegetables", image: "🥬", unit: "bunch", inStock: true },
];

interface ProductGridProps {
  onAddToCart: (product: Product) => void;
}

const ProductGrid = ({ onAddToCart }: ProductGridProps) => {
  const { user } = useAuth();
  const { toast } = useToast();
  const navigate = useNavigate();
  const [savingItem, setSavingItem] = useState<string | null>(null);

  const handleSaveItem = async (product: Product) => {
    if (!user) {
      toast({
        title: 'Sign in required',
        description: 'Please sign in to save items',
        variant: 'destructive',
      });
      navigate('/auth');
      return;
    }

    setSavingItem(product.id);
    try {
      const { error } = await supabase
        .from('saved_items')
        .insert({
          user_id: user.id,
          product_name: product.name,
          product_price: `$${product.price}`,
          product_image: product.image,
        });

      if (error) {
        if (error.message.includes('duplicate')) {
          toast({
            title: 'Already saved',
            description: 'This item is already in your saved items',
          });
        } else {
          throw error;
        }
      } else {
        toast({
          title: 'Saved!',
          description: 'Item added to your saved items',
        });
      }
    } catch (error: any) {
      toast({
        title: 'Error',
        description: error.message,
        variant: 'destructive',
      });
    } finally {
      setSavingItem(null);
    }
  };
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
      {mockProducts.map((product) => (
        <Card 
          key={product.id} 
          className="group overflow-hidden hover:shadow-hover transition-all duration-300 hover:-translate-y-1 bg-card border-border"
        >
          <div className="p-6">
            <div className="text-6xl mb-4 text-center">{product.image}</div>
            <Badge variant="secondary" className="mb-2 bg-muted text-muted-foreground">
              {product.category}
            </Badge>
            <h3 className="font-semibold text-lg mb-1 text-foreground">{product.name}</h3>
            <p className="text-sm text-muted-foreground mb-3">{product.unit}</p>
            <div className="flex items-center justify-between gap-2">
              <span className="text-2xl font-bold text-primary">${product.price}</span>
              <div className="flex gap-2">
                <Button
                  size="icon"
                  variant="outline"
                  onClick={() => handleSaveItem(product)}
                  disabled={savingItem === product.id}
                  className="hover:text-primary"
                >
                  <Heart className="w-4 h-4" />
                </Button>
                <Button 
                  size="sm" 
                  onClick={() => onAddToCart(product)}
                  className="bg-primary hover:bg-primary-hover"
                >
                  <Plus className="w-4 h-4 mr-1" />
                  Add
                </Button>
              </div>
            </div>
          </div>
        </Card>
      ))}
    </div>
  );
};

export default ProductGrid;