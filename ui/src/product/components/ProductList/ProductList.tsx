import React from "react";
import type { ProductType } from "../../types/product";
import { Grid } from "./Styled";
import { Product } from "../Product/Product";

const products: ProductType[] = [
  {
    id: "94ef6f2e-6bc2-4095-9a25-ce0c428ee27d",
    name: "Waffle with Berries",
    price: 6.5,
    priceInCents: 650,
    imageUrl:
      "https://images.unsplash.com/photo-1622970976780-c155a42ba86d?w=300&h=400",
    category: "Waffle",
  },
  {
    id: "bdb4b94f-655f-44ea-bbc1-0132c9dc2b18",
    name: "Vanilla Bean Crème Brulee",
    price: 7,
    priceInCents: 700,
    imageUrl:
      "https://images.unsplash.com/photo-1650419741906-1cdead9c9b4f?w=300&h=400",
    category: "Creme Brulee",
  },
];

export const ProductList: React.FC = () => {
  return (
    <div>
      <h2>Desserts</h2>
      <Grid>
        {products.map((product) => (
          <Product key={product.id} product={product} />
        ))}
      </Grid>
    </div>
  );
};

